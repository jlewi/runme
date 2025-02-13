//go:build !windows

package runnerv2service

import (
	"bytes"
	"context"
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/stateful/runme/v3/internal/command"
	runnerv2alpha1 "github.com/stateful/runme/v3/internal/gen/proto/go/runme/runner/v2alpha1"
)

func init() {
	// Set to false to disable sending signals to process groups in tests.
	// This can be turned on if setSysProcAttrPgid() is called in Start().
	command.SignalToProcessGroup = false

	command.EnvDumpCommand = "env -0"
}

func Test_conformsOpinionatedEnvVarNaming(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		assert.True(t, conformsOpinionatedEnvVarNaming("TEST"))
		assert.True(t, conformsOpinionatedEnvVarNaming("ABC"))
		assert.True(t, conformsOpinionatedEnvVarNaming("TEST_ABC"))
		assert.True(t, conformsOpinionatedEnvVarNaming("ABC_123"))
	})

	t.Run("lowercase is invalid", func(t *testing.T) {
		assert.False(t, conformsOpinionatedEnvVarNaming("test"))
		assert.False(t, conformsOpinionatedEnvVarNaming("abc"))
		assert.False(t, conformsOpinionatedEnvVarNaming("test_abc"))
		assert.False(t, conformsOpinionatedEnvVarNaming("abc_123"))
	})

	t.Run("too short", func(t *testing.T) {
		assert.False(t, conformsOpinionatedEnvVarNaming("AB"))
		assert.False(t, conformsOpinionatedEnvVarNaming("T"))
	})

	t.Run("numbers only is invalid", func(t *testing.T) {
		assert.False(t, conformsOpinionatedEnvVarNaming("123"))
		assert.False(t, conformsOpinionatedEnvVarNaming("8761123"))
		assert.False(t, conformsOpinionatedEnvVarNaming("138761123"))
	})

	t.Run("invalid characters", func(t *testing.T) {
		assert.False(t, conformsOpinionatedEnvVarNaming("ABC_%^!"))
		assert.False(t, conformsOpinionatedEnvVarNaming("&^%$"))
		assert.False(t, conformsOpinionatedEnvVarNaming("A@#$%"))
	})
}

func TestRunnerServiceServerExecute_Response(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	stream, err := client.Execute(context.Background())
	require.NoError(t, err)

	req := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo test | tee >(cat >&2)",
					},
				},
			},
		},
	}

	err = stream.Send(req)
	require.NoError(t, err)

	// Assert first response.
	resp, err := stream.Recv()
	assert.NoError(t, err)
	assert.Greater(t, resp.Pid.Value, uint32(1))
	assert.Nil(t, resp.ExitCode)

	// Assert second and third responses.
	var (
		out      bytes.Buffer
		mimeType string
	)

	resp, err = stream.Recv()
	assert.NoError(t, err)
	assert.Nil(t, resp.ExitCode)
	assert.Nil(t, resp.Pid)
	_, err = out.Write(resp.StdoutData)
	assert.NoError(t, err)
	_, err = out.Write(resp.StderrData)
	assert.NoError(t, err)
	if resp.MimeType != "" {
		mimeType = resp.MimeType
	}

	resp, err = stream.Recv()
	assert.NoError(t, err)
	assert.Nil(t, resp.ExitCode)
	assert.Nil(t, resp.Pid)
	_, err = out.Write(resp.StdoutData)
	assert.NoError(t, err)
	_, err = out.Write(resp.StderrData)
	assert.NoError(t, err)
	if resp.MimeType != "" {
		mimeType = resp.MimeType
	}

	assert.Contains(t, mimeType, "text/plain")
	assert.Equal(t, "test\ntest\n", out.String())

	// Assert fourth response.
	resp, err = stream.Recv()
	assert.NoError(t, err)
	assert.Equal(t, uint32(0), resp.ExitCode.Value)
	assert.Nil(t, resp.Pid)
}

