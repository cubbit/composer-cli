package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
)

func TestResetIAMUserPassword_UsesV3EndpointAndIAMPasswordChangeTokenBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("expected POST method, got %q", r.Method)
		}
		if r.URL.Path != "/v3/organizations/org-001/operators/operator-001/password/reset" {
			t.Fatalf("expected v3 password reset path, got %q", r.URL.Path)
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
		expectedBody := `{"iam_password_change_token":"reset-token","authentication_public_key":"public-key"}`
		if string(body) != expectedBody {
			t.Fatalf("expected body %q, got %q", expectedBody, string(body))
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	err := (&UserAPI{}).ResetIAMUserPassword(
		configuration_models.EndpointsV2{IAM: server.URL},
		"test-api-key",
		"org-001",
		"operator-001",
		&ResetIAMUserPasswordRequestBody{
			IAMPasswordChangeToken:  "reset-token",
			AuthenticationPublicKey: "public-key",
		},
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
