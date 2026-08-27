package cmd_location

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_List_Output_JSON(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      "physical",
				},
				{
					ClusterID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
					Name:      "staging-cluster",
					Type:      "virtual",
				},
			}, nil
		},
	}

	mockConfig := configuration_handler.NewMockConfigurationHandler()
	mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"list",
		"--output", "json",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "cluster_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "production-cluster",
    "type": "physical"
  },
  {
    "cluster_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
    "name": "staging-cluster",
    "type": "virtual"
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_List_Output_YAML(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      "physical",
				},
				{
					ClusterID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
					Name:      "staging-cluster",
					Type:      "virtual",
				},
			}, nil
		},
	}

	mockConfig := configuration_handler.NewMockConfigurationHandler()
	mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"list",
		"--output", "yaml",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- cluster_id: 550e8400-e29b-41d4-a716-446655440000
  name: production-cluster
  type: physical
- cluster_id: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
  name: staging-cluster
  type: virtual
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_List_Output_JSON_FromProfile(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      "physical",
				},
				{
					ClusterID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
					Name:      "staging-cluster",
					Type:      "virtual",
				},
			}, nil
		},
	}

	mockConfig := configuration_handler.NewMockConfigurationHandler()
	mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"list",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "cluster_id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "production-cluster",
    "type": "physical"
  },
  {
    "cluster_id": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
    "name": "staging-cluster",
    "type": "virtual"
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestLocationSubCmd_List_Output_YAML_FromProfile(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{
				{
					ClusterID: "550e8400-e29b-41d4-a716-446655440000",
					Name:      "production-cluster",
					Type:      "physical",
				},
				{
					ClusterID: "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
					Name:      "staging-cluster",
					Type:      "virtual",
				},
			}, nil
		},
	}

	mockConfig := configuration_handler.NewMockConfigurationHandler()
	mockConfig.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	locationService := service.NewLocationService(mockConfig, mockAPI, nil)

	locationCmd := NewLocationCmd(&locationService)
	locationCmd.PersistentFlags().String("output", "human", "Output format")
	locationCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	locationCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	locationCmd.SetOut(buf)
	locationCmd.SetErr(buf)
	locationCmd.SetArgs([]string{
		"list",
	})

	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- cluster_id: 550e8400-e29b-41d4-a716-446655440000
  name: production-cluster
  type: physical
- cluster_id: 6ba7b810-9dad-11d1-80b4-00c04fd430c8
  name: staging-cluster
  type: virtual
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

