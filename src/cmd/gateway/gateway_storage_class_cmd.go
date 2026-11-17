package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayStorageClassCmd(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	storageClassCmd := &cobra.Command{
		Use:   "storage-class",
		Short: "Manage gateway storage classes",
	}

	storageClassCmd.AddCommand(NewGatewayStorageClassSubCmdList(gatewayService))
	storageClassCmd.AddCommand(NewGatewayStorageClassSubCmdCreate(gatewayService))
	storageClassCmd.AddCommand(NewGatewayStorageClassSubCmdDelete(gatewayService))

	return storageClassCmd
}
