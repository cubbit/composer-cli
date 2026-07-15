package cmd_iam

import (
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserSubCmd_Import_Integration_HumanOutput(t *testing.T) {
	usersFile := writeIAMUsersIntegrationFile(t, `{"users":[{"username":"alice","password":"test-password","email":"alice@example.com"},{"username":"bob","password":"test-password"}]}`)
	email := "alice@example.com"
	mockAuthAPI := iamUserCommandChallengeAPI{
		GenerateChallengeFunc: func(endpoints configuration_models.EndpointsV2, email *string, username *string, organizationName *string) (*api.ChallengeResponseModel, error) {
			return &api.ChallengeResponseModel{Salt: "test-salt"}, nil
		},
	}
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			organizationName := "test-org"
			return &api.IAMUser{OrganizationName: &organizationName}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			if len(request.Users) != 2 {
				t.Fatalf("Expected two users in import request, got %+v", request)
			}

			return &api.BulkCreateIAMUsersResponse{
				Count: 2,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-001",
						Username: "alice",
						Email:    &email,
						Created:  true,
						Status:   "pending",
					},
					{
						ID:       "user-002",
						Username: "bob",
						Created:  false,
						Status:   "active",
					},
				},
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(mockAuthAPI, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"import",
		"--file", usersFile,
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
│ bob      │ user-002 │                   │ false   │ active  │
╰──────────┴──────────┴───────────────────┴─────────┴─────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user import output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
