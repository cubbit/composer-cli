package cmd_swarm

import (
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestSwarmSubCmd_Describe_Output_JSON(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com", Dash: "https://dash.example.com", CH: "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	status := api.EvaluatedStatusType("online")
	description := "Test swarm"

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-123", Name: "test-swarm"}},
				},
			}, nil
		},
		GetSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.SwarmV5Presentation, error) {
			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID: "swarm-123", Name: "test-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 512,
						CreatedAt: createdAt, NexusCount: 2, RedundancyClassCount: 1,
				},
				OrganizationID: "test-org-id", OwnerID: "owner-123",
					Description: &description, Configuration: map[string]interface{}{"tier": "hot"},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{EvaluatedStatus: &status},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"describe", "--swarm-name", "test-swarm", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "swarm-123",
  "name": "test-swarm",
  "total_storage_bytes": 1024,
  "used_storage_bytes": 512,
  "created_at": "2024-01-15T10:30:00Z",
  "nexus_count": 2,
  "redundancy_class_count": 1,
  "organization_id": "test-org-id",
  "owner_id": "owner-123",
  "description": "Test swarm",
  "configuration": {
    "tier": "hot"
  },
  "creation_status": "created",
  "evaluated_status": "online",
  "evaluated_status_last_updated_at": null
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Describe_Output_YAML(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com", Dash: "https://dash.example.com", CH: "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	status := api.EvaluatedStatusType("online")
	description := "Test swarm"

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-123", Name: "test-swarm"}},
				},
			}, nil
		},
		GetSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.SwarmV5Presentation, error) {
			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID: "swarm-123", Name: "test-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 512,
						CreatedAt: createdAt, NexusCount: 2, RedundancyClassCount: 1,
				},
				OrganizationID: "test-org-id", OwnerID: "owner-123",
					Description: &description, Configuration: map[string]interface{}{"tier": "hot"},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{EvaluatedStatus: &status},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"describe", "--swarm-name", "test-swarm", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: swarm-123
name: test-swarm
total_storage_bytes: 1024
used_storage_bytes: 512
created_at: 2024-01-15T10:30:00Z
nexus_count: 2
redundancy_class_count: 1
organization_id: test-org-id
owner_id: owner-123
description: Test swarm
configuration:
    tier: hot
creation_status: created
evaluated_status: online
evaluated_status_last_updated_at: null
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Describe_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com", Dash: "https://dash.example.com", CH: "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	status := api.EvaluatedStatusType("online")
	description := "Test swarm"

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-123", Name: "test-swarm"}},
				},
			}, nil
		},
		GetSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.SwarmV5Presentation, error) {
			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID: "swarm-123", Name: "test-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 512,
						CreatedAt: createdAt, NexusCount: 2, RedundancyClassCount: 1,
				},
				OrganizationID: "test-org-id", OwnerID: "owner-123",
					Description: &description, Configuration: map[string]interface{}{"tier": "hot"},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{EvaluatedStatus: &status},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"describe", "--swarm-name", "test-swarm", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "id": "swarm-123",
  "name": "test-swarm",
  "total_storage_bytes": 1024,
  "used_storage_bytes": 512,
  "created_at": "2024-01-15T10:30:00Z",
  "nexus_count": 2,
  "redundancy_class_count": 1,
  "organization_id": "test-org-id",
  "owner_id": "owner-123",
  "description": "Test swarm",
  "configuration": {
    "tier": "hot"
  },
  "creation_status": "created",
  "evaluated_status": "online",
  "evaluated_status_last_updated_at": null
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Describe_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com", Dash: "https://dash.example.com", CH: "https://ch.example.com",
			},
		}, nil
	}

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	status := api.EvaluatedStatusType("online")
	description := "Test swarm"

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{ListSwarmV5Item: api.ListSwarmV5Item{ID: "swarm-123", Name: "test-swarm"}},
				},
			}, nil
		},
		GetSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _, _ string) (*api.SwarmV5Presentation, error) {
			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID: "swarm-123", Name: "test-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 512,
						CreatedAt: createdAt, NexusCount: 2, RedundancyClassCount: 1,
				},
				OrganizationID: "test-org-id", OwnerID: "owner-123",
					Description: &description, Configuration: map[string]interface{}{"tier": "hot"},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{EvaluatedStatus: &status},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"describe", "--swarm-name", "test-swarm", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `id: swarm-123
name: test-swarm
total_storage_bytes: 1024
used_storage_bytes: 512
created_at: 2024-01-15T10:30:00Z
nexus_count: 2
redundancy_class_count: 1
organization_id: test-org-id
owner_id: owner-123
description: Test swarm
configuration:
    tier: hot
creation_status: created
evaluated_status: online
evaluated_status_last_updated_at: null
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

