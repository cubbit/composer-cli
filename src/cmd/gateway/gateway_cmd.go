package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayCmd(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	var gatewayCmd = &cobra.Command{
		Use:   "gateway",
		Short: "Execute commands in gateway sections",
	}

	gatewayDescribeCmd := NewGatewaySubCmdDescribe(gatewayService)
	gatewayCreateCmd := NewGatewaySubCmdCreate(gatewayService)
	gatewayListCmd := NewGatewaySubCmdList(gatewayService)
	gatewayUpdateCmd := NewGatewaySubCmdUpdate(gatewayService)
	gatewayStorageClassCmd := NewGatewayStorageClassCmd(gatewayService)
	gatewayConnectionCmd := NewGatewayConnectionCmd(gatewayService)
	gatewayRoutingRulesCmd := NewGatewayRoutingRulesCmd(gatewayService)
	gatewayCmd.AddCommand(
		gatewayCreateCmd,
		gatewayDescribeCmd,
		gatewayListCmd,
		gatewayUpdateCmd,
		gatewayStorageClassCmd,
		gatewayConnectionCmd,
		gatewayRoutingRulesCmd,
	)

	return gatewayCmd
}
