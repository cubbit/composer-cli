package cmd_iam

import (
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Create_Integration_HumanOutput(t *testing.T) {
	email := "alice@example.com"
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "alice", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if len(request.Users) != 1 || request.Users[0].Username != "alice" {
				t.Fatalf("Expected create request for alice, got %+v", request)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
					},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(iamUserCommandChallengeAPI{}, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"create",
		"--username", "alice",
		"--password", "test-password",
		"--email", email,
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────┬──────────┬───────────────────┬─────────┬─────────╮
│ Username │ ID       │ Email             │ Created │ Status  │
├──────────┼──────────┼───────────────────┼─────────┼─────────┤
│ alice    │ user-001 │ alice@example.com │ true    │ pending │
╰──────────┴──────────┴───────────────────┴─────────┴─────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user create output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
