package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewDomainSubCmdDelete(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainID string

	var domainDeleteCmd = &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Short:   "Delete a domain",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("domain-id")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := domainService.Delete(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	domainDeleteCmd.Flags().StringVar(&domainID, "domain-id", "", "Domain ID (required)")

	return domainDeleteCmd
}
