package status

import (
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func newTestProgressModel(title string) progressModel {
	return progressModel{
		progress: progress.New(progress.WithSolidFill("69"), progress.WithWidth(40)),
		title:    title,
		percent:  0.0,
	}
}

func TestProgressModel_Init(t *testing.T) {
	m := newTestProgressModel("Uploading")
	cmd := m.Init()
	if cmd != nil {
		t.Error("expected Init to return nil, got non-nil")
	}
}

func TestProgressModel_UpdateProgress(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, _ := m.Update(progressMsg{percent: 0.5})
	m1 := updated.(progressModel)

	if m1.percent != 0.5 {
		t.Errorf("expected percent 0.5, got %f", m1.percent)
	}
	if m1.done {
		t.Error("expected NOT done at 50%")
	}
}

func TestProgressModel_CompleteAtFull(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, cmd := m.Update(progressMsg{percent: 1.0})
	m1 := updated.(progressModel)

	if !m1.done {
		t.Error("expected done when percent reaches 1.0")
	}
	if cmd == nil {
		t.Error("expected a quit command on completion")
	}
}

func TestProgressModel_OverflowClampsToFull(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, cmd := m.Update(progressMsg{percent: 1.5})
	m1 := updated.(progressModel)

	if !m1.done {
		t.Error("expected done when percent exceeds 1.0")
	}
	if cmd == nil {
		t.Error("expected a quit command on overflow")
	}
}

func TestProgressModel_EscapeCancels(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEscape})
	m1 := updated.(progressModel)

	if !m1.done {
		t.Error("expected done to be true after escape")
	}
	if cmd == nil {
		t.Error("expected a quit command after escape")
	}
}

func TestProgressModel_CtrlCCancels(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	m1 := updated.(progressModel)

	if !m1.done {
		t.Error("expected done to be true after ctrl+c")
	}
	if cmd == nil {
		t.Error("expected a quit command after ctrl+c")
	}
}

func TestProgressModel_BarGrowsWithPercentage(t *testing.T) {
	m := newTestProgressModel("Uploading")

	m.percent = 0.0
	view0 := m.View()
	m.percent = 0.5
	view50 := m.View()
	m.percent = 1.0
	view100 := m.View()

	full0 := strings.Count(view0, "█")
	full50 := strings.Count(view50, "█")
	full100 := strings.Count(view100, "█")

	if full0 >= full50 {
		t.Errorf("expected more filled blocks at 50%% than 0%%: 0%%=%d, 50%%=%d", full0, full50)
	}
	if full50 >= full100 {
		t.Errorf("expected more filled blocks at 100%% than 50%%: 50%%=%d, 100%%=%d", full50, full100)
	}

	empty0 := strings.Count(view0, "░")
	empty50 := strings.Count(view50, "░")
	empty100 := strings.Count(view100, "░")

	if empty0 <= empty50 {
		t.Errorf("expected fewer empty blocks at 50%% than 0%%: 0%%=%d, 50%%=%d", empty0, empty50)
	}
	if empty50 <= empty100 {
		t.Errorf("expected fewer empty blocks at 100%% than 50%%: 50%%=%d, 100%%=%d", empty50, empty100)
	}
}

func TestProgressModel_ViewShowsTitle(t *testing.T) {
	m := newTestProgressModel("Uploading")
	m.percent = 0.42

	view := m.View()
	if !strings.Contains(view, "Uploading") {
		t.Errorf("expected view to contain title, got %q", view)
	}
}

func TestProgressModel_ViewEmptyWhenDone(t *testing.T) {
	m := newTestProgressModel("Uploading")
	m.done = true

	if m.View() != "" {
		t.Error("expected empty view when done")
	}
}

func TestProgressModel_ViewShowsStatusText(t *testing.T) {
	m := newTestProgressModel("Uploading")
	m.statusText = "Processing chunk 3/10"

	view := m.View()
	if !strings.Contains(view, "Processing chunk 3/10") {
		t.Errorf("expected view to contain status text, got %q", view)
	}
}

func TestProgressModel_StatusTextUpdates(t *testing.T) {
	m := newTestProgressModel("Uploading")

	updated, _ := m.Update(progressMsg{percent: 0.5, statusText: "Connecting..."})
	m1 := updated.(progressModel)

	if m1.statusText != "Connecting..." {
		t.Errorf("expected statusText 'Connecting...', got %q", m1.statusText)
	}

	updated, _ = m1.Update(progressMsg{percent: 0.8, statusText: "Uploading..."})
	m2 := updated.(progressModel)

	if m2.statusText != "Uploading..." {
		t.Errorf("expected statusText 'Uploading...', got %q", m2.statusText)
		if !strings.Contains(m2.View(), "Uploading...") {
			t.Error("expected view to show updated status text")
		}
	}
}

func TestProgressModel_EmptyStatusTextDoesNotClear(t *testing.T) {
	m := newTestProgressModel("Uploading")
	m.statusText = "Working..."

	updated, _ := m.Update(progressMsg{percent: 0.5})
	m1 := updated.(progressModel)

	if m1.statusText != "Working..." {
		t.Errorf("expected statusText preserved as 'Working...', got %q", m1.statusText)
	}
}

func TestClamp_LowerBound(t *testing.T) {
	if got := clamp(-0.5, 0.0, 1.0); got != 0.0 {
		t.Errorf("clamp(-0.5, 0, 1) = %f, want 0.0", got)
	}
}

func TestClamp_UpperBound(t *testing.T) {
	if got := clamp(1.5, 0.0, 1.0); got != 1.0 {
		t.Errorf("clamp(1.5, 0, 1) = %f, want 1.0", got)
	}
}

func TestClamp_InRange(t *testing.T) {
	if got := clamp(0.5, 0.0, 1.0); got != 0.5 {
		t.Errorf("clamp(0.5, 0, 1) = %f, want 0.5", got)
	}
}
