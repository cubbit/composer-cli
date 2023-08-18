package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
)

var (
	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Input struct {
	Placeholder string
	IsPassword  bool
}

type inputsModel struct {
	focusIndex int
	inputs     []textinput.Model
	quitting   bool
	submit     bool
	title      string
}

func initialModel(title string, submit bool, values []Input) inputsModel {
	m := inputsModel{
		inputs: make([]textinput.Model, len(values)),
		submit: submit,
		title:  title,
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

func (m inputsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.quitting = true
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
	var b strings.Builder
	b.WriteString((m.title) + "\n")
	if m.quitting {
		return ""
	}
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}
	if m.submit {
		button := &blurredButton
		if m.focusIndex == len(m.inputs) {
			button = &focusedButton
		}
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	}
	return b.String()
}

func Inputs(title string, submit bool, values ...Input) []string {
	m := initialModel(title, submit, values)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(0)
	}
	outs := make([]string, len(values))
	for i := range m.inputs {

		outs[i] = m.inputs[i].Value()
	}
	return outs
}
