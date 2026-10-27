package tui

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/truncate"
	"github.com/sahilm/fuzzy"
)

const (
	Unfiltered FilterState = iota
	Filtering
	FilterApplied
)

var (
	subduedColor    = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#5C5C5C"}
	noItems         = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#909090", Dark: "#626262"})
	statusEmpty     = lipgloss.NewStyle().Foreground(subduedColor)
	statusBar       = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).Padding(0, 0, 1, 2)
	paginationStyle = lipgloss.NewStyle().PaddingLeft(2)
	helpStyle       = lipgloss.NewStyle().Padding(1, 0, 0, 2)
	filter          = lipgloss.NewStyle().Padding(1, 0, 0, 2)
)

type filteredItems []filteredItem

type FilterMatchesMsg []filteredItem

type FilterFunc func(string, []string) []Rank

type Item interface {
	FilterValue() string
}

type ItemDelegate interface {
	Render(w io.Writer, m listModel, item string)
	Height() int
	Spacing() int
	Update(msg tea.Msg, m *listModel) tea.Cmd
}

type filteredItem struct {
	item    string
	matches []int
}

type Rank struct {
	Index          int
	MatchedIndexes []int
}

type KeyMap struct {
	NextPage             key.Binding
	PrevPage             key.Binding
	Filter               key.Binding
	ClearFilter          key.Binding
	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding
	Quit                 key.Binding
}

type statusMessageTimeoutMsg struct{}

type FilterState int

type listModel struct {
	showFilter              bool
	showStatusBar           bool
	showPagination          bool
	showHelp                bool
	filteringEnabled        bool
	itemNamePlural          string
	InfiniteScrolling       bool
	KeyMap                  KeyMap
	Filter                  FilterFunc
	disableQuitKeybindings  bool
	AdditionalShortHelpKeys func() []key.Binding
	width                   int
	height                  int
	Paginator               paginator.Model
	cursor                  int
	Help                    help.Model
	FilterInput             textinput.Model
	filterState             FilterState
	StatusMessageLifetime   time.Duration
	statusMessage           string
	statusMessageTimer      *time.Timer
	items                   []string
	filteredItems           filteredItems
	delegate                ItemDelegate
}

func (f filteredItems) items() []string {
	agg := make([]string, len(f))
	for i, v := range f {
		agg[i] = v.item
	}
	return agg
}

func DefaultFilter(term string, targets []string) []Rank {
	ranks := fuzzy.Find(term, targets)
	sort.Stable(ranks)
	result := make([]Rank, len(ranks))
	for i, r := range ranks {
		result[i] = Rank{
			Index:          r.Index,
			MatchedIndexes: r.MatchedIndexes,
		}
	}
	return result
}

func (f FilterState) String() string {
	return [...]string{
		"unfiltered",
		"filtering",
		"filter applied",
	}[f]
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithHelp("←/h/pgup", "prev page"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithHelp("→/l/pgdn", "next page"),
		),

		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),

		ClearFilter: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "clear filter"),
		),

		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),

		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("enter", "tab", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			key.WithHelp("enter", "apply filter"),
		),

		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
	}
}

func New(items []string, delegate ItemDelegate, width, height int) listModel {
	filterInput := textinput.New()
	filterInput.Prompt = filter.Render("Filter: ")
	filterInput.PromptStyle = focusedStyle
	filterInput.Cursor.Style = cursorStyle
	filterInput.CharLimit = 64
	filterInput.Focus()

	p := paginator.New()
	p.Type = paginator.Dots
	p.ActiveDot = subduedStyle.Render("•")
	p.InactiveDot = verySubduedStyle.Render("○")

	m := listModel{
		showFilter:            true,
		showStatusBar:         true,
		showPagination:        true,
		showHelp:              true,
		filteringEnabled:      true,
		KeyMap:                DefaultKeyMap(),
		Filter:                DefaultFilter,
		FilterInput:           filterInput,
		StatusMessageLifetime: time.Second,

		width:     width,
		height:    height,
		delegate:  delegate,
		items:     items,
		Paginator: p,
		Help:      help.New(),
	}

	m.updatePagination()
	m.updateKeybindings()
	return m
}

func (m listModel) VisibleItems() []string {
	if m.filterState != Unfiltered {
		return m.filteredItems.items()
	}
	return m.items
}

func (m listModel) Index() int {
	return m.Paginator.Page*m.Paginator.PerPage + m.cursor
}

func (m listModel) FilterState() FilterState {
	return m.filterState
}

func (m *listModel) SetWidth(v int) {
	m.setSize(v, m.height)
}

func (m *listModel) SetHeight(v int) {
	m.setSize(m.width, v)
}

func (m *listModel) setSize(width, height int) {
	m.width = width
	m.height = height
	m.Help.Width = width
	m.updatePagination()
}

func (m *listModel) resetFiltering() {
	if m.filterState == Unfiltered {
		return
	}

	m.filterState = Unfiltered
	m.FilterInput.Reset()
	m.filteredItems = nil
	m.updatePagination()
	m.updateKeybindings()
}

