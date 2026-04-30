package cmd

import (
	"clitemplate/app"
	"clitemplate/app/service/env"
	"clitemplate/registry"
	"context"

	"github.com/zeroSal/went-clio/clio"
	"github.com/zeroSal/went-command/command"
)

var _ command.Interface = (*ExampleCmd)(nil)
type ExampleCmd struct {
	command.Base
}

func init() {
    registry.Command.Register(&ExampleCmd{})
}

func (c *ExampleCmd) GetHeader() command.Header {
	return command.Header{
		Use:   "example",
		Short: "An example of command",
		Long:  "A small example of a command implementation in the went wireframe.",
	}
}

func (c *ExampleCmd) Invoke() any {
	return c.run
}

func (c *ExampleCmd) run(
	ctx context.Context,
	env *env.Env,
	clio *clio.Clio,
	specs *app.Specs,
) error {
	clio.Banner()

	clio.Info("The environment is '%s'.", env.Env)
	clio.Success("The application works!")

	return nil
}
