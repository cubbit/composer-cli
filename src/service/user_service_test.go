package service

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func TestUserService_ImportUsers_SampleSkipsProfileLoading(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		t.Fatal("GetActiveProfile should not be called when generating an import sample")
		return configuration_models.ProfileV2{}, nil
	}

	cmd := &cobra.Command{}
	cmd.SetOut(new(bytes.Buffer))
	cmd.Flags().String("sample", "", "Sample format")
	cmd.Flags().String("file", "", "Users file")
	cmd.Flags().Set("sample", "json")

	service := NewUserService(mockCfg, nil, nil)
	err := service.ImportUsers(cmd, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !strings.Contains(cmd.OutOrStdout().(*bytes.Buffer).String(), `"users": [`) {
		t.Fatalf("Expected sample JSON output, got %q", cmd.OutOrStdout().(*bytes.Buffer).String())
	}
}

func TestUserService_GetProfileOrganizationName(t *testing.T) {
	organizationName := "test-org"
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected API key to be propagated, got %q", apiKey)
			}
			return &api.IAMUser{
				OrganizationName: &organizationName,
			}, nil
		},
	}
	service := NewUserService(nil, nil, userAPI)

	actualOrganizationName, err := service.getProfileOrganizationName(configuration_models.ProfileV2{
		APIKey: "test-api-key",
		Endpoints: configuration_models.EndpointsV2{
			IAM: "https://iam.example.com",
		},
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if actualOrganizationName != organizationName {
		t.Fatalf("Expected organization name %q, got %q", organizationName, actualOrganizationName)
	}
}

func TestUserService_GetProfileOrganizationName_Missing(t *testing.T) {
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{}, nil
		},
	}
	service := NewUserService(nil, nil, userAPI)

	_, err := service.getProfileOrganizationName(configuration_models.ProfileV2{})
	if err == nil {
		t.Fatal("Expected missing organization name error, got nil")
	}
	if !strings.Contains(err.Error(), "current IAM user does not expose an organization name") {
		t.Fatalf("Expected missing organization name error, got %v", err)
	}
}
