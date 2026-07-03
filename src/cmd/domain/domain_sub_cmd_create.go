package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewDomainSubCmdCreate(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new domain",
		Long: `Create a new domain for the organization.

A domain represents a fully qualified domain name (FQDN) that the organization wants to verify ownership of.

Examples:
  # Create a domain
  cubbit domain create --domain-name "example.com"`,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("domain-name")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := domainService.Create(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	domainCreateCmd.Flags().String("domain-name", "", "Fully qualified domain name (FQDN) to create (required)")

	return domainCreateCmd
}
