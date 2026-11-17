package shared

import (
	"errors"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestResolveUserIDByID_UsesOrganizationScopedLookup(t *testing.T) {
	profile := testAPIKeyProfile()
	userAPI := &api.MockUserAPI{
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			if apiKey != "test-api-key" {
				t.Fatalf("Expected API key %q, got %q", "test-api-key", apiKey)
			}
			if organizationID != "test-org-id" {
				t.Fatalf("Expected organization ID %q, got %q", "test-org-id", organizationID)
			}
			if userID != "operator-002" {
				t.Fatalf("Expected user ID %q, got %q", "operator-002", userID)
			}

			return &api.IAMUser{ID: userID}, nil
		},
	}

	userID, err := ResolveUserIDByID(userAPI, profile, "operator-002")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if userID != "operator-002" {
		t.Fatalf("Expected user ID %q, got %q", "operator-002", userID)
	}
}

func TestResolveUserIDByID_RejectsMissingUser(t *testing.T) {
	userAPI := &api.MockUserAPI{
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			return &api.IAMUser{}, nil
		},
	}

	_, err := ResolveUserIDByID(userAPI, testAPIKeyProfile(), "operator-002")
	if err == nil {
		t.Fatal("Expected missing user error, got nil")
	}
	if err.Error() != "IAM user with ID 'operator-002' not found" {
		t.Fatalf("Expected missing user error, got %v", err)
	}
}

func TestResolveUserIDByID_RejectsDifferentOrganization(t *testing.T) {
	otherOrganizationID := "other-org-id"
	userAPI := &api.MockUserAPI{
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			return &api.IAMUser{
				ID:             userID,
				OrganizationID: &otherOrganizationID,
			}, nil
		},
	}

	_, err := ResolveUserIDByID(userAPI, testAPIKeyProfile(), "operator-002")
	if err == nil {
		t.Fatal("Expected organization mismatch error, got nil")
	}
	if err.Error() != "IAM user with ID 'operator-002' not found in organization 'test-org-id'" {
		t.Fatalf("Expected organization mismatch error, got %v", err)
	}
}

func TestResolveUserIDByID_WrapsLookupError(t *testing.T) {
	userAPI := &api.MockUserAPI{
		GetIAMUserByIDFunc: func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string, userID string) (*api.IAMUser, error) {
			return nil, errors.New("not found")
		},
	}

	_, err := ResolveUserIDByID(userAPI, testAPIKeyProfile(), "operator-002")
	if err == nil {
		t.Fatal("Expected lookup error, got nil")
	}
	if err.Error() != "failed to resolve user-id 'operator-002': not found" {
		t.Fatalf("Expected wrapped lookup error, got %v", err)
	}
}

func testAPIKeyProfile() configuration_models.ProfileV2 {
	return configuration_models.ProfileV2{
		APIKey:         "test-api-key",
		OrganizationID: "test-org-id",
		Endpoints: configuration_models.EndpointsV2{
			IAM: "https://iam.example.com",
		},
	}
}
