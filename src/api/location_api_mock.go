package api

import (
	"fmt"

	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/spf13/cobra"
)

// MockLocationAPI implements LocationAPIInterface for testing
type MockLocationAPI struct {
	ListFunc              func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]InfrastructureCluster, error)
	ListAggregatedFunc    func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]InfraAggregateCluster, error)
	CreateVirtualFunc     func(urlConfig configuration.URLs, apiKey string, organizationID string, name string, description *string) (*InfrastructureCluster, error)
	CreateVirtualNodeFunc func(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*InfraAggregateVirtualNodeDetail, error)
}

func (m *MockLocationAPI) List(urlConfig configuration.URLs, apiKey string, organizationID string) ([]InfrastructureCluster, error) {
	if m.ListFunc != nil {
		return m.ListFunc(urlConfig, apiKey, organizationID)
	}
	return []InfrastructureCluster{}, nil
}

func (m *MockLocationAPI) ListAggregated(urlConfig configuration.URLs, apiKey string, organizationID string) ([]InfraAggregateCluster, error) {
	if m.ListAggregatedFunc != nil {
		return m.ListAggregatedFunc(urlConfig, apiKey, organizationID)
	}
	return []InfraAggregateCluster{}, nil
}

func (m *MockLocationAPI) CreateVirtualCluster(urlConfig configuration.URLs, apiKey string, organizationID string, name string, description *string) (*InfrastructureCluster, error) {
	if m.CreateVirtualFunc != nil {
		return m.CreateVirtualFunc(urlConfig, apiKey, organizationID, name, description)
	}
	return &InfrastructureCluster{ClusterID: "test-cluster-id", Name: name}, nil
}

func (m *MockLocationAPI) CreateVirtualNode(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*InfraAggregateVirtualNodeDetail, error) {
	if m.CreateVirtualNodeFunc != nil {
		return m.CreateVirtualNodeFunc(urlConfig, apiKey, organizationID, clusterID, name, storageType, configuration)
	}
	return &InfraAggregateVirtualNodeDetail{
		NodeID:      "test-node-id",
		NodeName:    name,
		StorageType: InfraVirtualStorageType(storageType),
	}, nil
}

// mockConfig implements ConfigInterface for testing
type mockConfig struct {
	resolvedProfile *configuration.ResolvedProfile
	urls            *configuration.URLs
	err             error
}

func NewMockConfig(profile configuration.ProfileType, apiKey, orgID string) *mockConfig {
	return &mockConfig{
		resolvedProfile: &configuration.ResolvedProfile{
			Name:           "test",
			Type:           profile,
			APIKey:         apiKey,
			OrganizationID: orgID,
			Output:         configuration.OutputHuman,
		},
		urls: &configuration.URLs{
			BaseURL: "https://api.example.com",
			IamURL:  "https://api.example.com/iam",
			DashURL: "https://dash.example.com",
			ChURL:   "https://ch.example.com",
		},
	}
}

func (m *mockConfig) ResolveProfileAndURLs(cmd *cobra.Command, expectedProfileType configuration.ProfileType) (*configuration.ResolvedProfile, *configuration.URLs, error) {
	if m.err != nil {
		return nil, nil, m.err
	}
	if m.resolvedProfile.Type != expectedProfileType {
		return nil, nil, fmt.Errorf("profile '%s' has type '%s', expected '%s'", m.resolvedProfile.Name, m.resolvedProfile.Type, expectedProfileType)
	}
	return m.resolvedProfile, m.urls, nil
}
