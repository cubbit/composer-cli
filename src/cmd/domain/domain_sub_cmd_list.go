package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewDomainSubCmdList(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all domains",
		Run: func(cmd *cobra.Command, args []string) {
			if err := domainService.List(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	return domainListCmd
}
