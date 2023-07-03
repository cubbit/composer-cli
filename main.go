package main

import (
	"log"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

const DEFAULT_FILE_PATH = "./"

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "operator",
				Usage: "execute commands in operator sections",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "interactive",
						Aliases: []string{"i"},
						Usage:   "The deletion should be interactive",
					},
				},
				Action: func(cCtx *cli.Context) error {
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:  "signup",
						Usage: "create a new operator",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "api-server-url",
								Usage:       "Api server url",
								DefaultText: "https://api.cubbit.eu/iam",
							},
							&cli.StringFlag{
								Name:  "first-name",
								Usage: "First name",
							},
							&cli.StringFlag{
								Name:  "last-name",
								Usage: "Last name",
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
						},
						Action: actions.CreateOperator,
					},
					{
						Name:  "signin",
						Usage: "sign in the operator",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "api-server-url",
								Usage:       "Api server url",
								DefaultText: "https://api.cubbit.eu/iam",
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
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Configuration Name",
							},
							&cli.StringFlag{
								Name:        "config",
								Usage:       "Configuration path for file ./",
								DefaultText: "./",
							},
						},
						Action: actions.SignInOperator,
					},
					{
						Name:  "signout",
						Usage: "sign out the operator",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Configuration Name",
							},
							&cli.StringFlag{
								Name:        "config",
								Usage:       "Configuration path for file ./",
								DefaultText: "./",
							},
						},
						Action: actions.SignOutOperator,
					},
					{
						Name:  "access",
						Usage: "generate an access token",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Configuration Name",
							},
							&cli.StringFlag{
								Name:        "config",
								Usage:       "Configuration path for file ./",
								DefaultText: "./",
							},
						},
						Action: actions.GenerateAccessToken,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
