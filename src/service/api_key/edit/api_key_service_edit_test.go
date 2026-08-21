package edit

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/api_key/shared"
	"github.com/spf13/cobra"
)

func TestEditAPIKey_UsesUpdateByID(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	expiresAt := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	organizationName := "test-org"
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		UpdateIAMAPIKeyFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			apiKeyID string,
			request *api.UpdateIAMAPIKeyRequestBody,
		) (*api.OperatorAPIKey, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected api key %q, got %q", "test-api-key", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID %q, got %q", "test-org-id", organizationID)
			}
			if operatorID != "operator-001" {
				t.Fatalf("Expected operator ID %q, got %q", "operator-001", operatorID)
			}
			if apiKeyID != "key-001" {
				t.Fatalf("Expected API key ID %q, got %q", "key-001", apiKeyID)
			}
			if request.Name == nil || *request.Name != "automation-renamed" {
				t.Fatalf("Expected API key name %q, got %#v", "automation-renamed", request.Name)
			}
			if request.ExpiresAt == nil || !request.ExpiresAt.Equal(expiresAt) {
				t.Fatalf("Expected expires at %v, got %#v", expiresAt, request.ExpiresAt)
			}
			if request.Enabled == nil || *request.Enabled {
				t.Fatalf("Expected enabled false, got %#v", request.Enabled)
			}

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation-renamed",
				OperatorID: "operator-001",
				CreatedAt:  createdAt,
				ExpiresAt:  &expiresAt,
				Enabled:    false,
			}, nil
		},
	}

	cmd := &cobra.Command{}
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.Flags().String("id", "key-001", "ID")
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("expires-at", "", "Expires at")
	cmd.Flags().Bool("enabled", true, "Enabled")
	cmd.Flags().String("output", "human", "Output")
	cmd.Flags().Bool("quiet", false, "Quiet")
	_ = cmd.Flags().Set("name", "automation-renamed")
	_ = cmd.Flags().Set("expires-at", "2024-12-31T23:59:59Z")
	_ = cmd.Flags().Set("enabled", "false")

	err := EditAPIKey(
		shared.Dependencies{UserAPI: userAPI},
		cmd,
		configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
		},
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResult := strings.TrimSpace(`
API Key: automation-renamed
Enabled: false

Metadata:
  ID: key-001
  Operator ID: operator-001
  Created At: 2024-01-15 10:30:00
  Expires At: 2024-12-31 23:59:59
  Deleted At: N/A
  Banned At: N/A
`)

	actualResult := strings.TrimSpace(output.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key edit output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}

func TestEditAPIKey_RequiresUpdateField(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("id", "key-001", "ID")
	cmd.Flags().String("name", "", "Name")
	cmd.Flags().String("expires-at", "", "Expires at")
	cmd.Flags().Bool("enabled", true, "Enabled")

	err := EditAPIKey(
		shared.Dependencies{UserAPI: &api.MockUserAPI{}},
		cmd,
		configuration_models.ProfileV2{},
	)
	if err == nil {
		t.Fatal("Expected error when no update field is provided, got nil")
	}
	if err.Error() != "specify at least one field to update" {
		t.Fatalf("Expected update field error, got %v", err)
	}
}
