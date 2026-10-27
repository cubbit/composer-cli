package cmd_swarm

import (
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestSwarmSubCmd_Create_Output_Inline_JSON(t *testing.T) {
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

	mockProcessAPI := &api.MockProcessAPI{
		ListProcessesFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.ListProcessesOption) ([]api.Process, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "cluster", Name: "cluster",
					Details: api.InfraAggregateClusterDetail{
						Nodes: []api.InfraAggregateNodeDetail{
							{
								NodeID: "node", NodeName: "node",
								Disks: []api.InfraAggregateDiskDetail{
									{
										DiskUUID: "disk-1", Path: "/dev/sda",
										TotalStorageSizeBytes: 1073741824,
										Status:                api.InfraAggregateStatus{Code: string(api.StatusCodeOk)},
								},
							},
						},
					},
				},
				},
			}, nil
		},
	}

	mockSwarmAPI := &api.MockSwarmAPI{
		CreateSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _ *api.CreateSwarmV5Request) (*api.CreateSwarmV5Response, error) {
			return &api.CreateSwarmV5Response{ID: "process-456"}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, mockLocationAPI, mockProcessAPI, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"create", "--name", "test-swarm", "--nexus", "cluster:node", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Swarm creation started with Process ID: process-456\n"
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Create_Output_Inline_YAML(t *testing.T) {
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

	mockProcessAPI := &api.MockProcessAPI{
		ListProcessesFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.ListProcessesOption) ([]api.Process, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "cluster", Name: "cluster",
					Details: api.InfraAggregateClusterDetail{
						Nodes: []api.InfraAggregateNodeDetail{
							{
								NodeID: "node", NodeName: "node",
								Disks: []api.InfraAggregateDiskDetail{
									{
										DiskUUID: "disk-1", Path: "/dev/sda",
										TotalStorageSizeBytes: 1073741824,
										Status:                api.InfraAggregateStatus{Code: string(api.StatusCodeOk)},
								},
							},
						},
					},
				},
				},
			}, nil
		},
	}

	mockSwarmAPI := &api.MockSwarmAPI{
		CreateSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _ *api.CreateSwarmV5Request) (*api.CreateSwarmV5Response, error) {
			return &api.CreateSwarmV5Response{ID: "process-456"}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, mockLocationAPI, mockProcessAPI, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"create", "--name", "test-swarm", "--nexus", "cluster:node", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Swarm creation started with Process ID: process-456
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Create_InteractiveAndOutputConflict(t *testing.T) {
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

	swarmService := service.NewSwarmService(mockCfg, &api.MockSwarmAPI{}, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, _ := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{
		"create",
		"--interactive",
		"--output", "json",
	})

	err := swarmCmd.Execute()
	if err == nil {
		t.Fatal("Expected error for conflicting --interactive and --output flags, got nil")
	}

	if !strings.Contains(err.Error(), "mutually exclusive") {
		t.Errorf("Expected error message to mention flags are mutually exclusive, got %q", err.Error())
	}
}

func TestSwarmSubCmd_Create_Output_Inline_JSON_FromProfile(t *testing.T) {
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

	mockProcessAPI := &api.MockProcessAPI{
		ListProcessesFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.ListProcessesOption) ([]api.Process, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "cluster", Name: "cluster",
					Details: api.InfraAggregateClusterDetail{
						Nodes: []api.InfraAggregateNodeDetail{
							{
								NodeID: "node", NodeName: "node",
								Disks: []api.InfraAggregateDiskDetail{
									{
										DiskUUID: "disk-1", Path: "/dev/sda",
										TotalStorageSizeBytes: 1073741824,
										Status:                api.InfraAggregateStatus{Code: string(api.StatusCodeOk)},
								},
							},
						},
					},
				},
				},
			}, nil
		},
	}

	mockSwarmAPI := &api.MockSwarmAPI{
		CreateSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _ *api.CreateSwarmV5Request) (*api.CreateSwarmV5Response, error) {
			return &api.CreateSwarmV5Response{ID: "process-456"}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, mockLocationAPI, mockProcessAPI, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"create", "--name", "test-swarm", "--nexus", "cluster:node", "--output", "json"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "message": "Swarm creation started with Process ID: process-456\n"
}
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestSwarmSubCmd_Create_Output_Inline_YAML_FromProfile(t *testing.T) {
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

	mockProcessAPI := &api.MockProcessAPI{
		ListProcessesFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.ListProcessesOption) ([]api.Process, error) {
			return nil, nil
		},
	}

	mockLocationAPI := &api.MockLocationAPI{
		ListAggregatedFunc: func(_ configuration_models.EndpointsV2, _, _ string, _ ...api.LocationListOption) ([]api.InfraAggregateCluster, error) {
			return []api.InfraAggregateCluster{
				{
					ClusterID: "cluster", Name: "cluster",
					Details: api.InfraAggregateClusterDetail{
						Nodes: []api.InfraAggregateNodeDetail{
							{
								NodeID: "node", NodeName: "node",
								Disks: []api.InfraAggregateDiskDetail{
									{
										DiskUUID: "disk-1", Path: "/dev/sda",
										TotalStorageSizeBytes: 1073741824,
										Status:                api.InfraAggregateStatus{Code: string(api.StatusCodeOk)},
								},
							},
						},
					},
				},
				},
			}, nil
		},
	}

	mockSwarmAPI := &api.MockSwarmAPI{
		CreateSwarmV5Func: func(_ configuration_models.EndpointsV2, _, _ string, _ *api.CreateSwarmV5Request) (*api.CreateSwarmV5Response, error) {
			return &api.CreateSwarmV5Response{ID: "process-456"}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, mockLocationAPI, mockProcessAPI, service.NewRedundancyClassValidator())
	swarmCmd, buf := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{"create", "--name", "test-swarm", "--nexus", "cluster:node", "--output", "yaml"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `message: |
    Swarm creation started with Process ID: process-456
`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}
