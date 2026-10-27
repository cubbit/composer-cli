package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdDelete(
	locationServiceFn service.LocationServiceInterface,
) *cobra.Command {
	var locationDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete a location",
		Long: `Delete an infrastructure location.

The deletion is asynchronous: the command starts it and waits for the process to complete.
A location can only be deleted when no gateway or swarm is deployed on it.

Examples:
  cubbit infrastructure location delete --cluster-id <id>`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationServiceFn.Delete(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	locationDeleteCmd.Flags().String("cluster-id", "", "ID of the location to delete")
	_ = locationDeleteCmd.MarkFlagRequired("cluster-id")

	return locationDeleteCmd
}
