package command

import (
	"errors"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/urfave/cli/v2"
)

func Swarm() *cli.Command {
	command := cli.Command{
		Name:  "swarm",
		Usage: "execute commands in swarm sections",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "The operation should be interactive",
			},
		},
		Action: func(ctx *cli.Context) error {
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
						Name:     "description",
						Aliases:  []string{"desc"},
						Usage:    "A meaningful description of the swarm",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "configuration",
						Aliases:  []string{"c"},
						Usage:    "A Json object containing the swarm configuration",
						Required: true,
					},
				},
				Action: action.CreateSwarm,
			},
			{
				Name:    "describe",
				Aliases: []string{"info"},
				Usage:   "describes a swarm",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "id",
						Usage:    "id of the swarm",
						Required: true,
					},
				},
				Action: action.GetSwarm,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "lists all swarms for the current logged operator",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "lists all available information for swarms",
					},
				},
				Action: action.ListSwarms,
			},
		},
	}

	return &command
}
