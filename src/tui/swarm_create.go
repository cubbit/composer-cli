package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("160")).Bold(true)
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	infoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type paramInputModel struct {
	title         string
	question      string
	explanation   string
	showDerived   string
	value         string
	errorMsg      string
	quit          bool
	cancelled     bool
	isLastStep    bool
	spinnerDone   bool
	spinner       spinner.Model
	input         textinput.Model
	validateFn    func(string) string
	onChangeFn    func(string) (string, string)
}

func initialParamModel(title, question, explanation string, isLast bool, validateFn func(string) string, onChangeFn func(string) (string, string)) paramInputModel {
	ti := textinput.New()
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = focusedStyle
	ti.TextStyle = focusedStyle
	ti.CharLimit = 10
	ti.Focus()

	return paramInputModel{
		title:       title,
		question:    question,
		explanation: explanation,
		isLastStep:  isLast,
		validateFn:  validateFn,
		onChangeFn:  onChangeFn,
		input:       ti,
	}
}

func (m paramInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *paramInputModel) startSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot
}

func (m *paramInputModel) terminateSpinner() {
	time.Sleep(time.Millisecond * 500)
	p.Send(time.Now())
}

func (m paramInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.cancelled = true
			m.quit = true
			return m, tea.Quit

		case "enter":
			if m.value != "" && m.errorMsg == "" && m.validateFn != nil {
				if err := m.validateFn(m.value); err == "" {
					m.quit = true
					if m.isLastStep {
						m.startSpinner()
						return m, m.spinner.Tick
					}
					return m, tea.Quit
				}
			}
			if m.value != "" && m.errorMsg == "" {
				m.quit = true
				if m.isLastStep {
					m.startSpinner()
					return m, m.spinner.Tick
				}
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.value = m.input.Value()

	if m.onChangeFn != nil {
		m.showDerived, m.errorMsg = m.onChangeFn(m.value)
	} else if m.validateFn != nil {
		m.errorMsg = m.validateFn(m.value)
	}

	return m, cmd
}

func (m paramInputModel) View() string {
	var s strings.Builder

	if m.quit {
		if m.isLastStep && !m.spinnerDone {
			b := fmt.Sprintf("%s%s%s", m.spinner.View(), " ", textStyle("Processing..."))
			s.WriteString(b)
			return s.String()
		}
		return ""
	}

	s.WriteString(boldStyle.Render(m.title))
	s.WriteString("\n\n")
	s.WriteString(m.question)
	s.WriteString("\n")

	if m.explanation != "" {
		s.WriteString(subduedStyle.Render(m.explanation))
		s.WriteString("\n\n")
	}

	s.WriteString(m.input.View())
	s.WriteString("\n")

	if m.errorMsg != "" {
		s.WriteString(errorStyle.Render("✗ " + m.errorMsg))
		s.WriteString("\n")
	}

	if m.showDerived != "" {
		s.WriteString("\n")
		s.WriteString(infoStyle.Render(m.showDerived))
		s.WriteString("\n")
	}

	if m.isLastStep {
		button := &submitBlurredButton
		if m.value != "" && m.errorMsg == "" {
			button = &submitFocusedButton
		}
		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	} else {
		button := &continueBlurredButton
		if m.value != "" && m.errorMsg == "" {
			button = &continueFocusedButton
		}
		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	}

	return s.String()
}

type confirmModel struct {
	title     string
	body      string
	options   []string
	cursor    int
	quit      bool
	cancelled bool
	choice    string
}

func initialConfirmModel(title, body string, options []string) confirmModel {
	return confirmModel{
		title:   title,
		body:    body,
		options: options,
	}
}

func (m confirmModel) Init() tea.Cmd { return nil }

func (m confirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.cancelled = true
			m.quit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.choice = m.options[m.cursor]
			m.quit = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m confirmModel) View() string {
	if m.quit {
		return ""
	}

	var s strings.Builder
	s.WriteString(boldStyle.Render(m.title))
	s.WriteString("\n\n")

	if m.body != "" {
		s.WriteString(m.body)
		s.WriteString("\n\n")
	}

	for i, opt := range m.options {
		if i == m.cursor {
			s.WriteString(focusedStyle.Render("▸ " + opt))
		} else {
			s.WriteString(fmt.Sprintf("  %s", opt))
		}
		s.WriteString("\n")
	}

	return s.String()
}

// ConfirmDialog displays a selection list and returns the chosen option.
func ConfirmDialog(title, body string, options []string) (string, error) {
	m := initialConfirmModel(title, body, options)
	tm, err := tea.NewProgram(m).Run()
	if err != nil {
		return "", fmt.Errorf("failed to run confirmation: %w", err)
	}
	cm := tm.(confirmModel)
	if cm.cancelled {
		return "", fmt.Errorf("cancelled")
	}
	return cm.choice, nil
}

type phaseProgressModel struct {
	phase   string
	step    string
	status  string
	done    bool
	spinner spinner.Model
}

func (m phaseProgressModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m phaseProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case tea.KeyMsg:
		m.done = true
		return m, tea.Quit
	}
	return m, nil
}

