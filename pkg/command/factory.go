package command

import (
	"github.com/stateful/runme/v3/internal/command"
)

type FactoryOption command.FactoryOption
type Factory command.Factory

func NewFactory(opts ...FactoryOption) Factory {
	return command.NewFactory(opts...)
}
