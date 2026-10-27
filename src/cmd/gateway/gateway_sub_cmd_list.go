package cmd_gateway

import (
	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewGatewaySubCmdList(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	var gatewayListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all gateways using the v5 API",
		Long: `List all gateways using the Gateway V5 API.

Supports optional filtering with the syntax: field:operator(value).
Use --query multiple times for multiple filters, or a single comma-separated string.

Supported filters:
  - id: eq
  - name: eq, in
  - slug: eq, in
  - location: eq, in
  - type: eq, in (manual, singlecluster, multicluster_controller, multicluster_worker)
  - number_of_swarms: eq, lt, lte, gt, gte
  - number_of_tenants: eq, lt, lte, gt, gte

Examples:
  cubbit gateway list
  cubbit gateway list --query "name:eq(my-gateway)"
  cubbit gateway list --query "type:in(manual,singlecluster)"
  cubbit gateway list --query "name:eq(test)" --query "type:eq(manual)"
  cubbit gateway list --page 1 --items 50`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := gatewayService.List(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	gatewayListCmd.Flags().StringArray("query", []string{}, "Filter gateways (format: field:operator(value)). Can be specified multiple times.")
	gatewayListCmd.Flags().Int("page", 0, "Page number for pagination (if set, returns a single page instead of all gateways)")
	gatewayListCmd.Flags().Int("items", 100, "Number of items per page (default 100, max 1000)")

	return gatewayListCmd
}
