package cmd_tenant

import (
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewTenantSubCmdList(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all tenants using the v5 API",
		Long: `List all tenants using the Tenant V5 API.

Supports optional filtering with the syntax: field:operator(value).
Use --query multiple times for multiple filters, or a single comma-separated string.

Examples:
  cubbit tenant list
  cubbit tenant list --query "name:eq(my-tenant)"
  cubbit tenant list --page 1 --items 50`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := tenantService.List(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	tenantListCmd.Flags().StringArray("query", []string{}, "Filter tenants (format: field:operator(value)). Can be specified multiple times.")
	tenantListCmd.Flags().Int("page", 0, "Page number for pagination (if set, returns a single page instead of all tenants)")
	tenantListCmd.Flags().Int("items", 100, "Number of items per page (default 100, max 1000)")

	return tenantListCmd
}
