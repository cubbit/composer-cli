package cmd_iam

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Delete_Integration_HumanOutput(t *testing.T) {
	var capturedAPIKey string
	var capturedOrganizationID string
	var capturedUserID string
	var capturedDeleteToken string
	mockUserAPI := &api.MockUserAPI{
		DeleteIAMUserFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			deleteToken string,
		) error {
			capturedAPIKey = apiKey
			capturedOrganizationID = organizationID
			capturedUserID = userID
			capturedDeleteToken = deleteToken
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"delete",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "IAM user 550e8400-e29b-41d4-a716-446655440000 deleted successfully\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
	if capturedAPIKey != "test-api-key" {
		t.Fatalf("Expected api key %q, got %q", "test-api-key", capturedAPIKey)
	}
	if capturedOrganizationID != "test-org-id" {
		t.Fatalf("Expected organization ID %q, got %q", "test-org-id", capturedOrganizationID)
	}
	if capturedUserID != "550e8400-e29b-41d4-a716-446655440000" {
		t.Fatalf("Expected user ID %q, got %q", "550e8400-e29b-41d4-a716-446655440000", capturedUserID)
	}
	if capturedDeleteToken != "" {
		t.Fatalf("Expected empty delete token, got %q", capturedDeleteToken)
	}
}
