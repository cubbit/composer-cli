package create

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func setupUserBulkCreateTestCommand() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.SetErr(new(bytes.Buffer))
	cmd.Flags().String("file", "", "Users file")
	cmd.Flags().String("sample", "", "Sample format")
	cmd.Flags().String("username", "", "Username")
	cmd.Flags().String("password", "", "Password")
	cmd.Flags().String("email", "", "Email")
	cmd.Flags().StringArray("policy", []string{}, "Policy")
	cmd.Flags().String("profile", "", "Profile")
	cmd.Flags().String("output", "human", "Output format")
	cmd.Flags().Bool("no-headers", false, "No headers")
	cmd.Flags().Bool("quiet", false, "Quiet mode")
	return cmd
}

func userBulkCreateTestProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
		Output:         configuration_models.OutputHuman,
		Endpoints: configuration_models.EndpointsV2{
			IAM: "https://iam.example.com",
		},
	}
}

func TestImportUsers_SampleJSON(t *testing.T) {
	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("sample", "json")

	err := ImportUsers(Dependencies{}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	for _, expected := range []string{
		`"users": [`,
		`"username": "user-1"`,
		`"password": "password-1"`,
		`"attached_policies": []`,
		`"username": "user-2"`,
	} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected sample JSON to contain %q, got %q", expected, output)
		}
	}
}

func TestImportUsers_SampleCSV(t *testing.T) {
	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("sample", "csv")

	err := ImportUsers(Dependencies{}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := `username,password,first_name,last_name,email,attached_policies
user-1,password-1,first-name-1,last-name-1,email-1@example.com,
user-2,password-2,first-name-2,last-name-2,,
`
	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if output != expectedOutput {
		t.Fatalf("Expected sample CSV %q, got %q", expectedOutput, output)
	}
}

func TestImportUsers_InvalidSampleFormat(t *testing.T) {
	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("sample", "yaml")

	err := ImportUsers(Dependencies{}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err == nil {
		t.Fatal("Expected invalid sample format error, got nil")
	}
	if !strings.Contains(err.Error(), `invalid sample format "yaml": expected json or csv`) {
		t.Fatalf("Expected invalid sample format error, got %v", err)
	}
}

func TestImportUsers_MissingFile(t *testing.T) {
	cmd := setupUserBulkCreateTestCommand()

	err := ImportUsers(Dependencies{}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err == nil {
		t.Fatal("Expected missing file error, got nil")
	}
	if !strings.Contains(err.Error(), "file is required") {
		t.Fatalf("Expected missing file error, got %v", err)
	}
}

func TestImportUsers_Quiet(t *testing.T) {
	usersFile := writeUsersFile(t, `{
		"users": [
			{
				"username": "user1",
				"password": "test-password",
				"email": "user1@example.com",
				"attached_policies": ["695ed3dd-e77d-42b9-88ed-70bd3a1704ee"]
			}
		]
	}`)

	email := "user1@example.com"
	var actualEndpoints configuration_models.EndpointsV2
	var actualAPIKey string
	var actualOrganizationID string
	var actualRequest *api.BulkCreateIAMUsersRequestBody
	var actualSaltsRequest *api.BulkGenerateSaltsRequestBody
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			actualSaltsRequest = request
			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "user1", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			actualEndpoints = endpoints
			actualAPIKey = apiKey
			actualOrganizationID = organizationID
			actualRequest = request
			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-id-1",
						Username: "user1",
						Email:    &email,
						Created:  true,
						Status:   "pending",
					},
				},
			}, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)
	cmd.Flags().Set("quiet", "true")

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest := &api.BulkCreateIAMUsersRequestBody{
		Users: []api.BulkCreateIAMUserRequestBody{
			{
				Username:                "user1",
				AuthenticationPublicKey: testAuthenticationPublicKey("test-password", "test-salt"),
				Email:                   &email,
				AttachedPolicies:        []string{"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"},
			},
		},
	}
	if actualEndpoints.IAM != "https://iam.example.com" {
		t.Fatalf("Expected IAM endpoint to be propagated, got %q", actualEndpoints.IAM)
	}
	if actualAPIKey != "test-api-key" {
		t.Fatalf("Expected API key to be propagated, got %q", actualAPIKey)
	}
	if actualOrganizationID != "test-org-id" {
		t.Fatalf("Expected organization ID to be propagated, got %q", actualOrganizationID)
	}
	if actualSaltsRequest == nil || len(actualSaltsRequest.Operators) != 1 || actualSaltsRequest.Operators[0].Username != "user1" {
		t.Fatalf("Expected salts request with user1, got %+v", actualSaltsRequest)
	}
	if !reflect.DeepEqual(actualRequest, expectedRequest) {
		t.Fatalf("Expected request %+v, got %+v", expectedRequest, actualRequest)
	}

	expectedOutput := ""
	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	if output != expectedOutput {
		t.Fatalf("Expected quiet output %q, got %q", expectedOutput, output)
	}
}

