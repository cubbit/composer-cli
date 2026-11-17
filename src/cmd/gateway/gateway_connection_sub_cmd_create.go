package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayConnectionSubCmdCreate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	connectionCreateCmd := &cobra.Command{
		Use:   "create --gateway-id GATEWAY_ID --connection CONNECTION",
		Short: "Create connections on a gateway using the V5 API",
		Long: `Create one or more connections on a gateway using the Gateway V5 API.

A connection is a swarm and redundancy class pair. Each --connection flag groups
one swarm with one or more redundancy classes in the format:

  <swarm-id>:<redundancy-class-id>[,<redundancy-class-id>...]

The flag can be repeated to create multiple pairs in a single call. Swarm and
redundancy class IDs must be valid UUIDs. Repeating the same swarm merges its
redundancy classes into one request entry; repeating the same pair is rejected.
At least one --connection flag is required and validation happens before any
API call.

Examples:
  cubbit gateway connection create --gateway-id <gateway-id> \
    --connection <swarm-id>:<redundancy-class-id>

  cubbit gateway connection create --gateway-id <gateway-id> \
    --connection <swarm-id>:<redundancy-class-id-1>,<redundancy-class-id-2> \
    --connection <swarm-id-2>:<redundancy-class-id-3>`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.CreateConnections(cmd, args)
		},
	}

	connectionCreateCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = connectionCreateCmd.MarkFlagRequired("gateway-id")
	connectionCreateCmd.Flags().StringArray(
		"connection",
		[]string{},
		"Connection in format swarm-id:rc-id[,rc-id...]. Can be specified multiple times. (required)",
	)
	_ = connectionCreateCmd.MarkFlagRequired("connection")

	return connectionCreateCmd
}
