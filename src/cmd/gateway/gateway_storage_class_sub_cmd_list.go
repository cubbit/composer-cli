package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayStorageClassSubCmdList(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	storageClassListCmd := &cobra.Command{
		Use:     "list --gateway-id GATEWAY_ID",
		Aliases: []string{"ls"},
		Short:   "List the storage classes of a gateway using the V5 API",
		Long: `List the default and custom storage classes available on a gateway using the Gateway V5 API.

Default storage classes cannot be deleted. The status column reports
whether a storage class is active or fallback-only.

Examples:
  cubbit gateway storage-class list --gateway-id <gateway-id>
  cubbit gateway storage-class list --gateway-id <gateway-id> --output json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.ListStorageClasses(cmd, args)
		},
	}

	storageClassListCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = storageClassListCmd.MarkFlagRequired("gateway-id")

	return storageClassListCmd
}
