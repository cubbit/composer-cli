package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayStorageClassSubCmdDelete(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	storageClassDeleteCmd := &cobra.Command{
		Use:     "delete --gateway-id GATEWAY_ID --id STORAGE_CLASS_ID",
		Aliases: []string{"rm"},
		Short:   "Delete a custom storage class of a gateway using the V5 API",
		Long: `Delete a custom storage class of a gateway using the Gateway V5 API.

Default storage classes cannot be deleted. Use the ID column of the list
command to find the storage class ID.

Examples:
  cubbit gateway storage-class delete --gateway-id <gateway-id> --id <storage-class-id>`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.DeleteStorageClass(cmd, args)
		},
	}

	storageClassDeleteCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = storageClassDeleteCmd.MarkFlagRequired("gateway-id")
	storageClassDeleteCmd.Flags().String("id", "", "ID of the storage class to delete (required)")
	_ = storageClassDeleteCmd.MarkFlagRequired("id")

	return storageClassDeleteCmd
}
