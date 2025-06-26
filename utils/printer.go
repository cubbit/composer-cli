package utils

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/charmbracelet/lipgloss"
)

type OutputMode int

const (
	OutputQuiet OutputMode = iota
	OutputHuman
)

var (
	currentMode = OutputQuiet
)

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	RedBg     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	greenBg   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF80"))
	yellowBg  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFBA00"))
	grayStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	blueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#0080FF"))
)

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
	errStr, _ := strings.CutSuffix(err.Error(), "\n")
	lines := strings.Split(errStr, "\n")

	switch currentMode {
	case OutputHuman:
		for i, line := range lines {
			if i == 0 {
				l, _ := strings.CutSuffix(line, ": ")
				fmt.Fprintf(os.Stderr, "%s %s\n", style("ERR", RedBg), l)
			} else {
				fmt.Fprintf(os.Stderr, "%s %s\n", style("INF", blueStyle), line)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
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

func PrintVerbose(data interface{}, withLineSeparator bool) {
	value := reflect.ValueOf(data)
	if value.Len() == 0 {
		PrintEmptyList()
		return
	}
	switch currentMode {
	case OutputQuiet:
		printTabSeparatedTable(data)
		return
	default:
		printHumanTable(data, withLineSeparator)
	}
}

func printHumanTable(data interface{}, withLineSeparator bool) {
	value := reflect.ValueOf(data)
	if value.Len() == 0 {
		return
	}

	elemType := value.Index(0).Elem().Type()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// print header
	var headers []string
	for i := 0; i < elemType.NumField(); i++ {
		field := elemType.Field(i)
		if displayField(field.Type) ||
			(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = field.Name
			}
			headers = append(headers, strings.ToUpper(jsonTag))
		}
	}

	headerRow := strings.Join(headers, "\t")
	if currentMode == OutputHuman {
		headerRow = style(headerRow, boldStyle)
	}
	fmt.Fprintln(w, headerRow)

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i).Elem()
		var rowValues []string

		for j := 0; j < elem.NumField(); j++ {
			field := elem.Type().Field(j)
			fieldValue := elem.Field(j)

			if !displayField(field.Type) &&
				!(field.Type.Kind() == reflect.Ptr && displayField(field.Type.Elem())) {
				continue
			}

			displayValue := formatValue(fieldValue)
			rowValues = append(rowValues, displayValue)
		}

		fmt.Fprintln(w, strings.Join(rowValues, "\t"))
		if withLineSeparator && currentMode == OutputHuman {
			fmt.Fprintln(w, "")
		}
	}

	w.Flush()
}

func printTabSeparatedTable(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Len() == 0 {
		return
	}

	// print rows
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i).Elem()
		var rowValues []string

		for j := 0; j < elem.NumField(); j++ {
			field := elem.Type().Field(j)
			fieldValue := elem.Field(j)

			if !displayField(field.Type) {
				continue
			}

			displayValue := formatValue(fieldValue)
			rowValues = append(rowValues, displayValue)
		}

		fmt.Println(strings.Join(rowValues, "\t"))
	}
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
