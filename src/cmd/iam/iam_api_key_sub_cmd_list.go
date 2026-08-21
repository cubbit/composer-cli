package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMAPIKeySubCmdList(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List IAM API keys",
		Long: `List API keys for the current user or for a specified IAM user.

Examples:
  cubbit iam api-key list
  cubbit iam api-key list --sort-key created_at --sort-order desc
  cubbit iam api-key list --page 2 --items 50
  cubbit iam api-key list --user-id <uuid>
  cubbit iam api-key list --username jdoe`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.ListAPIKeys(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyListCmd.Flags().Int("page", 1, "Page number (1-indexed). If omitted, all pages are fetched")
	apiKeyListCmd.Flags().Int("items", 100, "Number of items per page")
	apiKeyListCmd.Flags().String("sort-key", "", "Sort key: name, created_at, expires_at, enabled")
	apiKeyListCmd.Flags().String("sort-order", "", "Sort order: asc, desc")
	apiKeyListCmd.Flags().String("user-id", "", "ID of the IAM user whose API keys should be listed")
	apiKeyListCmd.Flags().String("username", "", "Username of the IAM user whose API keys should be listed")

	return apiKeyListCmd
}
