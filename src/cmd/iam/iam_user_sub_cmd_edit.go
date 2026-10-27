package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdEdit(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userEditCmd := &cobra.Command{
		Use:     "edit [USER_ID]",
		Aliases: []string{"update"},
		Short:   "Edit an IAM user",
		Long: `Edit an IAM user in the current profile organization.

Examples:
  cubbit iam user edit <uuid> --first-name Alice --last-name Doe
  cubbit iam user edit --user-id <uuid> --email alice@example.com
  cubbit iam user edit --username jdoe --first-name John`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.EditUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userEditCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user to edit")
	userEditCmd.Flags().String("username", "", "Username of the IAM user to edit")
	userEditCmd.Flags().String("first-name", "", "New first name")
	userEditCmd.Flags().String("last-name", "", "New last name")
	userEditCmd.Flags().String("email", "", "New email to add to the IAM user")

	return userEditCmd
}
