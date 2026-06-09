package cmd

import (
	"encoding/json"
	"os"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
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

	configuration, err := configuration.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	authAPI := api.NewAuthAPI(configuration)
	operatorAPI := api.NewOperatorAPI(configuration)
	locationAPI := api.NewLocationAPI()
	userAPI := api.NewUserAPI()
	swarmAPI := api.NewSwarmAPI()
	processAPI := api.NewProcessAPI()

	agentService := service.NewAgentService(configuration)
	authService := service.NewAuthService(configuration, authAPI, userAPI)
	locationService := service.NewLocationService(configuration, locationAPI, userAPI)
	operatorService := service.NewOperatorService(configuration, operatorAPI, userAPI)
	configService := service.NewConfigService(configuration)
	redundancyClassValidator := service.NewRedundancyClassValidator()
	swarmService := service.NewSwarmService(configuration, swarmAPI, locationAPI, processAPI, redundancyClassValidator)

	rootCmd := NewRootCommand(agentService, authService, operatorService, locationService, configService, swarmService, pkg.Version)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	cleanupSilentMode()
}
