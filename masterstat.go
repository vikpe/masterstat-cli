package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vikpe/masterstat"
)

func run(args []string) int {
	cli.AppHelpTemplate = `{{.Name}} [{{.Version}}]
{{.Description}}

  Usage:   {{.UsageText}}
Example:   {{.Name}} master.quakeworld.nu:27000 qwmaster.ocrana.de:27000
`

	app := &cli.App{
		Name:        "masterstat",
		Description: "Fetch server addresses from QuakeWorld master servers.",
		UsageText:   "masterstat [<address> ...]",
		Version:     "__VERSION__", // updated during build workflow
		Action: func(c *cli.Context) error {
			masterAddresses := c.Args().Slice()
			serverAddresses, err := masterstat.GetServerAddressesFromMany(masterAddresses)

			if err != nil {
				return err
			}

			for _, serverAddress := range serverAddresses {
				fmt.Println(serverAddress)
			}

			return nil
		},
	}

	if 1 == len(args) {
		args = append(args, "--help")
	}

	err := app.Run(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args))
}
