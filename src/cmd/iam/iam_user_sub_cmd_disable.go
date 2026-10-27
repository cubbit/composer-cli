package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdDisable(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userDisableCmd := &cobra.Command{
		Use:   "disable [USER_ID]",
		Short: "Disable an IAM user",
		Long: `Disable an IAM user in the current profile organization.

Examples:
  cubbit iam user disable <uuid>
  cubbit iam user disable --user-id <uuid>
  cubbit iam user disable --username jdoe`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.DisableUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userDisableCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user to disable")
	userDisableCmd.Flags().String("username", "", "Username of the IAM user to disable")

	return userDisableCmd
}