func TestImportUsers_JSONIncludesCount(t *testing.T) {
	usersFile := writeUsersFile(t, `{"users":[{"username":"user1","password":"test-password"}]}`)
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "user1", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-id-1",
						Username: "user1",
						Created:  true,
						Status:   "pending",
					},
				},
			}, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)
	cmd.Flags().Set("output", "json")

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	output := cmd.OutOrStdout().(*bytes.Buffer).String()
	for _, expected := range []string{`"count": 1`, `"data":`, `"username": "user1"`} {
		if !strings.Contains(output, expected) {
			t.Fatalf("Expected JSON output to contain %q, got %q", expected, output)
		}
	}
}

func TestImportUsers_CSV(t *testing.T) {
	usersFile := writeUsersFileWithPattern(t, "users-*.csv", `username,password,first_name,last_name,email,attached_policies
user1,test-password,John,Doe,user1@example.com,695ed3dd-e77d-42b9-88ed-70bd3a1704ee;11111111-2222-3333-4444-555555555555
`)

	firstName := "John"
	lastName := "Doe"
	email := "user1@example.com"
	var actualRequest *api.BulkCreateIAMUsersRequestBody
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "user1", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			actualRequest = request
			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-id-1",
						Username: "user1",
						Email:    &email,
						Created:  true,
						Status:   "pending",
					},
				},
			}, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)
	cmd.Flags().Set("quiet", "true")

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest := &api.BulkCreateIAMUsersRequestBody{
		Users: []api.BulkCreateIAMUserRequestBody{
			{
				Username:                "user1",
				AuthenticationPublicKey: testAuthenticationPublicKey("test-password", "test-salt"),
				FirstName:               &firstName,
				LastName:                &lastName,
				Email:                   &email,
				AttachedPolicies:        []string{"695ed3dd-e77d-42b9-88ed-70bd3a1704ee", "11111111-2222-3333-4444-555555555555"},
			},
		},
	}
	if !reflect.DeepEqual(actualRequest, expectedRequest) {
		t.Fatalf("Expected request %+v, got %+v", expectedRequest, actualRequest)
	}
}

func TestCreateUser(t *testing.T) {
	email := "user1@example.com"
	var actualRequest *api.BulkCreateIAMUsersRequestBody
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			return &api.BulkGenerateSaltsResponse{
				Count: 1,
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "user1", Salt: "test-salt"},
				},
			}, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			actualRequest = request
			return &api.BulkCreateIAMUsersResponse{
				Count: 1,
				Data: []api.BulkCreateIAMUserResponseItem{
					{
						ID:       "user-id-1",
						Username: "user1",
						Email:    &email,
						Created:  true,
						Status:   "pending",
					},
				},
			}, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("username", "user1")
	cmd.Flags().Set("password", "test-password")
	cmd.Flags().Set("email", email)
	cmd.Flags().Set("policy", "695ed3dd-e77d-42b9-88ed-70bd3a1704ee")
	cmd.Flags().Set("quiet", "true")

	err := CreateUser(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedRequest := &api.BulkCreateIAMUsersRequestBody{
		Users: []api.BulkCreateIAMUserRequestBody{
			{
				Username:                "user1",
				AuthenticationPublicKey: testAuthenticationPublicKey("test-password", "test-salt"),
				Email:                   &email,
				AttachedPolicies:        []string{"695ed3dd-e77d-42b9-88ed-70bd3a1704ee"},
			},
		},
	}
	if !reflect.DeepEqual(actualRequest, expectedRequest) {
		t.Fatalf("Expected request %+v, got %+v", expectedRequest, actualRequest)
	}
}

