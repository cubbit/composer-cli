package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserPasswordSubCmdReset(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userPasswordResetCmd := &cobra.Command{
		Use:   "reset [USER_ID]",
		Short: "Reset an IAM user password",
		Long: `Reset an IAM user password in the current profile organization.
Specify exactly one target with USER_ID, --user-id or --username.

Examples:
  cubbit iam user password reset <uuid> --new-password new-secret
  cubbit iam user password reset --user-id <uuid> --new-password new-secret
  cubbit iam user password reset --username jdoe --password current-secret --new-password new-secret
  cubbit iam user password reset --username jdoe --new-password new-secret --tfa-code 123456`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.ResetUserPassword(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userPasswordResetCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user whose password should be reset")
	userPasswordResetCmd.Flags().String("username", "", "Username of the IAM user whose password should be reset")
	userPasswordResetCmd.Flags().String("password", "", "Current password of the authenticated IAM user")
	userPasswordResetCmd.Flags().String("new-password", "", "New password for the IAM user (required)")
	userPasswordResetCmd.Flags().String("tfa-code", "", "Two-factor authentication code for the admin operator")
	_ = userPasswordResetCmd.MarkFlagRequired("new-password")

	return userPasswordResetCmd
}
