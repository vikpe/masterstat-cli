package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikpe/qw-masterstat"
)

func init() {
	cli.AppHelpTemplate = `{{.Name}} [{{.Version}}]
{{.Description}}

  Usage:   {{.UsageText}}
Example:   {{.Name}} master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
`
}

func main() {
	app := &cli.App{
		Name:        "masterstat",
		Description: "Fetch server addresses from QuakeWorld master servers.",
		UsageText:   "masterstat [<address> ...]",
		Version:     "v0.1.0",
		Action: func(c *cli.Context) error {
			masterAddresses := c.Args().Slice()
			serverAddresses := masterstat.GetServerAddressesFromMany(masterAddresses)

			for _, serverAddress := range serverAddresses {
				fmt.Println(serverAddress)
			}

			return nil
		},
	}

	args := os.Args

	if 1 == len(os.Args) {
		args = append(args, "--help")
	}

	err := app.Run(args)
	if err != nil {
		fmt.Println(err)
	}
}
