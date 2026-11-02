package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayStorageClassSubCmdCreate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	storageClassCreateCmd := &cobra.Command{
		Use:   "create --gateway-id GATEWAY_ID --name NAME",
		Short: "Create a custom storage class on a gateway using the V5 API",
		Long: `Create a custom storage class on a gateway using the Gateway V5 API.

The name must be unique on the gateway, contain only ASCII letters, digits,
spaces, underscores and hyphens, start with a letter or digit and be at most
100 characters long. Default storage class names are reserved.

Examples:
  cubbit gateway storage-class create --gateway-id <gateway-id> --name "cold-data"`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.CreateStorageClass(cmd, args)
		},
	}

	storageClassCreateCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = storageClassCreateCmd.MarkFlagRequired("gateway-id")
	storageClassCreateCmd.Flags().String("name", "", "Name of the storage class (required)")
	_ = storageClassCreateCmd.MarkFlagRequired("name")

	return storageClassCreateCmd
}
