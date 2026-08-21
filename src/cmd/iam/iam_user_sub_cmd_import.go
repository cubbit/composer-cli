package cmd_iam

import (
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

func NewIAMUserSubCmdImport(
	userService iamUserService,
) *cobra.Command {
	userImportCmd := &cobra.Command{
		Use:   "import",
		Short: "Import IAM users from a JSON or CSV file",
		Long: `Import one or more IAM users from a JSON or CSV file.

JSON files must contain a users array with the following structure:

{
  "users": [
    {
      "username": "user1",
      "password": "<clear-text-password>",
      "first_name": "John",
      "last_name": "Doe",
      "email": "john@example.com",
      "attached_policies": ["<policy-uuid>"]
    }
  ]
}

All fields except username and password are optional.

CSV files must contain username and password columns. Optional columns are first_name, last_name, email, and attached_policies.
Use semicolons to separate multiple attached_policies values in CSV files.

Examples:
  cubbit iam user import --sample json > users.json
  cubbit iam user import --sample csv > users.csv
  cubbit iam user import --file users.json
  cubbit iam user import --file users.csv`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := userService.ImportUsers(cmd, args); err != nil {
				utils.PrintErrorWithWriter(cmd.ErrOrStderr(), err)
			}
		},
	}

	userImportCmd.Flags().String("file", "", "Path to JSON or CSV file containing users to import")
	userImportCmd.Flags().String("sample", "", "Print a sample import file to stdout. Accepted values: json, csv")

	return userImportCmd
}
