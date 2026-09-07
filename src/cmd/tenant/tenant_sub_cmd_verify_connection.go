package cmd_tenant

import (
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewTenantSubCmdVerifyConnection(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantVerifyConnectionCmd = &cobra.Command{
		Use:     "verify-connection",
		Aliases: []string{"verify-conn", "vconn"},
		Short:   "Verify DNS and TLS status for a tenant connection",
		Long: `Trigger a live verification of the DNS and TLS certificate status for a tenant connection.

Displays the verification state for the connection, showing the status
(unset, dns_ok, certificate_ok) of the S3, Console, and Wildcard S3 records
for both domain-level and per-gateway records.

Aggregated state (verified, partially_verified, unverified) is shown inline
at each level for quick scanning.

The connection is identified by --connection-id (required). Use --tenant-id
or --tenant-name to specify the tenant.

Examples:
  cubbit tenant verify-connection --tenant-id <tenant-id> --connection-id <connection-id>
  cubbit tenant verify-connection --tenant-name <tenant-name> --connection-id <connection-id>`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := tenantService.VerifyConnection(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	tenantVerifyConnectionCmd.Flags().String("tenant-id", "", "Tenant ID")
	tenantVerifyConnectionCmd.Flags().String("tenant-name", "", "Tenant name")
	tenantVerifyConnectionCmd.Flags().String("connection-id", "", "Connection ID (required)")
	tenantVerifyConnectionCmd.MarkFlagRequired("connection-id")

	return tenantVerifyConnectionCmd
}
