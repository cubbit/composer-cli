package cmd

import (
	"github.com/cubbit/cubbit/client/cli/src/action"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Execute commands in tenant sections",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var createTenantSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new tenant",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.CreateTenant); err != nil {
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
	Use:   "list",
	Short: "list tenants",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.ListTenant); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var removeTenantSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a tenant",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.RemoveTenant); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.DescribeTenant); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantDescriptionSubCmd = &cobra.Command{
	Use:   "edit-description",
	Short: "edit a tenant description",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.EditTenantDescription(cmd, args); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var editTenantImageSubCmd = &cobra.Command{
	Use:   "edit-image",
	Short: "edit a tenant image",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = action.EditTenantImage(cmd, args); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var listTenantAvailableSwarmsSubCmd = &cobra.Command{
	Use:   "list-available-swarms",
	Short: "lists the swarms that can be connected",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, action.ListAvailableSwarmsTenant); err != nil {
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

	tenantCmd.AddCommand(listTenantSubCmd)
	listTenantSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for tenants")
	listTenantSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different tenants")

	tenantCmd.AddCommand(describeTenantSubCmd)
	describeTenantSubCmd.Flags().String("id", "", "ID of the tenant")
	describeTenantSubCmd.Flags().String("name", "", "Name of the tenant")
	describeTenantSubCmd.Flags().String("format", "default", "Formats the output")

	tenantCmd.AddCommand(editTenantDescriptionSubCmd)

	tenantCmd.AddCommand(editTenantImageSubCmd)

	tenantCmd.AddCommand(listTenantAvailableSwarmsSubCmd)

	tenantCmd.AddCommand(removeTenantSubCmd)
	removeTenantSubCmd.Flags().String("id", "", "ID of the tenant")
	removeTenantSubCmd.Flags().String("name", "", "Name of the tenant")
	removeTenantSubCmd.Flags().String("email", "", "Email address")
	removeTenantSubCmd.Flags().String("password", "", "Password")
	removeTenantSubCmd.Flags().String("code", "", "Two factor authentication code")

	rootCmd.AddCommand(tenantCmd)
	tenantCmd.Flags().String("name", "", "Name of the tenant")
	tenantCmd.Flags().String("id", "", "ID of the tenant")
}
