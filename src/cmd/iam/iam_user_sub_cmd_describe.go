package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdDescribe(
	userService iamUserService,
) *cobra.Command {
	var userID string

	userDescribeCmd := &cobra.Command{
		Use:     "describe [USER_ID]",
		Aliases: []string{"info", "show"},
		Short:   "Describe an IAM user",
		Long: `Describe an IAM user in the current profile organization.

Examples:
  cubbit iam user describe --user-id <uuid>
  cubbit iam user describe <uuid>
  cubbit iam user describe --username jdoe`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.DescribeUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userDescribeCmd.Flags().StringVar(&userID, "user-id", "", "ID of the IAM user to describe")
	userDescribeCmd.Flags().String("username", "", "Username of the IAM user to describe")

	return userDescribeCmd
}
