package command

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/urfave/cli/v2"
)

func Access() *cli.Command {
	command := cli.Command{
		Name:  "access",
		Usage: "generate an access token",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "profile",
				Usage: "Configuration Profile",
			},
			&cli.StringFlag{
				Name:        "config",
				Usage:       "Configuration path for file ./",
				DefaultText: "./",
			},
		},
		Action: action.GenerateAccessToken,
	}

	return &command
}
