package commands

import (
	"errors"

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
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "The name of the tenant",
			},
			&cli.StringFlag{
				Name:  "id",
				Usage: "The id of the tenant",
			},
		},
		Action: func(cCtx *cli.Context) error {
			return errors.New("please specify a valid command")
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
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "removes tenants",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Usage: "removes the tenant with the specified id",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "removes the tenant with the specified name",
					},
					&cli.StringFlag{
						Name:     "email",
						Aliases:  []string{"e"},
						Usage:    "Email address",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "password",
						Aliases:  []string{"p"},
						Usage:    "Password",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "code",
						Aliases: []string{"2fa"},
						Usage:   "Two factor authentication code",
					},
				},
				Action: actions.RemoveTenant,
			},
			{
				Name:    "describe",
				Aliases: []string{"info"},
				Usage:   "describes a tenant",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "format",
						Usage:       "formats the description",
						DefaultText: "default",
						Value:       "default",
					},
				},
				Action: actions.DescribeTenant,
			},
			{
				Name:   "edit-description",
				Usage:  "changes the tenant description",
				Flags:  []cli.Flag{},
				Action: actions.EditTenantDescription,
			},
			{
				Name:   "edit-image",
				Usage:  "changes the tenant image",
				Flags:  []cli.Flag{},
				Action: actions.EditTenantImage,
			},
		},
	}

	return &command
}
