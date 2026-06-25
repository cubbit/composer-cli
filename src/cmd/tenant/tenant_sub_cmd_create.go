package cmd_tenant

import (
	servicetenant "github.com/cubbit/composer-cli/src/service/tenant"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewTenantSubCmdCreate(
	tenantService servicetenant.TenantServiceInterface,
) *cobra.Command {
	var tenantCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new tenant using the Tenant V5 API",
		Long: `Create a new tenant using the Tenant V5 API.

This command creates a tenant with the specified name, slug, and connections.

Connection Format:
   Each --connection flag specifies a connection in the format: domainID:gatewayID[,gatewayID...][:subdomain]

Examples:
   # Create a tenant with required flags
   cubbit tenant create \
     --name "Tenant Cubbit 1" \
     --slug "cubbit-1" \
     --connection "domainID:gatewayID1,gatewayID2:mysub"

   # Create a tenant with multiple connections
   cubbit tenant create \
     --name "Tenant Cubbit 1" \
     --slug "cubbit-1" \
     --connection "domainID:gatewayID1,gatewayID2" \
     --connection "domainID2:gatewayID3"

   # Create a tenant with optional description
   cubbit tenant create \
     --name "Tenant Cubbit 1" \
     --slug "cubbit-1" \
     --description "Production tenant" \
     --connection "domainID:gatewayID1:mysub"

   # Create a tenant in interactive mode
   cubbit tenant create --interactive`,
		PreRun: func(cmd *cobra.Command, args []string) {
			interactive, _ := cmd.Flags().GetBool("interactive")
			if !interactive {
				cmd.MarkFlagRequired("name")
				cmd.MarkFlagRequired("slug")
				cmd.MarkFlagRequired("connection")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := tenantService.Create(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	tenantCreateCmd.Flags().Bool("interactive", false, "Run in interactive mode")
	tenantCreateCmd.Flags().String("name", "", "Name of the tenant (required)")
	tenantCreateCmd.Flags().String("slug", "", "URL-friendly slug for the tenant (required)")
	tenantCreateCmd.Flags().String("description", "", "Optional description of the tenant")
	tenantCreateCmd.Flags().StringArray("connection", []string{}, "Connection in format domainID:gatewayID[,gatewayID...][:subdomain]. Can be specified multiple times. (required)")

	return tenantCreateCmd
}