func TestRunnerServiceServerExecute_MimeType(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	stream, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult := make(chan executeResult)
	go getExecuteResult(stream, execResult)

	req := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						// Echo JSON to stderr and plain text to stdout.
						// Only the plain text should be detected.
						">&2 echo '{\"field1\": \"value\", \"field2\": 2}'",
						"echo 'some plain text'",
					},
				},
			},
		},
	}

	err = stream.Send(req)
	assert.NoError(t, err)

	result := <-execResult

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "{\"field1\": \"value\", \"field2\": 2}\n", string(result.Stderr))
	assert.Equal(t, "some plain text\n", string(result.Stdout))
	assert.Contains(t, result.MimeType, "text/plain")
}

func TestRunnerServiceServerExecute_StoreLastStdout(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	sessionResp, err := client.CreateSession(context.Background(), &runnerv2alpha1.CreateSessionRequest{})
	require.NoError(t, err)
	require.NotNil(t, sessionResp.Session)

	stream1, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult1 := make(chan executeResult)
	go getExecuteResult(stream1, execResult1)

	req1 := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo test | tee >(cat >&2)",
					},
				},
			},
		},
		SessionId:        sessionResp.GetSession().GetId(),
		StoreStdoutInEnv: true,
	}

	err = stream1.Send(req1)
	assert.NoError(t, err)

	result := <-execResult1

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "test\n", string(result.Stdout))
	assert.Contains(t, result.MimeType, "text/plain")

	// subsequent req to check last stored value
	stream2, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult2 := make(chan executeResult)
	go getExecuteResult(stream2, execResult2)

	req2 := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo $__",
					},
				},
			},
		},
		SessionId: sessionResp.GetSession().GetId(),
	}

	err = stream2.Send(req2)
	assert.NoError(t, err)

	result = <-execResult2

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "test\n", string(result.Stdout))
	assert.Contains(t, result.MimeType, "text/plain")
}

func TestRunnerServiceServerExecute_StoreKnownName(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	sessionResp, err := client.CreateSession(context.Background(), &runnerv2alpha1.CreateSessionRequest{})
	require.NoError(t, err)
	require.NotNil(t, sessionResp.Session)

	stream1, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult1 := make(chan executeResult)
	go getExecuteResult(stream1, execResult1)

	req1 := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo test | tee >(cat >&2)",
					},
				},
			},
			KnownName: "TEST_VAR",
		},
		SessionId:        sessionResp.GetSession().GetId(),
		StoreStdoutInEnv: true,
	}

	err = stream1.Send(req1)
	assert.NoError(t, err)

	result := <-execResult1

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "test\n", string(result.Stdout))
	assert.Contains(t, result.MimeType, "text/plain")

	// subsequent req to check last stored value
	stream2, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult2 := make(chan executeResult)
	go getExecuteResult(stream2, execResult2)

	req2 := &runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo $TEST_VAR",
					},
				},
			},
		},
		SessionId: sessionResp.GetSession().GetId(),
	}

	err = stream2.Send(req2)
	assert.NoError(t, err)

	result = <-execResult2

	assert.NoError(t, result.Err)
	assert.EqualValues(t, 0, result.ExitCode)
	assert.Equal(t, "test\n", string(result.Stdout))
	assert.Contains(t, result.MimeType, "text/plain")
}

