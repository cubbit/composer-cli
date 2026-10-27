package cmd_swarm

import (
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestSwarmSubCmd_List_Output_JSON(t *testing.T) {
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

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-001", Name: "prod-swarm", TotalStorageBytes: 2048, UsedStorageBytes: 1024,
							CreatedAt: createdAt, NexusCount: 3, RedundancyClassCount: 2,
					},
				},
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-002", Name: "dev-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 256,
							CreatedAt: createdAt, NexusCount: 1, RedundancyClassCount: 1,
					},
				},
				},
				NextPage: nil, Count: 2,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.SetArgs([]string{"list", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "swarm-001",
    "name": "prod-swarm",
    "total_storage_bytes": 2048,
    "used_storage_bytes": 1024,
    "created_at": "2024-01-15T10:30:00Z",
    "nexus_count": 3,
    "redundancy_class_count": 2,
    "evaluated_status": null,
    "evaluated_status_last_updated_at": null
  },
  {
    "id": "swarm-002",
    "name": "dev-swarm",
    "total_storage_bytes": 1024,
    "used_storage_bytes": 256,
    "created_at": "2024-01-15T10:30:00Z",
    "nexus_count": 1,
    "redundancy_class_count": 1,
    "evaluated_status": null,
    "evaluated_status_last_updated_at": null
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_List_Output_YAML(t *testing.T) {
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

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-001", Name: "prod-swarm", TotalStorageBytes: 2048, UsedStorageBytes: 1024,
							CreatedAt: createdAt, NexusCount: 3, RedundancyClassCount: 2,
					},
				},
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-002", Name: "dev-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 256,
							CreatedAt: createdAt, NexusCount: 1, RedundancyClassCount: 1,
					},
				},
				},
				NextPage: nil, Count: 2,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.SetArgs([]string{"list", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: swarm-001
  name: prod-swarm
  total_storage_bytes: 2048
  used_storage_bytes: 1024
  created_at: 2024-01-15T10:30:00Z
  nexus_count: 3
  redundancy_class_count: 2
  evaluated_status: null
  evaluated_status_last_updated_at: null
- id: swarm-002
  name: dev-swarm
  total_storage_bytes: 1024
  used_storage_bytes: 256
  created_at: 2024-01-15T10:30:00Z
  nexus_count: 1
  redundancy_class_count: 1
  evaluated_status: null
  evaluated_status_last_updated_at: null
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_List_Output_JSON_FromProfile(t *testing.T) {
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

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-001", Name: "prod-swarm", TotalStorageBytes: 2048, UsedStorageBytes: 1024,
							CreatedAt: createdAt, NexusCount: 3, RedundancyClassCount: 2,
					},
				},
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-002", Name: "dev-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 256,
							CreatedAt: createdAt, NexusCount: 1, RedundancyClassCount: 1,
					},
				},
				},
				NextPage: nil, Count: 2,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.SetArgs([]string{"list", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `[
  {
    "id": "swarm-001",
    "name": "prod-swarm",
    "total_storage_bytes": 2048,
    "used_storage_bytes": 1024,
    "created_at": "2024-01-15T10:30:00Z",
    "nexus_count": 3,
    "redundancy_class_count": 2,
    "evaluated_status": null,
    "evaluated_status_last_updated_at": null
  },
  {
    "id": "swarm-002",
    "name": "dev-swarm",
    "total_storage_bytes": 1024,
    "used_storage_bytes": 256,
    "created_at": "2024-01-15T10:30:00Z",
    "nexus_count": 1,
    "redundancy_class_count": 1,
    "evaluated_status": null,
    "evaluated_status_last_updated_at": null
  }
]
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_List_Output_YAML_FromProfile(t *testing.T) {
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

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _, _ int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-001", Name: "prod-swarm", TotalStorageBytes: 2048, UsedStorageBytes: 1024,
							CreatedAt: createdAt, NexusCount: 3, RedundancyClassCount: 2,
					},
				},
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID: "swarm-002", Name: "dev-swarm", TotalStorageBytes: 1024, UsedStorageBytes: 256,
							CreatedAt: createdAt, NexusCount: 1, RedundancyClassCount: 1,
					},
				},
				},
				NextPage: nil, Count: 2,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.SetArgs([]string{"list", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `- id: swarm-001
  name: prod-swarm
  total_storage_bytes: 2048
  used_storage_bytes: 1024
  created_at: 2024-01-15T10:30:00Z
  nexus_count: 3
  redundancy_class_count: 2
  evaluated_status: null
  evaluated_status_last_updated_at: null
- id: swarm-002
  name: dev-swarm
  total_storage_bytes: 1024
  used_storage_bytes: 256
  created_at: 2024-01-15T10:30:00Z
  nexus_count: 1
  redundancy_class_count: 1
  evaluated_status: null
  evaluated_status_last_updated_at: null
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

