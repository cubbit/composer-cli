package cmd_iam

import (
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Edit_Integration_HumanOutput(t *testing.T) {
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
			if request.FirstName == nil || *request.FirstName != "Alice" {
				t.Fatalf("Expected first name %q, got %+v", "Alice", request.FirstName)
			}
			if request.LastName == nil || *request.LastName != "Updated" {
				t.Fatalf("Expected last name %q, got %+v", "Updated", request.LastName)
			}
			if request.Email == nil || *request.Email != "alice.updated@example.com" {
				t.Fatalf("Expected email %q, got %+v", "alice.updated@example.com", request.Email)
			}
			if request.Enabled != nil {
				t.Fatalf("Expected enabled to be omitted by edit, got %+v", request.Enabled)
			}

			return &api.IAMUser{
				ID:        "550e8400-e29b-41d4-a716-446655440000",
				Username:  "alice.wonder",
				FirstName: "Alice",
				LastName:  "Updated",
				Emails: []api.IAMUserEmail{
					{Email: "alice.updated@example.com", Default: true},
				},
				Status:    "active",
				Enabled:   false,
				CreatedAt: createdAt,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
		"--first-name", "Alice",
		"--last-name", "Updated",
		"--email", "alice.updated@example.com",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────┬──────────────────────────────────────┬───────────────────────────┬────────────┬───────────┬────────┬─────────┬─────────────────────╮
│ Username     │ ID                                   │ Email                     │ First Name │ Last Name │ Status │ Enabled │ Created At          │
├──────────────┼──────────────────────────────────────┼───────────────────────────┼────────────┼───────────┼────────┼─────────┼─────────────────────┤
│ alice.wonder │ 550e8400-e29b-41d4-a716-446655440000 │ alice.updated@example.com │ Alice      │ Updated   │ active │ false   │ 2024-03-10 14:30:00 │
╰──────────────┴──────────────────────────────────────┴───────────────────────────┴────────────┴───────────┴────────┴─────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user update output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestIAMUserSubCmd_Edit_Integration_WithUsername(t *testing.T) {
	var capturedSearch string
	var capturedUserID string
	mockUserAPI := &api.MockUserAPI{
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
			capturedSearch = search
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{ID: "user-001", Username: "alice"},
				},
			}, nil
		},
		UpdateIAMUserFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.UpdateIAMUserRequestBody,
		) (*api.IAMUser, error) {
			capturedUserID = userID
			return &api.IAMUser{
				ID:       userID,
				Username: "alice",
				Status:   "active",
				Enabled:  true,
			}, nil
		},
	}

	iamCmd, _ := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--username", "alice",
		"--first-name", "Alice",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if capturedSearch != "alice" {
		t.Fatalf("Expected username resolver search %q, got %q", "alice", capturedSearch)
	}
	if capturedUserID != "user-001" {
		t.Fatalf("Expected resolved user ID %q, got %q", "user-001", capturedUserID)
	}
}

func TestIAMUserSubCmd_Edit_Integration_RequiresUpdateField(t *testing.T) {
	mockUserAPI := &api.MockUserAPI{}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"edit",
		"--user-id", "550e8400-e29b-41d4-a716-446655440000",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no cobra execution error, got %v", err)
	}
	if !strings.Contains(commandOutput.String(), "specify at least one field to update") {
		t.Fatalf("Expected missing update field error, got %q", commandOutput.String())
	}
}
