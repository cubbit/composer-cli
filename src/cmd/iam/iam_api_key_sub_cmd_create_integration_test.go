package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMAPIKeySubCmd_Create_Integration_UsesAPIKeyAuthentication(t *testing.T) {
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		CreateIAMAPIKeyFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			request *api.CreateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if request.Name != "automation" {
				t.Fatalf("Expected API key name %q, got %q", "automation", request.Name)
			}
			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation",
				Key:        "secret-key",
				OperatorID: operatorID,
				CreatedAt:  time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC),
				Enabled:    true,
			}, nil
		},
	}
	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "create", "--name", "automation"})

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

Secret:
  Key: secret-key
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key create output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
