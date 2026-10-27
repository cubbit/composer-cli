package cmd_iam

import "github.com/spf13/cobra"

func NewIAMUserPasswordCmd(
	userService iamUserService,
) *cobra.Command {
	userPasswordCmd := &cobra.Command{
		Use:   "password",
		Short: "Manage IAM user passwords",
	}

	userPasswordCmd.AddCommand(NewIAMUserPasswordSubCmdReset(userService))

	return userPasswordCmd
}