func TestImportUsers_InvalidJSON(t *testing.T) {
	usersFile := writeUsersFile(t, `{`)
	bulkCreateCalled := false
	mockUserAPI := &api.MockUserAPI{
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			bulkCreateCalled = true
			return nil, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err == nil {
		t.Fatal("Expected malformed JSON error, got nil")
	}
	if !strings.Contains(err.Error(), "error parsing JSON from file") {
		t.Fatalf("Expected JSON parse error, got %v", err)
	}
	if bulkCreateCalled {
		t.Fatal("BulkCreateIAMUsers API should not be called for malformed JSON")
	}
}

func TestImportUsers_MissingPassword(t *testing.T) {
	usersFile := writeUsersFile(t, `{"users":[{"username":"user1"}]}`)
	saltsCalled := false
	bulkCreateCalled := false
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			saltsCalled = true
			return nil, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			bulkCreateCalled = true
			return nil, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err == nil {
		t.Fatal("Expected missing password error, got nil")
	}
	if !strings.Contains(err.Error(), `password is required for user "user1"`) {
		t.Fatalf("Expected missing password error, got %v", err)
	}
	if saltsCalled {
		t.Fatal("BulkGenerateSalts should not be called when password is missing")
	}
	if bulkCreateCalled {
		t.Fatal("BulkCreateIAMUsers API should not be called when password is missing")
	}
}

func TestImportUsers_InvalidAttachedPolicy(t *testing.T) {
	usersFile := writeUsersFile(t, `{"users":[{"username":"user1","password":"test-password","attached_policies":["$POLICY_ID"]}]}`)
	saltsCalled := false
	bulkCreateCalled := false
	mockUserAPI := &api.MockUserAPI{
		BulkGenerateSaltsFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkGenerateSaltsRequestBody) (*api.BulkGenerateSaltsResponse, error) {
			saltsCalled = true
			return nil, nil
		},
		BulkCreateIAMUsersFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, request *api.BulkCreateIAMUsersRequestBody) (*api.BulkCreateIAMUsersResponse, error) {
			bulkCreateCalled = true
			return nil, nil
		},
	}

	cmd := setupUserBulkCreateTestCommand()
	cmd.Flags().Set("file", usersFile)

	err := ImportUsers(Dependencies{UserAPI: mockUserAPI}, cmd, nil, userBulkCreateTestProfile(), "test-org")
	if err == nil {
		t.Fatal("Expected invalid attached policy error, got nil")
	}
	if !strings.Contains(err.Error(), `invalid attached policy "$POLICY_ID" for user "user1": expected UUID`) {
		t.Fatalf("Expected invalid attached policy error, got %v", err)
	}
	if saltsCalled {
		t.Fatal("BulkGenerateSalts should not be called when attached policy is invalid")
	}
	if bulkCreateCalled {
		t.Fatal("BulkCreateIAMUsers API should not be called when attached policy is invalid")
	}
}

func writeUsersFile(t *testing.T, content string) string {
	return writeUsersFileWithPattern(t, "users-*.json", content)
}

func writeUsersFileWithPattern(t *testing.T, pattern string, content string) string {
	t.Helper()

	file, err := os.CreateTemp(t.TempDir(), pattern)
	if err != nil {
		t.Fatalf("Failed to create users file: %v", err)
	}
	defer file.Close()

	if _, err := fmt.Fprint(file, content); err != nil {
		t.Fatalf("Failed to write users file: %v", err)
	}

	return file.Name()
}

func testAuthenticationPublicKey(password string, salt string) string {
	seed := sha256.Sum256([]byte(password + salt))
	privateKey := ed25519.NewKeyFromSeed(seed[:])
	return base64.StdEncoding.EncodeToString(privateKey.Public().(ed25519.PublicKey))
}
