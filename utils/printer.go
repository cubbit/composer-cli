package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	boldStyle = lipgloss.NewStyle().Bold(true)
	redBg     = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#FF0000"))
	greenBg   = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#00FF80"))
	yellowBg   = lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#FFBA00"))
)

func PrintSuccess(s string) {
	fmt.Printf("%s ✨🐝 %s\n", greenBg.Render("SUCCESS"), s)
}

func PrintError(err error) {
	fmt.Printf("%s %s\n", redBg.Render("ERROR"), err)
}

func PrintDelete(s string) {
	fmt.Printf("🗑️ 🚮 %s\n", s)
}

func PrintNotFound(s string) {
	fmt.Printf("%s %s\n", yellowBg.Render("WARN"), s)

}

func PrintEmptyList() {
	fmt.Print("🪣  [ ]\n")
}
func PrintList(s string) {
	fmt.Printf("📋 %s\n", boldStyle.Render(s))
}
