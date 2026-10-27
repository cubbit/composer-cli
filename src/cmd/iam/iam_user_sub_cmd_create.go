package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdCreate(
	userService iamUserService,
) *cobra.Command {
	userCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create one IAM user",
		Long: `Create one IAM user in the current profile organization.

Examples:
  cubbit iam user create --username user1 --password secret
  cubbit iam user create --username user1 --password secret --email user1@example.com --policy <policy-uuid>`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.CreateUser(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userCreateCmd.Flags().String("username", "", "Username of the IAM user to create (required)")
	userCreateCmd.Flags().String("password", "", "Password of the IAM user to create (required)")
	userCreateCmd.Flags().String("email", "", "Email of the IAM user to create")
	userCreateCmd.Flags().StringArray("policy", []string{}, "Policy UUID to attach. Can be specified multiple times.")
	_ = userCreateCmd.MarkFlagRequired("username")
	_ = userCreateCmd.MarkFlagRequired("password")

	return userCreateCmd
}
