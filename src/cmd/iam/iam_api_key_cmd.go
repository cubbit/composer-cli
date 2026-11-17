package cmd_iam

import "github.com/spf13/cobra"

type iamAPIKeyService interface {
	CreateAPIKey(cmd *cobra.Command, args []string) error
	ListAPIKeys(cmd *cobra.Command, args []string) error
	DescribeAPIKey(cmd *cobra.Command, args []string) error
	EditAPIKey(cmd *cobra.Command, args []string) error
	RevokeAPIKey(cmd *cobra.Command, args []string) error
}

func NewIAMAPIKeyCmd(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyCmd := &cobra.Command{
		Use:   "api-key",
		Short: "Manage IAM API keys",
	}

	apiKeyCmd.AddCommand(NewIAMAPIKeySubCmdCreate(apiKeyService))
	apiKeyCmd.AddCommand(NewIAMAPIKeySubCmdList(apiKeyService))
	apiKeyCmd.AddCommand(NewIAMAPIKeySubCmdDescribe(apiKeyService))
	apiKeyCmd.AddCommand(NewIAMAPIKeySubCmdEdit(apiKeyService))
	apiKeyCmd.AddCommand(NewIAMAPIKeySubCmdRevoke(apiKeyService))

	return apiKeyCmd
}
