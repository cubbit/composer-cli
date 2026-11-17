package utils

import (
	"encoding/json"
	"fmt"
)

func GetStructSchema(schemaTemplate map[string]interface{}) string {
	bytes, err := json.MarshalIndent(schemaTemplate, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error generating schema: %v", err)
	}

	return string(bytes)
}
