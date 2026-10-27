package configuration_handler

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func (h *ConfigurationHandler) Save() error {
	if h.loadType == LoadTypeNotLoaded {
		return configuration_models.ErrConfigurationIsNotLoaded
	}

	if h.configPath == "" {
		return fmt.Errorf("%w: configuration path is not set", configuration_models.ErrConfigurationIsNotLoaded)
	}

	if h.configFile != nil {
		if err := h.configFile.Save(h.configPath); err != nil {
			return fmt.Errorf("failed to save configuration to path %q: %w", h.configPath, err)
		}

		return nil
	}

	return fmt.Errorf("no target configuration entity found to be saved")
}
