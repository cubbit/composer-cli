package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Describe_Integration_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 3, 10, 14, 30, 0, 0, time.UTC)
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}

			email := "alice@example.com"
			defaultEmail := "alice.default@example.com"

			return &api.IAMUser{
				ID:               "550e8400-e29b-41d4-a716-446655440000",
				Username:         "alice.wonder",
				FirstName:        "Alice",
				LastName:         "Wonder",
				Enabled:          true,
				Internal:         false,
				Banned:           false,
				IsRoot:           false,
				TwoFactorEnabled: false,
				Status:           "active",
				CreatedAt:        createdAt,
				OrganizationName: &organizationName,
				Emails: []api.IAMUserEmail{
					{Email: email, Default: false},
					{Email: defaultEmail, Default: true},
				},
				Policies: []api.IAMUserPolicy{
					{ID: "policy-001", Name: "StorageAdmin"},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
User: Alice Wonder
Root: No
Username: alice.wonder
Status: ● Active

Organization:
  Name: test-org
  ID: test-org-id

Policies:
  StorageAdmin

Emails:
  alice@example.com
  alice.default@example.com (default)

Metadata:
  ID: 550e8400-e29b-41d4-a716-446655440000
  Two-Factor: Disabled
  Created At: 2024-03-10 14:30:00
  Deleted At: N/A
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user describe output does not match actual output\nExpected:\n%s\n\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestIAMUserSubCmd_Describe_Integration_CurrentUser(t *testing.T) {
	createdAt := time.Date(2024, 3, 10, 14, 30, 0, 0, time.UTC)
	organizationName := "test-org"

	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfV3Func: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			defaultEmail := "admin@example.com"

			return &api.IAMUser{
				ID:               "admin-id",
				Username:         "admin",
				FirstName:        "Admin",
				LastName:         "User",
				Enabled:          true,
				Internal:         false,
				Banned:           false,
				IsRoot:           true,
				TwoFactorEnabled: false,
				Status:           "active",
				CreatedAt:        createdAt,
				OrganizationName: &organizationName,
				Emails: []api.IAMUserEmail{
					{Email: defaultEmail, Default: true},
				},
				Policies: []api.IAMUserPolicy{
					{ID: "policy-001", Name: "Owner"},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"describe",
		"--self",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
User: Admin User
Root: Yes
Username: admin
Status: ● Active

Organization:
  Name: test-org
  ID: test-org-id

Policies:
  Owner

Emails:
  admin@example.com (default)

Metadata:
  ID: admin-id
  Two-Factor: Disabled
  Created At: 2024-03-10 14:30:00
  Deleted At: N/A
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user describe output does not match actual output\nExpected:\n%s\n\nActual:\n%s", expectedResult, actualResult)
	}
}
