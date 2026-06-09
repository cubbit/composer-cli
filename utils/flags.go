package utils

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/spf13/cobra"
)

func GetOptionalStringFlag(cmd *cobra.Command, flagName string) (*string, error) {
	if cmd.Flags().Changed(flagName) {
		val, err := cmd.Flags().GetString(flagName)
		if err != nil {
			return nil, fmt.Errorf("%s %s: %w", constants.ErrorRetrievingField, flagName, err)
		}

		return &val, nil
	}

	return nil, nil
}
