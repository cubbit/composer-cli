package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayConnectionSubCmdDelete(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	connectionDeleteCmd := &cobra.Command{
		Use:     "delete --gateway-id GATEWAY_ID --id CONNECTION_ID",
		Aliases: []string{"rm"},
		Short:   "Delete a connection of a gateway using the V5 API",
		Long: `Delete a connection of a gateway using the Gateway V5 API.

The default connection cannot be deleted because it is the smart data placement
fallback. Connections still referenced by routing rules can be deleted; the list
command shows the number of impacted rules (use --output json|yaml for their
names). Use the ID column of the list command to find the connection ID.

Examples:
  cubbit gateway connection delete --gateway-id <gateway-id> --id <connection-id>`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.DeleteConnection(cmd, args)
		},
	}

	connectionDeleteCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = connectionDeleteCmd.MarkFlagRequired("gateway-id")
	connectionDeleteCmd.Flags().String("id", "", "ID of the connection to delete (required)")
	_ = connectionDeleteCmd.MarkFlagRequired("id")

	return connectionDeleteCmd
}
