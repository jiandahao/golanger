package main

import (
	"fmt"
	"os"

	"github.com/jiandahao/golanger/pkg/generator/dbgen"
	"github.com/jiandahao/golanger/pkg/generator/jsongen"
	"github.com/jiandahao/golanger/pkg/impler"
	"github.com/jiandahao/golanger/pkg/template/validator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "golanger"
	app.Usage = "Tools for improving development efficiency"
	app.Commands = []*cli.Command{
		impler.NewCommand(),
		validator.NewCommand(),
		jsongen.NewCommand(),
		dbgen.NewCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
