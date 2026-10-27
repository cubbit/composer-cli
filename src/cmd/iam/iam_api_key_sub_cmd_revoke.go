package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMAPIKeySubCmdRevoke(apiKeyService iamAPIKeyService) *cobra.Command {
	apiKeyRevokeCmd := &cobra.Command{
		Use:     "revoke",
		Aliases: []string{"delete", "rm"},
		Short:   "Revoke an IAM API key",
		Long: `Revoke an API key for the current user or for a specified IAM user.

Examples:
  cubbit iam api-key revoke --id <uuid>
  cubbit iam api-key revoke --id <uuid> --user-id <user-uuid>
  cubbit iam api-key revoke --id <uuid> --username jdoe`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.RevokeAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyRevokeCmd.Flags().String("id", "", "ID of the IAM API key to revoke (required)")
	apiKeyRevokeCmd.Flags().String("user-id", "", "ID of the IAM user whose API key should be revoked")
	apiKeyRevokeCmd.Flags().String("username", "", "Username of the IAM user whose API key should be revoked")
	_ = apiKeyRevokeCmd.MarkFlagRequired("id")

	return apiKeyRevokeCmd
}
