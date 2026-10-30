package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserTFASubCmdDisable(
	userService iamUserService,
) *cobra.Command {
	userTFADisableCmd := &cobra.Command{
		Use:     "disable",
		Aliases: []string{"off"},
		Short:   "Disable TFA for the current IAM user",
		Long: `Disable TFA for the IAM user in the active profile.

Example:
  ds3-composer iam user tfa disable --password current-secret --tfa-code 123456`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.DisableUserTFA(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userTFADisableCmd.Flags().String("password", "", "Current password of the authenticated IAM user")
	userTFADisableCmd.Flags().String("tfa-code", "", "Two-factor authentication or recovery code used to sign in")

	return userTFADisableCmd
}
