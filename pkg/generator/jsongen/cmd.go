package jsongen

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// NewCommand creates an impler command
func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "jsongen",
		Usage: "Generates golang structure according to json data string.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "s",
				Usage: "json string",
			},
		},
		UsageText: ``,
		Action: func(c *cli.Context) (err error) {
			defer func() {
				if err != nil {
					err = cli.Exit(err.Error(), 1)
				}
			}()
			jsonString := c.String("s")

			codeGenerated, err := GenerateFromString(jsonString)
			if err != nil {
				return err
			}

			fmt.Println(codeGenerated)
			return nil
		},
	}
}
