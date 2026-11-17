package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesSubCmdGet(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRuleGetCmd := &cobra.Command{
		Use:   "get --gateway-id GATEWAY_ID --id RULE_ID",
		Short: "Show a routing rule of a gateway using the V5 API",
		Long: `Show a routing rule of a gateway using the Gateway V5 API.

By default the full policy document is printed as indented JSON, the same
document --output json emits, including fields this CLI does not model. Use
--output yaml for YAML. Redirect the output to a file and pass it to
update --file-path to round-trip the document.

Use "list" to discover rule IDs. The gateway is identified by the required
--gateway-id flag and the rule by the required --id flag.

Examples:
  cubbit gateway routing-rules get --gateway-id <gateway-id> --id <rule-id>
  cubbit gateway routing-rules get --gateway-id <gateway-id> --id <rule-id> --output json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.GetRoutingRule(cmd, args)
		},
	}

	routingRuleGetCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = routingRuleGetCmd.MarkFlagRequired("gateway-id")
	routingRuleGetCmd.Flags().String("id", "", "ID of the routing rule (required)")
	_ = routingRuleGetCmd.MarkFlagRequired("id")

	return routingRuleGetCmd
}
