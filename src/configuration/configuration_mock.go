package configuration

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MockConfig implements ConfigInterface for testing
type MockConfig struct {
	ResolvedProfile *ResolvedProfile
	Urls            *URLs
	Err             error
}

func NewMockConfig(profile ProfileType, apiKey, orgID string) *MockConfig {
	return &MockConfig{
		ResolvedProfile: &ResolvedProfile{
			Name:           "test",
			Type:           profile,
			APIKey:         apiKey,
			OrganizationID: orgID,
			Output:         OutputHuman,
		},
		Urls: &URLs{
			BaseURL: "https://api.example.com",
			IamURL:  "https://api.example.com/iam",
			DashURL: "https://dash.example.com",
			ChURL:   "https://ch.example.com",
		},
	}
}

func (m *MockConfig) ResolveProfileAndURLs(cmd *cobra.Command, expectedProfileType ProfileType) (*ResolvedProfile, *URLs, error) {
	if m.Err != nil {
		return nil, nil, m.Err
	}
	if m.ResolvedProfile.Type != expectedProfileType {
		return nil, nil, fmt.Errorf("profile '%s' has type '%s', expected '%s'", m.ResolvedProfile.Name, m.ResolvedProfile.Type, expectedProfileType)
	}
	return m.ResolvedProfile, m.Urls, nil
}
