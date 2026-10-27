package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdList(
	userService iamUserService,
) *cobra.Command {
	userListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List IAM users",
		Long: `List IAM users in the current profile organization.

Examples:
  cubbit iam user list
  cubbit iam user list --search alice
  cubbit iam user list --enabled true --sort-key username --sort-order asc
  cubbit iam user list --page 2 --items 50`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.ListUsers(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userListCmd.Flags().String("enabled", "", "Filter enabled (true) or disabled (false) users")
	userListCmd.Flags().String("search", "", "Search by username, first name, last name or email")
	userListCmd.Flags().Int("page", 1, "Page number (1-indexed). If omitted, all pages are fetched")
	userListCmd.Flags().Int("items", 100, "Number of items per page")
	userListCmd.Flags().String("sort-key", "", "Sort key: username, first_name, last_name, email, created_at, status")
	userListCmd.Flags().String("sort-order", "", "Sort order: asc, desc")

	return userListCmd
}
