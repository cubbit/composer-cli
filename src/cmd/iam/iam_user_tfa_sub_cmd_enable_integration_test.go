package cmd_iam

import (
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestIAMUserTFASubCmd_Enable_Integration_QuietOutput(t *testing.T) {
	organizationName := "test-org"
	setupKey := "JBSWY3DPEHPK3PXP"
	var capturedRequest *api.UpdateIAMUserTFARequestBody

	mockAuthAPI := iamUserCommandChallengeAPI{
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
			return "enable-tfa-token", nil
		},
	}
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			endpoints configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-id",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		EnableIAMUserTFAFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			enableTFAToken string,
			request *api.UpdateIAMUserTFARequestBody,
		) error {
			capturedRequest = request
			return nil
		},
	}

	iamCmd, commandOutput := setupIAMUserIntegrationCommand(mockAuthAPI, mockUserAPI)
	iamCmd.SetArgs([]string{
		"user",
		"tfa",
		"enable",
		"--password", "admin-password",
		"--setup-key", setupKey,
		"--quiet",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if capturedRequest == nil {
		t.Fatalf("Expected enable TFA request to be captured")
	}
	if len(capturedRequest.Codes) != 10 {
		t.Fatalf("Expected 10 recovery codes, got %d", len(capturedRequest.Codes))
	}

	expectedFields := append([]string{setupKey}, capturedRequest.Codes...)
	expectedOutput := strings.Join(expectedFields, "\t") + "\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
