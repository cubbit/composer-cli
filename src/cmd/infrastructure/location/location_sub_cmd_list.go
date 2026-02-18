package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdList(
	locationService service.LocationServiceInterface,
) *cobra.Command {
	var locationListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all locations",
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationService.List(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	return locationListCmd
}
