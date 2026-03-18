// Package cmd provides CLI commands for managing projects.
package cmd

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Execute commands in project sections",
}

var listTenantProjectsSubCmd = &cobra.Command{
	Use:   "list",
	Short: "lists tenant projects",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		allowedSortingKeys := []string{"project_id", "project_name", "project_created_at", "project_deleted_at", "project_banned_at", "project_tenant_id", "project_email", "root_account_email"}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			return fmt.Errorf("invalid sort key provided, allowed keys are: project_id, project_name, project_created_at, project_deleted_at, project_banned_at, project_tenant_id, project_email, root_account_email")
		}
		if err := action.ListTenantProjects(cmd, args); err != nil {
			utils.PrintError(err)
			return nil
		}
		return nil
	},
}

var describeTenantProjectSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describes tenant projects",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeTenantProjectSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a tenant project",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var banTenantProjectSubCmd = &cobra.Command{
	Use:   "freeze",
	Short: "freezes a tenant project",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.BanTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var unbanTenantProjectSubCmd = &cobra.Command{
	Use:   "unfreeze",
	Short: "unfreezes a tenant project",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.UnbanTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var restoreTenantProjectSubCmd = &cobra.Command{
	Use:   "restore",
	Short: "restores a tenant project",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RestoreTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var updateTenantProjectSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "updates a project in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("project-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.UpdateTenantProject(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

func init() {
	projectCmd.AddCommand(describeTenantProjectSubCmd)
	describeTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")

	projectCmd.AddCommand(updateTenantProjectSubCmd)
	updateTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")
	updateTenantProjectSubCmd.Flags().String("name", "", "Name of the project")
	updateTenantProjectSubCmd.Flags().String("description", "", "Description of the project")
	updateTenantProjectSubCmd.Flags().String("image-url", "", "Image URL of the project")

	projectCmd.AddCommand(banTenantProjectSubCmd)
	banTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")

	projectCmd.AddCommand(unbanTenantProjectSubCmd)
	unbanTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")

	projectCmd.AddCommand(removeTenantProjectSubCmd)
	removeTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")

	projectCmd.AddCommand(restoreTenantProjectSubCmd)
	restoreTenantProjectSubCmd.Flags().String("project-id", "", "ID of the project")

	projectCmd.AddCommand(listTenantProjectsSubCmd)
	listTenantProjectsSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantProjectsSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	tenantCmd.AddCommand(projectCmd)
	projectCmd.PersistentFlags().String("tenant-id", "", "ID of the tenant")

}
