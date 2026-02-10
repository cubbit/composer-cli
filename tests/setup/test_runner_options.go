package tests_setup

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func WithConfigSetup(config string, t *testing.T) RunnerOption {
	return func(base *RunnerBase) error {
		tempDir := t.TempDir()
		tempDirCubbit := filepath.Join(tempDir, "cubbit")
		configPath := filepath.Join(tempDirCubbit, "config.toml")

		err := os.MkdirAll(tempDirCubbit, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directories for config setup: %w", err)
		}

		err = os.WriteFile(configPath, []byte(config), 0644)
		if err != nil {
			return fmt.Errorf("failed to create config file: %v", err)
		}

		t.Setenv("XDG_CONFIG_HOME", tempDir)

		return nil
	}
}

var developmentConfig = `
	[default]
	output = "human"

	[active]
	profile = "dev"

	[profile.dev]
	inherits = "default"
	endpoint = ""
	type = "composer"
	output = "human"
	updated_at = "2026-01-28T12:00:00Z"

	[profile.dev.urls]
	base_url = "http://localhost"
	iam = "http://localhost:8181"
	dash = "http://localhost:3001"
	ch = "http://localhost:8380"
`

func WithDevelopmentConfig(t *testing.T) RunnerOption {
	return WithConfigSetup(developmentConfig, t)
}
