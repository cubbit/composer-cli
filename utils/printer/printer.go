package printer

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/text"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// OutputFormat represents the desired output serialization format.
type OutputFormat int

const (
	OutputHuman OutputFormat = iota
	OutputJSON
	OutputYAML
	OutputQuiet
)

// resolveFormat resolves the output format with the following priority:
// 1. --quiet flag (OutputQuiet)
// 2. --output flag when explicitly set (OutputJSON / OutputYAML)
// 3. handler-provided default from the active profile config
// 4. OutputHuman as ultimate fallback
func resolveFormat(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface) (OutputFormat, error) {
	f := cmd.Flag("quiet")
	if f != nil && f.Value.String() == "true" {
		return OutputQuiet, nil
	}

	if cmd.Flags().Changed("output") {
		f = cmd.Flag("output")
		switch f.Value.String() {
		case "json":
			return OutputJSON, nil
		case "yaml":
			return OutputYAML, nil
		}
		return OutputHuman, nil
	}

	if handler != nil {
		if profile, err := handler.GetActiveProfile(); err == nil {
			if o := outputFromString(string(profile.Output)); o != OutputHuman {
				return o, nil
			}
		}
	}

	return OutputHuman, nil
}

// outputFromString converts a string format to an OutputFormat.
// Returns OutputHuman for unrecognised values.
func outputFromString(s string) OutputFormat {
	switch s {
	case "json":
		return OutputJSON
	case "yaml":
		return OutputYAML
	default:
		return OutputHuman
	}
}

// writeJSON marshals data to indented JSON and writes it to w.
func writeJSON(w io.Writer, data any) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}
	if string(b) == "null" {
		return nil
	}
	_, err = fmt.Fprintln(w, string(b))
	return err
}

// writeYAML marshals data to YAML and writes it to w.
func writeYAML(w io.Writer, data any) error {
	b, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
	}
	_, err = fmt.Fprint(w, string(b))
	return err
}

func ShowSecrets(cmd *cobra.Command) bool {
	f := cmd.Flag("show-secrets")
	if f == nil {
		return false
	}
	return f.Value.String() == "true"
}

// PrintTree renders data as a human-friendly tree when output is human.
// When output is json or yaml it serialises data instead.
// buildNodes converts data into the tree nodes used for human-readable output.
func PrintTree[T any](cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, data T, buildNodes func(T) []tree.TreeNode, opts ...tree.Option) error {
	f, err := resolveFormat(cmd, handler)
	if err != nil || f == OutputQuiet {
		return err
	}

	switch f {
	case OutputJSON:
		return writeJSON(cmd.OutOrStdout(), data)
	case OutputYAML:
		return writeYAML(cmd.OutOrStdout(), data)
	default:
		result := tree.CreateTree(buildNodes(data), opts...)
		_, err = fmt.Fprint(cmd.OutOrStdout(), result)
		return err
	}
}

// PrintTable renders data as a lipgloss table when output is human.
// When output is json or yaml it serialises data instead.
func PrintTable[T any](cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, data []T, opts ...table.Option[T]) error {
	f, err := resolveFormat(cmd, handler)
	if err != nil || f == OutputQuiet {
		return err
	}

	switch f {
	case OutputJSON:
		return writeJSON(cmd.OutOrStdout(), data)
	case OutputYAML:
		return writeYAML(cmd.OutOrStdout(), data)
	default:
		result := table.CreateTable(data, opts...)
		_, err = fmt.Fprint(cmd.OutOrStdout(), result)
		return err
	}
}

// PrintText renders s as styled text when output is human.
// When output is json or yaml it wraps s in a {"message": s} envelope.
func PrintText(cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, s string, opts ...text.Option) error {
	f, err := resolveFormat(cmd, handler)
	if err != nil || f == OutputQuiet {
		return err
	}

	switch f {
	case OutputJSON:
		return writeJSON(cmd.OutOrStdout(), map[string]string{"message": s})
	case OutputYAML:
		return writeYAML(cmd.OutOrStdout(), map[string]string{"message": s})
	default:
		result := text.CreateText(s, opts...)
		_, err = fmt.Fprint(cmd.OutOrStdout(), result)
		return err
	}
}

// ComposeStructured runs printFuncs when output is human. When output
// is json or yaml it serialises data instead, skipping the printFuncs
// entirely.
func ComposeStructured[T any](cmd *cobra.Command, handler configuration_handler.ConfigurationHandlerInterface, data T, printFuncs ...func() error) error {
	f, err := resolveFormat(cmd, handler)
	if err != nil || f == OutputQuiet {
		return err
	}

	switch f {
	case OutputJSON:
		return writeJSON(cmd.OutOrStdout(), data)
	case OutputYAML:
		return writeYAML(cmd.OutOrStdout(), data)
	default:
		for _, fn := range printFuncs {
			if err := fn(); err != nil {
				return err
			}
		}
		return nil
	}
}
