package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMAPIKeySubCmd_Describe_Integration_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		GetIAMAPIKeyByIDFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			apiKeyID string,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID to be propagated, got %q", operatorID)
			}
			if apiKeyID != "key-001" {
				t.Fatalf("Expected API key ID to be propagated, got %q", apiKeyID)
			}

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation",
				OperatorID: "operator-001",
				CreatedAt:  createdAt,
				Enabled:    true,
			}, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			t.Fatal("Describe should fetch the API key by ID instead of listing API keys")
			return nil, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "describe", "--id", "key-001"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
API Key: automation
Enabled: true

Metadata:
  ID: key-001
  Operator ID: operator-001
  Created At: 2024-01-15 10:30:00
  Expires At: N/A
  Deleted At: N/A
  Banned At: N/A
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key describe output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
