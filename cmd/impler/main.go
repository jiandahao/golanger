package main

import (
	"fmt"
	"os"

	"github.com/jiandahao/golanger/pkg/impler"
	cliutils "github.com/jiandahao/golanger/pkg/utils/cli"
)

func main() {
	cmd := impler.NewCommand()
	app := cliutils.NewAppFromCommand(cmd)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
