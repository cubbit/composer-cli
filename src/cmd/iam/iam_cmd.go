package cmd_iam

import "github.com/spf13/cobra"

func NewIAMCmd(
	userService iamUserService,
	apiKeyServices ...iamAPIKeyService,
) *cobra.Command {
	iamCmd := &cobra.Command{
		Use:   "iam",
		Short: "Manage IAM resources",
	}

	iamCmd.AddCommand(NewIAMUserCmd(userService))
	if len(apiKeyServices) > 0 && apiKeyServices[0] != nil {
		iamCmd.AddCommand(NewIAMAPIKeyCmd(apiKeyServices[0]))
	}

	return iamCmd
}
