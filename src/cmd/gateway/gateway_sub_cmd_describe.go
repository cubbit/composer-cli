package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewGatewaySubCmdDescribe(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	var gatewayID string
	var gatewayName string

	var gatewayDescribeCmd = &cobra.Command{
		Use:     "describe [GATEWAY_ID]",
		Aliases: []string{"info", "show"},
		Short:   "Describe a gateway using the v5 API",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := gatewayService.Describe(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	gatewayDescribeCmd.Flags().StringVar(&gatewayID, "gateway-id", "", "Gateway ID")
	gatewayDescribeCmd.Flags().StringVar(&gatewayName, "gateway-name", "", "Gateway name")

	return gatewayDescribeCmd
}
