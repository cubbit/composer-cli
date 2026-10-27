package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserTFASubCmdEnable(
	userService iamUserService,
) *cobra.Command {
	userTFAEnableCmd := &cobra.Command{
		Use:     "enable",
		Aliases: []string{"on"},
		Short:   "Enable TFA for the current IAM user",
		Long: `Enable TFA for the IAM user in the active profile.
If --setup-key is omitted, a setup key is generated automatically.

Examples:
  cubbit iam user tfa enable --password current-secret
  cubbit iam user tfa enable --password current-secret --setup-key JBSWY3DPEHPK3PXP`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.EnableUserTFA(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userTFAEnableCmd.Flags().String("setup-key", "", "Base32 setup key to use")
	userTFAEnableCmd.Flags().String("password", "", "Current password of the authenticated IAM user")

	return userTFAEnableCmd
}
