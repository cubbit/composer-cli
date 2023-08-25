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
			if err = tui.Send(cmd, args, action.CreateTenant); err != nil {
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
			if err = tui.Send(cmd, args, action.ListTenant); err != nil {
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
	Use:   "remove",
	Short: "remove a tenant",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveTenant); err != nil {
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
			if err = tui.Send(cmd, args, action.DescribeTenant); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
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
	Use:   "list-available-swarms",
	Short: "lists the swarms that can be connected",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.ListAvailableSwarmsTenant); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.AddOperatorToTenant); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.ListTenantOperators); err != nil {
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
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if !interactive {
			if err = tui.Send(cmd, args, action.RemoveTenantOperator); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err = action.RemoveTenantOperatorInteractive(cmd); err != nil {
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
	listTenantAvailableSwarmsSubCmd.Flags().String("id", "", "ID of the tenant")
	listTenantAvailableSwarmsSubCmd.Flags().String("name", "", "Name of the tenant")

	tenantCmd.AddCommand(addOperatorToTenantSubCmd)
	addOperatorToTenantSubCmd.Flags().String("email", "", "Email of the operator")
	addOperatorToTenantSubCmd.Flags().String("role", "", "Role of the operator")

	tenantCmd.AddCommand(listTenantOperatorsSubCmd)
	listTenantOperatorsSubCmd.Flags().BoolP("verbose", "v", false, "Lists all available information for operators")
	listTenantOperatorsSubCmd.Flags().BoolP("line", "l", false, "Adds a line between the information about different operators")

	tenantCmd.AddCommand(removeTenantOperatorSubCmd)

	tenantCmd.AddCommand(removeTenantSubCmd)
	removeTenantSubCmd.Flags().String("id", "", "ID of the tenant")
	removeTenantSubCmd.Flags().String("name", "", "Name of the tenant")
	removeTenantSubCmd.Flags().String("email", "", "Email address")
	removeTenantSubCmd.Flags().String("password", "", "Password")
	removeTenantSubCmd.Flags().String("code", "", "Two factor authentication code")

	rootCmd.AddCommand(tenantCmd)
	tenantCmd.PersistentFlags().String("name", "", "Name of the tenant")
	tenantCmd.PersistentFlags().String("id", "", "ID of the tenant")
}
