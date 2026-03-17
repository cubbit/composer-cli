package cmd_location

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewLocationSubCmdCreateVirtualNode(
	locationService service.LocationServiceInterface,
) *cobra.Command {
	var locationCreateVirtualNodeCmd = &cobra.Command{
		Use:     "create-virtual-node",
		Aliases: []string{"cvn"},
		Short:   "Create a virtual node",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("cluster-id")
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("storage-type")
			cmd.MarkFlagRequired("configuration")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := locationService.CreateVirtualNode(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	locationCreateVirtualNodeCmd.Flags().String("cluster-id", "", "Cluster ID of the virtual node")
	locationCreateVirtualNodeCmd.Flags().String("name", "", "Name of the virtual node")
	locationCreateVirtualNodeCmd.Flags().String("storage-type", "", "Storage type of the virtual node (e.g., s3)")
	locationCreateVirtualNodeCmd.Flags().String("configuration", "", "Configuration of the virtual node (e.g. {\"endpoint\": \"https://s3.example.com\", \"bucket\": \"dev-bucket-1\", \"prefix\": \"folder\", \"region\": \"eu-west-1\", \"access_key\": \"my_access_key\", \"secret_key\": \"my_secret_key\"})")

	return locationCreateVirtualNodeCmd
}
