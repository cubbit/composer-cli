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
		Long: `Revoke an IAM API key for the current operator.

Examples:
  cubbit iam api-key revoke --id <uuid>`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := apiKeyService.RevokeAPIKey(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	apiKeyRevokeCmd.Flags().String("id", "", "ID of the IAM API key to revoke (required)")
	_ = apiKeyRevokeCmd.MarkFlagRequired("id")

	return apiKeyRevokeCmd
}
