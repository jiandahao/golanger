package impler

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"

	"github.com/urfave/cli/v2"
)

// NewCommand creates an impler command
func NewCommand() *cli.Command {
	return &cli.Command{
		Name:  "impler",
		Usage: "Implements/Extracts an golang interface",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "package source directory, useful for vendored code",
			},
		},
		UsageText: `impler [-dir directory] <recv> <type> <name> 
		Examples:
		- Implements an interface:
			impler 'f *File' io.Reader
			impler Murmur hash.Hash
			impler -dir $GOPATH/src/github.com/josharian/impl Murmur hash.Hash

			Don't forget the single quotes around the receiver type
			to prevent shell globbing.
		
		- Extracts an interface:
			impler myinterface struct time.Ticker`,
		Action: func(c *cli.Context) error {
			args := c.Args()
			if args.Len() != 3 {
				return cli.ShowSubcommandHelp(c)
			}

			recv, typ, name := args.Get(0), args.Get(1), args.Get(2)
			if !isValidReceiver(recv) {
				fatal(fmt.Sprintf("invalid receiver: %q", recv))
			}

			srcDir := c.String("dir")

			if srcDir == "" {
				if dir, err := os.Getwd(); err == nil {
					srcDir = dir
				}
			}

			var output []byte
			switch typ {
			case "struct":
				fns, err := Methods(name, srcDir)
				if err != nil {
					fatal(err)
				}

				output = GenInterfaceStubs(recv, fns)
			case "iface":
				fns, err := Funcs(name, srcDir)
				if err != nil {
					fatal(err)
				}

				// Get list of already implemented funcs
				implemented, err := ImplementedFuncs(fns, recv, srcDir)
				if err != nil {
					fatal(err)
				}

				output = GenImplStubs(recv, fns, implemented)
			default:
				return cli.Exit(fmt.Sprintf("invalid type: %s, only struct / iface is expected.", typ), 1)
			}

			fmt.Print(string(output))
			return nil
		},
	}
}

// isValidReceiver reports whether recv is a valid receiver expression.
func isValidReceiver(recv string) bool {
	if recv == "" {
		// The parse will parse empty receivers, but we don't want to accept them,
		// since it won't generate a usable code snippet.
		return false
	}
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", "package hack\nfunc ("+recv+") Foo()", 0)
	return err == nil
}
