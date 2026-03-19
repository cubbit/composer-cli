// Package cmd provides CLI commands for managing IAM operators.
package cmd

import (
	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var IAMUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Execute commands in iam user sections",
}

var addIAMOperatorSubCmd = &cobra.Command{
	Use:   "create",
	Short: "invites an operator into a tenant/swarm",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("email")
		cmd.MarkFlagRequired("policy-id")

		cmd.MarkFlagsOneRequired("tenant-id", "swarm-id")
		cmd.MarkFlagsMutuallyExclusive("tenant-id", "swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		tenantID, _ := cmd.Flags().GetString("tenant-id")
		if tenantID != "" {
			if err := action.AddOperatorToTenant(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}

		swarmID, _ := cmd.Flags().GetString("swarm-id")
		if swarmID != "" {
			if err := action.AddOperatorToSwarm(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}
	},
}

var listIAMOperatorsSubCmd = &cobra.Command{
	Use:   "list",
	Short: "lists tenant/swarm operators",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsOneRequired("tenant-id", "swarm-id")
		cmd.MarkFlagsMutuallyExclusive("tenant-id", "swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		tenantID, _ := cmd.Flags().GetString("tenant-id")
		if tenantID != "" {
			if err := action.ListTenantOperators(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}

		swarmID, _ := cmd.Flags().GetString("swarm-id")
		if swarmID != "" {
			if err := action.ListSwarmOperators(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}
	},
}

var removeIAMOperatorSubCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes tenant/swarm operator",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")

		cmd.MarkFlagsOneRequired("tenant-id", "swarm-id")
		cmd.MarkFlagsMutuallyExclusive("tenant-id", "swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		tenantID, _ := cmd.Flags().GetString("tenant-id")

		if tenantID != "" {
			if err := action.RemoveTenantOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}

		swarmID, _ := cmd.Flags().GetString("swarm-id")
		if swarmID != "" {
			if err := action.RemoveSwarmOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}
	},
}

var describeIAMOperatorsSubCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe tenant/swarm operator",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")

		cmd.MarkFlagsOneRequired("tenant-id", "swarm-id")
		cmd.MarkFlagsMutuallyExclusive("tenant-id", "swarm-id")
	},
	Run: func(cmd *cobra.Command, args []string) {
		tenantID, _ := cmd.Flags().GetString("tenant-id")

		if tenantID != "" {
			if err := action.DescribeTenantOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}

		swarmID, _ := cmd.Flags().GetString("swarm-id")
		if swarmID != "" {
			if err := action.DescribeSwarmOperator(cmd, args); err != nil {
				utils.PrintError(err)
			}

			return
		}
	},
}

func init() {

	IAMUserCmd.AddCommand(addIAMOperatorSubCmd)
	addIAMOperatorSubCmd.Flags().String("email", "", "Email of the operator")
	addIAMOperatorSubCmd.Flags().String("policy-id", "", "ID of the policy to assign to the operator")
	addIAMOperatorSubCmd.Flags().String("first-name", "", "First name of the operator")
	addIAMOperatorSubCmd.Flags().String("last-name", "", "Last name of the operator")

	IAMUserCmd.AddCommand(listIAMOperatorsSubCmd)
	listIAMOperatorsSubCmd.Flags().String("sort", "", "Sorts the output based on the given field")
	listIAMOperatorsSubCmd.Flags().String("filter", "", "Filters the output based on the given field, allowed format is key:value")

	IAMUserCmd.AddCommand(removeIAMOperatorSubCmd)
	removeIAMOperatorSubCmd.Flags().String("user-id", "", "ID of the operator")

	IAMUserCmd.AddCommand(describeIAMOperatorsSubCmd)
	describeIAMOperatorsSubCmd.Flags().String("user-id", "", "ID of the operator")

	iamCmd.AddCommand(IAMUserCmd)
	IAMUserCmd.PersistentFlags().String("tenant-id", "", "ID of the tenant")
	IAMUserCmd.PersistentFlags().String("swarm-id", "", "ID of the swarm")

}
