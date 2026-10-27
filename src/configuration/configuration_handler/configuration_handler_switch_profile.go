package configuration_handler

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func (h *ConfigurationHandler) SwitchProfile(profileName string) error {
	if h.loadType != LoadTypeStandard {
		return configuration_models.ErrConfigurationIsNotLoaded
	}

	if _, ok := h.configFile.Profile[profileName]; !ok {
		return fmt.Errorf("%w: profile %q does not exist", configuration_models.ErrProfileNotFound, profileName)
	}

	h.configFile.Active.Profile = profileName

	return h.Save()
}
