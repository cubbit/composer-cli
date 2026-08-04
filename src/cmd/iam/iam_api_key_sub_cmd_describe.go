package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMAPIKeySubCmdDescribe(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyDescribeCmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"info", "show"},
		Short:   "Describe an IAM API key",
		Long: `Describe an API key for the current user or for a specified IAM user.

Examples:
  cubbit iam api-key describe --id <uuid>
  cubbit iam api-key describe --id <uuid> --user-id <user-uuid>
  cubbit iam api-key describe --id <uuid> --username jdoe`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.DescribeAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyDescribeCmd.Flags().String("id", "", "ID of the IAM API key to describe (required)")
	apiKeyDescribeCmd.Flags().String("user-id", "", "ID of the IAM user whose API key should be described")
	apiKeyDescribeCmd.Flags().String("username", "", "Username of the IAM user whose API key should be described")
	_ = apiKeyDescribeCmd.MarkFlagRequired("id")

	return apiKeyDescribeCmd
}
