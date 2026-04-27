package cmd

import (
	"clitemplate/app"
	"clitemplate/app/config"
	"clitemplate/registry"
	"context"
	"fmt"

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
	env *config.Env,
	clio *clio.Clio,
	specs *app.Specs,
) error {
	clio.Banner()

	clio.Info(fmt.Sprintf("The environment is '%s'.", env.Env))
	clio.Success("The application works!")

	return nil
}
