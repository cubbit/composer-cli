// Package cmd provides CLI commands for managing tenants.
package cmd

import (
	"fmt"
	"os"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Execute commands in tenant sections",
}

var createTenantSubCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("name")
		cmd.MarkFlagRequired("distributor-code")

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateTenant(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listTenantSubCmd = &cobra.Command{
	Use:     "list",
	Short:   "list tenants",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		allowedSortingKeys := []string{"id", "name", "owner_id", "coupon_id", "created_at", "deleted_at"}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			fmt.Println("Error: invalid sort key provided, allowed keys are: id, name, owner_id, coupon_id, created_at, deleted_at")
			return fmt.Errorf("invalid sort key: %s", sort)
		}

		filter, _ := cmd.Flags().GetString("filter")
		if filter != "" {
			if !utils.IsValidFilter(filter) {
				fmt.Println("Error: invalid filter provided, allowed format is: key:value key:value ...")
				return fmt.Errorf("invalid filter: %s", filter)
			}
		}

		if err := action.ListTenant(cmd, args); err != nil {
			utils.PrintError(err)
			return err
		}
		return nil
	},
}

var removeTenantSubCmd = &cobra.Command{
	Use:     "remove",
	Short:   "remove a tenant",
	Aliases: []string{"rm"},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveTenant(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeTenantSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeTenant(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var editTenantSubCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit tenant properties or connect swarm to tenant",
	Long: `Edit tenant properties or connect swarm to tenant.

This command supports two distinct operations:

1. Edit tenant properties:
   --tenant-id --description --settings

2. Connect swarm to tenant:
   --tenant-id --swarm-id --rc-id [--default]

Note: These operations cannot be mixed. Use either property flags or swarm connection flags.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")

		hasPropertyFlags := cmd.Flags().Changed("description") || cmd.Flags().Changed("settings")
		hasSwarmFlags := cmd.Flags().Changed("swarm-id") || cmd.Flags().Changed("rc-id") || cmd.Flags().Changed("default")

		if hasPropertyFlags && hasSwarmFlags {
			cmd.PrintErr("Error: Cannot mix property flags (--description, --settings) with swarm flags (--swarm-id, --rc-id, --default)\n")
			cmd.PrintErr("Use either property editing or swarm connection, not both\n")
			os.Exit(1)
		}

		if hasSwarmFlags {
			if swarmID := cmd.Flags().Lookup("swarm-id"); swarmID != nil {
				if swarmID.Value.String() != "" {
					cmd.MarkFlagRequired("rc-id")
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.EditTenant(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var reportTenantSubCmd = &cobra.Command{
	Use:   "report",
	Short: "downloads/prints a full report for the tenant",
	PreRun: func(cmd *cobra.Command, args []string) {

		cmd.MarkFlagRequired("tenant-id")
		cmd.MarkFlagRequired("from")
		cmd.MarkFlagRequired("to")

		isChanged := cmd.Flags().Changed("output-dir")
		if isChanged {
			outputDir, _ := cmd.Flags().GetString("output-dir")

			if outputDir == "" {
				fmt.Println("Error: output cannot be empty.Use a dot (.) to indicate the current directory.")
				cmd.Usage()
				os.Exit(1)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.GetTenantReport(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var configureTenantDNSSubCmd = &cobra.Command{
	Use:   "configure-dns",
	Short: "configures DNS for a tenant",
	Long:  "This command prints the value of the TXT record that needs to be added with the name '_acme-challenge'",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsOneRequired("interactive", "tenant-id")
		cmd.MarkFlagsRequiredTogether("tenant-id", "domain")
		cmd.MarkFlagsMutuallyExclusive("interactive", "domain")
		cmd.MarkFlagsMutuallyExclusive("interactive", "tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		interactive, _ := cmd.Flags().GetBool("interactive")
		if !interactive {
			if err := action.ConfigureTenantDNS(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.ConfigureAndVerifyDNSInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

var verifyTenantDNSSubCmd = &cobra.Command{
	Use:   "verify-dns",
	Short: "verifies DNS for a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsOneRequired("interactive", "tenant-id")
		cmd.MarkFlagsMutuallyExclusive("interactive", "tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		interactive, _ := cmd.Flags().GetBool("interactive")
		if !interactive {
			if err := action.VerifyTenantDNS(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.ConfigureAndVerifyDNSInteractive(cmd); err != nil {
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
	listTenantSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listTenantSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is: key:value")

	tenantCmd.AddCommand(describeTenantSubCmd)
	describeTenantSubCmd.Flags().String("tenant-id", "", "ID of the tenant")

	tenantCmd.AddCommand(editTenantSubCmd)
	editTenantSubCmd.Flags().String("tenant-id", "", "ID of the tenant")
	editTenantSubCmd.Flags().String("description", "", "Description of the tenant")
	editTenantSubCmd.Flags().String("settings", "", "A Json object containing the tenant settings")
	editTenantSubCmd.Flags().String("swarm-id", "", "ID of the swarm to connect to the tenant")
	editTenantSubCmd.Flags().String("rc-id", "", "ID of the redundancy class to use for the swarm")
	editTenantSubCmd.Flags().Bool("default", false, "Sets the swarm as the default swarm for the tenant")

	tenantCmd.AddCommand(removeTenantSubCmd)
	removeTenantSubCmd.Flags().String("tenant-id", "", "ID of the tenant")

	tenantCmd.AddCommand(reportTenantSubCmd)
	reportTenantSubCmd.Flags().String("tenant-id", "", "ID of the tenant")
	reportTenantSubCmd.Flags().String("from", "", "Start date and time in DD/MM/YYYY+HH:mm:ss format")
	reportTenantSubCmd.Flags().String("to", "", "End date and time in DD/MM/YYYY+HH:mm:ss format")
	reportTenantSubCmd.Flags().String("output-dir", "", "Directory to save the report file, if not provided, the report will be printed to the console")

	tenantCmd.AddCommand(configureTenantDNSSubCmd)
	configureTenantDNSSubCmd.Flags().String("tenant-id", "", "ID of the tenant")
	configureTenantDNSSubCmd.Flags().String("domain", "", "Domain to configure for the tenant")
	configureTenantDNSSubCmd.Flags().Bool("force", false, "Force the configuration of DNS even if it already exists")
	configureTenantDNSSubCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")

	tenantCmd.AddCommand(verifyTenantDNSSubCmd)
	verifyTenantDNSSubCmd.Flags().String("tenant-id", "", "ID of the tenant")
	verifyTenantDNSSubCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")

	rootCmd.AddCommand(tenantCmd)
}
