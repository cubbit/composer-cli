package input

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
)

func newSelectModel(title string, items []string, multi bool) selectModel {
	p := paginator.New()
	p.PerPage = 10
	p.SetTotalPages((len(items) + 9) / 10)
	p.Type = paginator.Dots
	p.ActiveDot = "•"
	p.InactiveDot = "○"

	return selectModel{
		title:    title,
		items:    items,
		selected: make(map[int]bool),
		multi:    multi,
		pager:    p,
	}
}

func TestSelectModel_Init(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, false)
	cmd := m.Init()
	if cmd != nil {
		t.Error("expected Init to return nil, got non-nil")
	}
}

func TestSelectModel_CursorNavigation(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m1 := updated.(selectModel)
	if m1.cursor != 1 {
		t.Errorf("expected cursor 1 after down, got %d", m1.cursor)
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyUp})
	m2 := updated.(selectModel)
	if m2.cursor != 0 {
		t.Errorf("expected cursor 0 after up, got %d", m2.cursor)
	}
}

func TestSelectModel_SelectItem(t *testing.T) {
	m := newSelectModel("Test", []string{"alpha", "beta", "gamma"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m1 := updated.(selectModel)

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m2 := updated.(selectModel)

	if !m2.selected[1] {
		t.Error("expected item at cursor 1 to be selected")
	}
	if !m2.done {
		t.Error("expected single select to complete immediately on space")
	}
}

func TestSelectModel_EnterSelectsItem(t *testing.T) {
	m := newSelectModel("Test", []string{"alpha", "beta"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m1 := updated.(selectModel)

	if !m1.selected[0] {
		t.Error("expected item at cursor 0 to be selected")
	}
	if !m1.done {
		t.Error("expected done to be true after enter")
	}
}

func TestSelectModel_Cancel(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEscape})
	m1 := updated.(selectModel)

	if !m1.done {
		t.Error("expected done to be true after escape")
	}
}

func TestSelectModel_CancelWithCtrlC(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	m1 := updated.(selectModel)

	if !m1.done {
		t.Error("expected done to be true after ctrl+c")
	}
}

func TestSelectModel_QuitWithQ(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	m1 := updated.(selectModel)

	if !m1.done {
		t.Error("expected done to be true after q")
	}
}

func TestSelectModel_CursorBounds(t *testing.T) {
	m := newSelectModel("Test", []string{"x", "y"}, false)
	m.cursor = 1

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m1 := updated.(selectModel)
	if m1.cursor != 1 {
		t.Errorf("expected cursor to stay at last item, got %d", m1.cursor)
	}
}

func TestSelectModel_CursorBoundsUp(t *testing.T) {
	m := newSelectModel("Test", []string{"x", "y"}, false)
	m.cursor = 0

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m1 := updated.(selectModel)
	if m1.cursor != 0 {
		t.Errorf("expected cursor to stay at first item on up, got %d", m1.cursor)
	}
}

func TestSelectModel_ViewShowsTitle(t *testing.T) {
	m := newSelectModel("Pick One", []string{"Option A", "Option B"}, false)

	view := m.View()
	if !strings.Contains(view, "Pick One") {
		t.Error("expected view to contain title")
	}
	if !strings.Contains(view, "Option A") {
		t.Error("expected view to contain Option A")
	}
	if !strings.Contains(view, "Option B") {
		t.Error("expected view to contain Option B")
	}
}

func TestSelectModel_ViewEmptyWhenDone(t *testing.T) {
	m := newSelectModel("Test", []string{"a"}, false)
	m.done = true

	if m.View() != "" {
		t.Error("expected empty view when done")
	}
}

func TestSelectModel_MultiSelectToggle(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, true)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m1 := updated.(selectModel)

	if !m1.selected[0] {
		t.Error("expected item 0 to be selected")
	}
	if m1.done {
		t.Error("expected multi-select NOT to complete on space")
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m2 := updated.(selectModel)

	if m2.selected[0] {
		t.Error("expected item 0 to be deselected after second space")
	}
}

func TestSelectModel_MultiSelectSubmit(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, true)

	// Select item 0 and item 2
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m1 := updated.(selectModel)

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyDown})
	m2 := updated.(selectModel)

	updated, _ = m2.Update(tea.KeyMsg{Type: tea.KeyDown})
	m3 := updated.(selectModel)

	updated, _ = m3.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m4 := updated.(selectModel)

	// Press enter to submit. In multi-select, enter submits without
	// toggling the current item. Items selected with space remain.
	updated, _ = m4.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m5 := updated.(selectModel)

	if !m5.done {
		t.Error("expected done after enter in multi-select")
	}
	if !m5.selected[0] {
		t.Error("expected item 0 to be selected")
	}
	if !m5.selected[2] {
		t.Error("expected item 2 to remain selected")
	}
}

func TestSelectModel_MultiSelectSubmitEmptyShowsError(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, true)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m1 := updated.(selectModel)

	if m1.done {
		t.Error("expected multi-select NOT to complete on enter with nothing selected")
	}
	if !strings.Contains(m1.View(), "at least one option must be selected") {
		t.Error("expected error message in view when submitting empty multi-select")
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(" ")})
	m2 := updated.(selectModel)

	if m2.errMsg != "" {
		t.Error("expected error to clear after selecting an item")
	}

	updated, _ = m2.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m3 := updated.(selectModel)

	if !m3.done {
		t.Error("expected done after enter with selection")
	}
}

func TestSelectModel_KKeyNavigation(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b", "c"}, false)

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	m1 := updated.(selectModel)
	if m1.cursor != 1 {
		t.Errorf("expected cursor 1 after j, got %d", m1.cursor)
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")})
	m2 := updated.(selectModel)
	if m2.cursor != 0 {
		t.Errorf("expected cursor 0 after k, got %d", m2.cursor)
	}
}

func TestSelectModel_ViewNoCheckmarkWhenEmpty(t *testing.T) {
	m := newSelectModel("Test", []string{"alpha", "beta", "gamma"}, false)

	view := m.View()
	if strings.Contains(view, "✓") {
		t.Error("expected no checkmark with nothing selected")
	}
}

func TestSelectModel_CursorMark(t *testing.T) {
	m := newSelectModel("Test", []string{"a", "b"}, false)

	view := m.View()
	if !strings.Contains(view, "> a") {
		t.Errorf("expected cursor to show > marker, got:\n%s", view)
	}
}
