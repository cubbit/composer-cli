package routing_rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

// gatewayIDFromFlags reads and validates the required --gateway-id flag.
func gatewayIDFromFlags(cmd *cobra.Command) (string, error) {
	gatewayID, err := cmd.Flags().GetString("gateway-id")
	if err != nil {
		return "", fmt.Errorf("%s gateway-id: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(gatewayID) == "" {
		return "", fmt.Errorf("--gateway-id must not be empty")
	}

	return strings.TrimSpace(gatewayID), nil
}

// ruleIDFromFlags reads and validates the required --id flag.
func ruleIDFromFlags(cmd *cobra.Command) (string, error) {
	ruleID, err := cmd.Flags().GetString("id")
	if err != nil {
		return "", fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(ruleID) == "" {
		return "", fmt.Errorf("--id must not be empty")
	}

	return normalizeRuleID(ruleID)
}

// List prints the routing rules of a gateway. A gateway can have zero or more
// rules; the list is metadata only, use get for the full document of a rule.
func List(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.GatewayAPI.ListGatewayRoutingRulesV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
	)
	if err != nil {
		return fmt.Errorf("failed to list routing rules: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, response,
		func() error { return PrintRoutingRuleList(cmd, handler, *response) },
	)
}

// Get prints a single routing rule selected by --id. Structured output is the
// raw server document, so fields this CLI does not model are preserved.
func Get(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	ruleID, err := ruleIDFromFlags(cmd)
	if err != nil {
		return err
	}

	rawRule, err := deps.GatewayAPI.GetGatewayRoutingRuleV5Raw(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		ruleID,
	)
	if err != nil {
		return fmt.Errorf("failed to get routing rule: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, ruleDocument{raw: rawRule},
		func() error { return PrintRoutingRuleDocument(cmd, handler, rawRule) },
	)
}

// Create creates a routing rule from the JSON document read from --file-path
// (or stdin for "-"). Unknown fields are forwarded to the server verbatim.
func Create(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	_, raw, err := readRuleDocument(cmd)
	if err != nil {
		return err
	}

	created, err := deps.GatewayAPI.CreateGatewayRoutingRuleV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		raw,
	)
	if err != nil {
		return fmt.Errorf("failed to create routing rule: %w", err)
	}

	rule, err := decodeRawRule(created)
	if err != nil {
		return err
	}

	return PrintRoutingRuleCreated(cmd, handler, rule.ID)
}

// Update replaces a routing rule with the JSON document read from --file-path
// (or stdin for "-"). The raw document is forwarded verbatim so unmodeled
// fields obtained from get survive the round-trip.
func Update(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	document, raw, err := readRuleDocument(cmd)
	if err != nil {
		return err
	}

	ruleID, err := resolveInputRuleID(cmd, document)
	if err != nil {
		return err
	}

	if _, err := deps.GatewayAPI.UpdateGatewayRoutingRuleV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		ruleID,
		raw,
	); err != nil {
		return fmt.Errorf("failed to update routing rule: %w", err)
	}

	return PrintRoutingRuleUpdated(cmd, handler, ruleID)
}

// Delete removes a routing rule from a gateway.
func Delete(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	ruleID, err := ruleIDFromFlags(cmd)
	if err != nil {
		return err
	}

	if err := deps.GatewayAPI.DeleteGatewayRoutingRuleV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		ruleID,
	); err != nil {
		return fmt.Errorf("failed to delete routing rule: %w", err)
	}

	return printer.PrintText(cmd, handler, buildRoutingRuleDeletedOutput(ruleID))
}

func decodeRawRule(raw json.RawMessage) (*api.GatewayRoutingRuleV5, error) {
	var rule api.GatewayRoutingRuleV5
	if err := json.Unmarshal(raw, &rule); err != nil {
		return nil, fmt.Errorf("failed to parse routing rule JSON: %w", err)
	}

	return &rule, nil
}

