package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesSubCmdList(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRuleListCmd := &cobra.Command{
		Use:     "list --gateway-id GATEWAY_ID",
		Aliases: []string{"ls"},
		Short:   "List the routing rules of a gateway using the V5 API",
		Long: `List the routing rules of a gateway using the Gateway V5 API.

A gateway can have zero or more routing rules. The list prints metadata only
(ID, name, enabled, input storage class and description); use get to show the
full document of a rule, including conditions and action.

Examples:
  cubbit gateway routing-rules list --gateway-id <gateway-id>
  cubbit gateway routing-rules list --gateway-id <gateway-id> --output json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.ListRoutingRules(cmd, args)
		},
	}

	routingRuleListCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = routingRuleListCmd.MarkFlagRequired("gateway-id")

	return routingRuleListCmd
}
