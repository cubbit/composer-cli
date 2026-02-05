package utils

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

type JSONMap map[string]interface{}

type JSONMapValidator func(JSONMap) error

func JSONMapFromCommand(cmd *cobra.Command, flagName string, validators ...JSONMapValidator) (JSONMap, error) {
	if !cmd.Flags().Changed(flagName) {
		return nil, nil
	}

	rawStr, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return nil, err
	}

	var jm JSONMap
	if err := jm.Set(rawStr); err != nil {
		return nil, fmt.Errorf("failed to parse flag '%s': %w", flagName, err)
	}

	for _, validate := range validators {
		if err := validate(jm); err != nil {
			return nil, fmt.Errorf("validation for flag '%s' failed: %w", flagName, err)
		}
	}

	return jm, nil
}

func (j *JSONMap) String() string {
	b, _ := json.Marshal(*j)
	return string(b)
}

func (j *JSONMap) Set(value string) error {
	if err := json.Unmarshal([]byte(value), j); err != nil {
		return fmt.Errorf("invalid json map: %w", err)
	}
	return nil
}

func (j *JSONMap) Type() string {
	return "json-map"
}
