package cmd_tenant

import (
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewTenantSubCmdListConnections(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantListConnectionsCmd = &cobra.Command{
		Use:     "list-connections",
		Aliases: []string{"ls-connections", "list-conn"},
		Short:   "List all connections for a tenant using the v5 API",
		Long: `List all connections for a tenant using the Tenant Connections V5 API.

A connection is a relation between a tenant and a domain with the associated gateways and their verification statuses.

By default each connection is live-verified before the list is refreshed.
Use --skip-verify to skip this process and return cached data.

Examples:
  cubbit tenant list-connections --tenant-id <tenant-id>
  cubbit tenant list-connections --tenant-name <tenant-name>
  cubbit tenant list-connections --tenant-id <tenant-id> --page 1 --items 50
  cubbit tenant list-connections --tenant-id <tenant-id> --skip-verify`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := tenantService.ListConnections(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	tenantListConnectionsCmd.Flags().String("tenant-id", "", "Tenant ID")
	tenantListConnectionsCmd.Flags().String("tenant-name", "", "Tenant name")
	tenantListConnectionsCmd.Flags().Int("page", 0, "Page number for pagination (if set, returns a single page instead of all connections)")
	tenantListConnectionsCmd.Flags().Int("items", 100, "Number of items per page (default 100, max 1000)")
	tenantListConnectionsCmd.Flags().Bool("skip-verify", false, "Skip live verification and show cached data")

	return tenantListConnectionsCmd
}
