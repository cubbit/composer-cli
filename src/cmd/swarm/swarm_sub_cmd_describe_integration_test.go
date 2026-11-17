package cmd_swarm

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func setupDescribeIntegrationCommand(swarmService service.SwarmServiceInterface) (*cobra.Command, *bytes.Buffer) {
	swarmCmd := NewSwarmCmd(swarmService)
	swarmCmd.PersistentFlags().String("output", "human", "Output format")
	swarmCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)

	return swarmCmd, commandOutput
}

func TestSwarmSubCmd_Describe_Integration_Success_WithSwarmName(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
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

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	description := "Test swarm"
	status := api.EvaluatedStatusType("online")

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if page != 1 || items != 1000 {
				t.Fatalf("Expected pagination defaults page=1 items=1000, got page=%d items=%d", page, items)
			}

			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID:   "swarm-123",
							Name: "test-swarm",
						},
					},
				},
			}, nil
		},
		GetSwarmV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			if swarmID != "swarm-123" {
				t.Fatalf("Expected resolved swarm ID swarm-123, got %q", swarmID)
			}

			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "swarm-123",
						Name:                 "test-swarm",
						TotalStorageBytes:    1024,
						UsedStorageBytes:     512,
						CreatedAt:            createdAt,
						NexusCount:           2,
						RedundancyClassCount: 1,
					},
					OrganizationID: "test-org-id",
					OwnerID:        "owner-123",
					Description:    &description,
					Configuration: map[string]interface{}{
						"tier": "hot",
					},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{
					EvaluatedStatus: &status,
				},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, commandOutput := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{
		"describe",
		"--swarm-name", "test-swarm",
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := commandOutput.String()
	expectedOutput := strings.TrimSpace(`
test-swarm
├── Status: ● Online
├── Last Update: N/A
├── Organization: test-org-id
├── Owner: owner-123
├── Storage Usage
│   ├── Usage: [█████░░░░░] 50%
│   ├── Total Used: 512 B
│   ├── Total Assigned: 1.0 KB
│   └── Total Unused: 512 B
├── Metadata
│   ├── ID: swarm-123
│   ├── Description: Test swarm
│   ├── Created At: 2024-01-15 10:30:00
│   └── Creation Status: created
├── Composition
│   ├── Nexus Count: 2
│   └── Redundancy Class Count: 1
└── Configuration
    └── tier: hot`)

	if strings.TrimSpace(output) != expectedOutput {
		t.Fatalf("Expected output:\n%s\nactual:\n%s", expectedOutput, strings.TrimSpace(output))
	}
}

func TestSwarmSubCmd_Describe_Integration_Error(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
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

	mockSwarmAPI := &api.MockSwarmAPI{
		GetSwarmV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			return nil, errors.New("network error")
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, commandOutput := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{
		"describe",
		"swarm-123",
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute, got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to describe swarm") {
		t.Fatalf("Expected output to contain describe error, got %q", output)
	}
}

func TestSwarmSubCmd_Describe_Integration_Success_WithPositionalID(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
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

	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	status := api.EvaluatedStatusType("online")

	mockSwarmAPI := &api.MockSwarmAPI{
		GetSwarmV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, swarmID string) (*api.SwarmV5Presentation, error) {
			if swarmID != "swarm-123" {
				t.Fatalf("Expected swarm ID swarm-123, got %q", swarmID)
			}

			return &api.SwarmV5Presentation{
				SwarmV5: api.SwarmV5{
					ListSwarmV5Item: api.ListSwarmV5Item{
						ID:                   "swarm-123",
						Name:                 "test-swarm",
						TotalStorageBytes:    2048,
						UsedStorageBytes:     1024,
						CreatedAt:            createdAt,
						NexusCount:           3,
						RedundancyClassCount: 2,
					},
					OrganizationID: "test-org-id",
					OwnerID:        "owner-123",
					Configuration: map[string]interface{}{
						"replication": 3,
					},
					CreationStatus: "created",
				},
				SummaryStatusNullable: api.SummaryStatusNullable{
					EvaluatedStatus: &status,
				},
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())
	swarmCmd, commandOutput := setupDescribeIntegrationCommand(swarmService)
	swarmCmd.SetArgs([]string{
		"describe",
		"swarm-123",
	})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := commandOutput.String()
	expectedOutput := strings.TrimSpace(`
test-swarm
├── Status: ● Online
├── Last Update: N/A
├── Organization: test-org-id
├── Owner: owner-123
├── Storage Usage
│   ├── Usage: [█████░░░░░] 50%
│   ├── Total Used: 1.0 KB
│   ├── Total Assigned: 2.0 KB
│   └── Total Unused: 1.0 KB
├── Metadata
│   ├── ID: swarm-123
│   ├── Description: N/A
│   ├── Created At: 2024-01-15 10:30:00
│   └── Creation Status: created
├── Composition
│   ├── Nexus Count: 3
│   └── Redundancy Class Count: 2
└── Configuration
    └── replication: 3`)

	if strings.TrimSpace(output) != expectedOutput {
		t.Fatalf("Expected output:\n%s\nactual:\n%s", expectedOutput, strings.TrimSpace(output))
	}
}
