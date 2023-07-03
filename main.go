package main

import (
	"log"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/commands"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func main() {
	app := &cli.App{
		Name:        "Cubbit",
		Description: "The official Cubbit CLI (Command-Line Interface) for operators",
		Usage:       "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
		Commands: []*cli.Command{
			commands.Login(),
			commands.Logout(),
			commands.Access(),
			commands.Operator(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
