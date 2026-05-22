package interactive

import "github.com/charmbracelet/lipgloss"

var (
	successMark = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	errorMark   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).SetString("✗")
	infoMark    = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).SetString("ℹ")
)
