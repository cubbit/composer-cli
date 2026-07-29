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
		Long: `Describe an IAM API key for the current operator.

Examples:
  cubbit iam api-key describe --id <uuid>`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.DescribeAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyDescribeCmd.Flags().String("id", "", "ID of the IAM API key to describe (required)")
	_ = apiKeyDescribeCmd.MarkFlagRequired("id")

	return apiKeyDescribeCmd
}
