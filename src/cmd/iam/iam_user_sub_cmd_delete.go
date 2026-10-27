package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdDelete(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userDeleteCmd := &cobra.Command{
		Use:     "delete [USER_ID]",
		Aliases: []string{"rm"},
		Short:   "Delete an IAM user",
		Long: `Delete an IAM user from the current profile organization.

Examples:
  cubbit iam user delete <uuid>
  cubbit iam user delete --user-id <uuid>
  cubbit iam user delete --username jdoe`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.DeleteUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userDeleteCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user to delete")
	userDeleteCmd.Flags().String("username", "", "Username of the IAM user to delete")

	return userDeleteCmd
}
