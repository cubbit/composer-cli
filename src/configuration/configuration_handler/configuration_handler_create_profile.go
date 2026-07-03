package configuration_handler

import (
	"time"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func (h *ConfigurationHandler) CreateProfile(
	profileName string,
	apiKey string,
	organizationID string,
) error {
	if h.loadType == LoadTypeNotLoaded || h.loadType == LoadTypeForConfigCommands {
		return configuration_models.ErrConfigurationIsNotLoaded
	}

	if h.configFile == nil {
		newConfig := configuration_models.CreateConfigV2(
			profileName,
			configuration_models.OutputHuman,
			apiKey,
			organizationID,
			*h.endpoints,
			time.Now(),
		)

		h.configFile = &newConfig
		h.requestedProfile = &profileName
		h.apiKey = &apiKey
	} else {
		h.configFile.CreateProfile(
			profileName,
			configuration_models.OutputHuman,
			apiKey,
			organizationID,
			*h.endpoints,
			time.Now(),
		)
	}

	return h.Save()
}