func TestRunnerServiceServerExecute_Configs(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	testCases := []struct {
		name           string
		programConfig  *runnerv2alpha1.ProgramConfig
		inputData      []byte
		expectedOutput string
	}{
		{
			name: "Basic",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"echo test",
						},
					},
				},
			},
			expectedOutput: "test\n",
		},
		{
			name: "BasicInteractive",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"echo test",
						},
					},
				},
				Interactive: true,
			},
			expectedOutput: "test\r\n",
		},
		{
			name: "BasicSleep",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"echo 1",
							"sleep 1",
							"echo 2",
						},
					},
				},
			},
			expectedOutput: "1\n2\n",
		},
		{
			name: "BasicInput",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"read name",
							"echo \"My name is $name\"",
						},
					},
				},
				Interactive: false,
			},
			inputData:      []byte("Frank\n"),
			expectedOutput: "My name is Frank\n",
		},
		{
			name: "BasicInputInteractive",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"read name",
							"echo \"My name is $name\"",
						},
					},
				},
				Interactive: true,
			},
			inputData:      []byte("Frank\n"),
			expectedOutput: "My name is Frank\r\n",
		},
		{
			name: "Python",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "",
				LanguageId:  "py",
				Source: &runnerv2alpha1.ProgramConfig_Script{
					Script: "print('test')",
				},
				Interactive: true,
			},
			expectedOutput: "test\r\n",
		},
		{
			name: "Javascript",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "node",
				Source: &runnerv2alpha1.ProgramConfig_Script{
					Script: "console.log('1');\nconsole.log('2');",
				},
			},
			expectedOutput: "1\n2\n",
		},
		{
			name: "Javascript_inferred",
			programConfig: &runnerv2alpha1.ProgramConfig{
				ProgramName: "",
				LanguageId:  "js",
				Source: &runnerv2alpha1.ProgramConfig_Script{
					Script: "console.log('1');\nconsole.log('2');",
				},
			},
			expectedOutput: "1\n2\n",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			stream, err := client.Execute(context.Background())
			require.NoError(t, err)

			execResult := make(chan executeResult)
			go getExecuteResult(stream, execResult)

			req := &runnerv2alpha1.ExecuteRequest{
				Config: tc.programConfig,
			}

			if tc.inputData != nil {
				req.InputData = tc.inputData
			}

			err = stream.Send(req)
			assert.NoError(t, err)

			result := <-execResult

			assert.NoError(t, result.Err)
			assert.Equal(t, tc.expectedOutput, string(result.Stdout))
			assert.EqualValues(t, 0, result.ExitCode)
		})
	}
}

func TestRunnerServiceServerExecute_CommandMode_Terminal(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	sessResp, err := client.CreateSession(context.Background(), &runnerv2alpha1.CreateSessionRequest{})
	require.NoError(t, err)

	// Step 1: execute the first command in the terminal mode with bash,
	// then write a line that exports an environment variable.
	{
		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"bash",
						},
					},
				},
				Mode: runnerv2alpha1.CommandMode_COMMAND_MODE_TERMINAL,
			},
			SessionId: sessResp.GetSession().GetId(),
		})
		require.NoError(t, err)

		time.Sleep(time.Second)

		// Export some variables so that it can be tested if they are collected.
		req := &runnerv2alpha1.ExecuteRequest{InputData: []byte("export TEST_ENV=TEST_VALUE\n")}
		err = stream.Send(req)
		require.NoError(t, err)
		// Signal the end of input.
		req = &runnerv2alpha1.ExecuteRequest{InputData: []byte{0x04}}
		err = stream.Send(req)
		require.NoError(t, err)

		result := <-execResult
		require.NoError(t, result.Err)
	}

	// Step 2: execute the second command which will try to get the value of
	// the exported environment variable from the step 1.
	{
		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"echo -n $TEST_ENV",
						},
					},
				},
				Mode: runnerv2alpha1.CommandMode_COMMAND_MODE_INLINE,
			},
			SessionId: sessResp.GetSession().GetId(),
		})
		require.NoError(t, err)

		result := <-execResult
		require.NoError(t, result.Err)
		require.Equal(t, "TEST_VALUE", string(result.Stdout))
	}
}

