// Package cmd provides CLI commands for managing tenant users.
package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Execute commands in user sections",
}

var createTenantUsersSubCmd = &cobra.Command{
	Use:   "create",
	Short: "creates users in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("emails")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateTenantAccounts(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeTenantUserSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describes tenant users",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var updateTenantUserSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "updates a user in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.UpdateTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeTenantUserSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a tenant user",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var restoreTenantUserSubCmd = &cobra.Command{
	Use:     "restore",
	Short:   "restores a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RestoreTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var banTenantUserSubCmd = &cobra.Command{
	Use:   "freeze",
	Short: "freezes a tenant user",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.BanTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var unbanTenantUserSubCmd = &cobra.Command{
	Use:     "unfreeze",
	Short:   "unfreezes a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.UnbanTenantAccount(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var deleteTenantUserSessionsSubCmd = &cobra.Command{
	Use:     "force-logout",
	Short:   "deletes all sessions of a tenant user",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DeleteTenantAccountSessions(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listTenantUsersSubCmd = &cobra.Command{
	Use:   "list",
	Short: "lists tenant users",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")
		allowedSortingKeys := []string{"id", "first_name", "last_name", "max_allowed_projects", "created_at", "deleted_at", "tenant_id"}
		sort, _ := cmd.Flags().GetString("sort")

		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			fmt.Println("Error: invalid sort key provided, allowed keys are: id, name, owner_id, coupon_id, created_at, deleted_at")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.ListTenantAccounts(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	userCmd.AddCommand(createTenantUsersSubCmd)
	createTenantUsersSubCmd.Flags().StringSlice("emails", []string{}, "list of users emails to create")

	userCmd.AddCommand(describeTenantUserSubCmd)
	describeTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(updateTenantUserSubCmd)
	updateTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")
	updateTenantUserSubCmd.Flags().String("first-name", "", "First name of the user")
	updateTenantUserSubCmd.Flags().String("last-name", "", "Last name of the user")
	updateTenantUserSubCmd.Flags().String("endpoint-gateway", "", "Endpoint gateway of the user")
	updateTenantUserSubCmd.Flags().Bool("internal", false, "Defines if the user is internal")
	updateTenantUserSubCmd.Flags().Int("max-allowed-projects", 1, "Max allowed projects for the user")

	userCmd.AddCommand(removeTenantUserSubCmd)
	removeTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(restoreTenantUserSubCmd)
	restoreTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(banTenantUserSubCmd)
	banTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(unbanTenantUserSubCmd)
	unbanTenantUserSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(deleteTenantUserSessionsSubCmd)
	deleteTenantUserSessionsSubCmd.Flags().String("user-id", "", "ID of the user")

	userCmd.AddCommand(listTenantUsersSubCmd)
	listTenantUsersSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantUsersSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	tenantCmd.AddCommand(userCmd)
	userCmd.PersistentFlags().String("tenant-id", "", "ID of the tenant")

}
