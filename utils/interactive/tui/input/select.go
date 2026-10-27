package input

import (
	"errors"
	"strings"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
)

type selectModel struct {
	title    string
	cursor   int
	items    []string
	selected map[int]bool
	multi    bool
	pager    paginator.Model
	done     bool
	errMsg   string
}

func (m selectModel) hasSelection() bool {
	for _, v := range m.selected {
		if v {
			return true
		}
	}
	return false
}

func (m selectModel) Init() tea.Cmd { return nil }

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	start, end := m.pager.GetSliceBounds(len(m.items))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			m.done = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}

		case " ":
			if m.multi {
				m.selected[m.cursor] = !m.selected[m.cursor]
				if m.hasSelection() {
					m.errMsg = ""
				}
			} else {
				m.selected = map[int]bool{m.cursor: true}
				m.done = true
				return m, tea.Quit
			}

		case "enter":
			if m.multi {
				if !m.hasSelection() {
					m.errMsg = ErrAtLeastOneRequired.Error()
					return m, nil
				}
			} else if m.cursor < end {
				m.selected = map[int]bool{m.cursor: true}
			}
			m.done = true
			return m, tea.Quit
		}

		if m.cursor > end {
			m.cursor = start
		} else if m.cursor < start {
			m.cursor = end
		}
	}

	var cmd tea.Cmd
	m.pager, cmd = m.pager.Update(msg)
	return m, cmd
}

func (m selectModel) View() string {
	if m.done {
		return ""
	}

	var s strings.Builder
	s.WriteString(bold.Render(m.title) + "\n")

	start, end := m.pager.GetSliceBounds(len(m.items))

	for i, item := range m.items[start:end] {
		idx := start + i
		prefix := "  "
		if idx == m.cursor {
			prefix = "> "
		}
		check := ""
		if m.multi {
			if m.selected[idx] {
				check = "[x] "
			} else {
				check = "[ ] "
			}
		} else {
			if m.selected[idx] {
				check = "✓ "
			}
		}
		line := prefix + check + item
		if idx == m.cursor {
			s.WriteString(focus.Render(line) + "\n")
		} else {
			s.WriteString(line + "\n")
		}
	}

	if m.pager.TotalPages > 1 {
		s.WriteString("\n  " + m.pager.View() + "\n")
	}

	if m.errMsg != "" {
		s.WriteString("\n" + errSty.Render(m.errMsg) + "\n")
	}

	return s.String()
}

func runProgram(title string, options []string, multi bool, progOpts ...tea.ProgramOption) ([]string, error) {
	p := paginator.New()
	p.PerPage = 10
	p.SetTotalPages((len(options) + 9) / 10)
	p.Type = paginator.Dots
	p.ActiveDot = "•"
	p.InactiveDot = "○"

	prog := tea.NewProgram(selectModel{
		title:    title,
		items:    options,
		selected: make(map[int]bool),
		multi:    multi,
		pager:    p,
	}, progOpts...)

	m, err := prog.Run()
	if err != nil {
		return nil, err
	}

	sm := m.(selectModel)
	var result []string
	for i := 0; i < len(options); i++ {
		if sm.selected[i] {
			result = append(result, options[i])
		}
	}
	return result, nil
}

// Select shows a list of options and returns the chosen one.
func Select(title string, options []string) (string, error) {
	return SelectWithOpts(title, options)
}

// SelectWithOpts is like Select but accepts tea.ProgramOption values for testing.
func SelectWithOpts(title string, options []string, progOpts ...tea.ProgramOption) (string, error) {
	if len(options) == 0 {
		return "", errors.New("no options")
	}

	result, err := runProgram(title, options, false, progOpts...)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", ErrCancelled
	}
	return result[0], nil
}

// MultiSelect shows a list of options with checkboxes and returns all selected.
func MultiSelect(title string, options []string) ([]string, error) {
	return MultiSelectWithOpts(title, options)
}

// MultiSelectWithOpts is like MultiSelect but accepts tea.ProgramOption values for testing.
func MultiSelectWithOpts(title string, options []string, progOpts ...tea.ProgramOption) ([]string, error) {
	if len(options) == 0 {
		return nil, errors.New("no options")
	}

	result, err := runProgram(title, options, true, progOpts...)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, ErrCancelled
	}
	return result, nil
}

// Confirm shows a yes/no prompt and returns the choice.
func Confirm(title string) (bool, error) {
	return ConfirmWithOpts(title)
}

// ConfirmWithOpts is like Confirm but accepts tea.ProgramOption values for testing.
func ConfirmWithOpts(title string, progOpts ...tea.ProgramOption) (bool, error) {
	choice, err := SelectWithOpts(title, []string{"Yes", "No"}, progOpts...)
	if err != nil {
		return false, err
	}
	return choice == "Yes", nil
}
