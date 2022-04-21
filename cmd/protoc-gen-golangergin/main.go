package main

import (
	"flag"

	"github.com/jiandahao/golanger/pkg/generator/gingen"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(gingen.NewProtocPlugin())
}
