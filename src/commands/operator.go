package commands

import (
	"errors"

	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

func Operator() *cli.Command {
	command := cli.Command{
		Name:  "operator",
		Usage: "execute commands in operator sections",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "The operation should be interactive",
			},
		},
		Action: func(cCtx *cli.Context) error {
			return errors.New("please specify a valid command")
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
					&cli.StringFlag{
						Name:    "secret",
						Aliases: []string{"s"},
						Usage:   "Secret",
					},
				},
				Action: actions.CreateOperator,
			},
		},
	}

	return &command
}
