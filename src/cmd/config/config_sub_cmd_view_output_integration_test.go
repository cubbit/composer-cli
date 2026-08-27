package cmd_config

import (
	"bytes"
	"testing"
	"time"

	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service"
)

func TestConfigSubCmd_View_Output_JSON(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockCfg.GetConfigFileFunc = func() (configuration_models.ConfigV2, error) {
		return configuration_models.ConfigV2{
			Version: "v2",
			Active:  configuration_models.ActiveConfigV2{Profile: "default"},
			Profile: map[string]configuration_models.ProfileV2{
				"default": {
					Output:         "human",
					APIKey:         "test-api-key",
					UpdatedAt:      frozen,
					OrganizationID: "test-org-id",
					Endpoints: configuration_models.EndpointsV2{
						IAM:  "https://iam.example.com",
						Dash: "https://dash.example.com",
						CH:   "https://ch.example.com",
					},
				},
			},
		}, nil
	}

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         "human",
			APIKey:         "test-api-key",
			UpdatedAt:      frozen,
			OrganizationID: "test-org-id",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	svc := service.NewConfigService(mockCfg)

	configCmd := NewConfigCmd(svc)
	configCmd.PersistentFlags().String("profile", "", "Profile")
	configCmd.PersistentFlags().String("output", "human", "Output format")
	configCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	configCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"view",
		"--output", "json",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "version": "v2",
  "active": {
    "profile": "default"
  },
  "profile": {
    "default": {
      "output": "human",
      "api_key": "test-api-key",
      "updated_at": "2026-08-19T10:00:00Z",
      "organization_id": "test-org-id",
      "endpoints": {
        "iam": "https://iam.example.com",
        "dash": "https://dash.example.com",
        "ch": "https://ch.example.com"
      }
    }
  }
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestConfigSubCmd_View_Output_YAML(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockCfg.GetConfigFileFunc = func() (configuration_models.ConfigV2, error) {
		return configuration_models.ConfigV2{
			Version: "v2",
			Active:  configuration_models.ActiveConfigV2{Profile: "default"},
			Profile: map[string]configuration_models.ProfileV2{
				"default": {
					Output:         "human",
					APIKey:         "test-api-key",
					UpdatedAt:      frozen,
					OrganizationID: "test-org-id",
					Endpoints: configuration_models.EndpointsV2{
						IAM:  "https://iam.example.com",
						Dash: "https://dash.example.com",
						CH:   "https://ch.example.com",
					},
				},
			},
		}, nil
	}

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         "human",
			APIKey:         "test-api-key",
			UpdatedAt:      frozen,
			OrganizationID: "test-org-id",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	svc := service.NewConfigService(mockCfg)

	configCmd := NewConfigCmd(svc)
	configCmd.PersistentFlags().String("profile", "", "Profile")
	configCmd.PersistentFlags().String("output", "human", "Output format")
	configCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	configCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"view",
		"--output", "yaml",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `version: v2
active:
    profile: default
profile:
    default:
        output: human
        api_key: test-api-key
        updated_at: 2026-08-19T10:00:00Z
        organization_id: test-org-id
        endpoints:
            iam: https://iam.example.com
            dash: https://dash.example.com
            ch: https://ch.example.com
`
	if commandOutput.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestConfigSubCmd_View_Output_JSON_FromProfile(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockCfg.GetConfigFileFunc = func() (configuration_models.ConfigV2, error) {
		return configuration_models.ConfigV2{
			Version: "v2",
			Active:  configuration_models.ActiveConfigV2{Profile: "default"},
			Profile: map[string]configuration_models.ProfileV2{
				"default": {
					Output:         "json",
					APIKey:         "test-api-key",
					UpdatedAt:      frozen,
					OrganizationID: "test-org-id",
					Endpoints: configuration_models.EndpointsV2{
						IAM:  "https://iam.example.com",
						Dash: "https://dash.example.com",
						CH:   "https://ch.example.com",
					},
				},
			},
		}, nil
	}

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         "json",
			APIKey:         "test-api-key",
			UpdatedAt:      frozen,
			OrganizationID: "test-org-id",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	svc := service.NewConfigService(mockCfg)

	configCmd := NewConfigCmd(svc)
	configCmd.PersistentFlags().String("profile", "", "Profile")
	configCmd.PersistentFlags().String("output", "human", "Output format")
	configCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	configCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"view",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := `{
  "version": "v2",
  "active": {
    "profile": "default"
  },
  "profile": {
    "default": {
      "output": "json",
      "api_key": "test-api-key",
      "updated_at": "2026-08-19T10:00:00Z",
      "organization_id": "test-org-id",
      "endpoints": {
        "iam": "https://iam.example.com",
        "dash": "https://dash.example.com",
        "ch": "https://ch.example.com"
      }
    }
  }
}
`
	if commandOutput.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, commandOutput.String())
	}
}

func TestConfigSubCmd_View_Output_Quiet(t *testing.T) {
	mockCfg := configuration_handler.NewMockConfigurationHandler()
	frozen := time.Date(2026, 8, 19, 10, 0, 0, 0, time.UTC)

	mockCfg.GetConfigFileFunc = func() (configuration_models.ConfigV2, error) {
		return configuration_models.ConfigV2{
			Version: "v2",
			Active:  configuration_models.ActiveConfigV2{Profile: "default"},
			Profile: map[string]configuration_models.ProfileV2{
				"default": {
					Output:         "human",
					APIKey:         "test-api-key",
					UpdatedAt:      frozen,
					OrganizationID: "test-org-id",
					Endpoints: configuration_models.EndpointsV2{
						IAM:  "https://iam.example.com",
						Dash: "https://dash.example.com",
						CH:   "https://ch.example.com",
					},
				},
			},
		}, nil
	}

	mockCfg.GetActiveProfileFunc = func() (configuration_models.ProfileV2, error) {
		return configuration_models.ProfileV2{
			Output:         "human",
			APIKey:         "test-api-key",
			UpdatedAt:      frozen,
			OrganizationID: "test-org-id",
			Endpoints: configuration_models.EndpointsV2{
				IAM:  "https://iam.example.com",
				Dash: "https://dash.example.com",
				CH:   "https://ch.example.com",
			},
		}, nil
	}

	svc := service.NewConfigService(mockCfg)

	configCmd := NewConfigCmd(svc)
	configCmd.PersistentFlags().String("profile", "", "Profile")
	configCmd.PersistentFlags().String("output", "human", "Output format")
	configCmd.PersistentFlags().Bool("no-headers", false, "No headers")
	configCmd.PersistentFlags().Bool("quiet", false, "Quiet mode")

	commandOutput := new(bytes.Buffer)
	configCmd.SetOut(commandOutput)
	configCmd.SetErr(commandOutput)
	configCmd.SetArgs([]string{
		"view",
		"--quiet",
	})

	err := configCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if commandOutput.String() != "" {
		t.Errorf("Expected empty output in quiet mode, got %q", commandOutput.String())
	}
}
