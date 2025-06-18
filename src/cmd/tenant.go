package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Execute commands in tenant sections",
}

var createTenantSubCmd = &cobra.Command{
	Use:     "create",
	Short:   "create a new tenant",
	Aliases: []string{"new"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			cmd.MarkFlagRequired("name")
			cmd.MarkFlagRequired("settings")
			cmd.MarkFlagRequired("distributor-code")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.CreateTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantSubCmd = &cobra.Command{
	Use:     "list",
	Short:   "list tenants",
	Aliases: []string{"ls"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			allowed_sorting_keys := []string{"id", "name", "owner_id", "coupon_id", "created_at", "deleted_at"}
			sort, _ := cmd.Flags().GetString("sort")

			if sort != "" && !utils.Contains(allowed_sorting_keys, sort) {
				fmt.Println("Error: invalid sort key provided, allowed keys are: id, name, owner_id, coupon_id, created_at, deleted_at")
				cmd.Usage()
				os.Exit(1)
			}

			filter, _ := cmd.Flags().GetString("filter")
			if filter != "" {
				if !utils.IsValidFilter(filter) {
					fmt.Println("Error: invalid filter provided, allowed format is: key:value key:value ...")
					cmd.Usage()
					os.Exit(1)
				}
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ListTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeTenantSubCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a tenant",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.RemoveTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeTenantSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.DescribeTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantDescriptionSubCmd = &cobra.Command{
	Use:   "edit-description",
	Short: "edit a tenant description",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no new description argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.EditTenantDescription(cmd, args...); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditTenantDescriptionInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantImageSubCmd = &cobra.Command{
	Use:   "edit-image",
	Short: "edit a tenant image",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error:a t least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no new image url argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.EditTenantImage(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditTenantImageInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantAvailableSwarmsSubCmd = &cobra.Command{
	Use:   "list-swarms",
	Short: "lists the swarms that can be connected",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ListAvailableSwarmsTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListAvailableSwarmsTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}
var addOperatorToTenantSubCmd = &cobra.Command{
	Use:   "add-operator",
	Short: "invites an operator",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("role")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.AddOperatorToTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.AddOperatorToTenantInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantOperatorsSubCmd = &cobra.Command{
	Use:   "list-operators",
	Short: "lists tenant operators",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ListTenantOperators(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListTenantOperatorsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeTenantOperatorSubCmd = &cobra.Command{
	Use:   "remove-operator",
	Short: "removes tenant operator by email or id",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no operator argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.RemoveTenantOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveTenantOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var connectSwarmSubCmd = &cobra.Command{
	Use:   "connect-swarm",
	Short: "connects a swarm with a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no swarm argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ConnectSwarm(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ConnectSwarmInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantSettingsSubCmd = &cobra.Command{
	Use:   "edit-settings",
	Short: "edit a tenant settings",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no new settings argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.EditTenantSettings(cmd, args...); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditTenantSettingsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeTenantOperatorsSubCmd = &cobra.Command{
	Use:   "describe-operator",
	Short: "describe tenant operator",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no operator name or id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.DescribeTenantOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeTenantOperatorInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var EditTenantOperatorRoleSubCmd = &cobra.Command{
	Use:   "edit-operator",
	Short: "edit tenant operator role",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no operator id or name argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.EditTenantOperatorRole(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.EditTenantOperatorRoleInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantAccountsSubCmd = &cobra.Command{
	Use:   "list-users",
	Short: "lists tenant users",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			allowed_sorting_keys := []string{"id", "first_name", "last_name", "max_allowed_projects", "created_at", "deleted_at", "tenant_id"}
			sort, _ := cmd.Flags().GetString("sort")

			if sort != "" && !utils.Contains(allowed_sorting_keys, sort) {
				fmt.Println("Error: invalid sort key provided, allowed keys are: id, name, owner_id, coupon_id, created_at, deleted_at")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ListTenantAccounts(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListTenantAccountsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeTenantAccountSubCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "describes tenant users",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.DescribeTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeTenantAccountSubCmd = &cobra.Command{
	Use:     "remove-user",
	Short:   "removes a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.RemoveTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var banTenantAccountSubCmd = &cobra.Command{
	Use:     "freeze-user",
	Short:   "freezes a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.BanTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.BanTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var unbanTenantAccountSubCmd = &cobra.Command{
	Use:     "unfreeze-user",
	Short:   "unfreezes a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.UnbanTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.UnbanTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var restoreTenantAccountSubCmd = &cobra.Command{
	Use:     "restore-user",
	Short:   "restores a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.RestoreTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RestoreTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var deleteTenantAccountSessionsSubCmd = &cobra.Command{
	Use:     "delete-user-sessions",
	Short:   "deletes all sessions of a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.DeleteTenantAccountSessions(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DeleteTenantAccountSessionsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var createTenantAccountsSubCmd = &cobra.Command{
	Use:   "create-users",
	Short: "creates users in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("emails")

			swarms, _ := cmd.Flags().GetStringSlice("emails")
			if len(swarms) == 0 {
				fmt.Println("Error: no emails provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.CreateTenantAccounts(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.CreateTenantAccountsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var updateTenantAccountSubCmd = &cobra.Command{
	Use:   "edit-user",
	Short: "updates a user in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no user id argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.UpdateTenantAccount(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.UpdateTenantAccountInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantProjectsSubCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "lists tenant projects",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			allowed_sorting_keys := []string{"project_id", "project_name", "project_created_at", "project_deleted_at", "project_banned_at", "project_tenant_id", "project_email", "root_account_email"}
			sort, _ := cmd.Flags().GetString("sort")

			if sort != "" && !utils.Contains(allowed_sorting_keys, sort) {
				fmt.Println("Error: invalid sort key provided, allowed keys are: project_id, project_name, project_created_at, project_deleted_at, project_banned_at, project_tenant_id, project_email, root_account_email")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.ListTenantProjects(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.ListTenantProjectsInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var describeTenantProjectSubCmd = &cobra.Command{
	Use:   "describe-project",
	Short: "describes tenant projects",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.DescribeTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.DescribeTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeTenantProjectSubCmd = &cobra.Command{
	Use:     "remove-project",
	Short:   "removes a tenant project",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			cmd.MarkFlagRequired("email")
			cmd.MarkFlagRequired("password")

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.RemoveTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var banTenantProjectSubCmd = &cobra.Command{
	Use:     "freeze-project",
	Short:   "freezes a tenant project",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.BanTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.BanTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var unbanTenantProjectSubCmd = &cobra.Command{
	Use:     "unfreeze-project",
	Short:   "unfreezes a tenant project",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.UnbanTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.UnbanTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var restoreTenantProjectSubCmd = &cobra.Command{
	Use:     "restore-project",
	Short:   "restores a tenant project",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.RestoreTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RestoreTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var updateTenantProjectSubCmd = &cobra.Command{
	Use:   "edit-project",
	Short: "updates a project in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no project id argument provided")
				cmd.Usage()
				os.Exit(1)
			}

		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.UpdateTenantProject(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.UpdateTenantProjectInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantDistributorCodeSubCmd = &cobra.Command{
	Use:   "edit-distributor-code",
	Short: "assigns a tenant to a new distributor code",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !interactive {
			id, _ := cmd.Flags().GetString("id")
			name, _ := cmd.Flags().GetString("name")
			if id == "" && name == "" {
				fmt.Println("Error: at least one of the two required flags --id or --name should be provided.")
				cmd.Usage()
				os.Exit(1)
			}

			if len(args) == 0 {
				fmt.Println("Error: no distributor code argument provided")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		action.SetupOutput(cmd)

		if !interactive {
			if err = action.AssignTenantToCoupon(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.AssignTenantToCouponInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	tenantCmd.AddCommand(createTenantSubCmd)
	createTenantSubCmd.Flags().String("name", "", "Name of the tenant")
	createTenantSubCmd.Flags().String("description", "", "Description of the tenant")
	createTenantSubCmd.Flags().String("image-url", "", "Image URL of the tenant")
	createTenantSubCmd.Flags().String("settings", "", "A Json object containing the tenant settings")
	createTenantSubCmd.Flags().String("distributor-code", "", "A code provided by the Distributor that authorizes the tenant creation")
	createTenantSubCmd.Flags().String("zone", "", "Zone of the tenant creation")

	tenantCmd.AddCommand(listTenantSubCmd)
	listTenantSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for tenants")
	listTenantSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different tenants")
	listTenantSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantSubCmd.Flags().String("filter", "", "Filters the output based on the given field")

	tenantCmd.AddCommand(describeTenantSubCmd)
	describeTenantSubCmd.Flags().String("format", "default", "Formats the output")

	tenantCmd.AddCommand(editTenantDescriptionSubCmd)

	tenantCmd.AddCommand(editTenantImageSubCmd)

	tenantCmd.AddCommand(listTenantAvailableSwarmsSubCmd)
	listTenantAvailableSwarmsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for swarms")
	listTenantAvailableSwarmsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different swarms")

	tenantCmd.AddCommand(addOperatorToTenantSubCmd)
	addOperatorToTenantSubCmd.Flags().String("email", "", "Email of the operator")
	addOperatorToTenantSubCmd.Flags().String("role", "", "Role of the operator")
	addOperatorToTenantSubCmd.Flags().String("first-name", "", "First name of the operator")
	addOperatorToTenantSubCmd.Flags().String("last-name", "", "Last name of the operator")

	tenantCmd.AddCommand(listTenantOperatorsSubCmd)
	listTenantOperatorsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for operators")
	listTenantOperatorsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different operators")

	tenantCmd.AddCommand(removeTenantOperatorSubCmd)

	tenantCmd.AddCommand(connectSwarmSubCmd)

	tenantCmd.AddCommand(removeTenantSubCmd)
	removeTenantSubCmd.Flags().String("email", "", "Email address")
	removeTenantSubCmd.Flags().String("password", "", "Password")
	removeTenantSubCmd.Flags().String("code", "", "Two factor authentication code")

	tenantCmd.AddCommand(editTenantSettingsSubCmd)

	tenantCmd.AddCommand(describeTenantOperatorsSubCmd)
	describeTenantOperatorsSubCmd.Flags().String("format", "default", "Formats the output")

	tenantCmd.AddCommand(EditTenantOperatorRoleSubCmd)
	EditTenantOperatorRoleSubCmd.Flags().String("role", "", "Role of the operator")

	tenantCmd.AddCommand(listTenantAccountsSubCmd)
	listTenantAccountsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for users")
	listTenantAccountsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different users")
	listTenantAccountsSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantAccountsSubCmd.Flags().String("filter", "", "Filters the output based on the given field")

	tenantCmd.AddCommand(describeTenantAccountSubCmd)
	describeTenantAccountSubCmd.Flags().String("format", "default", "Formats the output")

	tenantCmd.AddCommand(removeTenantAccountSubCmd)
	removeTenantAccountSubCmd.Flags().String("email", "", "Email address")
	removeTenantAccountSubCmd.Flags().String("password", "", "Password")
	removeTenantAccountSubCmd.Flags().String("code", "", "Two factor authentication code")

	tenantCmd.AddCommand(banTenantAccountSubCmd)
	tenantCmd.AddCommand(unbanTenantAccountSubCmd)

	tenantCmd.AddCommand(restoreTenantAccountSubCmd)

	tenantCmd.AddCommand(deleteTenantAccountSessionsSubCmd)

	tenantCmd.AddCommand(createTenantAccountsSubCmd)
	createTenantAccountsSubCmd.Flags().StringSlice("emails", []string{}, "list of users emails to create")

	tenantCmd.AddCommand(updateTenantAccountSubCmd)
	updateTenantAccountSubCmd.Flags().String("first-name", "", "First name of the user")
	updateTenantAccountSubCmd.Flags().String("last-name", "", "Last name of the user")
	updateTenantAccountSubCmd.Flags().String("endpoint-gateway", "", "Endpoint gateway of the user")
	updateTenantAccountSubCmd.Flags().Bool("internal", false, "Defines if the user is internal")
	updateTenantAccountSubCmd.Flags().Int("max-allowed-projects", 1, "Max allowed projects for the user")

	tenantCmd.AddCommand(listTenantProjectsSubCmd)
	listTenantProjectsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for projects")
	listTenantProjectsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different projects")
	listTenantProjectsSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantProjectsSubCmd.Flags().String("filter", "", "Filters the output based on the given field")

	tenantCmd.AddCommand(describeTenantProjectSubCmd)
	describeTenantProjectSubCmd.Flags().String("format", "default", "Formats the output")

	tenantCmd.AddCommand(removeTenantProjectSubCmd)
	removeTenantProjectSubCmd.Flags().String("email", "", "Email address")
	removeTenantProjectSubCmd.Flags().String("password", "", "Password")
	removeTenantProjectSubCmd.Flags().String("code", "", "Two factor authentication code")

	tenantCmd.AddCommand(banTenantProjectSubCmd)
	tenantCmd.AddCommand(unbanTenantProjectSubCmd)
	tenantCmd.AddCommand(restoreTenantProjectSubCmd)

	tenantCmd.AddCommand(updateTenantProjectSubCmd)
	updateTenantProjectSubCmd.Flags().String("description", "", "Description of the project")
	updateTenantProjectSubCmd.Flags().String("image-url", "", "Image URL of the project")
	tenantCmd.AddCommand(editTenantDistributorCodeSubCmd)

	rootCmd.AddCommand(tenantCmd)
	tenantCmd.PersistentFlags().String("name", "", "Name of the tenant")
	tenantCmd.PersistentFlags().String("id", "", "ID of the tenant")
}
