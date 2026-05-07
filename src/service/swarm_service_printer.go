package service

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/utils"
	"github.com/spf13/cobra"
)

func PrintSwarmDetails(cmd *cobra.Command, swarm api.SwarmV5Presentation) error {
	return printer.PrintText(cmd, buildSwarmDetailsOutput(swarm))
}

func buildSwarmDetailsOutput(swarm api.SwarmV5Presentation) string {
	description := "N/A"
	if swarm.Description != nil && *swarm.Description != "" {
		description = *swarm.Description
	}

	creationStatus := swarm.CreationStatus
	if creationStatus == "" {
		creationStatus = "N/A"
	}

	evaluatedStatus := "N/A"
	if swarm.EvaluatedStatus != nil {
		evaluatedStatus = string(*swarm.EvaluatedStatus)
	}

	evaluatedStatusUpdatedAt := "N/A"
	if swarm.EvaluatedStatusLastUpdatedAt != nil {
		evaluatedStatusUpdatedAt = utils.FormatTime(*swarm.EvaluatedStatusLastUpdatedAt)
	}

	availableStorage := swarm.TotalStorageBytes - swarm.UsedStorageBytes
	if availableStorage < 0 {
		availableStorage = 0
	}

	lines := []string{
		fmt.Sprintf("Swarm: %s", swarm.Name),
		fmt.Sprintf("Status: %s", formatSwarmStatus(evaluatedStatus)),
		fmt.Sprintf("Last Update: %s", evaluatedStatusUpdatedAt),
		fmt.Sprintf("Organization: %s", swarm.OrganizationID),
		fmt.Sprintf("Owner: %s", swarm.OwnerID),
		"",
		"Storage Usage:",
		fmt.Sprintf("  Usage: %s %s", formatStorageBar(swarm.UsedStorageBytes, swarm.TotalStorageBytes, 10), formatStoragePercent(swarm.UsedStorageBytes, swarm.TotalStorageBytes)),
		fmt.Sprintf("  Total Used: %s", utils.FormatBytes(swarm.UsedStorageBytes)),
		fmt.Sprintf("  Total Assigned: %s", utils.FormatBytes(swarm.TotalStorageBytes)),
		fmt.Sprintf("  Total Unused: %s", utils.FormatBytes(availableStorage)),
		"",
		"Metadata:",
		fmt.Sprintf("  ID: %s", swarm.ID),
		fmt.Sprintf("  Description: %s", description),
		fmt.Sprintf("  Created At: %s", utils.FormatTime(swarm.CreatedAt)),
		fmt.Sprintf("  Creation Status: %s", creationStatus),
		"",
		"Composition:",
		fmt.Sprintf("  Nexus Count: %d", swarm.NexusCount),
		fmt.Sprintf("  Redundancy Class Count: %d", swarm.RedundancyClassCount),
		"",
		"Configuration:",
	}

	configurationLines := formatSwarmConfiguration(swarm.Configuration)
	for _, line := range configurationLines {
		lines = append(lines, "  "+line)
	}

	return strings.Join(lines, "\n")
}

func formatSwarmStatus(status string) string {
	if status == "" || status == "N/A" {
		return "N/A"
	}

	return "● " + strings.ToUpper(status[:1]) + status[1:]
}

func formatStorageBar(usedBytes, totalBytes int64, width int) string {
	if width <= 0 {
		width = 10
	}

	if totalBytes <= 0 {
		return "[" + strings.Repeat("░", width) + "]"
	}

	usedRatio := float64(usedBytes) / float64(totalBytes)
	if usedRatio < 0 {
		usedRatio = 0
	}
	if usedRatio > 1 {
		usedRatio = 1
	}

	filled := int(usedRatio * float64(width))
	if filled > width {
		filled = width
	}

	return "[" + strings.Repeat("█", filled) + strings.Repeat("░", width-filled) + "]"
}

func formatStoragePercent(usedBytes, totalBytes int64) string {
	if totalBytes <= 0 {
		return "0%"
	}

	percent := int((float64(usedBytes) / float64(totalBytes)) * 100)
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	return fmt.Sprintf("%d%%", percent)
}

func formatSwarmConfiguration(configuration map[string]interface{}) []string {
	if len(configuration) == 0 {
		return []string{"N/A"}
	}

	keys := make([]string, 0, len(configuration))
	for key := range configuration {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		lines = append(lines, fmt.Sprintf("%s: %s", key, formatConfigValue(configuration[key])))
	}

	return lines
}

func formatConfigValue(value interface{}) string {
	if value == nil {
		return "null"
	}

	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Map:
		keys := rv.MapKeys()
		keyStrings := make([]string, 0, len(keys))
		keyMap := make(map[string]reflect.Value, len(keys))
		for _, key := range keys {
			keyStr := fmt.Sprintf("%v", key.Interface())
			keyStrings = append(keyStrings, keyStr)
			keyMap[keyStr] = rv.MapIndex(key)
		}
		sort.Strings(keyStrings)

		parts := make([]string, 0, len(keyStrings))
		for _, key := range keyStrings {
			parts = append(parts, fmt.Sprintf("%s: %s", key, formatConfigValue(keyMap[key].Interface())))
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case reflect.Slice, reflect.Array:
		parts := make([]string, 0, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			parts = append(parts, formatConfigValue(rv.Index(i).Interface()))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	default:
		return fmt.Sprintf("%v", value)
	}
}