func (m listModel) itemsAsFilterItems() filteredItems {
	fi := make([]filteredItem, len(m.items))
	for i, item := range m.items {
		fi[i] = filteredItem{
			item: item,
		}
	}
	return fi
}

func (m *listModel) updateKeybindings() {
	switch m.filterState {
	case Filtering:
		m.KeyMap.NextPage.SetEnabled(false)
		m.KeyMap.PrevPage.SetEnabled(false)
		m.KeyMap.Filter.SetEnabled(false)
		m.KeyMap.ClearFilter.SetEnabled(false)
		m.KeyMap.CancelWhileFiltering.SetEnabled(true)
		m.KeyMap.AcceptWhileFiltering.SetEnabled(m.FilterInput.Value() != "")
		m.KeyMap.Quit.SetEnabled(false)

	default:
		hasItems := len(m.items) != 0
		hasPages := m.Paginator.TotalPages > 1
		m.KeyMap.NextPage.SetEnabled(hasPages)
		m.KeyMap.PrevPage.SetEnabled(hasPages)
		m.KeyMap.Filter.SetEnabled(m.filteringEnabled && hasItems)
		m.KeyMap.ClearFilter.SetEnabled(m.filterState == FilterApplied)
		m.KeyMap.CancelWhileFiltering.SetEnabled(false)
		m.KeyMap.AcceptWhileFiltering.SetEnabled(false)
		m.KeyMap.Quit.SetEnabled(!m.disableQuitKeybindings)
	}
}

func (m *listModel) updatePagination() {
	index := m.Index()

	m.Paginator.PerPage = 10

	if pages := len(m.VisibleItems()); pages < 1 {
		m.Paginator.SetTotalPages(1)
	} else {
		m.Paginator.SetTotalPages(pages)
	}

	m.Paginator.Page = index / m.Paginator.PerPage
	m.cursor = index % m.Paginator.PerPage

	if m.Paginator.Page >= m.Paginator.TotalPages-1 {
		m.Paginator.Page = max(0, m.Paginator.TotalPages-1)
	}
}

func (m *listModel) hideStatusMessage() {
	m.statusMessage = ""
	if m.statusMessageTimer != nil {
		m.statusMessageTimer.Stop()
	}
}

func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case FilterMatchesMsg:
		m.filteredItems = filteredItems(msg)
		return m, nil

	case statusMessageTimeoutMsg:
		m.hideStatusMessage()
	}

	if m.filterState == Filtering {
		cmds = append(cmds, m.handleFiltering(msg))
	} else {
		cmds = append(cmds, m.handleBrowsing(msg))
	}

	return m, tea.Batch(cmds...)
}

func (m *listModel) handleBrowsing(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.KeyMap.Quit):
			return tea.Quit

		case key.Matches(msg, m.KeyMap.PrevPage):
			m.Paginator.PrevPage()

		case key.Matches(msg, m.KeyMap.NextPage):
			m.Paginator.NextPage()

		case key.Matches(msg, m.KeyMap.Filter):
			m.hideStatusMessage()
			if m.FilterInput.Value() == "" {
				m.filteredItems = m.itemsAsFilterItems()
			}
			m.Paginator.Page = 0
			m.cursor = 0
			m.filterState = Filtering
			m.FilterInput.CursorEnd()
			m.FilterInput.Focus()
			m.updateKeybindings()
			return textinput.Blink

		}
	}

	cmd := m.delegate.Update(msg, m)
	cmds = append(cmds, cmd)

	itemsOnPage := m.Paginator.ItemsOnPage(len(m.VisibleItems()))
	if m.cursor > itemsOnPage-1 {
		m.cursor = max(0, itemsOnPage-1)
	}

	return tea.Batch(cmds...)
}

func (m *listModel) handleFiltering(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	if msg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(msg, m.KeyMap.CancelWhileFiltering):
			m.resetFiltering()
			m.KeyMap.Filter.SetEnabled(true)
			m.KeyMap.ClearFilter.SetEnabled(false)

		case key.Matches(msg, m.KeyMap.AcceptWhileFiltering):
			m.hideStatusMessage()

			if len(m.items) == 0 {
				break
			}

			h := m.VisibleItems()

			if len(h) == 0 {
				m.resetFiltering()
				break
			}

			m.FilterInput.Blur()
			m.filterState = FilterApplied
			m.updateKeybindings()

			if m.FilterInput.Value() == "" {
				m.resetFiltering()
			}
		}
	}

	newFilterInputModel, inputCmd := m.FilterInput.Update(msg)
	filterChanged := m.FilterInput.Value() != newFilterInputModel.Value()
	m.FilterInput = newFilterInputModel
	cmds = append(cmds, inputCmd)

	if filterChanged {
		cmds = append(cmds, filterItems(*m))
		m.KeyMap.AcceptWhileFiltering.SetEnabled(m.FilterInput.Value() != "")
	}

	m.updatePagination()

	return tea.Batch(cmds...)
}

