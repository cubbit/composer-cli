package describe

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

func TestDescribeAPIKey_UsesGetByID(t *testing.T) {
	createdAt := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	organizationName := "test-org"
	userAPI := &api.MockUserAPI{
		GetIAMUserSelfFunc: func(endpoints configuration_models.EndpointsV2, accessToken string, apiKey string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:               "operator-001",
				Username:         "admin",
				OrganizationName: &organizationName,
			}, nil
		},
		GetIAMAPIKeyByIDFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			apiKeyID string,
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

			return &api.OperatorAPIKey{
				ID:         "key-001",
				Name:       "automation",
				OperatorID: "operator-001",
				CreatedAt:  createdAt,
				Enabled:    true,
			}, nil
		},
		ListIAMAPIKeysFunc: func(
			endpoints configuration_models.EndpointsV2,
			apiKey string,
			organizationID string,
			operatorID string,
			page int,
			items int,
			sortKey string,
			sortOrder string,
		) (*api.GenericPaginatedResponse[api.OperatorAPIKey], error) {
			t.Fatal("DescribeAPIKey should not list API keys")
			return nil, nil
		},
	}

	cmd := &cobra.Command{}
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.Flags().String("id", "key-001", "ID")
	cmd.Flags().String("output", "human", "Output")
	cmd.Flags().Bool("quiet", false, "Quiet")

	err := DescribeAPIKey(
		shared.Dependencies{UserAPI: userAPI},
		cmd,
		nil,
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
API Key: automation
Enabled: true

Metadata:
  ID: key-001
  Operator ID: operator-001
  Created At: 2024-01-15 10:30:00
  Expires At: N/A
  Deleted At: N/A
  Banned At: N/A
`)

	actualResult := strings.TrimSpace(output.String())
	if actualResult != expectedResult {
		t.Fatalf("Expected IAM API key describe output does not match actual output\nExpected:\n%s\nActual:\n%s", expectedResult, actualResult)
	}
}
