package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdCreateVirtual(
	locationServiceFn service.LocationServiceInterface,
) *cobra.Command {
	var locationCreateVirtualCmd = &cobra.Command{
		Use:     "create-virtual",
		Aliases: []string{"cv"},
		Short:   "Create a virtual location",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("name")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationServiceFn.CreateVirtual(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	locationCreateVirtualCmd.Flags().String("name", "", "Name of the virtual location")
	locationCreateVirtualCmd.Flags().String("description", "", "Description of the virtual location")

	return locationCreateVirtualCmd
}
