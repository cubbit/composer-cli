package tui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var p *tea.Program

var (
	submitFocusedButton   = focusedStyle.Copy().Render("[ Submit ]")
	submitBlurredButton   = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	continueFocusedButton = focusedStyle.Copy().Render("Continue ->")
	continueBlurredButton = fmt.Sprintf(" %s ", blurredStyle.Render("Continue ->"))
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))
	blurredStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle           = focusedStyle.Copy()
	noStyle               = lipgloss.NewStyle()
	boldStyle             = lipgloss.NewStyle().Bold(true)
)

type Input struct {
	Placeholder string
	IsPassword  bool
	Value       *string
}

type inputsModel struct {
	title       string
	focusIndex  int
	isLastStep  bool
	cancel      bool
	quit        bool
	spinnerDone bool
	inputs      []textinput.Model
	spinner     spinner.Model
}

func initialInputsModel(title string, isLastStep bool, values []Input) inputsModel {
	m := inputsModel{
		inputs:     make([]textinput.Model, len(values)),
		isLastStep: isLastStep,
		title:      title,
	}

	var t textinput.Model

	for i, v := range values {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32
		switch v.IsPassword {
		case false:
			t.Placeholder = v.Placeholder
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.CharLimit = 0
		case true:
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Placeholder = v.Placeholder
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		switch i {
		case 0:
			t.Focus()
		}
		m.inputs[i] = t
	}

	return m
}

func (m inputsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *inputsModel) startSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot
}

func (m *inputsModel) terminateSpinner() {
	time.Sleep(time.Millisecond * 500)
	p.Send(time.Now())
}

func (m inputsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case "ctrl+c", "esc":
			m.quit = true
			m.cancel = true
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
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

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *inputsModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m inputsModel) View() string {
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

	for i := range m.inputs {

		s.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			s.WriteRune('\n')
		}
	}

	if m.isLastStep {
		button := &submitBlurredButton

		if m.focusIndex == len(m.inputs) {
			button = &submitFocusedButton
		}

		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	} else {
		button := &continueBlurredButton

		if m.focusIndex == len(m.inputs) {
			button = &continueFocusedButton
		}

		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	}

	return s.String()
}

func TextInputs(title string, isLastStep bool, values ...Input) ([]Input, error) {
	var err error
	var tm tea.Model

	p = tea.NewProgram(initialInputsModel(title, isLastStep, values))

	if tm, err = p.Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(0)
	}

	m := tm.(inputsModel)

	if m.cancel {
		return values, fmt.Errorf("cancelled")
	}

	for i := range m.inputs {
		*values[i].Value = m.inputs[i].Value()
	}

	return values, err
}
