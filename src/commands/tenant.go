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
				Name:  "show",
				Usage: "lists all the available tenants for the current logged operator",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "list",
						Aliases: []string{"ls"},
						Usage:   "lists available tenants for the operator",
					},
				},
				Action: actions.ListTenant,
			},
		},
	}

	return &command
}
