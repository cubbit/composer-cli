package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestCreateIAMAPIKey_DoesNotSendTokenQueryParam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		if r.Header.Get("Authorization") != "ApiKey test-api-key" {
			t.Fatalf("expected API key authorization header, got %q", r.Header.Get("Authorization"))
		}

		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(OperatorAPIKey{ID: "key-001", Name: "automation", Enabled: true})
	}))
	defer server.Close()

	_, err := (&UserAPI{}).CreateIAMAPIKey(
		configuration_models.EndpointsV2{IAM: server.URL},
		"test-api-key",
		"org-001",
		"operator-001",
		&CreateIAMAPIKeyRequestBody{Name: "automation"},
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteIAMAPIKey_DoesNotSendTokenQueryParam(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		if r.Header.Get("Authorization") != "ApiKey test-api-key" {
			t.Fatalf("expected API key authorization header, got %q", r.Header.Get("Authorization"))
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := (&UserAPI{}).DeleteIAMAPIKey(
		configuration_models.EndpointsV2{IAM: server.URL},
		"test-api-key",
		"org-001",
		"operator-001",
		"key-001",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUpdateIAMAPIKey_UsesPatchV3Endpoint(t *testing.T) {
	expiresAt := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Fatalf("expected PATCH method, got %q", r.Method)
		}
		if r.URL.Path != "/v3/organizations/org-001/operators/operator-001/api-keys/key-001" {
			t.Fatalf("expected v3 API key path, got %q", r.URL.Path)
		}
		if r.URL.RawQuery != "" {
			t.Fatalf("expected no query string, got %q", r.URL.RawQuery)
		}
		if r.Header.Get("Authorization") != "ApiKey test-api-key" {
			t.Fatalf("expected API key authorization header, got %q", r.Header.Get("Authorization"))
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("failed to read request body: %v", err)
		}
		expectedBody := `{"name":"automation-renamed","expires_at":"2024-12-31T23:59:59Z","enabled":false}`
		if string(body) != expectedBody {
			t.Fatalf("expected body %q, got %q", expectedBody, string(body))
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(OperatorAPIKey{ID: "key-001", Name: "automation-renamed", Enabled: false})
	}))
	defer server.Close()

	name := "automation-renamed"
	enabled := false
	_, err := (&UserAPI{}).UpdateIAMAPIKey(
		configuration_models.EndpointsV2{IAM: server.URL},
		"test-api-key",
		"org-001",
		"operator-001",
		"key-001",
		&UpdateIAMAPIKeyRequestBody{
			Name:      &name,
			ExpiresAt: &expiresAt,
			Enabled:   &enabled,
		},
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
