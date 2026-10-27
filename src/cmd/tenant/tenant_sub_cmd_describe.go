package cmd_tenant

import (
	"fmt"

	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewTenantSubCmdDescribe(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantDescribeCmd = &cobra.Command{
		Use:   "describe <TENANT_ID>",
		Short: "Describe a tenant using the v5 API",
		Long: `Describe a tenant using the Tenant V5 API.

Shows detailed information about a tenant including metadata, resource usage, and settings.

Examples:
  cubbit tenant describe <tenant-id>`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg(s), received %d", len(args))
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := tenantService.Describe(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	tenantDescribeCmd.Flags().Bool("show-secrets", false, "Shows secret values")

	return tenantDescribeCmd
}
