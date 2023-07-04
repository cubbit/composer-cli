package commands

import (
	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

func Login() *cli.Command {
	command := cli.Command{
		Name:  "login",
		Usage: "log in the operator",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "The operation should be interactive",
			},
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
				Name:  "profile",
				Usage: "Configuration Profile",
			},
			&cli.StringFlag{
				Name:        "config",
				Usage:       "Configuration path for file ./",
				DefaultText: "./",
			},
		},
		Action: actions.SignInOperator,
	}

	return &command
}
