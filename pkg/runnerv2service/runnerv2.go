package runnerv2service

import (
	"github.com/stateful/runme/v3/internal/command"
	runnerInternal "github.com/stateful/runme/v3/internal/runnerv2service"
	runnerv2 "github.com/stateful/runme/v3/pkg/api/gen/proto/go/runme/runner/v2"
	"go.uber.org/zap"
)

// NewRunnerService creates a new runner service.
func NewRunnerService(factory command.Factory, logger *zap.Logger) (runnerv2.RunnerServiceServer, error) {
	// N.B. The purpose of this function is to make it accessible outside of the runme package.
	return runnerInternal.NewRunnerService(factory, logger)
}
