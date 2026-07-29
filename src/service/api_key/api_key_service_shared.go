package apikey

import (
	"fmt"
	"time"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func resolveOutput(cmd *cobra.Command, defaultOutput configuration_models.OutputFormat) (string, error) {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return "", fmt.Errorf("%s output: %w", constants.ErrorRetrievingField, err)
	}
	if defaultOutput != "" &&
		!cmd.Flags().Changed("output") &&
		!cmd.Flags().Changed("quiet") {
		output = string(defaultOutput)
	}
	return output, nil
}

func parseOptionalExpiresAt(cmd *cobra.Command) (*time.Time, error) {
	expiresAt, err := cmd.Flags().GetString("expires-at")
	if err != nil {
		return nil, fmt.Errorf("%s expires-at: %w", constants.ErrorRetrievingField, err)
	}
	if expiresAt == "" {
		return nil, nil
	}

	parsed, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("invalid value for --expires-at: expected RFC3339 timestamp, got %q", expiresAt)
	}

	return &parsed, nil
}
