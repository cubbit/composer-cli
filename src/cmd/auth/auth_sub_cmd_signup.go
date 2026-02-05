package cmd_auth

import (
	"github.com/cubbit/composer-cli/src/service"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewAuthSubCmdSignUp(
	authService service.AuthServiceInterface,
) *cobra.Command {
	var authSignUpCmd = &cobra.Command{
		Use:   "signup",
		Short: "Sign up a new user and organization",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("username")
			cmd.MarkFlagRequired("organization")

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := authService.SignUp(cmd, args); err != nil {
				utils.PrintError(err)
			}
		},
	}

	var basePolicy, settings utils.JSONMap
	authSignUpCmd.Flags().String("email", "", "Email address of the new user (required)")
	authSignUpCmd.Flags().String("username", "", "Username of the new user (required)")
	authSignUpCmd.Flags().String("first-name", "", "First name of the new user")
	authSignUpCmd.Flags().String("last-name", "", "Last name of the new user")
	authSignUpCmd.Flags().String("password", "", "Password of the new user")
	authSignUpCmd.Flags().String("organization", "", "Name of the new organization (required)")
	authSignUpCmd.Flags().Var(&basePolicy, "base-policy", "Base policy for the new organization (as JSON string object)")
	authSignUpCmd.Flags().Var(&settings, "settings", "Settings for the new organization (as JSON string object)")

	return authSignUpCmd
}
