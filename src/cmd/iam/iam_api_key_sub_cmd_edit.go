package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMAPIKeySubCmdEdit(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyEditCmd := &cobra.Command{
		Use:     "edit",
		Aliases: []string{"update"},
		Short:   "Edit an IAM API key",
		Long: `Edit an API key for the current user or for a specified IAM user.

Examples:
  cubbit iam api-key edit --id <uuid> --name automation
  cubbit iam api-key edit --id <uuid> --expires-at 2026-12-31T23:59:59Z
  cubbit iam api-key edit --id <uuid> --enabled=false
  cubbit iam api-key edit --id <uuid> --user-id <user-uuid> --name automation
  cubbit iam api-key edit --id <uuid> --username jdoe --enabled=false`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.EditAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyEditCmd.Flags().String("id", "", "ID of the IAM API key to edit (required)")
	apiKeyEditCmd.Flags().String("name", "", "New name")
	apiKeyEditCmd.Flags().String("expires-at", "", "New expiration timestamp in RFC3339 format")
	apiKeyEditCmd.Flags().Bool("enabled", true, "Whether the IAM API key is enabled")
	apiKeyEditCmd.Flags().String("user-id", "", "ID of the IAM user whose API key should be edited")
	apiKeyEditCmd.Flags().String("username", "", "Username of the IAM user whose API key should be edited")
	_ = apiKeyEditCmd.MarkFlagRequired("id")

	return apiKeyEditCmd
}