func (m listModel) ShortHelp() []key.Binding {
	kb := []key.Binding{}

	filtering := m.filterState == Filtering

	if !filtering {
		if b, ok := m.delegate.(help.KeyMap); ok {
			kb = append(kb, b.ShortHelp()...)
		}
	}

	kb = append(kb,
		m.KeyMap.Filter,
	)

	if !filtering && m.AdditionalShortHelpKeys != nil {
		kb = append(kb, m.AdditionalShortHelpKeys()...)
	}

	return append(kb,
		m.KeyMap.Quit,
	)
}

func (m listModel) FullHelp() [][]key.Binding {
	kb := [][]key.Binding{{
		m.KeyMap.NextPage,
		m.KeyMap.PrevPage,
	}}

	listLevelBindings := []key.Binding{
		m.KeyMap.Filter,
	}

	return append(kb,
		listLevelBindings,
		[]key.Binding{
			m.KeyMap.Quit,
		})
}

func (m listModel) View() string {
	var (
		sections    []string
		availHeight = m.height
	)

	if m.showFilter && m.filteringEnabled {
		v := m.filterView()
		sections = append(sections, v)
		availHeight -= lipgloss.Height(v)
	}

	if m.showStatusBar {
		v := m.statusView()
		sections = append(sections, v)
		availHeight -= lipgloss.Height(v)
	}

	var pagination string
	if m.showPagination {
		pagination = m.paginationView()
		availHeight -= lipgloss.Height(pagination)
	}

	var help string
	if m.showHelp {
		help = m.helpView()
		availHeight -= lipgloss.Height(help)
	}

	content := lipgloss.NewStyle().Height(availHeight).Render(m.populatedView())
	sections = append(sections, content)

	if m.showPagination {
		sections = append(sections, pagination)
	}

	if m.showHelp {
		sections = append(sections, help)
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m listModel) filterView() string {
	var view string

	if m.showFilter && m.filterState == Filtering {
		view += m.FilterInput.View()
	}

	return view
}

func (m listModel) statusView() string {
	var status string

	visibleItems := len(m.VisibleItems())

	itemsDisplay := fmt.Sprintf("count: %d", visibleItems)

	if m.filterState == Filtering {
		if visibleItems == 0 {
			status = statusEmpty.Render("Nothing matched")
		} else {
			status = itemsDisplay
		}
	} else if len(m.items) == 0 {
		status = statusEmpty.Render("No " + m.itemNamePlural)
	} else {
		filtered := m.FilterState() == FilterApplied

		if filtered {
			f := strings.TrimSpace(m.FilterInput.Value())
			f = truncate.StringWithTail(f, 10, "…")
			status += fmt.Sprintf("“%s” ", f)
		}

		status += itemsDisplay
	}

	return statusBar.Render(status)
}

func (m listModel) paginationView() string {
	if m.Paginator.TotalPages < 2 {
		return ""
	}

	s := m.Paginator.View()

	style := paginationStyle
	if m.delegate.Spacing() == 0 && style.GetMarginTop() == 0 {
		style = style.Copy().MarginTop(1)
	}

	return style.Render(s)
}

func (m listModel) populatedView() string {
	items := m.VisibleItems()

	var b strings.Builder

	if len(items) == 0 {
		if m.filterState == Filtering {
			return ""
		}
		return noItems.Render("No " + m.itemNamePlural + ".")
	}

	if len(items) > 0 {
		start, end := m.Paginator.GetSliceBounds(len(items))
		docs := items[start:end]

		for i, item := range docs {
			m.delegate.Render(&b, m, item)
			if i != len(docs)-1 {
				fmt.Fprint(&b, strings.Repeat("\n", m.delegate.Spacing()+1))
			}
		}
	}

	itemsOnPage := m.Paginator.ItemsOnPage(len(items))
	if itemsOnPage < m.Paginator.PerPage {
		n := (m.Paginator.PerPage - itemsOnPage) * (m.delegate.Height() + m.delegate.Spacing())
		if len(items) == 0 {
			n -= m.delegate.Height() - 1
		}
		fmt.Fprint(&b, strings.Repeat("\n", n))
	}

	return b.String()
}

func (m listModel) helpView() string {
	return helpStyle.Render(m.Help.View(m))
}

func filterItems(m listModel) tea.Cmd {
	return func() tea.Msg {
		if m.FilterInput.Value() == "" || m.filterState == Unfiltered {
			return FilterMatchesMsg(m.itemsAsFilterItems())
		}

		items := m.items
		targets := make([]string, len(items))

		copy(targets, items)

		filterMatches := []filteredItem{}
		for _, r := range m.Filter(m.FilterInput.Value(), targets) {
			filterMatches = append(filterMatches, filteredItem{
				item:    items[r.Index],
				matches: r.MatchedIndexes,
			})
		}

		return FilterMatchesMsg(filterMatches)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                            { return 1 }
func (d itemDelegate) Spacing() int                           { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *listModel) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m listModel, listItem string) {

	str := fmt.Sprintf(" %s", listItem)

	fmt.Fprint(w, str)
}

type ListModel struct {
	list     listModel
	quitting bool
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListModel) View() string {

	return "\n" + m.list.View()
}

func List(items []string) error {

	l := New(items, itemDelegate{}, 0, 0)

	m := ListModel{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	if m.quitting {
		return fmt.Errorf("cancelled")
	}

	return nil
}
