package configuration_handler

import "github.com/cubbit/composer-cli/src/configuration/configuration_models"

func (h *ConfigurationHandler) DeleteProfile(profileName string) error {
	if h.loadType != LoadTypeStandard {
		return configuration_models.ErrConfigurationIsNotLoaded
	}

	if err := h.configFile.DeleteProfile(profileName); err != nil {
		return err
	}

	return h.Save()
}

func (h *ConfigurationHandler) DeleteAllProfiles() error {
	if h.loadType != LoadTypeStandard {
		return configuration_models.ErrConfigurationIsNotLoaded
	}

	h.configFile.DeleteAllProfiles()

	return h.Save()
}
