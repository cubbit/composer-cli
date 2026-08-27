package cmd

import (
	cmd_auth "github.com/cubbit/composer-cli/src/cmd/auth"
	cmd_config "github.com/cubbit/composer-cli/src/cmd/config"
	cmd_docs "github.com/cubbit/composer-cli/src/cmd/docs"
	cmd_domain "github.com/cubbit/composer-cli/src/cmd/domain"
	cmd_gateway "github.com/cubbit/composer-cli/src/cmd/gateway"
	cmd_iam "github.com/cubbit/composer-cli/src/cmd/iam"
	cmd_infrastructure "github.com/cubbit/composer-cli/src/cmd/infrastructure"
	cmd_operator "github.com/cubbit/composer-cli/src/cmd/operator"
	cmd_swarm "github.com/cubbit/composer-cli/src/cmd/swarm"
	cmd_tenant "github.com/cubbit/composer-cli/src/cmd/tenant"
	cmd_version "github.com/cubbit/composer-cli/src/cmd/version"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service"
	api_key_service "github.com/cubbit/composer-cli/src/service/api_key"
	gateway_service "github.com/cubbit/composer-cli/src/service/gateway"
	tenant_service "github.com/cubbit/composer-cli/src/service/tenant"
	user_service "github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func NewRootCommand(
	configurationHandler configuration_handler.ConfigurationHandlerInterface,
	authService service.AuthServiceInterface,
	operatorService service.OperatorServiceInterface,
	userService user_service.UserServiceInterface,
	apiKeyService api_key_service.APIKeyServiceInterface,
	locationService service.LocationServiceInterface,
	configService service.ConfigServiceInterface,
	swarmService service.SwarmServiceInterface,
	domainService service.DomainServiceInterface,
	tenantService tenant_service.TenantServiceInterface,
	gatewayService gateway_service.GatewayServiceInterface,
	version string,
) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "cubbit",
		Short: "The official Cubbit CLI (Command-Line Interface) for operators",
		Long:  "The CLI for managing operators, tenants and swarms in Cubbit distributed datacenter",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			setupSilentMode(cmd)

			if err := configurationHandler.Load(cmd, args); err != nil {
				return err
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Persistent flags (available to all subcommands)
	rootCommand.PersistentFlags().String("config-path", "", "Path to configuration file (default is $HOME/.config/cubbit/config.toml)")
	rootCommand.PersistentFlags().String("profile", "", "Profile Configuration")
	rootCommand.PersistentFlags().String("output", "human", "Output format: human (default), json, yaml")
	rootCommand.PersistentFlags().Bool("no-headers", false, "Suppress table headers in human output (for easier scripting)")
	rootCommand.PersistentFlags().Bool("quiet", false, "Minimize stdout for CI/CD workflows (no table output, just essentials)")
	rootCommand.PersistentFlags().Bool("silent", false, "Redirect all output to /dev/null")

	authCmd := cmd_auth.NewAuthCmd(authService)
	rootCommand.AddCommand(authCmd)

	operatorCmd := cmd_operator.NewOperatorCmd(operatorService)
	rootCommand.AddCommand(operatorCmd)

	iamCmd := cmd_iam.NewIAMCmd(userService, apiKeyService)
	rootCommand.AddCommand(iamCmd)

	infrastructureCmd := cmd_infrastructure.NewInfrastructureCmd(locationService)
	rootCommand.AddCommand(infrastructureCmd)

	configCmd := cmd_config.NewConfigCmd(configService)
	rootCommand.AddCommand(configCmd)

	docsCmd := cmd_docs.NewDocsCmd()
	rootCommand.AddCommand(docsCmd)

	versionCmd := cmd_version.NewVersionCmd(version)
	rootCommand.AddCommand(versionCmd)

	swarmCmd := cmd_swarm.NewSwarmCmd(swarmService)
	rootCommand.AddCommand(swarmCmd)

	domainCmd := cmd_domain.NewDomainCmd(domainService)
	rootCommand.AddCommand(domainCmd)

	tenantCmd := cmd_tenant.NewTenantCmd(tenantService)
	rootCommand.AddCommand(tenantCmd)

	gatewayCmd := cmd_gateway.NewGatewayCmd(gatewayService)
	rootCommand.AddCommand(gatewayCmd)

	return rootCommand
}
