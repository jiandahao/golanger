package utils

import "github.com/urfave/cli/v2"

// NewAppFromCommand new cli app from a cli command.
func NewAppFromCommand(cmd *cli.Command) *cli.App {
	app := cli.NewApp()
	app.Name = cmd.Name
	app.Usage = cmd.Usage
	app.UsageText = cmd.UsageText
	app.After = cmd.After
	app.Before = cmd.Before
	app.Description = cmd.Description
	app.Flags = cmd.Flags
	app.Commands = cmd.Subcommands
	app.Action = cmd.Action
	return app
}
