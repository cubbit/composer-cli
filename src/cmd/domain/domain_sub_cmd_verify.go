package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewDomainSubCmdVerify(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainID string

	var domainVerifyCmd = &cobra.Command{
		Use:   "verify",
		Short: "Verify a domain ownership via DNS TXT record check",
		Long: `Verify a domain ownership by checking the DNS TXT record.

After creating a domain, you must add the DNS TXT record to prove ownership.
This command checks if the DNS challenge has been configured correctly.

Examples:
  # Verify a domain
  cubbit domain verify --domain-id "domain-uuid"`,
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("domain-id")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := domainService.Verify(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	domainVerifyCmd.Flags().StringVar(&domainID, "domain-id", "", "Domain ID (required)")

	return domainVerifyCmd
}
