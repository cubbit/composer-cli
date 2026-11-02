package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayConnectionSubCmdList(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	connectionListCmd := &cobra.Command{
		Use:     "list --gateway-id GATEWAY_ID",
		Aliases: []string{"ls"},
		Short:   "List the connections of a gateway using the V5 API",
		Long: `List the swarm and redundancy class pairs a gateway can offload to using
the Gateway V5 API.

Each row shows the swarm, the redundancy class, whether the connection is the
gateway default, whether it is active and how many routing rules reference it.
The ID column is the connection ID accepted by the delete command.

Examples:
  cubbit gateway connection list --gateway-id <gateway-id>
  cubbit gateway connection list --gateway-id <gateway-id> --output json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.ListConnections(cmd, args)
		},
	}

	connectionListCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = connectionListCmd.MarkFlagRequired("gateway-id")

	return connectionListCmd
}
