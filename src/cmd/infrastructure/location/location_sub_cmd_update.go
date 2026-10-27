package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdUpdate(
	locationServiceFn service.LocationServiceInterface,
) *cobra.Command {
	var locationUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a location",
		Long: `Update the metadata of an infrastructure location.

Only the provided fields are updated; the others are left unchanged.
An empty value clears the field. The location name is the user-facing name of the location.

Examples:
  cubbit infrastructure location update --cluster-id <id> --name "Milan datacenter"
  cubbit infrastructure location update --cluster-id <id> --description "Primary datacenter"
  cubbit infrastructure location update --cluster-id <id> --position "Milan, Italy"
  cubbit infrastructure location update --cluster-id <id> --contact-user jdoe
  cubbit infrastructure location update --cluster-id <id> --description ""`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationServiceFn.Update(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	locationUpdateCmd.Flags().String("cluster-id", "", "ID of the location to update")
	locationUpdateCmd.Flags().String("name", "", "New name of the location (3-256 characters, empty to clear)")
	locationUpdateCmd.Flags().String("description", "", "New description of the location (max 120 characters, empty to clear)")
	locationUpdateCmd.Flags().String("position", "", "New physical location as a plain string (max 80 characters, empty to clear)")
	locationUpdateCmd.Flags().String("contact-user", "", "IAM user (id or username) acting as contact for the location (empty to clear)")
	_ = locationUpdateCmd.MarkFlagRequired("cluster-id")

	return locationUpdateCmd
}
