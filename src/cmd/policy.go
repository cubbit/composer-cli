// Package cmd provides CLI commands for managing IAM policies.
package cmd

import (
	"github.com/cubbit/composer-cli/src/action"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

var IAMPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Execute commands in iam policy sections",
}

var promoteSubCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote an operator to a higher policy",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("email")
		cmd.MarkFlagRequired("policy-name")
		cmd.MarkFlagRequired("secret")

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := action.PromoteOperator(cmd, args); err != nil {
			utils.PrintError(err)
		}

	},
}

var EditIAMOperatorPolicySubCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit tenant/swarm operator policy",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagRequired("user-id")
		cmd.MarkFlagRequired("policy-id")

		cmd.MarkFlagsOneRequired("tenant-id", "swarm-id")
		cmd.MarkFlagsMutuallyExclusive("tenant-id", "swarm-id")

	},
	Run: func(cmd *cobra.Command, args []string) {
		var tenantID, swarmID string
		tenantID, _ = cmd.Flags().GetString("tenant-id")
		swarmID, _ = cmd.Flags().GetString("swarm-id")

		if tenantID != "" {
			if err := action.EditTenantOperatorRole(cmd, args); err != nil {
				utils.PrintError(err)
			}
		} else if swarmID != "" {
			if err := action.EditSwarmOperatorRole(cmd, args); err != nil {
				utils.PrintError(err)
			}
		}
	},
}

func init() {

	IAMPolicyCmd.AddCommand(promoteSubCmd)
	promoteSubCmd.Flags().String("api-server-url", "https://api.cubbit.eu/iam", "Api server URL")
	promoteSubCmd.Flags().String("email", "", "Email Address")
	promoteSubCmd.Flags().String("policy-name", "system-admin", "Policy name")
	promoteSubCmd.Flags().String("secret", "", "Secret")

	IAMPolicyCmd.AddCommand(EditIAMOperatorPolicySubCmd)
	EditIAMOperatorPolicySubCmd.Flags().String("user-id", "", "ID of the operator")
	EditIAMOperatorPolicySubCmd.Flags().String("policy-id", "", "ID of the policy")
	EditIAMOperatorPolicySubCmd.Flags().String("tenant-id", "", "ID of the tenant")
	EditIAMOperatorPolicySubCmd.Flags().String("swarm-id", "", "ID of the swarm")

	iamCmd.AddCommand(IAMPolicyCmd)
}
