package cmd

import (
	"os"

	cmd_agent "github.com/cubbit/composer-cli/src/cmd/agent"
	cmd_auth "github.com/cubbit/composer-cli/src/cmd/auth"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewRootCommand(
	newAgentService service.AgentServiceInterface,
	authService service.AuthServiceInterface,
) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "cubbit",
		Short: "The official Cubbit CLI (Command-Line Interface) for operators",
		Long:  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			setupSilentMode(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			endpoint, _ := cmd.Flags().GetString("endpoint")
			if endpoint != "" {
				err := configuration.SetAPIEndpoint(endpoint)
				if err != nil {
					os.Exit(1)
				}
				return
			}

			cmd.Help()
		},
	}

	// Persistent flags (available to all subcommands)
	rootCommand.PersistentFlags().BoolVarP(&interactive, "interactive", "i", false, "Run in interactive mode")
	rootCommand.PersistentFlags().String("profile", "", "Profile Configuration")
	rootCommand.PersistentFlags().String("output", "human", "Output format: human (default), json, yaml, xml")
	rootCommand.PersistentFlags().Bool("no-headers", false, "Suppress table headers in human output (for easier scripting)")
	rootCommand.PersistentFlags().Bool("quiet", false, "Minimize stdout for CI/CD workflows (no table output, just essentials)")
	rootCommand.PersistentFlags().Bool("silent", false, "Redirect all output to /dev/null")

	// Local flags (only available to root command)
	rootCommand.Flags().String("endpoint", "", "Override the default API endpoint URL")

	agentCmd := cmd_agent.NewAgentCmd(newAgentService)
	rootCommand.AddCommand(agentCmd)

	authCmd := cmd_auth.NewAuthCmd(authService)
	rootCommand.AddCommand(authCmd)

	return rootCommand
}
