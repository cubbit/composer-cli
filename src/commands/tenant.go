package commands

import (
	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

func Tenant() *cli.Command {
	command := cli.Command{
		Name:  "tenant",
		Usage: "execute commands in tenant sections",
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
		Action: func(cCtx *cli.Context) error {
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create a new tenant",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "The name of the tenant",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "description",
						Aliases: []string{"desc"},
						Usage:   "A meaningful desciription of the tenant",
					},
					&cli.StringFlag{
						Name:  "image-url",
						Usage: "The url of an image",
					},
					&cli.StringFlag{
						Name:    "settings",
						Aliases: []string{"s"},
						Usage:   "A Json object containing the settings",
					},
				},
				Action: actions.CreateTenant,
			},
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "lists all the available tenants for the current logged operator",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "lists all available information for tenants",
					},
					&cli.BoolFlag{
						Name:    "line",
						Aliases: []string{"l"},
						Usage:   "adds a line between the information about different tentants",
					},
				},
				Action: actions.ListTenant,
			},
			{
				Name:  "list-available-swarms",
				Usage: "lists the swarms that can be connected",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Usage: "lists all available information for swarms",
					},
					&cli.BoolFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "lists all available information for swarms",
					},
					&cli.BoolFlag{
						Name:    "line",
						Aliases: []string{"l"},
						Usage:   "adds a line between the information about different swarms",
					},
				},
				Action: actions.ListAvailableSwarmsTenant,
			},
		},
	}

	return &command
}
