package cmd

import (
	"os"

	cmd_auth "github.com/cubbit/composer-cli/src/cmd/auth"
	cmd_config "github.com/cubbit/composer-cli/src/cmd/config"
	cmd_docs "github.com/cubbit/composer-cli/src/cmd/docs"
	cmd_infrastructure "github.com/cubbit/composer-cli/src/cmd/infrastructure"
	cmd_operator "github.com/cubbit/composer-cli/src/cmd/operator"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func NewRootCommand(
	newAgentService service.AgentServiceInterface,
	authService service.AuthServiceInterface,
	operatorService service.OperatorServiceInterface,
	locationService service.LocationServiceInterface,
	configService service.ConfigServiceInterface,
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
	rootCommand.PersistentFlags().String("profile", "", "Profile Configuration")
	rootCommand.PersistentFlags().String("output", "human", "Output format: human (default), json, yaml, xml")
	rootCommand.PersistentFlags().Bool("no-headers", false, "Suppress table headers in human output (for easier scripting)")
	rootCommand.PersistentFlags().Bool("quiet", false, "Minimize stdout for CI/CD workflows (no table output, just essentials)")
	rootCommand.PersistentFlags().Bool("silent", false, "Redirect all output to /dev/null")

	// Local flags (only available to root command)
	rootCommand.Flags().String("endpoint", "", "Override the default API endpoint URL")

	authCmd := cmd_auth.NewAuthCmd(authService)
	rootCommand.AddCommand(authCmd)

	operatorCmd := cmd_operator.NewOperatorCmd(operatorService)
	rootCommand.AddCommand(operatorCmd)

	infrastructureCmd := cmd_infrastructure.NewInfrastructureCmd(locationService)
	rootCommand.AddCommand(infrastructureCmd)

	configCmd := cmd_config.NewConfigCmd(configService)
	rootCommand.AddCommand(configCmd)

	docsCmd := cmd_docs.NewDocsCmd()
	rootCommand.AddCommand(docsCmd)

	return rootCommand
}
