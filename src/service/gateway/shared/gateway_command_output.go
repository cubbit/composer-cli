package shared

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func ResolveCommandOutput(cmd *cobra.Command, defaultOutput configuration_models.OutputFormat) (string, error) {
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