func TestRunnerServiceServerExecute_WithInput(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	t.Run("ContinuousInput", func(t *testing.T) {
		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"cat - | tr a-z A-Z",
						},
					},
				},
				Interactive: true,
			},
			InputData: []byte("a\n"),
		})
		require.NoError(t, err)

		for _, data := range [][]byte{[]byte("b\n"), []byte("c\n"), []byte("d\n"), {0x04}} {
			req := &runnerv2alpha1.ExecuteRequest{InputData: data}
			err = stream.Send(req)
			assert.NoError(t, err)
		}

		result := <-execResult

		assert.NoError(t, result.Err)
		assert.Equal(t, "A\r\nB\r\nC\r\nD\r\n", string(result.Stdout))
		assert.EqualValues(t, 0, result.ExitCode)
	})

	t.Run("SimulateCtrlC", func(t *testing.T) {
		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"bash",
						},
					},
				},
				Interactive: true,
			},
		})
		require.NoError(t, err)

		time.Sleep(time.Millisecond * 500)
		err = stream.Send(&runnerv2alpha1.ExecuteRequest{InputData: []byte("sleep 30")})
		assert.NoError(t, err)

		// cancel sleep
		time.Sleep(time.Millisecond * 500)
		err = stream.Send(&runnerv2alpha1.ExecuteRequest{InputData: []byte{0x03}})
		assert.NoError(t, err)

		// terminate shell
		time.Sleep(time.Millisecond * 500)
		err = stream.Send(&runnerv2alpha1.ExecuteRequest{InputData: []byte{0x04}})
		assert.NoError(t, err)

		result := <-execResult

		// TODO(adamb): This should be a specific gRPC error rather than Unknown.
		assert.Contains(t, result.Err.Error(), "exit status 130")
		assert.Equal(t, 130, result.ExitCode)
	})

	t.Run("CloseSendDirection", func(t *testing.T) {
		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "sleep",
				Arguments:   []string{"30"},
				Mode:        runnerv2alpha1.CommandMode_COMMAND_MODE_INLINE,
			},
		})
		require.NoError(t, err)

		// Close the send direction.
		assert.NoError(t, stream.CloseSend())

		result := <-execResult
		// TODO(adamb): This should be a specific gRPC error rather than Unknown.
		require.NotNil(t, result.Err)
		assert.Contains(t, result.Err.Error(), "signal: interrupt")
		assert.Equal(t, 130, result.ExitCode)
	})
}

func TestRunnerServiceServerExecute_WithSession(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	t.Run("WithEnvAndMostRecentSessionStrategy", func(t *testing.T) {
		{
			stream, err := client.Execute(context.Background())
			require.NoError(t, err)

			execResult := make(chan executeResult)
			go getExecuteResult(stream, execResult)

			err = stream.Send(&runnerv2alpha1.ExecuteRequest{
				Config: &runnerv2alpha1.ProgramConfig{
					ProgramName: "bash",
					Source: &runnerv2alpha1.ProgramConfig_Commands{
						Commands: &runnerv2alpha1.ProgramConfig_CommandList{
							Items: []string{
								"echo -n \"$TEST_ENV\"",
								"export TEST_ENV=hello-2",
							},
						},
					},
					Env: []string{"TEST_ENV=hello"},
				},
			})
			require.NoError(t, err)

			result := <-execResult

			assert.NoError(t, result.Err)
			assert.Equal(t, "hello", string(result.Stdout))
		}

		{
			// Execute again with most recent session.
			stream, err := client.Execute(context.Background())
			require.NoError(t, err)

			execResult := make(chan executeResult)
			go getExecuteResult(stream, execResult)

			err = stream.Send(&runnerv2alpha1.ExecuteRequest{
				Config: &runnerv2alpha1.ProgramConfig{
					ProgramName: "bash",
					Source: &runnerv2alpha1.ProgramConfig_Commands{
						Commands: &runnerv2alpha1.ProgramConfig_CommandList{
							Items: []string{
								"echo -n \"$TEST_ENV\"",
							},
						},
					},
				},
				SessionStrategy: runnerv2alpha1.SessionStrategy_SESSION_STRATEGY_MOST_RECENT,
			})
			require.NoError(t, err)

			result := <-execResult

			assert.NoError(t, result.Err)
			assert.Equal(t, "hello-2", string(result.Stdout))
		}
	})
}

func TestRunnerServiceServerExecute_WithStop(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	stream, err := client.Execute(context.Background())
	require.NoError(t, err)

	execResult := make(chan executeResult)
	go getExecuteResult(stream, execResult)

	err = stream.Send(&runnerv2alpha1.ExecuteRequest{
		Config: &runnerv2alpha1.ProgramConfig{
			ProgramName: "bash",
			Source: &runnerv2alpha1.ProgramConfig_Commands{
				Commands: &runnerv2alpha1.ProgramConfig_CommandList{
					Items: []string{
						"echo 1",
						"sleep 30",
					},
				},
			},
			Interactive: true,
		},
		InputData: []byte("a\n"),
	})
	require.NoError(t, err)

	errc := make(chan error)
	go func() {
		defer close(errc)
		time.Sleep(500 * time.Millisecond)
		err := stream.Send(&runnerv2alpha1.ExecuteRequest{
			Stop: runnerv2alpha1.ExecuteStop_EXECUTE_STOP_INTERRUPT,
		})
		errc <- err
	}()
	require.NoError(t, <-errc)

	result := <-execResult

	// TODO(adamb): There should be no error.
	assert.Contains(t, result.Err.Error(), "signal: interrupt")
	assert.Equal(t, 130, result.ExitCode)
}

