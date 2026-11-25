package service

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestAuthService_Login_WithoutProfileFlagUsesActiveProfile(t *testing.T) {
	activeEndpoints := configuration_models.EndpointsV2{
		IAM:  "https://iam.active.example.com",
		Dash: "https://dash.active.example.com",
		CH:   "https://ch.active.example.com",
	}

	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetEndpointsFunc = func() (configuration_models.EndpointsV2, error) {
		return activeEndpoints, nil
	}
	mockCfg.GetSelectedProfileNameFunc = func() (string, error) {
		return "prod", nil
	}

	var createdProfileName, createdAPIKey, createdOrganizationID string
	mockCfg.CreateProfileFunc = func(profileName string, apiKey string, organizationID string) error {
		createdProfileName = profileName
		createdAPIKey = apiKey
		createdOrganizationID = organizationID

		return nil
	}

	organizationID := "prod-org-id"
	var requestedEndpoints configuration_models.EndpointsV2
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			endpoints configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			requestedEndpoints = endpoints

			return &api.IAMUser{OrganizationID: &organizationID}, nil
		},
	}

	authService := NewAuthService(mockCfg, &api.MockAuthAPI{}, mockUserAPI)

	cmd := setupTestCommand()
	cmd.Flags().String("username", "", "")
	cmd.Flags().String("organization", "", "")
	cmd.Flags().String("api-key", "", "")
	cmd.Flags().String("password", "", "")
	cmd.Flags().String("tfa-code", "", "")
	cmd.Flags().Set("api-key", "test-api-key")

	if err := authService.Login(cmd, []string{}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if requestedEndpoints.IAM != activeEndpoints.IAM {
		t.Errorf("expected authentication against the active profile endpoints, got %q", requestedEndpoints.IAM)
	}
	if createdProfileName != "prod" {
		t.Errorf("expected profile 'prod' to be created, got %q", createdProfileName)
	}
	if createdAPIKey != "test-api-key" {
		t.Errorf("expected API key to be stored, got %q", createdAPIKey)
	}
	if createdOrganizationID != organizationID {
		t.Errorf("expected organization ID %q, got %q", organizationID, createdOrganizationID)
	}
}

func TestAuthService_Login_WithProfileFlagUsesSelectedProfileEndpoints(t *testing.T) {
	selectedEndpoints := configuration_models.EndpointsV2{
		IAM:  "https://iam.dev.example.com",
		Dash: "https://dash.dev.example.com",
		CH:   "https://ch.dev.example.com",
	}

	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetEndpointsFunc = func() (configuration_models.EndpointsV2, error) {
		return selectedEndpoints, nil
	}
	mockCfg.GetSelectedProfileNameFunc = func() (string, error) {
		return "dev", nil
	}

	var createdProfileName string
	mockCfg.CreateProfileFunc = func(profileName string, apiKey string, organizationID string) error {
		createdProfileName = profileName

		return nil
	}

	organizationID := "dev-org-id"
	var requestedEndpoints configuration_models.EndpointsV2
	mockUserAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(
			endpoints configuration_models.EndpointsV2,
			accessToken string,
			apiKey string,
		) (*api.IAMUser, error) {
			requestedEndpoints = endpoints

			return &api.IAMUser{OrganizationID: &organizationID}, nil
		},
	}

	authService := NewAuthService(mockCfg, &api.MockAuthAPI{}, mockUserAPI)

	cmd := setupTestCommand()
	cmd.Flags().String("username", "", "")
	cmd.Flags().String("organization", "", "")
	cmd.Flags().String("api-key", "", "")
	cmd.Flags().String("password", "", "")
	cmd.Flags().String("tfa-code", "", "")
	cmd.Flags().Set("profile", "dev")
	cmd.Flags().Set("api-key", "dev-api-key")

	if err := authService.Login(cmd, []string{}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if requestedEndpoints.IAM != selectedEndpoints.IAM {
		t.Errorf("expected authentication against the requested profile endpoints, got %q", requestedEndpoints.IAM)
	}
	if createdProfileName != "dev" {
		t.Errorf("expected profile 'dev' to be updated, got %q", createdProfileName)
	}
}

func TestAuthService_Logout_WithoutProfileFlagUsesActiveProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileNameFunc = func() (string, error) {
		return "prod", nil
	}

	var deletedProfileName string
	mockCfg.DeleteProfileFunc = func(profileName string) error {
		deletedProfileName = profileName

		return nil
	}

	authService := NewAuthService(mockCfg, &api.MockAuthAPI{}, &api.MockUserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Bool("all", false, "")

	if err := authService.Logout(cmd, []string{}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deletedProfileName != "prod" {
		t.Errorf("expected active profile 'prod' to be deleted, got %q", deletedProfileName)
	}
}

func TestAuthService_Logout_WithProfileFlagDeletesRequestedProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()

	var deletedProfileName string
	mockCfg.DeleteProfileFunc = func(profileName string) error {
		deletedProfileName = profileName

		return nil
	}

	authService := NewAuthService(mockCfg, &api.MockAuthAPI{}, &api.MockUserAPI{})

	cmd := setupTestCommand()
	cmd.Flags().Bool("all", false, "")
	cmd.Flags().Set("profile", "dev")

	if err := authService.Logout(cmd, []string{}); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deletedProfileName != "dev" {
		t.Errorf("expected requested profile 'dev' to be deleted, got %q", deletedProfileName)
	}
}
