package cmd_location

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestLocationSubCmd_List_Integration_Success(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
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
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"list",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────────────────────────────┬────────────────────┬──────────╮
│ Cluster ID                           │ Name               │ Type     │
├──────────────────────────────────────┼────────────────────┼──────────┤
│ 550e8400-e29b-41d4-a716-446655440000 │ production-cluster │ physical │
│ 6ba7b810-9dad-11d1-80b4-00c04fd430c8 │ staging-cluster    │ virtual  │
╰──────────────────────────────────────┴────────────────────┴──────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected clusters output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_List_Integration_Empty(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return []api.InfrastructureCluster{}, nil
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
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"list",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭────────────┬──────┬──────╮
│ Cluster ID │ Name │ Type │
├────────────┼──────┼──────┤
╰────────────┴──────┴──────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Error("Expected empty clusters output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestLocationSubCmd_List_Integration_Error(t *testing.T) {
	mockAPI := &api.MockLocationAPI{
		ListFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, opts ...api.LocationListOption) ([]api.InfrastructureCluster, error) {
			return nil, errors.New("failed to list locations")
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
	locationCmd.PersistentFlags().Bool("quiet", false, "quiet mode")

	commandOutput := new(bytes.Buffer)
	locationCmd.SetOut(commandOutput)
	locationCmd.SetErr(commandOutput)
	locationCmd.SetArgs([]string{
		"list",
	})
	err := locationCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error from Execute (Run doesn't return errors), got %v", err)
	}

	output := commandOutput.String()
	expectedError := "failed to list locations"
	if !strings.Contains(output, expectedError) {
		t.Errorf("Expected output to contain error message %q, got %q", expectedError, output)
	}
}
