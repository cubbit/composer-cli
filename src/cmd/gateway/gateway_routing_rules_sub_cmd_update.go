package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

func NewGatewayRoutingRulesSubCmdUpdate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	routingRuleUpdateCmd := &cobra.Command{
		Use:   "update --gateway-id GATEWAY_ID --id RULE_ID --file-path PATH",
		Short: "Update a gateway routing rule using the V5 API",
		Long: `Update a routing rule of a gateway using the Gateway V5 API.

The rule document is a JSON object read from --file-path; use "-" to read it
from stdin. It is sent with PUT to /routing-rules/{rule_id}: the server
replaces the whole policy, it does not merge it with the existing one.

The document may contain an id field. When it does, it must match --id.

The raw document is forwarded verbatim, so a document obtained with
"get --output json" can be edited and passed back without losing fields this
CLI does not model. On success only a confirmation message is printed; use
"get --id" to inspect the updated document.

On a 422 the server returns the validator message, which is printed as-is.

The gateway is identified by the required --gateway-id flag.

Examples:
  cubbit gateway routing-rules get --gateway-id <gateway-id> --id <rule-id> --output json > rule.json
  cubbit gateway routing-rules update --gateway-id <gateway-id> --id <rule-id> --file-path rule.json
  cat rule.json | cubbit gateway routing-rules update --gateway-id <gateway-id> --id <rule-id> --file-path -`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.UpdateRoutingRule(cmd, args)
		},
	}

	routingRuleUpdateCmd.Flags().String("gateway-id", "", "ID of the gateway (required)")
	_ = routingRuleUpdateCmd.MarkFlagRequired("gateway-id")
	routingRuleUpdateCmd.Flags().String("id", "", "ID of the routing rule (required)")
	_ = routingRuleUpdateCmd.MarkFlagRequired("id")
	routingRuleUpdateCmd.Flags().String("file-path", "", "Path to the routing rule JSON document, or - for stdin (required)")
	_ = routingRuleUpdateCmd.MarkFlagRequired("file-path")

	return routingRuleUpdateCmd
}
