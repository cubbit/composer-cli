package cmd_tenant

import (
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/spf13/cobra"
)

func NewTenantCmd(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantCmd = &cobra.Command{
		Use:   "tenant",
		Short: "Execute commands in tenant sections",
	}

	tenantCreateCmd := NewTenantSubCmdCreate(tenantService)
	tenantCmd.AddCommand(tenantCreateCmd)

	return tenantCmd
}
