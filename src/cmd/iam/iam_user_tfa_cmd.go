package cmd_iam

import "github.com/spf13/cobra"

func NewIAMUserTFACmd(
	userService iamUserService,
) *cobra.Command {
	userTFACmd := &cobra.Command{
		Use:     "tfa",
		Aliases: []string{"2fa"},
		Short:   "Manage IAM user TFA",
	}

	userTFACmd.AddCommand(NewIAMUserTFASubCmdEnable(userService))
	userTFACmd.AddCommand(NewIAMUserTFASubCmdDisable(userService))

	return userTFACmd
}
