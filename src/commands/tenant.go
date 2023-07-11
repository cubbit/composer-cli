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
						Name:    "email",
						Aliases: []string{"e"},
						Usage:   "Email address",
					},
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p"},
						Usage:   "Password",
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
<<<<<<< HEAD
<<<<<<< HEAD
				Name:    "describe",
				Aliases: []string{"info"},
				Usage:   "describes a tenant",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Usage: "shows information about the tenant with the specified id",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "shows information about the tenant with the specified name",
					},
					&cli.StringFlag{
						Name:        "format",
						Usage:       "formats the description",
						DefaultText: "default",
						Value:       "default",
					},
				},
				Action: actions.DescribeTenant,
=======
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "removes tenants",
=======
				Name:    "describe",
				Aliases: []string{"desc"},
				Usage:   "describes a tenant",
>>>>>>> 09ed954 (feat(tenant): gives information abpout tenant)
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Usage: "shows information about the tenant with the specified id",
					},
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "shows information about the tenant with the specified name",
					},
					&cli.StringFlag{
						Name:        "format",
						Usage:       "formats the description",
						DefaultText: "default",
						Value:       "default",
					},
				},
<<<<<<< HEAD
				Action: actions.RemoveTenant,
>>>>>>> a796b82 (feat(tenant): remove tenants command)
=======
				Action: actions.DescribeTenant,
>>>>>>> 09ed954 (feat(tenant): gives information abpout tenant)
			},
		},
	}

	return &command
}
