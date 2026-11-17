package input

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	focus            = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))
	bold             = lipgloss.NewStyle().Bold(true)
	normal           = lipgloss.NewStyle()
	errSty           = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	placeholderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
	cursorStyle      = lipgloss.NewStyle().Background(lipgloss.Color("111")).Foreground(lipgloss.Color("0"))
	errRequired      = fmt.Errorf("required")
)
