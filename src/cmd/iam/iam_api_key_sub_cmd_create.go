package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMAPIKeySubCmdCreate(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an IAM API key",
		Long: `Create an API key for the current user or for a specified IAM user.

Examples:
  cubbit iam api-key create --name automation
  cubbit iam api-key create --name automation --expires-at 2026-12-31T23:59:59Z
  cubbit iam api-key create --name automation --user-id <uuid>
  cubbit iam api-key create --name automation --username jdoe`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.CreateAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyCreateCmd.Flags().String("name", "", "Name of the IAM API key to create (required)")
	apiKeyCreateCmd.Flags().String("expires-at", "", "Expiration timestamp in RFC3339 format")
	apiKeyCreateCmd.Flags().String("user-id", "", "ID of the IAM user whose API key should be created")
	apiKeyCreateCmd.Flags().String("username", "", "Username of the IAM user whose API key should be created")
	_ = apiKeyCreateCmd.MarkFlagRequired("name")

	return apiKeyCreateCmd
}