func (m phaseProgressModel) View() string {
	if m.done {
		return successStyle.Render("✓ " + m.status)
	}
	return fmt.Sprintf("%s %s - %s", m.spinner.View(), m.phase, m.step)
}

type tickMsg time.Time

type monitorModel struct {
	processID    string
	totalSteps   int
	currentStep  int
	currentTitle string
	status       string
	errorMsg     string
	done         bool
	quit         bool
	cancelled    bool
	spinner      spinner.Model
	pollFn       func() (string, string, int, error)
}

func initialMonitorModel(processID string, pollFn func() (string, string, int, error)) monitorModel {
	s := spinner.New()
	s.Style = spinnerStyle
	s.Spinner = spinner.MiniDot

	return monitorModel{
		processID:  processID,
		spinner:    s,
		pollFn:     pollFn,
		totalSteps: 5,
	}
}

func (m monitorModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, m.tickCmd())
}

func (m monitorModel) tickCmd() tea.Cmd {
	return tea.Tick(5*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m monitorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.cancelled = true
			m.quit = true
			return m, tea.Quit
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tickMsg:
		if m.pollFn != nil {
			status, step, progress, err := m.pollFn()
			if err != nil {
				m.status = fmt.Sprintf("Error: %v", err)
				m.errorMsg = err.Error()
				m.quit = true
				return m, tea.Quit
			}
			m.status = status
			m.currentTitle = step
			if progress > 0 {
				m.currentStep = int(float64(m.totalSteps) * float64(progress) / 100)
			}

			switch status {
			case "success":
				m.done = true
				m.quit = true
				return m, tea.Quit
			case "failed":
				m.done = true
				m.quit = true
				return m, tea.Quit
			}
		}
		return m, m.tickCmd()
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func statusLabel(status string) string {
	switch status {
	case "running":
		return "In progress"
	case "success":
		return "Completed"
	case "failed":
		return "Failed"
	default:
		return status
	}
}

func (m monitorModel) View() string {
	var s strings.Builder

	if m.quit && m.done {
		if m.status == "success" {
			s.WriteString(successStyle.Render("✓ Swarm creation completed successfully!"))
		} else if m.status == "failed" {
			s.WriteString(errorStyle.Render("✗ Swarm creation failed"))
			if m.errorMsg != "" {
				s.WriteString("\n" + m.errorMsg)
			}
		} else if m.cancelled {
			s.WriteString(boldStyle.Render("Monitoring paused"))
			s.WriteString("\n")
			s.WriteString(subduedStyle.Render("Run 'swarm create --interactive' to monitor progress again"))
		}
		return s.String()
	}

	s.WriteString(boldStyle.Render("Swarm Creation in Progress"))
	s.WriteString("\n\n")
	s.WriteString(fmt.Sprintf("Process ID: %s", m.processID))
	s.WriteString("\n\n")

	s.WriteString(fmt.Sprintf("%s %s", m.spinner.View(), statusLabel(m.status)))
	s.WriteString("\n")

	if m.currentTitle != "" {
		s.WriteString("\n")
		s.WriteString(fmt.Sprintf("Current operation: %s", boldStyle.Render(m.currentTitle)))
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString(subduedStyle.Render("Press q or Ctrl+C to stop monitoring"))
	s.WriteString("\n")

	return s.String()
}

// MonitorResult holds the outcome of a swarm creation monitoring session.
type MonitorResult struct {
	Status    string
	ErrorMsg  string
	Done      bool
	Cancelled bool
}

// MonitorSession wraps the internal monitor model and provides an exported Run method.
type MonitorSession struct {
	model monitorModel
}

// NewMonitorSession creates a new MonitorSession for swarm creation monitoring.
func NewMonitorSession(processID string, pollFn func() (string, string, int, error)) MonitorSession {
	return MonitorSession{model: initialMonitorModel(processID, pollFn)}
}

// Run starts the monitor TUI and returns the result when it completes.
func (s MonitorSession) Run() (MonitorResult, error) {
	tm, err := tea.NewProgram(s.model).Run()
	if err != nil {
		return MonitorResult{}, fmt.Errorf("failed to run monitor: %w", err)
	}
	mm := tm.(monitorModel)
	return MonitorResult{
		Status:    mm.status,
		ErrorMsg:  mm.errorMsg,
		Done:      mm.done,
		Cancelled: mm.cancelled,
	}, nil
}

// CollectIntParameter runs a TUI to collect a numeric parameter with real-time validation.
func CollectIntParameter(title, question, explanation string, validateFn func(string) string, onChangeFn func(string) (string, string)) int {
	m := initialParamModel(title, question, explanation, false, validateFn, onChangeFn)
	tm, err := tea.NewProgram(m).Run()
	if err != nil {
		return -1
	}
	pm := tm.(paramInputModel)
	if pm.cancelled {
		return -1
	}
	return parseInt(pm.value)
}

func parseInt(s string) int {
	if s == "" {
		return 0
	}
	val := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
		} else {
			return -1
		}
	}
	return val
}

// ShowError prints an error message to stdout.
func ShowError(msg string) {
	fmt.Println(errorStyle.Render("✗ Error: " + msg))
}
