package utils

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	redBg     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))
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
			fmt.Printf("%s %s\n", redBg.Render("ERR"), l)
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
