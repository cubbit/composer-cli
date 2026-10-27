package cmd_operator

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

type mockOperatorAPI struct {
	ConnectFunc func(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) (string, error)
}

func (m *mockOperatorAPI) Connect(endpoints configuration_models.EndpointsV2, apiKey string, organizationID string) (string, error) {
	if m.ConnectFunc != nil {
		return m.ConnectFunc(endpoints, apiKey, organizationID)
	}
	return "", nil
}

var _ api.OperatorAPIInterface = (*mockOperatorAPI)(nil)

func TestOperatorSubCmd_Connect_Output_JSON(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH: "https://ch.example.com",
			},
		}, nil
	}

	mockOpAPI := &mockOperatorAPI{
		ConnectFunc: func(_ configuration_models.EndpointsV2, apiKey, orgID string) (string, error) {
			return "kubectl apply -f https://example.com/operator.yaml", nil
		},
	}

	operatorService := service.NewOperatorService(mockCfg, mockOpAPI, &api.MockUserAPI{})
	operatorCmd := NewOperatorCmd(operatorService)
	operatorCmd.PersistentFlags().String("output", "human", "Output format")
	operatorCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	operatorCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	operatorCmd.SetOut(buf)
	operatorCmd.SetErr(buf)
	operatorCmd.SetArgs([]string{"generate-connect-command", "--output", "json"})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `"kubectl apply -f https://example.com/operator.yaml"
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestOperatorSubCmd_Connect_Output_YAML(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH: "https://ch.example.com",
			},
		}, nil
	}

	mockOpAPI := &mockOperatorAPI{
		ConnectFunc: func(_ configuration_models.EndpointsV2, apiKey, orgID string) (string, error) {
			return "kubectl apply -f https://example.com/operator.yaml", nil
		},
	}

	operatorService := service.NewOperatorService(mockCfg, mockOpAPI, &api.MockUserAPI{})
	operatorCmd := NewOperatorCmd(operatorService)
	operatorCmd.PersistentFlags().String("output", "human", "Output format")
	operatorCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	operatorCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	operatorCmd.SetOut(buf)
	operatorCmd.SetErr(buf)
	operatorCmd.SetArgs([]string{"generate-connect-command", "--output", "yaml"})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `kubectl apply -f https://example.com/operator.yaml

`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestOperatorSubCmd_Connect_Output_Human(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         configuration_models.OutputHuman,
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH: "https://ch.example.com",
			},
		}, nil
	}

	mockOpAPI := &mockOperatorAPI{
		ConnectFunc: func(_ configuration_models.EndpointsV2, apiKey, orgID string) (string, error) {
			return "kubectl apply -f https://example.com/operator.yaml", nil
		},
	}

	operatorService := service.NewOperatorService(mockCfg, mockOpAPI, &api.MockUserAPI{})
	operatorCmd := NewOperatorCmd(operatorService)
	operatorCmd.PersistentFlags().String("output", "human", "Output format")
	operatorCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	operatorCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	operatorCmd.SetOut(buf)
	operatorCmd.SetErr(buf)
	operatorCmd.SetArgs([]string{"generate-connect-command", "--output", "human"})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `kubectl apply -f https://example.com/operator.yaml
`
	if buf.String() != expected {
		t.Errorf("Human snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestOperatorSubCmd_Connect_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "json",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH: "https://ch.example.com",
			},
		}, nil
	}

	mockOpAPI := &mockOperatorAPI{
		ConnectFunc: func(_ configuration_models.EndpointsV2, apiKey, orgID string) (string, error) {
			return "kubectl apply -f https://example.com/operator.yaml", nil
		},
	}

	operatorService := service.NewOperatorService(mockCfg, mockOpAPI, &api.MockUserAPI{})
	operatorCmd := NewOperatorCmd(operatorService)
	operatorCmd.PersistentFlags().String("output", "human", "Output format")
	operatorCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	operatorCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	operatorCmd.SetOut(buf)
	operatorCmd.SetErr(buf)
	operatorCmd.SetArgs([]string{"generate-connect-command", "--output", "json"})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `"kubectl apply -f https://example.com/operator.yaml"
`
	if buf.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

func TestOperatorSubCmd_Connect_Output_YAML_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			APIKey:         "test-api-key",
			OrganizationID: "test-org-id",
			Output:         "yaml",
			Endpoints: configuration_models.EndpointsV2{
				IAM: "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH: "https://ch.example.com",
			},
		}, nil
	}

	mockOpAPI := &mockOperatorAPI{
		ConnectFunc: func(_ configuration_models.EndpointsV2, apiKey, orgID string) (string, error) {
			return "kubectl apply -f https://example.com/operator.yaml", nil
		},
	}

	operatorService := service.NewOperatorService(mockCfg, mockOpAPI, &api.MockUserAPI{})
	operatorCmd := NewOperatorCmd(operatorService)
	operatorCmd.PersistentFlags().String("output", "human", "Output format")
	operatorCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	operatorCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	buf := new(bytes.Buffer)
	operatorCmd.SetOut(buf)
	operatorCmd.SetErr(buf)
	operatorCmd.SetArgs([]string{"generate-connect-command", "--output", "yaml"})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `kubectl apply -f https://example.com/operator.yaml

`
	if buf.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, buf.String())
	}
}

