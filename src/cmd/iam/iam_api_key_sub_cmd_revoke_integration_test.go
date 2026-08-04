package cmd_iam

import (
	"fmt"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMAPIKeySubCmd_Revoke_Integration_UsesAPIKeyAuthentication(t *testing.T) {
	organizationName := "test-org"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		DeleteIAMAPIKeyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) error {
			if apiKeyID != "key-001" {
				t.Fatalf("Expected API key ID %q, got %q", "key-001", apiKeyID)
			}
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "revoke", "--id", "key-001"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := fmt.Sprintf("IAM API key %s revoked successfully\n", "key-001")
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_Revoke_Integration_WithUserID(t *testing.T) {
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			t.Fatal("Revoke with --user-id should not resolve the current IAM user")
			return nil, nil
		},
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			if userID != "operator-002" {
				t.Fatalf("Expected user ID %q, got %q", "operator-002", userID)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID %q, got %q", "test-org-id", organizationID)
			}

			return &api.IAMUser{ID: userID}, nil
		},
		DeleteIAMAPIKeyFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, operatorID string, apiKeyID string) error {
			if operatorID != "operator-002" {
				t.Fatalf("Expected target operator ID %q, got %q", "operator-002", operatorID)
			}
			if apiKeyID != "key-002" {
				t.Fatalf("Expected API key ID %q, got %q", "key-002", apiKeyID)
			}
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMAPIKeyIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"api-key", "revoke", "--id", "key-002", "--user-id", "operator-002"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := fmt.Sprintf("IAM API key %s revoked successfully\n", "key-002")
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
