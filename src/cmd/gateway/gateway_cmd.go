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

	gatewayCreateCmd := NewGatewaySubCmdCreate(gatewayService)
	gatewayListCmd := NewGatewaySubCmdList(gatewayService)
	gatewayCmd.AddCommand(gatewayCreateCmd, gatewayListCmd)

	return gatewayCmd
}
