package input

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Field is a single input in a form.
type Field struct {
	Name        string
	Placeholder string
	Password    bool
	Required    bool
	Validate    func(string) error
}

// FieldOption configures a Field.
type FieldOption func(*Field)

// WithValidate sets a validation function for the field.
func WithValidate(fn func(string) error) FieldOption {
	return func(f *Field) {
		f.Validate = fn
	}
}

// WithPlaceholder sets a placeholder for the field.
func WithPlaceholder(s string) FieldOption {
	return func(f *Field) {
		f.Placeholder = s
	}
}

// WithRequired marks the field as required (cannot be empty on submit).
func WithRequired() FieldOption {
	return func(f *Field) {
		f.Required = true
	}
}

type formModel struct {
	title     string
	fields    []Field
	inputs    []textinput.Model
	focus     int
	quitting  bool
	submitted bool
	errs      []error
	initCmds  []tea.Cmd
}

func (m formModel) Init() tea.Cmd {
	cmds := append(m.initCmds, textinput.Blink)
	return tea.Batch(cmds...)
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			if s == "enter" {
				if m.hasErrors() {
					return m, nil
				}
				m.quitting = true
				m.submitted = true
				return m, tea.Quit
			}
			if s == "up" || s == "shift+tab" {
				m.focus--
			} else {
				m.focus++
			}
			if m.focus >= len(m.inputs) {
				m.focus = 0
			} else if m.focus < 0 {
				m.focus = len(m.inputs) - 1
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				if i == m.focus {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focus
					m.inputs[i].TextStyle = focus
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = normal
					m.inputs[i].TextStyle = normal
				}
			}
			return m, tea.Batch(cmds...)
		}
	}

	cmd := m.updateInputs(msg)
	m.validate()
	return m, cmd
}

func (m *formModel) validate() {
	m.errs = make([]error, len(m.inputs))
	for i, f := range m.fields {
		val := m.inputs[i].Value()
		if val == "" {
			if f.Required {
				m.errs[i] = errRequired
			}
			continue
		}
		if f.Validate != nil {
			m.errs[i] = f.Validate(val)
		}
	}
}

func (m formModel) hasErrors() bool {
	for _, e := range m.errs {
		if e != nil {
			return true
		}
	}
	return false
}

func (m *formModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return tea.Batch(cmds...)
}

func (m formModel) View() string {
	if m.quitting {
		return ""
	}
	var s strings.Builder
	s.WriteString(bold.Render(m.title))
	s.WriteString("\n")
	for i, v := range m.inputs {
		s.WriteString(v.View() + "\n")
		if i < len(m.errs) && m.errs[i] != nil {
			s.WriteString(errSty.Render("  "+m.errs[i].Error()) + "\n")
		}
	}
	return s.String()
}

// MultiInput shows a form with multiple fields and returns a map of name to value.
func MultiInput(title string, fields ...Field) (map[string]string, error) {
	return MultiInputWithOpts(title, fields)
}

// MultiInputWithOpts is like MultiInput but accepts optional tea.ProgramOption
// values for testing (e.g. tea.WithInput, tea.WithOutput).
func MultiInputWithOpts(title string, fields []Field, progOpts ...tea.ProgramOption) (map[string]string, error) {
	inputs := make([]textinput.Model, len(fields))
	var initCmds []tea.Cmd
	for i, f := range fields {
		t := textinput.New()
		t.Placeholder = f.Placeholder
		t.PlaceholderStyle = placeholderStyle
		t.Cursor.Style = focus
		t.Cursor.TextStyle = cursorStyle
		t.PromptStyle = focus
		t.TextStyle = focus
		if f.Password {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		if i == 0 {
			initCmds = append(initCmds, t.Focus())
		}
		inputs[i] = t
	}

	p := tea.NewProgram(formModel{
		title:    title,
		fields:   fields,
		inputs:   inputs,
		errs:     make([]error, len(fields)),
		initCmds: initCmds,
	}, progOpts...)

	m, err := p.Run()
	if err != nil {
		return nil, err
	}

	fm := m.(formModel)
	if fm.submitted {
		values := make(map[string]string, len(fields))
		for i, f := range fields {
			values[f.Name] = fm.inputs[i].Value()
		}
		return values, nil
	}

	return nil, ErrCancelled
}

// Input shows a single-line text input and returns the entered value.
func Input(title string, opts ...FieldOption) (string, error) {
	return InputWithOpts(title, nil, opts...)
}

// InputWithOpts is like Input but accepts tea.ProgramOption values for testing.
func InputWithOpts(title string, progOpts []tea.ProgramOption, opts ...FieldOption) (string, error) {
	f := Field{Name: "input", Placeholder: title}
	for _, o := range opts {
		o(&f)
	}
	res, err := MultiInputWithOpts(title, []Field{f}, progOpts...)
	if err != nil {
		return "", err
	}
	return res["input"], nil
}
