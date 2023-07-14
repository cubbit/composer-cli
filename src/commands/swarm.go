package commands

import (
	"errors"

	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

func Swarm() *cli.Command {
	command := cli.Command{
		Name:  "swarm",
		Usage: "execute commands in swarm sections",
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
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "The name of the swarm",
			},
			&cli.StringFlag{
				Name:  "id",
				Usage: "The id of the swarm",
			},
		},
		Action: func(cCtx *cli.Context) error {
			return errors.New("please specify a valid command")
		},
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create a new swarm",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "The name of the swarm",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "description",
						Aliases: []string{"desc"},
						Usage:   "A meaningful desciription of the swarm",
					},
					&cli.StringFlag{
						Name:  "configuration",
						Usage: "A Json object containing the configuration",
					},
				},
				Action: actions.CreateSwarm,
			},
		},
	}

	return &command
}
