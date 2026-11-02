package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayConnectionCmd(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	connectionCmd := &cobra.Command{
		Use:   "connection",
		Short: "Manage gateway connections",
		Long: `Manage the connections of a gateway.

A connection is a swarm and redundancy class pair the gateway can offload to.
Use the connection commands to inspect, create and delete these pairs.`,
	}

	connectionCmd.AddCommand(NewGatewayConnectionSubCmdList(gatewayService))
	connectionCmd.AddCommand(NewGatewayConnectionSubCmdCreate(gatewayService))
	connectionCmd.AddCommand(NewGatewayConnectionSubCmdDelete(gatewayService))

	return connectionCmd
}
