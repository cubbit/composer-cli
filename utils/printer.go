package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	RedBg     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
	greenBg   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF80"))
	yellowBg  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFBA00"))
)

func PrintSuccess(s string) {
	fmt.Printf("%s ✨🐝 %s\n", greenBg.Render("SUCCESS"), s)
}

func PrintError(err error) {
	errStr, _ := strings.CutSuffix(err.Error(), "\n")
	lines := strings.Split(errStr, "\n")
	for i, line := range lines {
		if i == 0 {
			l, _ := strings.CutSuffix(line, ": ")
			fmt.Printf("%s %s\n", RedBg.Render("ERR"), l)
		} else {
			fmt.Println(line)
		}
	}
}

func PrintDelete(s string) {
	fmt.Printf("🗑️ 🚮 %s\n", s)
}

func PrintNotFound(s string) {
	fmt.Printf("%s %s\n", yellowBg.Render("WRN"), s)
}

func PrintEmptyList() {
	fmt.Print("🪣  [ ]\n")
}

func PrintList(s string) {
	fmt.Printf("📋 %s\n", boldStyle.Render(s))
}

func PrintVerbose(data interface{}, withLineSeparator bool) {
	value := reflect.ValueOf(data)
	elemType := value.Index(0).Elem().Type()
	numFields := elemType.NumField()
	maxWidths := calculateWidths(value)

	// print header
	header := make([]interface{}, 0)
	for i := 0; i < numFields; i++ {
		field := elemType.Field(i)
		fieldType := field.Type
		if displayField(field.Type.Kind().String()) {
			if fieldType.Kind() == reflect.Ptr {
				if displayField(fieldType.Elem().Kind().String()) {
					header = append(header, field.Tag.Get("json"))
				}
			} else {
				header = append(header, field.Tag.Get("json"))
			}

		}
	}

	headerFormat := ""
	for i := range header {
		headerFormat += fmt.Sprintf("%%-%dv ", maxWidths[i])
	}
	headerFormat += "\n"
	fmt.Printf(headerFormat, header...)

	// print separator
	width, _, err := terminal.GetSize(0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(strings.Repeat("-", width) + "\n")

	// print rows
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i).Elem()
		numFields := elem.Type().NumField()
		row := make([]interface{}, 0)
		for j := 0; j < numFields; j++ {
			field := elem.Type().Field(j)
			fieldType := field.Type
			if displayField(field.Type.Kind().String()) {
				if fieldType.Kind() == reflect.Ptr {
					if displayField(fieldType.Elem().Kind().String()) {
						header = append(header, field.Tag.Get("json"))
						if elem.Field(j).IsNil() {
							row = append(row, fmt.Sprintf("%v", "null"))
						} else if !elem.Field(j).IsNil() {
							value := elem.Field(j).Elem()
							row = append(row, value.Interface())
						}
					}
				} else {
					row = append(row, elem.Field(j).Interface())
				}
			}
		}

		rowFormat := ""
		for i := range row {
			rowFormat += fmt.Sprintf("%%-%dv	", maxWidths[i])
		}
		fmt.Printf(fmt.Sprintf("• %s \n", rowFormat), row...)

		if withLineSeparator {
			fmt.Println()
		}
	}
}

func displayField(str string) bool {
	types := []string{"string", "int", "float64", "bool", "time.Time", "int64", "int32", "ptr"}
	for _, s := range types {
		if s == str {
			return true
		}
	}
	return false
}

func calculateWidths(value reflect.Value) []int {
	elemType := value.Index(0).Elem().Type()
	numFields := elemType.NumField()
	maxWidths := make([]int, numFields)

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i).Elem()
		for j := 0; j < numFields; j++ {
			if !displayField(elem.Type().String()) {
				fieldValue := fmt.Sprintf("%v", elem.Field(j).Interface())
				maxWidths[j] = len(fieldValue)
			}
		}
	}

	return maxWidths
}
