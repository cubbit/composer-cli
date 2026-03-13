// Package cmd provides CLI commands for managing tenant gateways.
package cmd

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Execute commands in gateway sections",
}

var createGatewaySubCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a gateway for a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("name")
		cmd.MarkFlagRequired("location")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.CreateGateway(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var describeGatewaySubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describes tenant gateways",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("gateway-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.DescribeGateway(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var updateGatewaySubCmd = &cobra.Command{
	Use:   "edit",
	Short: "updates a gateway in a tenant",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("gateway-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.UpdateGateway(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var listGatewaysSubCmd = &cobra.Command{
	Use:   "list",
	Short: "lists tenant gateways",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		allowedSortingKeys := []string{"id, name"}
		sort, _ := cmd.Flags().GetString("sort")
		if sort != "" && !utils.Contains(allowedSortingKeys, sort) {
			fmt.Println("Error: invalid sort key provided, allowed keys are: id, name")
			return
		}

		if err := action.ListGateways(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var removeGatewaySubCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a tenant gateway",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("gateway-id")
		cmd.MarkFlagRequired("tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.RemoveGateway(cmd, args); err != nil {
			utils.PrintError(err)
		}
	},
}

var installGatewaySubCmd = &cobra.Command{
	Use:     "install",
	Short:   "installs a gateway for a tenant",
	Example: "cubbit gateway install --interactive\ncubbit gateway install --tenant-id <tenant-id> --gateway-id <gateway-id>",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsOneRequired("interactive", "gateway-id")

		cmd.MarkFlagsRequiredTogether("tenant-id", "gateway-id")
		cmd.MarkFlagsMutuallyExclusive("interactive", "gateway-id")
		cmd.MarkFlagsMutuallyExclusive("interactive", "tenant-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		interactive, _ := cmd.Flags().GetBool("interactive")
		if !interactive {
			if err := action.InstallGateway(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else {
			if err := action.InstallTenantGatewayInteractive(cmd); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {
	gatewayCmd.AddCommand(createGatewaySubCmd)
	createGatewaySubCmd.Flags().String("name", "", "Name of the gateway")
	createGatewaySubCmd.Flags().String("location", "", "Location of the gateway")

	gatewayCmd.AddCommand(updateGatewaySubCmd)
	updateGatewaySubCmd.Flags().String("gateway-id", "", "ID of the gateway")
	updateGatewaySubCmd.Flags().String("name", "", "Name of the gateway")
	updateGatewaySubCmd.Flags().String("location", "", "Location of the gateway")
	updateGatewaySubCmd.Flags().StringP("default-redundancy-class-id", "r", "", "Default redundancy class ID of the gateway")
	updateGatewaySubCmd.Flags().BoolP("smart-data-placement-enabled", "e", false, "Enable or disable smart data placement for the gateway")
	updateGatewaySubCmd.Flags().StringP("smart-data-placement-policies", "p", "", "The policies for smart data placement in JSON format")

	gatewayCmd.AddCommand(listGatewaysSubCmd)
	listGatewaysSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listGatewaysSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	gatewayCmd.AddCommand(describeGatewaySubCmd)
	describeGatewaySubCmd.Flags().String("gateway-id", "", "ID of the gateway")

	gatewayCmd.AddCommand(removeGatewaySubCmd)
	removeGatewaySubCmd.Flags().String("gateway-id", "", "ID of the gateway")

	gatewayCmd.AddCommand(installGatewaySubCmd)
	installGatewaySubCmd.Flags().String("gateway-id", "", "ID of the gateway")
	installGatewaySubCmd.Flags().String("cache", "", "Cache path")
	installGatewaySubCmd.Flags().String("cert-root", "./cert", "Certificate root path")
	installGatewaySubCmd.Flags().Bool("no-tls", false, "Disable TLS")
	installGatewaySubCmd.Flags().Bool("no-init", false, "Skip node initialization")
	installGatewaySubCmd.Flags().Bool("no-infra", false, "Skip infrastructure setup")
	installGatewaySubCmd.Flags().Bool("no-app", false, "Skip application setup")
	installGatewaySubCmd.Flags().Bool("no-console", false, "Skip console setup")
	installGatewaySubCmd.Flags().Bool("no-offloader", false, "Skip offloader setup")
	installGatewaySubCmd.Flags().Bool("no-s3", false, "Skip S3 setup")
	installGatewaySubCmd.Flags().Bool("ingress", false, "Install only ingress")
	installGatewaySubCmd.Flags().BoolP("interactive", "i", false, "Run in interactive mode")

	rootCmd.AddCommand(gatewayCmd)
	gatewayCmd.PersistentFlags().String("tenant-id", "", "ID of the tenant")
}
