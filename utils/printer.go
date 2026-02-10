package utils

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/spf13/cobra"
)

type OutputMode int

const (
	OutputQuiet OutputMode = iota
	OutputHuman
)

var (
	currentMode = OutputHuman
)

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	RedBg     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	greenBg   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF80"))
	yellowBg  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFBA00"))
	grayStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	blueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0080FF"))
)

type SmartOutputConfig[T any] struct {
	SingleResource              bool
	SingleResourceCompactOutput bool
	DefaultOutput               configuration.OutputFormat
}

func SetOutputMode(mode OutputMode) {
	currentMode = mode
}

func IsQuietMode() bool {
	return currentMode == OutputQuiet
}

func IsHumanMode() bool {
	return currentMode == OutputHuman
}

func PrintSuccess(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s ✨🐝 %s\n", style("SUCCESS", greenBg), s)
	default:
		return
	}
}

func PrintInfo(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s %s\n", style("INFO", blueStyle), s)
	default:
		return
	}
}

func PrintWarn(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s %s\n", style("WARN", yellowBg), s)
	default:
		return
	}
}

func PrintHint(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s 💡 %s\n", style("HINT", grayStyle), s)
	default:
		return
	}
}

func PrintEmptyLine() {
	switch currentMode {
	case OutputHuman:
		fmt.Println()
	default:
		return
	}
}

func PrintCreateSuccess(resourceType, id string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s ✨🐝 %s %s created\n",
			style("SUCCESS", greenBg),
			style(resourceType, boldStyle),
			style(id, blueStyle))
	default:
		return
	}
}

func PrintDelete(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("🗑️ 🚮 %s\n", s)
	default:
		return
	}
}

func PrintError(err error) {
	PrintErrorWithWriter(os.Stderr, err)
}

func PrintErrorWithWriter(writer io.Writer, err error) {
	errStr, _ := strings.CutSuffix(err.Error(), "\n")
	lines := strings.Split(errStr, "\n")

	switch currentMode {
	case OutputHuman:
		for i, line := range lines {
			if i == 0 {
				l, _ := strings.CutSuffix(line, ": ")
				fmt.Fprintf(writer, "%s %s\n", style("ERR", RedBg), l)
			} else {
				fmt.Fprintf(writer, "%s %s\n", style("INF", blueStyle), line)
			}
		}
	default:
		fmt.Fprintf(writer, "%s\n", err.Error())
	}
}

func PrintNotFound(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("%s %s\n", style("WRN", yellowBg), s)
	default:
		return
	}
}

func PrintEmptyList() {
	switch currentMode {
	case OutputHuman:
		fmt.Print("🪣 [ ]\n")
	default:
		return
	}
}

func PrintList(s string) {
	switch currentMode {
	case OutputHuman:
		fmt.Printf("📋 %s\n", style(s, boldStyle))
	default:
		return
	}
}

func PrintSimpleList(items []string) {
	switch currentMode {
	case OutputHuman:
		for _, item := range items {
			fmt.Printf(" • %s\n", item)
		}
	default:
		for _, item := range items {
			fmt.Println(item)
		}
	}
}

func PrintVerbose(writer io.Writer, data interface{}, noHeaders bool) {
	printMarkdownTable(writer, data, noHeaders)
}

