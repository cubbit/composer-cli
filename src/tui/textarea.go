package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextArea struct {
	Placeholder string
	Value       *string
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
}

func initialTextAreaModel(title string, isLastStep bool, ta TextArea) textAreaModel {
	m := textAreaModel{
		isLastStep: isLastStep,
		title:      title,
	}
	t := textarea.New()
	t.Cursor.Style = cursorStyle
	t.CharLimit = 32
	t.CharLimit = 0
	t.Focus()
	t.InsertString("{}")
	t.Prompt = ""
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
		case "ctrl+l":
			if m.input.Focused() {
				m.input.InsertString("\n")
			}
			return m, m.input.Focus()

		case "ctrl+c", "esc":
			m.quit = true
			m.cancel = true
			return m, tea.Quit

		case "tab", "shift+tab", "up", "enter", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == 1 {
				m.quit = true
				if m.isLastStep {
					m.startSpinner()
					return m, m.spinner.Tick
				}
				return m, tea.Quit
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

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *textAreaModel) updateInputs(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
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
