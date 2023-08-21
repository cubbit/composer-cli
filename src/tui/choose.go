package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

var (
	subduedStyle     = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#847A85", Dark: "#979797"})
	verySubduedStyle = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#DDDADA", Dark: "#3C3C3C"})
)

type keyMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Enter  key.Binding
	Select key.Binding
	Quit   key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
	),
	Select: key.NewBinding(
		key.WithKeys(" ", "tab", "x"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
	),
}

type item struct {
	text     string
	selected bool
	order    int
}

type chooseModel struct {
	title             string
	height            int
	cursor            string
	selectedPrefix    string
	unselectedPrefix  string
	cursorPrefix      string
	items             []item
	quit              bool
	isLastStep        bool
	spinnerDone       bool
	index             int
	limit             int
	numSelected       int
	currentOrder      int
	keys              keyMap
	paginator         paginator.Model
	cancelled         bool
	cursorStyle       lipgloss.Style
	itemStyle         lipgloss.Style
	selectedItemStyle lipgloss.Style
	spinner           spinner.Model
}

func (m chooseModel) Init() tea.Cmd { return nil }

func (m *chooseModel) startSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinner.MiniDot
}

func (m *chooseModel) terminateSpinner() {
	time.Sleep(time.Millisecond * 500)
	p.Send(time.Now())
}

func (m chooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		return m, nil

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
		start, end := m.paginator.GetSliceBounds(len(m.items))
		switch {
		case key.Matches(msg, m.keys.Down):
			m.index++

			if m.index >= len(m.items) {
				m.index = 0
				m.paginator.Page = 0
			}

			if m.index >= end {
				m.paginator.NextPage()
			}

		case key.Matches(msg, m.keys.Up):
			m.index--

			if m.index < 0 {
				m.index = len(m.items) - 1
				m.paginator.Page = m.paginator.TotalPages - 1
			}

			if m.index < start {
				m.paginator.PrevPage()
			}

		case key.Matches(msg, m.keys.Right):
			m.index = clamp(m.index+m.height, 0, len(m.items)-1)
			m.paginator.NextPage()

		case key.Matches(msg, m.keys.Left):
			m.index = clamp(m.index-m.height, 0, len(m.items)-1)
			m.paginator.PrevPage()

		case key.Matches(msg, m.keys.Quit):
			m.cancelled = true
			m.quit = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Select):
			if m.limit == 1 {
				break
			}

			if m.items[m.index].selected {
				m.items[m.index].selected = false
				m.numSelected--
			} else if m.numSelected < m.limit {
				m.items[m.index].selected = true
				m.items[m.index].order = m.currentOrder
				m.numSelected++
				m.currentOrder++
			}

		case key.Matches(msg, m.keys.Enter):
			m.quit = true

			if m.numSelected < 1 {
				m.items[m.index].selected = true
			}

			if m.isLastStep {
				m.startSpinner()
				return m, m.spinner.Tick
			}
			
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.paginator, cmd = m.paginator.Update(msg)
	return m, cmd
}

func (m chooseModel) View() string {
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

	start, end := m.paginator.GetSliceBounds(len(m.items))

	for i, item := range m.items[start:end] {
		if i == m.index%m.height {
			s.WriteString(m.cursorStyle.Render(m.cursor))
		} else {
			s.WriteString(strings.Repeat(" ", runewidth.StringWidth(m.cursor)))
		}

		if item.selected {
			s.WriteString(m.selectedItemStyle.Render(m.selectedPrefix + item.text))
		} else if i == m.index%m.height {
			s.WriteString(m.cursorStyle.Render(m.cursorPrefix + item.text))
		} else {
			s.WriteString(m.itemStyle.Render(m.unselectedPrefix + item.text))
		}

		if i != m.height {
			s.WriteRune('\n')
		}
	}

	if m.isLastStep {
		button := &submitBlurredButton
		if m.numSelected == 1 {
			button = &submitFocusedButton
		}
		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	} else {
		button := &continueBlurredButton
		if m.numSelected == 1 {
			button = &continueFocusedButton
		}
		fmt.Fprintf(&s, "\n\n%s\n\n", boldStyle.Render(*button))
	}

	if m.paginator.TotalPages <= 1 {
		return s.String()
	}

	s.WriteString(strings.Repeat("\n", m.height-m.paginator.ItemsOnPage(len(m.items))+1))
	s.WriteString("  " + m.paginator.View())

	return s.String()
}

func clamp(x, min, max int) int {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func ChooseOne(title string, isLastStep bool, options []string) (string, error) {
	var err error
	var choice []string

	if choice, err = choose(title, isLastStep, options, 1); err != nil {
		return "", err
	}

	if len(choice) == 0 {
		return "", fmt.Errorf("no option was selected")
	}

	return choice[0], err
}

func choose(title string, isLastStep bool, options []string, limit int) ([]string, error) {
	var err error
	var tm tea.Model

	if limit == 0 {
		limit = len(options)
	}

	items := make([]item, len(options))
	height := 10
	pager := paginator.New()
	pager.SetTotalPages((len(items) + height - 1) / height)
	pager.PerPage = height
	pager.Type = paginator.Dots
	pager.ActiveDot = subduedStyle.Render("•")
	pager.InactiveDot = verySubduedStyle.Render("•")

	for i, option := range options {
		items[i] = item{text: option, selected: false, order: i}
	}

	p = tea.NewProgram(chooseModel{
		title:             title,
		index:             0,
		currentOrder:      0,
		height:            height,
		cursor:            "> ",
		selectedPrefix:    "- ",
		unselectedPrefix:  " ",
		cursorPrefix:      "",
		items:             items,
		limit:             limit,
		keys:              keys,
		paginator:         pager,
		cursorStyle:       focusedStyle,
		selectedItemStyle: focusedStyle,
		numSelected:       0,
		isLastStep:        isLastStep,
	})

	if tm, err = p.Run(); err != nil {
		return []string{}, fmt.Errorf("failed to start tea program: %w", err)
	}

	m := tm.(chooseModel)

	if m.cancelled {
		return []string{}, fmt.Errorf("cancelled")
	}

	if limit > 1 {
		sort.Slice(m.items, func(i, j int) bool {
			return m.items[i].order < m.items[j].order
		})
	}

	var results []string

	for _, item := range m.items {
		if item.selected {
			results = append(results, item.text)
		}
	}

	return results, nil
}
