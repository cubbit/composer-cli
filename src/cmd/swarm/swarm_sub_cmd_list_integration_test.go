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
)

func TestSwarmSubCmd_List_Integration_Success(t *testing.T) {
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
	firstSync := createdAt.Add(30 * time.Minute)
	secondSync := createdAt.Add(90 * time.Minute)

	mockSwarmAPI := &api.MockSwarmAPI{
		ListSwarmsV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data: []api.ListSwarmV5ItemPresentation{
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID:                   "swarm-001",
							Name:                 "prod-swarm",
							TotalStorageBytes:    1099511627776,
							UsedStorageBytes:     549755813888,
							CreatedAt:            createdAt,
							NexusCount:           3,
							RedundancyClassCount: 2,
						},
						SummaryStatusNullable: api.SummaryStatusNullable{
							EvaluatedStatus:              &status,
							EvaluatedStatusLastUpdatedAt: &firstSync,
						},
					},
					{
						ListSwarmV5Item: api.ListSwarmV5Item{
							ID:                   "swarm-002",
							Name:                 "dev-swarm",
							TotalStorageBytes:    107374182400,
							UsedStorageBytes:     0,
							CreatedAt:            createdAt,
							NexusCount:           1,
							RedundancyClassCount: 1,
						},
						SummaryStatusNullable: api.SummaryStatusNullable{
							EvaluatedStatusLastUpdatedAt: &secondSync,
						},
					},
				},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())

	swarmCmd := NewSwarmCmd(swarmService)
	swarmCmd.PersistentFlags().String("profile", "", "Profile")
	swarmCmd.PersistentFlags().String("output", "human", "Output format")
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"list",
	})
	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭────────────┬─────────────────┬─────────┬───────────┬─────────────────────┬─────────────────────┬────────╮
│ Name       │ Used Capacity   │ % Usage │ Locations │ Created On          │ Last Sync           │ Status │
├────────────┼─────────────────┼─────────┼───────────┼─────────────────────┼─────────────────────┼────────┤
│ prod-swarm │ 512.0 GB/1.0 TB │ 50%     │ 3         │ 2024-01-15 10:30:00 │ 2024-01-15 11:00:00 │ Online │
│ dev-swarm  │ 0/100.0 GB      │ 0%      │ 1         │ 2024-01-15 10:30:00 │ 2024-01-15 12:00:00 │ N/A    │
╰────────────┴─────────────────┴─────────┴───────────┴─────────────────────┴─────────────────────┴────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected swarm list output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestSwarmSubCmd_List_Integration_Empty(t *testing.T) {
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
		ListSwarmsV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return &api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation]{
				Data:     []api.ListSwarmV5ItemPresentation{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())

	swarmCmd := NewSwarmCmd(swarmService)
	swarmCmd.PersistentFlags().String("profile", "", "Profile")
	swarmCmd.PersistentFlags().String("output", "human", "Output format")
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"list",
	})
	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No swarms found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Errorf("Expected output %q, got %q", expectedResult, actualResult)
	}
}

func TestSwarmSubCmd_List_Integration_Error(t *testing.T) {
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
		ListSwarmsV5Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, page int, items int) (*api.GenericPaginatedResponse[api.ListSwarmV5ItemPresentation], error) {
			return nil, errors.New("network error")
		},
	}

	swarmService := service.NewSwarmService(mockCfg, mockSwarmAPI, &api.MockLocationAPI{}, &api.ProcessAPI{}, service.NewRedundancyClassValidator())

	swarmCmd := NewSwarmCmd(swarmService)
	swarmCmd.PersistentFlags().String("profile", "", "Profile")
	swarmCmd.PersistentFlags().String("output", "human", "Output format")
	swarmCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	swarmCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{
		"list",
	})
	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	if !strings.Contains(output, "failed to list swarms") {
		t.Errorf("Expected output to contain error message, got %q", output)
	}
}
