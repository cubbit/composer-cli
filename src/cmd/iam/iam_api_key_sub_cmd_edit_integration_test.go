package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMAPIKeySubCmd_Edit_Integration_UsesAPIKeyAuthentication(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	expiresAt := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		UpdateIAMAPIKeyFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			apiKeyID string,
			request *api.UpdateIAMAPIKeyRequestBody,
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
			if request.Name == nil || *request.Name != "automation-renamed" {
				t.Fatalf("Expected API key name %q, got %#v", "automation-renamed", request.Name)
			}

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation-renamed",
				OperatorID: operatorID,
				CreatedAt:  createdAt,
				ExpiresAt:  &expiresAt,
				Enabled:    true,
			}, nil
		},
	}
	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "edit", "--id", "key-001", "--name", "automation-renamed"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
API Key: automation-renamed
Enabled: true

Metadata:
  ID: key-001
  Operator ID: operator-001
  Created At: 2024-01-15 10:30:00
  Expires At: 2024-12-31 23:59:59
  Deleted At: N/A
  Banned At: N/A
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key edit output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestIAMAPIKeySubCmd_Edit_Integration_WithUsername(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			t.Fatal("Edit with --username should not resolve the current IAM user")
			return nil, nil
		},
		ListIAMUsersFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			enabled *bool,
			search string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.IAMUserListItem], error) {
			if search != "alice" {
				t.Fatalf("Expected username search %q, got %q", "alice", search)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{ID: "operator-002", Username: "alice"},
				},
			}, nil
		},
		UpdateIAMAPIKeyFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			apiKeyID string,
			request *api.UpdateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if operatorID != "operator-002" {
				t.Fatalf("Expected target operator ID %q, got %q", "operator-002", operatorID)
			}
			if apiKeyID != "key-002" {
				t.Fatalf("Expected API key ID %q, got %q", "key-002", apiKeyID)
			}
			if request.Enabled == nil || *request.Enabled {
				t.Fatalf("Expected enabled false, got %#v", request.Enabled)
			}

			return &api.OperatorAPIKey{
				ID:         apiKeyID,
				Name:       "automation",
				OperatorID: operatorID,
				CreatedAt:  createdAt,
				Enabled:    false,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "edit", "--id", "key-002", "--username", "alice", "--enabled=false", "--quiet"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "key-002\tautomation\tfalse\n"
	if commandOutput.String() != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, commandOutput.String())
	}
}
