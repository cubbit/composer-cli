package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdEnable(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userEnableCmd := &cobra.Command{
		Use:   "enable [USER_ID]",
		Short: "Enable an IAM user",
		Long: `Enable an IAM user in the current profile organization.

Examples:
  cubbit iam user enable <uuid>
  cubbit iam user enable --user-id <uuid>
  cubbit iam user enable --username jdoe`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.EnableUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userEnableCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user to enable")
	userEnableCmd.Flags().String("username", "", "Username of the IAM user to enable")

	return userEnableCmd
}
