package cmd

import (
	"encoding/json"
	"os"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

var interactive bool
var packageJSON []byte
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

var rootCmd = func() *cobra.Command {
	configuration, err := configuration.LoadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	authAPI := api.NewAuthAPI(configuration)

	agentService := service.NewAgentService(configuration)
	authService := service.NewAuthService(configuration, authAPI)

	return NewRootCommand(agentService, authService)
}()

func Execute(packageJSON []byte) {
	var pkg PackageData
	if err := json.Unmarshal(packageJSON, &pkg); err != nil {
		os.Exit(1)
	}
	rootCmd.Version = pkg.Version
	rootCmd.SetVersionTemplate("{{.Use}} version {{.Version}}\n")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	cleanupSilentMode()
}
