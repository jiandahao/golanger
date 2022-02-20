package validator

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"
)

// NewCommand creates a template validator.
func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "tmpl_validator",
		Usage: "Go template validation",
		Description: `When working with the "text/template" and "html/template" packages, it's really hard to locate and understand the template's error. tmpl_validator tries to
		simplify it, makes it more easier to figure out what and where the validation errors are happening.
		`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "choose a file that holds the template code",
			},
			&cli.StringFlag{
				Name:    "text",
				Aliases: []string{"s"},
				Usage:   "provides the template code directly as a string text",
			},
		},
		Action: func(c *cli.Context) (err error) {
			defer func() {
				if err != nil {
					err = cli.Exit(err.Error(), 1)
				}
			}()
			tmplFile := c.String("file")
			tmplText := c.String("text")
			if tmplFile != "" {
				fd, err := os.Open(tmplFile)
				if err != nil {
					return err
				}

				data, err := ioutil.ReadAll(fd)
				if err != nil {
					return err
				}

				tmplText = string(data)
			} else if tmplText == "" {
				return errors.New("no template code found, plz choose a file or pass your template code directly")
			}

			_, errs := Validate(tmplText)
			if errs != nil {
				PrintErrorDetails(tmplText, errs)
			}
			return nil
		},
	}
}
