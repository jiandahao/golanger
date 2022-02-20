package main

import (
	"fmt"
	"os"

	"github.com/jiandahao/golanger/pkg/template/validator"
	clituils "github.com/jiandahao/golanger/pkg/utils/cli"
)

func main() {
	cmd := validator.NewCommand()
	app := clituils.NewAppFromCommand(cmd)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
