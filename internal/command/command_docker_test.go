//go:build test_with_docker

package command

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stateful/runme/v3/internal/dockerexec"
	runnerv2alpha1 "github.com/stateful/runme/v3/internal/gen/proto/go/runme/runner/v2alpha1"
)

func TestDockerCommand(t *testing.T) {
	t.Parallel()

	docker, err := dockerexec.New(&dockerexec.Options{Debug: false, Image: "alpine:3.19"})
	require.NoError(t, err)

	// This test case is treated as a warm up. Do not parallelize.
	t.Run("NoOutput", func(t *testing.T) {
		cmd := NewDocker(testConfigBasicProgram, docker, Options{})
		require.NoError(t, cmd.Start(context.Background()))
		require.NoError(t, cmd.Wait())
	})

	t.Run("Output", func(t *testing.T) {
		t.Parallel()
		stdout := bytes.NewBuffer(nil)
		cmd := NewDocker(
			testConfigBasicProgram,
			docker,
			Options{Stdout: stdout},
		)
		require.NoError(t, cmd.Start(context.Background()))
		require.NoError(t, cmd.Wait())
		assert.Equal(t, "test", stdout.String())
	})

	t.Run("Running", func(t *testing.T) {
		t.Parallel()
		cmd := NewDocker(
			&Config{
				ProgramName: "sleep",
				Arguments:   []string{"1"},
				Mode:        runnerv2alpha1.CommandMode_COMMAND_MODE_INLINE,
			},
			docker,
			Options{},
		)
		err := cmd.Start(context.Background())
		require.NoError(t, err)
		require.True(t, cmd.Running())
		require.Greater(t, cmd.Pid(), 0)
		require.NoError(t, cmd.Wait())
	})

	t.Run("NonZeroExit", func(t *testing.T) {
		t.Parallel()
		cmd := NewDocker(
			&Config{
				ProgramName: "sh",
				Arguments:   []string{"-c", "exit 11"},
			},
			docker,
			Options{},
		)
		require.NoError(t, err)
		require.NoError(t, cmd.Start(context.Background()))
		// TODO(adamb): wait should return non-nil error due to non-zero exit code.
		require.NoError(t, cmd.Wait())
		require.Equal(t, 11, cmd.cmd.ProcessState.ExitCode)
	})
}
