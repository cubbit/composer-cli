package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesSubCmdDelete(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRuleDeleteCmd := &cobra.Command{
		Use:     "delete --gateway-id GATEWAY_ID --id RULE_ID",
		Aliases: []string{"rm"},
		Short:   "Delete a gateway routing rule using the V5 API",
		Long: `Delete a routing rule of a gateway using the Gateway V5 API.

Use "list" to discover rule IDs.

Examples:
  cubbit gateway routing-rules delete --gateway-id <gateway-id> --id <rule-id>`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.DeleteRoutingRule(cmd, args)
		},
	}

	routingRuleDeleteCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = routingRuleDeleteCmd.MarkFlagRequired("gateway-id")
	routingRuleDeleteCmd.Flags().String("id", "", "ID of the routing rule to delete (required)")
	_ = routingRuleDeleteCmd.MarkFlagRequired("id")

	return routingRuleDeleteCmd
}
