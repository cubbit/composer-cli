package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Enable_Integration_UpdateEnabledTrue(t *testing.T) {
	createdAt := time.Date(2024, 3, 10, 14, 30, 0, 0, time.UTC)
	mockUserAPI := &api.MockUserAPI{
		UpdateIAMUserFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}
			if request.Enabled == nil || *request.Enabled != true {
				t.Fatalf("Expected enabled=true, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        userID,
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Wonder",
				Emails: []api.IAMUserEmail{
					{Email: "alice.updated@example.com", Default: true},
				},
				Status:    "active",
				Enabled:   true,
				CreatedAt: createdAt,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"enable",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────┬──────────────────────────────────────┬───────────────────────────┬────────────┬───────────┬────────┬─────────┬─────────────────────╮
│ Username     │ ID                                   │ Email                     │ First Name │ Last Name │ Status │ Enabled │ Created At          │
├──────────────┼──────────────────────────────────────┼───────────────────────────┼────────────┼───────────┼────────┼─────────┼─────────────────────┤
│ alice.wonder │ 550e8400-e29b-41d4-a716-446655440000 │ alice.updated@example.com │ Alice      │ Wonder    │ active │ true    │ 2024-03-10 14:30:00 │
╰──────────────┴──────────────────────────────────────┴───────────────────────────┴────────────┴───────────┴────────┴─────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user enable output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
