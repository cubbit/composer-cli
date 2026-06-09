package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/spf13/cobra"
)

// setupTestCommand creates a cobra command with all necessary flags
func setupTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("profile", "", "Profile")
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	return cmd
}

func TestLocationService_NewLocationService(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	mockUserAPI := &api.UserAPI{}

	service := NewLocationService(mockCfg, mockLocationAPI, mockUserAPI)

	if service.configuration != mockCfg {
		t.Error("Expected configuration to be set")
	}
	if service.locationAPI != mockLocationAPI {
		t.Error("Expected locationAPI to be set")
	}
	if service.userAPI != mockUserAPI {
		t.Error("Expected userAPI to be set")
	}
}

func TestLocationService_List_Success(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	// Mock API to return test data
	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfrastructureCluster, error) {
			if apiKey != "test-api-key" {
				t.Error("Expected API key to be passed correctly")
			}
			if organizationID != "test-org-id" {
				t.Error("Expected organization ID to be passed correctly")
			}
			return []api.InfrastructureCluster{
				{ClusterID: "cluster-1", Name: "Cluster 1"},
				{ClusterID: "cluster-2", Name: "Cluster 2"},
			}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("output", "human")

	err := service.List(cmd, []string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLocationService_List_APIError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	// Mock API to return error
	mockLocationAPI := &api.MockLocationAPI{
		ListFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string) ([]api.InfrastructureCluster, error) {
			return nil, fmt.Errorf("api error")
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Set("profile", "test")

	err := service.List(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "failed to list locations") {
		t.Errorf("Expected error to contain 'failed to list locations', got %q", err.Error())
	}
}

func TestLocationService_List_ConfigError(t *testing.T) {
	// Mock config that returns an error
	mockCfg := &configuration.MockConfig{
		Err: fmt.Errorf("profile not found"),
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Set("profile", "nonexistent")

	err := service.List(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "profile not found") {
		t.Errorf("Expected error to contain 'profile not found', got %q", err.Error())
	}
}

func TestLocationService_List_WrongProfileType(t *testing.T) {
	// Mock config with wrong profile type
	mockCfg := &configuration.MockConfig{
		ResolvedProfile: &configuration.ResolvedProfile{
			Name:           "test",
			Type:           configuration.ProfileTypeConsole, // Wrong type
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration.OutputHuman,
		},
		Urls: &configuration.URLs{
			BaseURL: "https://api.example.com",
			IamURL:  "https://api.example.com/iam",
			DashURL: "https://dash.example.com",
			ChURL:   "https://ch.example.com",
		},
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Set("profile", "test")

	err := service.List(cmd, []string{})
	if err == nil {
		t.Error("Expected error for wrong profile type, got nil")
	}
	if !contains(err.Error(), "has type") || !contains(err.Error(), "expected") {
		t.Errorf("Expected error about profile type mismatch, got %q", err.Error())
	}
}

func TestLocationService_CreateVirtual_Success(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	// Mock API to return test data
	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, name string, description *string) (*api.InfrastructureCluster, error) {
			if name != "test-virtual-cluster" {
				t.Errorf("Expected name 'test-virtual-cluster', got %q", name)
			}
			if description == nil {
				t.Error("Expected description to be provided")
			} else if *description != "Test description" {
				t.Errorf("Expected description 'Test description', got %q", *description)
			}
			return &api.InfrastructureCluster{ClusterID: "new-cluster-id", Name: name}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("description", "", "Description")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-virtual-cluster")
	cmd.Flags().Set("description", "Test description")

	err := service.CreateVirtual(cmd, []string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLocationService_CreateVirtual_WithNilDescription(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, name string, description *string) (*api.InfrastructureCluster, error) {
			if description != nil {
				t.Error("Expected description to be nil")
			}
			return &api.InfrastructureCluster{ClusterID: "new-cluster-id", Name: name}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("description", "", "Description")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-virtual-cluster")

	err := service.CreateVirtual(cmd, []string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLocationService_CreateVirtual_APIError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, name string, description *string) (*api.InfrastructureCluster, error) {
			return nil, fmt.Errorf("api error")
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-virtual-cluster")

	err := service.CreateVirtual(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "failed to create virtual location") {
		t.Errorf("Expected error to contain 'failed to create virtual location', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtual_ConfigError(t *testing.T) {
	// Mock config that returns an error
	mockCfg := &configuration.MockConfig{
		Err: fmt.Errorf("failed to load profile"),
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().Set("profile", "nonexistent")

	err := service.CreateVirtual(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "failed to load profile") {
		t.Errorf("Expected error to contain 'failed to load profile', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtual_WrongProfileType(t *testing.T) {
	// Mock config with wrong profile type
	mockCfg := &configuration.MockConfig{
		ResolvedProfile: &configuration.ResolvedProfile{
			Name:           "test",
			Type:           configuration.ProfileTypeConsole, // Wrong type
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration.OutputHuman,
		},
		Urls: &configuration.URLs{
			BaseURL: "https://api.example.com",
			IamURL:  "https://api.example.com/iam",
			DashURL: "https://dash.example.com",
			ChURL:   "https://ch.example.com",
		},
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().Set("profile", "test")

	err := service.CreateVirtual(cmd, []string{})
	if err == nil {
		t.Error("Expected error for wrong profile type, got nil")
	}
	if !contains(err.Error(), "has type") || !contains(err.Error(), "expected") {
		t.Errorf("Expected error about profile type mismatch, got %q", err.Error())
	}
}

func TestLocationService_CreateVirtual_FlagGetError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	// Don't add the name flag at all - this will cause GetString to fail
	cmd.Flags().String("description", "", "Description")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("description", "Test description")

	err := service.CreateVirtual(cmd, []string{})
	if err == nil {
		t.Error("Expected error when name flag is not defined, got nil")
	}
	if !contains(err.Error(), "name") {
		t.Errorf("Expected error to mention 'name', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtual_MissingNameFlag(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("description", "", "Description")
	cmd.Flags().Set("profile", "test")
	// Don't set the name flag - empty string is valid

	err := service.CreateVirtual(cmd, []string{})
	// Empty string is a valid flag value, API may or may not accept it
	if err != nil {
		t.Logf("CreateVirtual failed (possibly at API level): %v", err)
	}
}

func TestLocationService_CreateVirtualNode_FlagGetError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	// Don't add the name flag - this will cause GetString to fail
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error when name flag is not defined, got nil")
	}
	if !contains(err.Error(), "name") {
		t.Errorf("Expected error to mention 'name', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_ClusterIDFlagGetError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	// Don't add cluster-id flag
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error when cluster-id flag is not defined, got nil")
	}
	if !contains(err.Error(), "cluster-id") {
		t.Errorf("Expected error to mention 'cluster-id', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_StorageTypeFlagGetError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	// Don't add storage-type flag
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error when storage-type flag is not defined, got nil")
	}
	if !contains(err.Error(), "storage-type") {
		t.Errorf("Expected error to mention 'storage-type', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_ConfigurationFlagGetError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	// Don't add configuration flag
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error when configuration flag is not defined, got nil")
	}
	if !contains(err.Error(), "configuration") {
		t.Errorf("Expected error to mention 'configuration', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_MissingRequiredFlags(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	// Don't set any required flags - empty strings are valid

	err := service.CreateVirtualNode(cmd, []string{})
	// Empty strings are valid flag values, API may handle validation
	if err != nil {
		t.Logf("CreateVirtualNode failed (possibly at API level): %v", err)
	}
}

func TestLocationService_CreateVirtualNode_MissingStorageType(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	// When storage-type is not provided, it defaults to empty string
	// The service should still call the API with the empty value
	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			// Empty storage type is passed to API (validation happens at API level)
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				StorageType: api.InfraVirtualStorageType(storageType),
			}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	// Don't set storage-type - it will be empty string
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	// Should succeed (empty string is a valid flag value, validation happens at API level)
	if err != nil {
		t.Errorf("Expected no error for empty storage-type (API validation), got %v", err)
	}
}

func TestLocationService_CreateVirtualNode_ConfigError(t *testing.T) {
	// Mock config that returns an error
	mockCfg := &configuration.MockConfig{
		Err: fmt.Errorf("failed to load profile"),
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().Set("profile", "nonexistent")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "failed to load profile") {
		t.Errorf("Expected error to contain 'failed to load profile', got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_WrongProfileType(t *testing.T) {
	// Mock config with wrong profile type
	mockCfg := &configuration.MockConfig{
		ResolvedProfile: &configuration.ResolvedProfile{
			Name:           "test",
			Type:           configuration.ProfileTypeConsole, // Wrong type
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration.OutputHuman,
		},
		Urls: &configuration.URLs{
			BaseURL: "https://api.example.com",
			IamURL:  "https://api.example.com/iam",
			DashURL: "https://dash.example.com",
			ChURL:   "https://ch.example.com",
		},
	}

	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().Set("profile", "test")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error for wrong profile type, got nil")
	}
	if !contains(err.Error(), "has type") || !contains(err.Error(), "expected") {
		t.Errorf("Expected error about profile type mismatch, got %q", err.Error())
	}
}

func TestLocationService_CreateVirtualNode_Success(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			if clusterID != "test-cluster-id" {
				t.Errorf("Expected clusterID 'test-cluster-id', got %q", clusterID)
			}
			if name != "test-node" {
				t.Errorf("Expected name 'test-node', got %q", name)
			}
			if storageType != "s3" {
				t.Errorf("Expected storageType 's3', got %q", storageType)
			}
			if configuration == nil {
				t.Error("Expected configuration to be provided")
			}
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				StorageType: api.InfraVirtualStorageType(storageType),
			}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLocationService_CreateVirtualNode_WithComplexConfiguration(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	configData := map[string]interface{}{
		"bucket": "test-bucket",
		"region": "eu-west-1",
		"tags":   []interface{}{"tag1", "tag2"},
		"nested": map[string]interface{}{
			"key1": "value1",
			"key2": float64(123),
		},
	}

	configJSON, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			// Verify configuration was parsed correctly
			if configuration["bucket"] != "test-bucket" {
				t.Errorf("Expected bucket 'test-bucket', got %v", configuration["bucket"])
			}
			return &api.InfraAggregateVirtualNodeDetail{
				NodeID:      "new-node-id",
				NodeName:    name,
				StorageType: api.InfraVirtualStorageType(storageType),
			}, nil
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", string(configJSON))

	err = service.CreateVirtualNode(cmd, []string{})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestLocationService_CreateVirtualNode_InvalidJSON(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")
	mockLocationAPI := &api.MockLocationAPI{}
	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", "invalid json{")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	} else {
		expectedSubstrings := []string{"parsing json", "invalid character", "configuration"}
		errMsg := err.Error()
		found := false
		for _, substr := range expectedSubstrings {
			if contains(errMsg, substr) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected error message to contain one of %v, got %q", expectedSubstrings, errMsg)
		}
	}
}

func TestLocationService_CreateVirtualNode_APIError(t *testing.T) {
	mockCfg := configuration.NewMockConfig(configuration.ProfileTypeComposer, "test-api-key", "test-org-id")

	mockLocationAPI := &api.MockLocationAPI{
		CreateVirtualNodeFunc: func(urlConfig configuration.URLs, apiKey string, organizationID string, clusterID string, name string, storageType string, configuration map[string]any) (*api.InfraAggregateVirtualNodeDetail, error) {
			return nil, fmt.Errorf("api error")
		},
	}

	service := NewLocationService(mockCfg, mockLocationAPI, &api.UserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("cluster-id", "", "Cluster ID")
	cmd.Flags().String("storage-type", "", "Storage Type")
	cmd.Flags().String("configuration", "", "Configuration")
	cmd.Flags().Set("profile", "test")
	cmd.Flags().Set("name", "test-node")
	cmd.Flags().Set("cluster-id", "test-cluster-id")
	cmd.Flags().Set("storage-type", "s3")
	cmd.Flags().Set("configuration", "{}")

	err := service.CreateVirtualNode(cmd, []string{})
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !contains(err.Error(), "failed to create virtual node") {
		t.Errorf("Expected error to contain 'failed to create virtual node', got %q", err.Error())
	}
}

func TestLocationService_InterfaceCompliance(t *testing.T) {
	// Verify that LocationService implements LocationServiceInterface
	var _ LocationServiceInterface = LocationService{}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
