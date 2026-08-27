package configuration_handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

// getConfigurationPath retrieves the configuration path
// It first checks the --config-path flag
// then the XDG_CONFIG_HOME environment variable
// and finally defaults to ~/.config/cubbit/config.toml
func (h *ConfigurationHandler) getConfigurationPath(cmd *cobra.Command, args []string) (string, error) {
	configPathFromFlag, err := cmd.Flags().GetString("config-path")
	if err != nil {
		return "", fmt.Errorf("failed to get 'config-path' flag: %w", err)
	}

	if configPathFromFlag != "" {
		return configPathFromFlag, nil
	}

	configPathFromEnv := os.Getenv(configuration_models.ConfigurationPathEnvVariable)
	if configPathFromEnv != "" {
		return filepath.Join(configPathFromEnv, filepath.Base(configuration_models.ConfigurationDefaultDirName), configuration_models.ConfigurationFileName), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, configuration_models.ConfigurationDefaultDirName, configuration_models.ConfigurationFileName), nil
}
