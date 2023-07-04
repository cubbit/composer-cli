package tenantlist

import (
	"github.com/cubbit/cubbit/client/cli/src/actions"
	"github.com/urfave/cli/v2"
)

func TenantList() *cli.Command {
	command := cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "Lists all the available tenants for the current logged operator",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Configuration Name",
			},
		},
		Action: actions.GenerateAccessToken,
	}

	return &command
}
