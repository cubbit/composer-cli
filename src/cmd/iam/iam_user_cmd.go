package cmd_iam

import "github.com/spf13/cobra"

type iamUserService interface {
	ImportUsers(cmd *cobra.Command, args []string) error
	CreateUser(cmd *cobra.Command, args []string) error
	ListUsers(cmd *cobra.Command, args []string) error
}

func NewIAMUserCmd(
	userService iamUserService,
) *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage IAM users",
	}

	userCmd.AddCommand(NewIAMUserSubCmdImport(userService))
	userCmd.AddCommand(NewIAMUserSubCmdCreate(userService))
	userCmd.AddCommand(NewIAMUserSubCmdList(userService))

	return userCmd
}
