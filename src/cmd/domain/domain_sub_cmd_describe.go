package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewDomainSubCmdDescribe(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainID string

	var domainDescribeCmd = &cobra.Command{
		Use:     "describe",
		Aliases: []string{"info", "show"},
		Short:   "Describe a domain",
		PreRun: func(cmd *cobra.Command, args []string) {
			cmd.MarkFlagRequired("domain-id")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := domainService.Describe(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	domainDescribeCmd.Flags().StringVar(&domainID, "domain-id", "", "Domain ID (required)")

	return domainDescribeCmd
}
