package dbgen

import (
	"github.com/urfave/cli/v2"
)

// NewCommand creates a dbgen command.
func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "dbgen",
		Usage: "Tool for generating CURD code based on database ddl",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "set the output location",
			},
			&cli.StringFlag{
				Name:    "module",
				Aliases: []string{"m"},
				Usage:   "set the project name",
			},
			&cli.StringFlag{
				Name:    "ddl",
				Aliases: []string{"s"},
				Usage:   "specify the source ddl file",
			},
			&cli.BoolFlag{
				Name:    "cache",
				Aliases: []string{"c"},
				Usage:   "enable cache",
			},
		},
		Action: func(c *cli.Context) (err error) {
			defer func() {
				if err != nil {
					err = cli.Exit(err.Error(), 1)
				}
			}()

			output := c.String("output")
			ddlFilePath := c.String("ddl")
			moduleName := c.String("module")
			withCache := c.Bool("cache")

			if err := GenerateFromDDL(ddlFilePath, moduleName, output, withCache); err != nil {
				return err
			}

			return nil
		},
	}
}