// readRuleDocument reads and validates the routing rule JSON document from
// --file-path, or from stdin when the flag value is "-". It returns both the
// typed document (for validation and human output) and the original raw bytes
// (forwarded verbatim to the server so unmodeled fields are preserved).
func readRuleDocument(cmd *cobra.Command) (*api.GatewayRoutingRuleV5, json.RawMessage, error) {
	filePath, err := cmd.Flags().GetString("file-path")
	if err != nil {
		return nil, nil, fmt.Errorf("%s file-path: %w", constants.ErrorRetrievingField, err)
	}
	filePath = strings.TrimSpace(filePath)

	if filePath == "" {
		return nil, nil, fmt.Errorf("--file-path must not be empty")
	}

	var data []byte
	if filePath == "-" {
		data, err = io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read routing rule from stdin: %w", err)
		}
	} else {
		data, err = os.ReadFile(filePath)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to read file %q: %w", filePath, err)
		}
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return nil, nil, fmt.Errorf("routing rule input is empty")
	}

	// The document must be a JSON object; reject arrays and scalars early with
	// a clear message instead of a decoder type error.
	if firstByte := bytes.TrimSpace(data)[0]; firstByte != '{' {
		return nil, nil, fmt.Errorf("failed to parse routing rule JSON: document must be a JSON object")
	}

	// Unknown fields are deliberately allowed: the raw document is forwarded
	// to the server, so fields this CLI does not model are not dropped.
	decoder := json.NewDecoder(bytes.NewReader(data))

	var document api.GatewayRoutingRuleV5
	if err := decoder.Decode(&document); err != nil {
		return nil, nil, fmt.Errorf("failed to parse routing rule JSON: %w", err)
	}

	// Reject trailing JSON values or garbage: {..}{..} must not be silently
	// accepted as a single document.
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return nil, nil, fmt.Errorf("failed to parse routing rule JSON: unexpected trailing data")
	}

	if err := validateRuleDocument(&document); err != nil {
		return nil, nil, err
	}

	return &document, json.RawMessage(data), nil
}

// validateRuleDocument rejects obviously incomplete documents before any API
// call. Deeper SDP validation is server-side and is surfaced from the 422
// response.
func validateRuleDocument(document *api.GatewayRoutingRuleV5) error {
	if strings.TrimSpace(document.Name) == "" {
		return fmt.Errorf("routing rule name is required")
	}
	if strings.TrimSpace(document.InputStorageClass) == "" {
		return fmt.Errorf("routing rule inputStorageClass is required")
	}
	if strings.TrimSpace(document.Action.Type) == "" {
		return fmt.Errorf("routing rule action.type is required")
	}

	return nil
}

// resolveInputRuleID returns the normalized rule id from the required --id flag
// and checks that it matches the id carried by the input document when the
// document has one.
func resolveInputRuleID(cmd *cobra.Command, document *api.GatewayRoutingRuleV5) (string, error) {
	ruleID, err := ruleIDFromFlags(cmd)
	if err != nil {
		return "", err
	}

	documentID := strings.TrimSpace(document.ID)
	if documentID == "" {
		return ruleID, nil
	}

	normalizedDocumentID, err := normalizeRuleID(documentID)
	if err != nil {
		return "", err
	}
	if normalizedDocumentID != ruleID {
		return "", fmt.Errorf(
			"--id %s does not match the input document id %s",
			ruleID,
			normalizedDocumentID,
		)
	}

	return ruleID, nil
}

func normalizeRuleID(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "", fmt.Errorf("routing rule id must not be empty")
	}

	parsed, err := uuid.Parse(trimmed)
	if err != nil {
		return "", fmt.Errorf("routing rule id must be a valid UUID, got %q", trimmed)
	}

	return parsed.String(), nil
}

// ruleDocument serialises the raw server document verbatim for JSON and YAML
// output, so fields this CLI does not model are preserved.
type ruleDocument struct {
	raw json.RawMessage
}

func (document ruleDocument) MarshalJSON() ([]byte, error) {
	return document.raw, nil
}

// MarshalYAML converts the raw JSON document into a generic value first, so
// unmodeled fields survive YAML output instead of being dropped by the typed
// struct. Numbers are decoded as json.Number and re-typed, so large integers
// are not rounded through float64.
func (document ruleDocument) MarshalYAML() (any, error) {
	decoded, err := decodeJSONDocument(document.raw)
	if err != nil {
		return nil, fmt.Errorf("failed to encode routing rule YAML: %w", err)
	}

	return decoded, nil
}

// decodeJSONDocument decodes a JSON document into generic Go values while
// preserving number tokens exactly: integers stay int64/uint64 instead of
// being rounded to float64.
func decodeJSONDocument(raw []byte) (any, error) {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()

	var decoded any
	if err := decoder.Decode(&decoded); err != nil {
		return nil, err
	}

	return convertJSONNumbers(decoded), nil
}

func convertJSONNumbers(value any) any {
	switch typed := value.(type) {
	case map[string]any:
		for key, item := range typed {
			typed[key] = convertJSONNumbers(item)
		}

		return typed
	case []any:
		for index, item := range typed {
			typed[index] = convertJSONNumbers(item)
		}

		return typed
	case json.Number:
		if number, err := typed.Int64(); err == nil {
			return number
		}
		if number, err := strconv.ParseUint(typed.String(), 10, 64); err == nil {
			return number
		}
		if number, err := typed.Float64(); err == nil {
			return number
		}

		return typed.String()
	default:
		return value
	}
}
