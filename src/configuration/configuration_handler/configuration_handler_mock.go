package configuration_handler

import (
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

type MockConfigurationHandler struct {
	LoadFunc                 func(cmd *cobra.Command, args []string) error
	SaveFunc                 func() error
	CreateProfileFunc        func(profileName string, apiKey string, organizationID string) error
	SwitchProfileFunc        func(profileName string) error
	DeleteProfileFunc        func(profileName string) error
	DeleteAllProfilesFunc    func() error
	GetConfigFileFunc        func() (configuration_models.ConfigV2, error)
	GetConfigFilePathFunc    func() (string, error)
	GetProfilesFunc          func() (map[string]configuration_models.ProfileV2, error)
	GetActiveProfileFunc     func() (configuration_models.ProfileV2, error)
	GetActiveProfileNameFunc func() (string, error)
	GetEndpointsFunc         func() (configuration_models.EndpointsV2, error)
	GetAPIKeyFunc            func() (string, error)
}

func NewMockConfigurationHandler() *MockConfigurationHandler {
	return &MockConfigurationHandler{
		LoadFunc: func(cmd *cobra.Command, args []string) error {
			return nil
		},
		SaveFunc: func() error {
			return nil
		},
		CreateProfileFunc: func(profileName string, apiKey string, organizationID string) error {
			return nil
		},
		SwitchProfileFunc: func(profileName string) error {
			return nil
		},
		DeleteProfileFunc: func(profileName string) error {
			return nil
		},
		DeleteAllProfilesFunc: func() error {
			return nil
		},
		GetConfigFileFunc: func() (configuration_models.ConfigV2, error) {
			return configuration_models.ConfigV2{}, nil
		},
		GetConfigFilePathFunc: func() (string, error) {
			return "", nil
		},
		GetProfilesFunc: func() (map[string]configuration_models.ProfileV2, error) {
			return nil, nil
		},
		GetActiveProfileFunc: func() (configuration_models.ProfileV2, error) {
			return configuration_models.ProfileV2{}, nil
		},
		GetActiveProfileNameFunc: func() (string, error) {
			return "", nil
		},
		GetEndpointsFunc: func() (configuration_models.EndpointsV2, error) {
			return configuration_models.EndpointsV2{}, nil
		},
		GetAPIKeyFunc: func() (string, error) {
			return "", nil
		},
	}
}

func (m *MockConfigurationHandler) Load(cmd *cobra.Command, args []string) error {
	if m.LoadFunc != nil {
		return m.LoadFunc(cmd, args)
	}
	return nil
}

func (m *MockConfigurationHandler) Save() error {
	if m.SaveFunc != nil {
		return m.SaveFunc()
	}
	return nil
}

func (m *MockConfigurationHandler) CreateProfile(profileName string, apiKey string, organizationID string) error {
	if m.CreateProfileFunc != nil {
		return m.CreateProfileFunc(profileName, apiKey, organizationID)
	}
	return nil
}

func (m *MockConfigurationHandler) SwitchProfile(profileName string) error {
	if m.SwitchProfileFunc != nil {
		return m.SwitchProfileFunc(profileName)
	}
	return nil
}

func (m *MockConfigurationHandler) DeleteProfile(profileName string) error {
	if m.DeleteProfileFunc != nil {
		return m.DeleteProfileFunc(profileName)
	}
	return nil
}

func (m *MockConfigurationHandler) DeleteAllProfiles() error {
	if m.DeleteAllProfilesFunc != nil {
		return m.DeleteAllProfilesFunc()
	}
	return nil
}

func (m *MockConfigurationHandler) GetConfigFile() (configuration_models.ConfigV2, error) {
	if m.GetConfigFileFunc != nil {
		return m.GetConfigFileFunc()
	}
	return configuration_models.ConfigV2{}, nil
}

func (m *MockConfigurationHandler) GetConfigFilePath() (string, error) {
	if m.GetConfigFilePathFunc != nil {
		return m.GetConfigFilePathFunc()
	}
	return "", nil
}

func (m *MockConfigurationHandler) GetProfiles() (map[string]configuration_models.ProfileV2, error) {
	if m.GetProfilesFunc != nil {
		return m.GetProfilesFunc()
	}
	return nil, nil
}

func (m *MockConfigurationHandler) GetActiveProfile() (configuration_models.ProfileV2, error) {
	if m.GetActiveProfileFunc != nil {
		return m.GetActiveProfileFunc()
	}
	return configuration_models.ProfileV2{}, nil
}

func (m *MockConfigurationHandler) GetActiveProfileName() (string, error) {
	if m.GetActiveProfileNameFunc != nil {
		return m.GetActiveProfileNameFunc()
	}
	return "", nil
}

func (m *MockConfigurationHandler) GetEndpoints() (configuration_models.EndpointsV2, error) {
	if m.GetEndpointsFunc != nil {
		return m.GetEndpointsFunc()
	}
	return configuration_models.EndpointsV2{}, nil
}

func (m *MockConfigurationHandler) GetAPIKey() (string, error) {
	if m.GetAPIKeyFunc != nil {
		return m.GetAPIKeyFunc()
	}
	return "", nil
}
