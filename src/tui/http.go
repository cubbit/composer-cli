package tui

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
)

type model struct {
	spinner  spinner.Model
	quitting bool
}

func newModel() model {
	s := spinner.New()
	s.Style = spinnerStyle
	s.Spinner = spinner.MiniDot
	return model{
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m model) View() (s string) {
	if m.quitting {
		return
	}
	s += fmt.Sprintf("%s%s%s", m.spinner.View(), " ", textStyle("Sending request..."))
	return
}

func Send(cmd *cobra.Command, action func(cmd *cobra.Command) error) error {
	var err error
	p := tea.NewProgram(newModel())
	go func() {
		for {
			pause := time.Duration(rand.Int63n(899)+100) * time.Millisecond
			time.Sleep(pause)
			err = action(cmd)
			p.Quit()
		}
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
	return err
}
