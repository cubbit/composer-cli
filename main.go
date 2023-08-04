package main

import (
	"log"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/command"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "Cubbit",
		Description:            "The official Cubbit CLI (Command-Line Interface) for operators",
		Usage:                  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			command.Login(),
			command.Logout(),
			command.Access(),
			command.Operator(),
			command.Tenant(),
			command.Swarm(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
