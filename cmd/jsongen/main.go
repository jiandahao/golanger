package main

import (
	"fmt"
	"os"

	"github.com/jiandahao/golanger/pkg/generator/jsongen"
	cliutils "github.com/jiandahao/golanger/pkg/utils/cli"
)

func main() {
	cmd := jsongen.NewCommand()
	app := cliutils.NewAppFromCommand(cmd)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
