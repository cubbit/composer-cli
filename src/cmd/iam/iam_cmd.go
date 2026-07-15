package cmd_iam

import "github.com/spf13/cobra"

func NewIAMCmd(
	userService iamUserService,
) *cobra.Command {
	iamCmd := &cobra.Command{
		Use:   "iam",
		Short: "Manage IAM resources",
	}

	iamCmd.AddCommand(NewIAMUserCmd(userService))

	return iamCmd
}