func TestRunnerServiceServerExecute_Winsize(t *testing.T) {
	lis, stop := testStartRunnerServiceServer(t)
	t.Cleanup(stop)
	_, client := testCreateRunnerServiceClient(t, lis)

	t.Run("DefaultWinsize", func(t *testing.T) {
		t.Parallel()

		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"tput lines",
							"tput cols",
						},
					},
				},
				Env:         []string{"TERM=linux"},
				Interactive: true,
			},
		})
		require.NoError(t, err)

		result := <-execResult

		assert.NoError(t, result.Err)
		assert.Equal(t, "24\r\n80\r\n", string(result.Stdout))
		assert.EqualValues(t, 0, result.ExitCode)
	})

	t.Run("SetWinsizeInInitialRequest", func(t *testing.T) {
		t.Parallel()

		stream, err := client.Execute(context.Background())
		require.NoError(t, err)

		execResult := make(chan executeResult)
		go getExecuteResult(stream, execResult)

		err = stream.Send(&runnerv2alpha1.ExecuteRequest{
			Config: &runnerv2alpha1.ProgramConfig{
				ProgramName: "bash",
				Source: &runnerv2alpha1.ProgramConfig_Commands{
					Commands: &runnerv2alpha1.ProgramConfig_CommandList{
						Items: []string{
							"sleep 3", // wait for the winsize to be set
							"tput lines",
							"tput cols",
						},
					},
				},
				Interactive: true,
				Env:         []string{"TERM=linux"},
			},
			Winsize: &runnerv2alpha1.Winsize{
				Cols: 200,
				Rows: 64,
			},
		})
		require.NoError(t, err)

		result := <-execResult

		assert.NoError(t, result.Err)
		assert.Equal(t, "64\r\n200\r\n", string(result.Stdout))
		assert.EqualValues(t, 0, result.ExitCode)
	})
}

func testStartRunnerServiceServer(t *testing.T) (
	interface{ Dial() (net.Conn, error) },
	func(),
) {
	t.Helper()

	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	lis := bufconn.Listen(1024 << 10)
	server := grpc.NewServer()
	runnerService, err := NewRunnerService(logger)
	require.NoError(t, err)

	runnerv2alpha1.RegisterRunnerServiceServer(server, runnerService)
	go server.Serve(lis)

	return lis, server.Stop
}

func testCreateRunnerServiceClient(
	t *testing.T,
	lis interface{ Dial() (net.Conn, error) },
) (*grpc.ClientConn, runnerv2alpha1.RunnerServiceClient) {
	t.Helper()

	conn, err := grpc.Dial(
		"",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)

	return conn, runnerv2alpha1.NewRunnerServiceClient(conn)
}

type executeResult struct {
	Stdout   []byte
	Stderr   []byte
	MimeType string
	ExitCode int
	Err      error
}

func getExecuteResult(
	stream runnerv2alpha1.RunnerService_ExecuteClient,
	resultc chan<- executeResult,
) {
	result := executeResult{ExitCode: -1}

	for {
		r, rerr := stream.Recv()
		if rerr != nil {
			if rerr == io.EOF {
				rerr = nil
			}
			result.Err = rerr
			break
		}
		result.Stdout = append(result.Stdout, r.StdoutData...)
		result.Stderr = append(result.Stderr, r.StderrData...)
		if r.MimeType != "" {
			result.MimeType = r.MimeType
		}
		if r.ExitCode != nil {
			result.ExitCode = int(r.ExitCode.Value)
		}
	}

	resultc <- result
}
