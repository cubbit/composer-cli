package configuration_handler

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func (h *ConfigurationHandler) GetConfigFilePath() (string, error) {
	if h.loadType != LoadTypeForConfigCommands {
		return "", configuration_models.ErrConfigurationIsNotLoaded
	}

	return h.configPath, nil
}

func (h *ConfigurationHandler) GetConfigFile() (configuration_models.ConfigV2, error) {
	if h.loadType != LoadTypeStandard {
		return configuration_models.ConfigV2{}, configuration_models.ErrConfigurationIsNotLoaded
	}

	return *h.configFile, nil
}

func (h *ConfigurationHandler) GetAPIKey() (string, error) {
	if h.loadType != LoadTypeStandard {
		return "", configuration_models.ErrConfigurationIsNotLoaded
	}

	return *h.apiKey, nil
}

func (h *ConfigurationHandler) GetEndpoints() (configuration_models.EndpointsV2, error) {
	if h.loadType != LoadTypeStandard && h.loadType != LoadTypeForAuthCommands {
		return configuration_models.EndpointsV2{}, configuration_models.ErrConfigurationIsNotLoaded
	}

	return *h.endpoints, nil
}

func (h *ConfigurationHandler) GetProfiles() (map[string]configuration_models.ProfileV2, error) {
	if h.loadType != LoadTypeStandard {
		return nil, configuration_models.ErrConfigurationIsNotLoaded
	}

	return h.configFile.Profile, nil
}

func (h *ConfigurationHandler) GetActiveProfile() (configuration_models.ProfileV2, error) {
	if h.loadType != LoadTypeStandard {
		return configuration_models.ProfileV2{}, configuration_models.ErrConfigurationIsNotLoaded
	}

	return h.configFile.GetActiveProfile(), nil
}

func (h *ConfigurationHandler) GetActiveProfileName() (string, error) {
	if h.loadType != LoadTypeStandard {
		return "", configuration_models.ErrConfigurationIsNotLoaded
	}

	return h.configFile.GetActiveProfileName(), nil
}
