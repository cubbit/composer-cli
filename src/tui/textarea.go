package tui

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	greenBorderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	redBorderStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
)

type TextArea struct {
	Placeholder  string
	Value        *string
	InitialValue string
}

type textAreaModel struct {
	title       string
	focusIndex  int
	isLastStep  bool
	cancel      bool
	quit        bool
	spinnerDone bool
	input       textarea.Model
	spinner     spinner.Model
	err         error
}

func initialTextAreaModel(title string, isLastStep bool, ta TextArea) textAreaModel {
	m := textAreaModel{
		isLastStep: isLastStep,
		title:      title,
	}
	t := textarea.New()
	t.Cursor.Style = cursorStyle
	t.CharLimit = 0
	t.Focus()
	t.InsertString(ta.InitialValue)
	t.Prompt = greenBorderStyle.Render(lipgloss.ThickBorder().Left + " ")
	t.SetWidth(80)
	t.SetHeight(20)
	m.input = t

	return m
}

func (m textAreaModel) Init() tea.Cmd {
	return textarea.Blink
}
func (m *textAreaModel) startSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot
}

func (m *textAreaModel) terminateSpinner() {
	time.Sleep(time.Millisecond * 500)
	p.Send(time.Now())
}

func (m *textAreaModel) validateJSONCmd(input string) tea.Cmd {
	return func() tea.Msg {
		if !json.Valid([]byte(input)) {
			m.err = fmt.Errorf("invalid JSON format")
			m.input.Prompt = redBorderStyle.Render(lipgloss.ThickBorder().Left + " ")
			return m.err
		}

		var data map[string]interface{}
		if err := json.Unmarshal([]byte(input), &data); err != nil {
			m.err = err
			return err
		}

		// Clear error if JSON is valid
		m.err = nil
		m.input.Prompt = greenBorderStyle.Render(lipgloss.ThickBorder().Left + " ")
		return data
	}
}
func (m textAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case time.Time:
		m.quit = true
		m.spinnerDone = true
		return m, tea.Quit

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		go m.terminateSpinner()
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			m.input.Focus()

		case "ctrl+enter":
			if m.input.Focused() {
				m.input.InsertString("\n")
			}
			return m, m.input.Focus()
		case "tab":
			if m.input.Focused() {
				m.input.InsertString("\t")
			}

		case "ctrl+c":
			m.quit = true
			m.cancel = true
			return m, tea.Quit

		case "esc":
			m.input.Blur()
		case "shift+tab", "up", "enter", "down":
			if !m.input.Focused() {
				s := msg.String()

				if s == "enter" && m.focusIndex == 1 {
					if m.err == nil {
						m.quit = true
						if m.isLastStep {
							m.startSpinner()
							return m, m.spinner.Tick
						}
						return m, tea.Quit
					}
				}

				if s == "up" || s == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > 1 {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = 1
				}

				cmds := make([]tea.Cmd, 1)
				return m, tea.Batch(cmds...)
			}
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *textAreaModel) updateInputs(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.validateJSONCmd(m.input.Value())()
	m.input, cmd = m.input.Update(msg)
	return cmd
}

func (m textAreaModel) View() string {
	var s strings.Builder

	if m.quit {
		if m.isLastStep && !m.spinnerDone {
			b := fmt.Sprintf("%s%s%s", m.spinner.View(), " ", textStyle("Sending request..."))
			s.WriteString(b)
			return s.String()
		}
		return ""
	}

	s.WriteString(boldStyle.Render(m.title))
	s.WriteString("\n")
	if m.err != nil {
		s.WriteString("\n" + "❌ " + m.err.Error() + "\n")
	}
	s.WriteString("\n")
	s.WriteString(m.input.View())

	if m.isLastStep {
		button := &submitBlurredButton

		if m.focusIndex == 1 {
			button = &submitFocusedButton
		}

		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	} else {
		button := &continueBlurredButton

		if m.focusIndex == 1 {
			button = &continueFocusedButton
		}

		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	}

	return s.String()
}

func TextAreas(title string, isLastStep bool, ta TextArea) (TextArea, error) {
	var err error
	var tm tea.Model

	p = tea.NewProgram(initialTextAreaModel(title, isLastStep, ta))

	if tm, err = p.Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(0)
	}

	m := tm.(textAreaModel)

	if m.cancel {
		return ta, fmt.Errorf("cancelled")
	}

	*ta.Value = m.input.Value()

	return ta, err
}
