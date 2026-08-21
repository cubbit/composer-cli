package cmd_iam

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
)

func TestIAMUserPasswordSubCmd_Reset_Integration_HumanOutput(t *testing.T) {
	t.Setenv("PASSWORD", "admin-password")

	organizationName := "test-org"
	var capturedTokenType string
	var capturedForgeOperatorID string
	var capturedListSearch string
	var capturedSaltUsername string
	var capturedResetRequest *api.ResetIAMUserPasswordRequestBody

	mockAuthAPI := iamUserCommandChallengeAPI{
		SignInFunc: func(
			endpoints configuration_models.EndpointsV2,
			username string,
			organization string,
			password string,
			tfaCode string,
		) (*api.SignInToken, error) {
			if username != "admin" {
				t.Fatalf("Expected admin username, got %q", username)
			}
			if organization != organizationName {
				t.Fatalf("Expected organization name %q, got %q", organizationName, organization)
			}
			if password != "admin-password" {
				t.Fatalf("Expected admin password to be propagated, got %q", password)
			}
			if tfaCode != "123456" {
				t.Fatalf("Expected TFA code to be propagated, got %q", tfaCode)
			}

			return &api.SignInToken{AccessToken: "access-token", RefreshToken: "refresh-token"}, nil
		},
		ForgeTokenFunc: func(
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
			capturedForgeOperatorID = operatorID
			capturedTokenType = tokenType
			if username != "admin" {
				t.Fatalf("Expected admin username, got %q", username)
			}
			if organizationName != "test-org" {
				t.Fatalf("Expected organization name, got %q", organizationName)
			}
			if password != "admin-password" {
				t.Fatalf("Expected admin password to be propagated, got %q", password)
			}
			if tfaCode != "123456" {
				t.Fatalf("Expected TFA code to be propagated, got %q", tfaCode)
			}
			if token != "access-token" {
				t.Fatalf("Expected access token to be propagated, got %q", token)
			}
			if refreshToken != "refresh-token" {
				t.Fatalf("Expected refresh token to be propagated, got %q", refreshToken)
			}

			return "reset-token", nil
		},
	}
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
			capturedListSearch = search

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{
						ID:       "550e8400-e29b-41d4-a716-446655440000",
						Username: "target-user",
					},
				},
			}, nil
		},
		GetIAMUserByIDFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
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

			return &api.IAMUser{
				ID:       userID,
				Username: "target-user",
			}, nil
		},
		GetIAMUserSelfFunc: func(
			endpoints configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}

			return &api.IAMUser{
				ID:               "admin-id",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		BulkGenerateSaltsFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			capturedSaltUsername = request.Operators[0].Username
			return &api.BulkGenerateSaltsResponse{
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "target-user", Salt: "target-salt"},
				},
				Count: 1,
			}, nil
		},
		ResetIAMUserPasswordFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.ResetIAMUserPasswordRequestBody,
		) error {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key to be propagated, got %q", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID to be propagated, got %q", organizationID)
			}
			if userID != "550e8400-e29b-41d4-a716-446655440000" {
				t.Fatalf("Expected user ID to be propagated, got %q", userID)
			}

			capturedResetRequest = request
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(mockAuthAPI, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--username", "target-user",
		"--new-password", "new-password",
		"--tfa-code", "123456",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "IAM user 550e8400-e29b-41d4-a716-446655440000 password reset successfully\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
	if capturedListSearch != "target-user" {
		t.Fatalf("Expected username search to be propagated, got %q", capturedListSearch)
	}
	if capturedForgeOperatorID != "550e8400-e29b-41d4-a716-446655440000" {
		t.Fatalf("Expected forged token target user ID, got %q", capturedForgeOperatorID)
	}
	if capturedTokenType != "iam_password_change" {
		t.Fatalf("Expected IAM password change token type, got %q", capturedTokenType)
	}
	if capturedSaltUsername != "target-user" {
		t.Fatalf("Expected salt username, got %q", capturedSaltUsername)
	}
	if capturedResetRequest == nil {
		t.Fatal("Expected reset request to be captured")
	}
	if capturedResetRequest.IAMPasswordChangeToken != "reset-token" {
		t.Fatalf("Expected IAM password change token in iam_password_change_token field, got %q", capturedResetRequest.IAMPasswordChangeToken)
	}
	expectedAuthenticationPublicKey, err := utils.AuthenticationPublicKeyFromPassword("new-password", "target-salt")
	if err != nil {
		t.Fatalf("Failed to generate expected authentication public key: %v", err)
	}
	if capturedResetRequest.AuthenticationPublicKey != expectedAuthenticationPublicKey {
		t.Fatalf("Expected authentication public key %q, got %q", expectedAuthenticationPublicKey, capturedResetRequest.AuthenticationPublicKey)
	}
}

func TestIAMUserPasswordSubCmd_Reset_Integration_PasswordFlagOverridesEnv(t *testing.T) {
	t.Setenv("PASSWORD", "env-password")

	organizationName := "test-org"
	var capturedSignInPassword string
	var capturedForgePassword string

	mockAuthAPI := iamUserCommandChallengeAPI{
		SignInFunc: func(
			endpoints configuration_models.EndpointsV2,
			username string,
			organization string,
			password string,
			tfaCode string,
		) (*api.SignInToken, error) {
			capturedSignInPassword = password
			return &api.SignInToken{AccessToken: "access-token", RefreshToken: "refresh-token"}, nil
		},
		ForgeTokenFunc: func(
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
			capturedForgePassword = password
			return "reset-token", nil
		},
	}
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			endpoints configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "frank-id",
				Username:         "frank",
				OrganizationName: &organizationName,
			}, nil
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
			if search != "frank" {
				t.Fatalf("Expected username search to be propagated, got %q", search)
			}

			return &api.GenericPaginatedResponse[api.IAMUserListItem]{
				Data: []api.IAMUserListItem{
					{ID: "frank-id", Username: "frank"},
				},
			}, nil
		},
		GetIAMUserByIDFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
		) (*api.IAMUser, error) {
			if userID != "frank-id" {
				t.Fatalf("Expected target user ID, got %q", userID)
			}

			return &api.IAMUser{
				ID:       "frank-id",
				Username: "frank",
			}, nil
		},
		BulkGenerateSaltsFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			request *api.BulkGenerateSaltsRequestBody,
		) (*api.BulkGenerateSaltsResponse, error) {
			return &api.BulkGenerateSaltsResponse{
				Data: []api.BulkGenerateSaltResponseItem{
					{Username: "frank", Salt: "frank-salt"},
				},
				Count: 1,
			}, nil
		},
		ResetIAMUserPasswordFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			userID string,
			request *api.ResetIAMUserPasswordRequestBody,
		) error {
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(mockAuthAPI, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"password",
		"reset",
		"--username", "frank",
		"--password", "flag-password",
		"--new-password", "new-password",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := "IAM user frank-id password reset successfully\n"
	actualResult := commandOutput.String()
	if actualResult != expectedResult {
		t.Fatalf("Expected output %q, got %q", expectedResult, actualResult)
	}
	if capturedSignInPassword != "flag-password" {
		t.Fatalf("Expected sign in password from flag, got %q", capturedSignInPassword)
	}
	if capturedForgePassword != "flag-password" {
		t.Fatalf("Expected forge password from flag, got %q", capturedForgePassword)
	}
}
