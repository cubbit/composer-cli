package cmd

import (
	"encoding/json"
	"os"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/service"
	service_gateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/spf13/cobra"
)

var devNull *os.File

type PackageData struct {
	Version string `json:"version"`
}

func setupSilentMode(cmd *cobra.Command) {
	silent, _ := cmd.Flags().GetBool("silent")
	if silent {
		f, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
		if err == nil {
			devNull = f
			os.Stdout = f
			os.Stderr = f
		}
	}
}

func cleanupSilentMode() {
	if devNull != nil {
		devNull.Close()
		devNull = nil
	}
}

func Execute(packageJSON []byte) {
	var pkg PackageData
	if err := json.Unmarshal(packageJSON, &pkg); err != nil {
		os.Exit(1)
	}

	configurationHandler := configuration_handler.NewConfigurationHandler()

	authAPI := api.NewAuthAPI()
	domainAPI := api.NewDomainAPI()
	operatorAPI := api.NewOperatorAPI()
	locationAPI := api.NewLocationAPI()
	userAPI := api.NewUserAPI()
	swarmAPI := api.NewSwarmAPI()
	processAPI := api.NewProcessAPI()
	gatewayAPI := api.NewGatewayAPI()
	redundancyClassAPI := api.NewRedundancyClassAPI()

	authService := service.NewAuthService(configurationHandler, authAPI, userAPI)
	configService := service.NewConfigService(configurationHandler)
	locationService := service.NewLocationService(configurationHandler, locationAPI, userAPI)
	operatorService := service.NewOperatorService(configurationHandler, operatorAPI, userAPI)
	redundancyClassValidator := service.NewRedundancyClassValidator()
	swarmService := service.NewSwarmService(configurationHandler, swarmAPI, locationAPI, processAPI, redundancyClassValidator)
	domainService := service.NewDomainService(configurationHandler, domainAPI, userAPI)
	gatewayService := service_gateway.NewGatewayService(configurationHandler, gatewayAPI, swarmAPI, redundancyClassAPI, processAPI, locationAPI)

	rootCmd := NewRootCommand(
		configurationHandler,
		authService,
		operatorService,
		locationService,
		configService,
		swarmService,
		domainService,
		gatewayService,
		pkg.Version,
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	cleanupSilentMode()
}
