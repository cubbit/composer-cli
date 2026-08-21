package cmd_iam

import "github.com/spf13/cobra"

type iamUserService interface {
	ImportUsers(cmd *cobra.Command, args []string) error
	CreateUser(cmd *cobra.Command, args []string) error
	ListUsers(cmd *cobra.Command, args []string) error
	DescribeUser(cmd *cobra.Command, args []string) error
	EditUser(cmd *cobra.Command, args []string) error
	EnableUser(cmd *cobra.Command, args []string) error
	DisableUser(cmd *cobra.Command, args []string) error
	DeleteUser(cmd *cobra.Command, args []string) error
	ResetUserPassword(cmd *cobra.Command, args []string) error
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
	userCmd.AddCommand(NewIAMUserSubCmdDescribe(userService))
	userCmd.AddCommand(NewIAMUserSubCmdEdit(userService))
	userCmd.AddCommand(NewIAMUserSubCmdEnable(userService))
	userCmd.AddCommand(NewIAMUserSubCmdDisable(userService))
	userCmd.AddCommand(NewIAMUserSubCmdDelete(userService))
	userCmd.AddCommand(NewIAMUserPasswordCmd(userService))

	return userCmd
}