func printMarkdownTable(writer io.Writer, data interface{}, noHeaders bool) {
	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}

	var rows [][]string
	var headers []string

	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		if value.Len() == 0 {
			return
		}

		firstElem := value.Index(0)

		if firstElem.Kind() == reflect.Ptr {
			if firstElem.IsNil() {
				firstElem = reflect.New(firstElem.Type().Elem()).Elem()
			} else {
				firstElem = firstElem.Elem()
			}
		}

		if firstElem.Kind() == reflect.Struct {
			elemType := firstElem.Type()
			for i := 0; i < elemType.NumField(); i++ {
				field := elemType.Field(i)
				if !displayField(field.Type) &&
					!(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
					continue
				}

				jsonTag := field.Tag.Get("json")
				if jsonTag == "" {
					jsonTag = field.Name
				}
				headers = append(headers, jsonTag)
			}

			for i := 0; i < value.Len(); i++ {
				elem := value.Index(i)

				if elem.Kind() == reflect.Ptr {
					if elem.IsNil() {
						row := make([]string, len(headers))
						for j := range row {
							row[j] = ""
						}
						rows = append(rows, row)
						continue
					}
					elem = elem.Elem()
				}
				var row []string
				for j := 0; j < elem.NumField(); j++ {
					field := elem.Type().Field(j)
					fieldValue := elem.Field(j)
					if !displayField(field.Type) &&
						!(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
						continue
					}
					row = append(row, formatValue(fieldValue))
				}
				rows = append(rows, row)
			}
		} else {
			headers = []string{"Value"}

			for i := 0; i < value.Len(); i++ {
				elem := value.Index(i)
				row := []string{formatValue(elem)}
				rows = append(rows, row)
			}
		}
	} else if value.Kind() == reflect.Struct {
		elemType := value.Type()

		for i := 0; i < elemType.NumField(); i++ {
			field := elemType.Field(i)
			if !displayField(field.Type) &&
				!(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
				continue
			}

			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			headers = append(headers, jsonTag)
		}

		var row []string
		for j := 0; j < value.NumField(); j++ {
			field := elemType.Field(j)
			fieldValue := value.Field(j)
			if !displayField(field.Type) &&
				!(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
				continue
			}
			row = append(row, formatValue(fieldValue))
		}
		rows = append(rows, row)
	} else {
		headers = []string{"Value"}
		row := []string{formatValue(value)}
		rows = append(rows, row)
	}

	colWidths := make([]int, len(headers))
	for i, h := range headers {
		colWidths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	if !noHeaders {
		fmt.Fprint(writer, "| ")
		for i, h := range headers {
			if i > 0 {
				fmt.Fprint(writer, " | ")
			}
			fmt.Fprint(writer, padRight(h, colWidths[i]))
		}
		fmt.Fprintln(writer, " |")

		fmt.Fprint(writer, "| ")
		for i, w := range colWidths {
			if i > 0 {
				fmt.Fprint(writer, " | ")
			}
			fmt.Fprint(writer, strings.Repeat("-", w))
		}
		fmt.Fprintln(writer, " |")
	}

	for _, row := range rows {
		fmt.Fprint(writer, "| ")
		for i, cell := range row {
			if i > 0 {
				fmt.Fprint(writer, " | ")
			}
			fmt.Fprint(writer, padRight(cell, colWidths[i]))
		}
		fmt.Fprintln(writer, " |")
	}
}

func padRight(s string, width int) string {
	if len(s) < width {
		return s + strings.Repeat(" ", width-len(s))
	}
	return s
}

func displayField(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Int, reflect.Int64, reflect.Int32, reflect.Float64, reflect.Bool:
		return true
	case reflect.Struct:
		return t.Name() == "Time" || true
	case reflect.Ptr:
		return displayField(t.Elem())
	case reflect.Slice:
		return displayField(t.Elem())
	default:
		return false
	}
}

func style(text string, style lipgloss.Style) string {
	return style.Render(text)
}

func formatValue(v reflect.Value) string {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		return formatValue(v.Elem())
	}

	if v.Kind() == reflect.Struct {
		var parts []string

		if v.Type().Name() == "Time" {
			t := v.Interface().(time.Time)
			return t.Format(time.RFC3339)
		}

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.CanInterface() {
				continue
			}
			parts = append(parts, formatValue(field))
		}
		return fmt.Sprintf("{%s}", strings.Join(parts, " "))
	}

	if v.Kind() == reflect.Slice {
		var parts []string
		for i := 0; i < v.Len(); i++ {
			parts = append(parts, formatValue(v.Index(i)))
		}
		return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
	}

	if v.CanInterface() {
		return fmt.Sprintf("%v", v.Interface())
	}

	return ""
}

func PrintQuiet(writer io.Writer, fields ...string) {
	fmt.Fprintln(writer, strings.Join(fields, "\t"))
}

func PrintSmartOutput[T any](
	cmd *cobra.Command,
	items []T,
	fieldsFunc func(T) []string,
	config *SmartOutputConfig[T],
) error {
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("failed to get output flag: %w", err)
	}

	if config != nil && config.DefaultOutput != "" &&
		!cmd.Flags().Changed("output") &&
		!cmd.Flags().Changed("quiet") {
		output = string(config.DefaultOutput)
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to get no-headers flag: %w", err)
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("failed to get quiet flag: %w", err)
	}

	isSingleResource := config != nil && config.SingleResource && len(items) == 1
	compactOutput := config != nil && config.SingleResourceCompactOutput && len(items) == 1

	switch output {
	case "human":
		if quiet {
			if fieldsFunc == nil || (compactOutput && isSingleResource) {
				return nil
			}
			for _, item := range items {
				PrintQuiet(cmd.OutOrStdout(), fieldsFunc(item)...)
			}
			return nil
		}
		if compactOutput && fieldsFunc != nil {
			PrintQuiet(cmd.OutOrStdout(), fieldsFunc(items[0])...)
			return nil
		}

		printMarkdownTable(cmd.OutOrStdout(), items, noHeaders)
		return nil

	default:
		if isSingleResource {
			PrintFormattedData(cmd.OutOrStdout(), items[0], output)
		} else {
			PrintFormattedData(cmd.OutOrStdout(), items, output)
		}
		return nil
	}
}
