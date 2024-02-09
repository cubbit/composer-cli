package tui

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorLineStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("57")).
			Foreground(lipgloss.Color("230"))

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("238"))

	endOfBufferStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("235"))

	focusedPlaceholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99"))

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("238"))

	blurredBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.HiddenBorder())
)

type EmailBulkTextArea struct {
	Placeholder string
	Value       *string
}

type emailBulkTextAreaModel struct {
	title       string
	focusIndex  int
	isLastStep  bool
	cancel      bool
	quit        bool
	spinnerDone bool
	input       textarea.Model
	spinner     spinner.Model
	width       int
	height      int
	err         error
}

func initialEmailBulkTextAreaModel(title string, isLastStep bool, ta EmailBulkTextArea) emailBulkTextAreaModel {
	m := emailBulkTextAreaModel{
		isLastStep: isLastStep,
		title:      title,
	}
	t := textarea.New()
	t.Prompt = ""
	t.Placeholder = ta.Placeholder
	t.ShowLineNumbers = true
	t.Cursor.Style = cursorStyle
	t.BlurredStyle.Placeholder = placeholderStyle
	t.FocusedStyle.Base = focusedBorderStyle
	t.BlurredStyle.Base = focusedBorderStyle
	t.FocusedStyle.EndOfBuffer = endOfBufferStyle
	t.BlurredStyle.EndOfBuffer = endOfBufferStyle
	t.KeyMap.DeleteWordBackward.SetEnabled(false)
	t.KeyMap.LineNext = key.NewBinding(key.WithKeys("down"))
	t.KeyMap.LinePrevious = key.NewBinding(key.WithKeys("up"))
	t.Focus()

	m.input = t

	return m
}

func (m emailBulkTextAreaModel) Init() tea.Cmd {
	return textarea.Blink
}
func (m *emailBulkTextAreaModel) startSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot
}

func (m *emailBulkTextAreaModel) terminateSpinner() {
	time.Sleep(time.Millisecond * 500)
	p.Send(time.Now())
}

func (m emailBulkTextAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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

				if s == "enter" && m.focusIndex == 1 && m.err == nil {
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
					m.input.Blur()
				}

				cmds := make([]tea.Cmd, 1)
				return m, tea.Batch(cmds...)
			}
		}
	}

	m.sizeInputs()
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *emailBulkTextAreaModel) updateInputs(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.validateEmailBulkStringCmd()
	return cmd
}

func (m *emailBulkTextAreaModel) sizeInputs() {
	m.input.SetWidth(m.width - 2)

}

func validateEmailString(emailString string) bool {
	regex := `^[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,}(?:\s*,\s*[\w\.-]+@[a-zA-Z\d\.-]+\.[a-zA-Z]{2,})*$`
	match, err := regexp.MatchString(regex, emailString)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	return match
}

func (m *emailBulkTextAreaModel) validateEmailBulkStringCmd() bool {
	if !validateEmailString(m.input.Value()) {
		m.err = fmt.Errorf("invalid input format, please try again")
		return false
	}

	m.err = nil
	return true

}

func (m emailBulkTextAreaModel) View() string {
	var s strings.Builder

	if m.quit {
		if m.isLastStep && !m.spinnerDone {
			b := fmt.Sprintf("%s%s%s", m.spinner.View(), " ", textStyle("Sending request..."))
			s.WriteString(b)
			return s.String()
		}
		return ""
	}

	s.WriteString(boldStyle.Render(m.title) + "\n")
	s.WriteString("\n")
	s.WriteString("💡 Hint: Please make sure emails are separated by commas.\n")

	if m.input.Value() != "" && m.focusIndex == 1 {
		isValid := m.validateEmailBulkStringCmd()
		if !isValid {
			s.WriteString("\n🚫 " + m.err.Error() + "\n")
		}
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

func EmailBulkTextAreas(title string, isLastStep bool, ta EmailBulkTextArea) (EmailBulkTextArea, error) {
	var err error
	var tm tea.Model

	p = tea.NewProgram(initialEmailBulkTextAreaModel(title, isLastStep, ta))

	if tm, err = p.Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(0)
	}

	m := tm.(emailBulkTextAreaModel)

	if m.cancel {
		return ta, fmt.Errorf("cancelled")
	}

	*ta.Value = m.input.Value()

	return ta, err
}
