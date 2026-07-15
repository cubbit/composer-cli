package cmd_iam

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func setupIAMUserIntegrationCommand(
	authAPI api.AuthAPIInterface,
	userAPI api.UserAPIInterface,
) (*cobra.Command, *bytes.Buffer) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
			},
		}, nil
	}

	userService := service.NewUserService(mockCfg, authAPI, userAPI)
	iamCmd := NewIAMCmd(userService)
	iamCmd.PersistentFlags().String("profile", "", "Profile")
	iamCmd.PersistentFlags().String("output", "human", "Output format")
	iamCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	iamCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)

	return iamCmd, commandOutput
}

func writeIAMUsersIntegrationFile(t *testing.T, content string) string {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), "users-*.json")
	if err != nil {
		t.Fatalf("Failed to create users file: %v", err)
	}
	defer file.Close()

	if _, err := fmt.Fprint(file, content); err != nil {
		t.Fatalf("Failed to write users file: %v", err)
	}

	return file.Name()
}

type iamUserCommandChallengeAPI struct {
	GenerateChallengeFunc func(
		endpoints configuration_models.EndpointsV2,
		email *string,
		username *string,
		organizationName *string,
	) (*api.ChallengeResponseModel, error)
}

func (m iamUserCommandChallengeAPI) GenerateChallenge(
	endpoints configuration_models.EndpointsV2,
	email *string,
	username *string,
	organizationName *string,
) (*api.ChallengeResponseModel, error) {
	return m.GenerateChallengeFunc(endpoints, email, username, organizationName)
}

func (m iamUserCommandChallengeAPI) Activate(
	endpoints configuration_models.EndpointsV2,
	token string,
) error {
	return nil
}

func (m iamUserCommandChallengeAPI) SignUp(
	endpoints configuration_models.EndpointsV2,
	email string,
	username string,
	firstName *string,
	lastName *string,
	authenticationPublicKey *string,
	organizationName string,
	organizationBasePolicy map[string]interface{},
	organizationSettings map[string]interface{},
) error {
	return nil
}

func (m iamUserCommandChallengeAPI) SignIn(
	endpoints configuration_models.EndpointsV2,
	username string,
	organization string,
	password string,
	tfaCode string,
) (*api.SignInToken, error) {
	return nil, nil
}

func (m iamUserCommandChallengeAPI) ForgeToken(
	endpoints configuration_models.EndpointsV2,
	operatorID string,
	username string,
	organizationName string,
	password string,
	tfaCode string,
	tokenType string,
	token string,
	refreshToken string,
) (string, error) {
	return "", nil
}

func (m iamUserCommandChallengeAPI) CreateApiKey(
	endpoints configuration_models.EndpointsV2,
	operatorID string,
	name string,
	token string,
	forgeApiKeyToken string,
) (string, error) {
	return "", nil
}

func TestIAMUserSubCmd_List_Integration_HumanOutput(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	firstName := "Alice"
	lastName := "Wonder"
	disabledFirstName := "Bob"
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
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if page != 1 || items != 100 {
				t.Fatalf("Expected page=1 items=100, got page=%d items=%d", page, items)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:        "user-001",
						Username:  "alice.wonder",
						FirstName: &firstName,
						LastName:  &lastName,
						Enabled:   true,
						CreatedAt: createdAt,
						Emails: []api.IAMUserEmail{
							{Email: "alice.secondary@example.com"},
							{Email: "alice@example.com", Default: true},
						},
						Status: "active",
					},
					{
						ID:        "user-002",
						Username:  "bob.builder",
						FirstName: &disabledFirstName,
						Enabled:   false,
						CreatedAt: createdAt.Add(time.Hour),
						Emails: []api.IAMUserEmail{
							{Email: "bob@example.com"},
						},
						Status: "pending",
					},
				},
				NextPage: nil,
				Count:    2,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{"user", "list"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────────────┬──────────┬───────────────────┬────────────┬───────────┬─────────┬─────────┬─────────────────────╮
│ Username     │ ID       │ Email             │ First Name │ Last Name │ Status  │ Enabled │ Created At          │
├──────────────┼──────────┼───────────────────┼────────────┼───────────┼─────────┼─────────┼─────────────────────┤
│ alice.wonder │ user-001 │ alice@example.com │ Alice      │ Wonder    │ active  │ true    │ 2024-01-15 10:30:00 │
│ bob.builder  │ user-002 │ bob@example.com   │ Bob        │           │ pending │ false   │ 2024-01-15 11:30:00 │
╰──────────────┴──────────┴───────────────────┴────────────┴───────────┴─────────┴─────────┴─────────────────────╯
`)

	actualResult := strings.TrimSpace(commandOutput.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM user list output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestIAMUserSubCmd_List_Integration_WithFilters(t *testing.T) {
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
			if enabled == nil || *enabled != true {
				t.Fatalf("Expected enabled=true, got %v", enabled)
			}
			if search != "alice" {
				t.Fatalf("Expected search alice, got %q", search)
			}
			if page != 2 || items != 50 {
				t.Fatalf("Expected page=2 items=50, got page=%d items=%d", page, items)
			}
			if sortKey != "username" {
				t.Fatalf("Expected sort-key username, got %q", sortKey)
			}
			if sortOrder != "asc" {
				t.Fatalf("Expected sort-order asc, got %q", sortOrder)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data:     []api.IAMUserListItem{},
				NextPage: nil,
				Count:    0,
			}, nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(nil, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"list",
		"--enabled", "true",
		"--search", "alice",
		"--page", "2",
		"--items", "50",
		"--sort-key", "username",
		"--sort-order", "asc",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "No IAM users found.\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
}
