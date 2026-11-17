package cmd_domain

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewDomainCmd(
	domainService service.DomainServiceInterface,
) *cobra.Command {
	var domainCmd = &cobra.Command{
		Use:   "domain",
		Short: "Execute commands in domain sections",
	}

	domainCreateCmd := NewDomainSubCmdCreate(domainService)
	domainDescribeCmd := NewDomainSubCmdDescribe(domainService)
	domainListCmd := NewDomainSubCmdList(domainService)
	domainDeleteCmd := NewDomainSubCmdDelete(domainService)
	domainVerifyCmd := NewDomainSubCmdVerify(domainService)
	domainCmd.AddCommand(domainCreateCmd, domainDescribeCmd, domainListCmd, domainDeleteCmd, domainVerifyCmd)

	return domainCmd
}
