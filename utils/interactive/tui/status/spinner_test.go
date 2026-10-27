package status

import (
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func newTestSpinModel(title string) spinModel {
	s := spinner.New()
	s.Style = spinnerColor
	s.Spinner = spinner.MiniDot

	return spinModel{spin: s, title: title}
}

func TestSpinModel_Init(t *testing.T) {
	m := newTestSpinModel("Loading")
	cmd := m.Init()
	if cmd == nil {
		t.Error("expected Init to return a command, got nil")
	}
}

func TestSpinModel_TickUpdatesSpinner(t *testing.T) {
	m := newTestSpinModel("Loading")
	before := m.spin.View()

	updated, _ := m.Update(spinner.TickMsg{})
	m1 := updated.(spinModel)

	if m1.spin.View() == before {
		t.Error("expected spinner frame to change after TickMsg")
	}
}

func TestSpinModel_DoneMsgQuits(t *testing.T) {
	m := newTestSpinModel("Loading")

	updated, cmd := m.Update(doneMsg{})
	m1 := updated.(spinModel)

	if !m1.done {
		t.Error("expected done to be true after doneMsg")
	}
	if cmd == nil {
		t.Error("expected a quit command after doneMsg")
	}
}

func TestSpinModel_EscapeCancels(t *testing.T) {
	m := newTestSpinModel("Loading")

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEscape})
	m1 := updated.(spinModel)

	if !m1.done {
		t.Error("expected done to be true after escape")
	}
	if cmd == nil {
		t.Error("expected a quit command after escape")
	}
}

func TestSpinModel_CtrlCCancels(t *testing.T) {
	m := newTestSpinModel("Loading")

	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	m1 := updated.(spinModel)

	if !m1.done {
		t.Error("expected done to be true after ctrl+c")
	}
	if cmd == nil {
		t.Error("expected a quit command after ctrl+c")
	}
}

func TestSpinModel_ViewShowsTitle(t *testing.T) {
	m := newTestSpinModel("Loading data")

	view := m.View()
	if !strings.Contains(view, "Loading data") {
		t.Errorf("expected view to contain title, got %q", view)
	}
}

func TestSpinModel_ViewEmptyWhenDone(t *testing.T) {
	m := newTestSpinModel("Loading")
	m.done = true

	if m.View() != "" {
		t.Error("expected empty view when done")
	}
}

func TestSpinModel_ViewShowsSpinnerFrame(t *testing.T) {
	m := newTestSpinModel("Loading")

	view := m.View()
	frame := m.spin.View()
	if !strings.Contains(view, frame) {
		t.Errorf("expected view to contain spinner frame %q, got %q", frame, view)
	}
}

func TestSpinModel_StatusMsgUpdatesTitle(t *testing.T) {
	m := newTestSpinModel("Loading")

	updated, _ := m.Update(statusMsg{text: "Connecting to server..."})
	m1 := updated.(spinModel)

	if m1.title != "Connecting to server..." {
		t.Errorf("expected title 'Connecting to server...', got %q", m1.title)
	}

	view := m1.View()
	if !strings.Contains(view, "Connecting to server...") {
		t.Errorf("expected view to contain updated title, got %q", view)
	}
}

func TestSpinModel_MultipleStatusMsgUpdates(t *testing.T) {
	m := newTestSpinModel("Loading")

	updated, _ := m.Update(statusMsg{text: "Step 1"})
	m1 := updated.(spinModel)
	updated, _ = m1.Update(statusMsg{text: "Step 2"})
	m2 := updated.(spinModel)

	if m2.title != "Step 2" {
		t.Errorf("expected final title 'Step 2', got %q", m2.title)
	}
}

func TestObserveWithOpts_ReturnsValue(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()

	type res struct {
		val string
		err error
	}
	ch := make(chan res, 1)

	go func() {
		v, err := ObserveWithOpts("Working", func(update func(string)) (string, error) {
			update("Starting...")
			return "hello world", nil
		}, tea.WithInput(inR))
		ch <- res{v, err}
	}()

	time.Sleep(5 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}
	if r.val != "hello world" {
		t.Errorf("expected 'hello world', got %q", r.val)
	}
}

func TestObserveWithOpts_ReturnsError(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()

	type res struct {
		val string
		err error
	}
	ch := make(chan res, 1)

	go func() {
		v, err := ObserveWithOpts("Working", func(update func(string)) (string, error) {
			update("Failing...")
			return "", errors.New("something went wrong")
		}, tea.WithInput(inR))
		ch <- res{v, err}
	}()

	time.Sleep(5 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
	r := <-ch

	if r.err == nil || r.err.Error() != "something went wrong" {
		t.Errorf("expected 'something went wrong', got %v", r.err)
	}
	if r.val != "" {
		t.Errorf("expected empty value on error, got %q", r.val)
	}
}

func TestObserveWithOpts_StatusUpdatesInOutput(t *testing.T) {
	var term terminalWriter
	inR, inW := io.Pipe()
	defer inW.Close()

	type res struct {
		val string
		err error
	}
	ch := make(chan res, 1)

	go func() {
		v, err := ObserveWithOpts("Working",
			func(update func(string)) (string, error) {
				update("Processing...")
				time.Sleep(30 * time.Millisecond)
				update("Finalizing...")
				time.Sleep(10 * time.Millisecond)
				return "done", nil
			},
			tea.WithInput(inR), tea.WithOutput(&term))
		ch <- res{v, err}
	}()

	time.Sleep(50 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}
	if r.val != "done" {
		t.Errorf("expected 'done', got %q", r.val)
	}

	frames := term.renderedFrames()
	foundProcessing := false
	foundFinalizing := false
	for _, frame := range frames {
		if strings.Contains(frame, "Processing...") {
			foundProcessing = true
		}
		if strings.Contains(frame, "Finalizing...") {
			foundFinalizing = true
		}
	}
	if !foundProcessing {
		t.Error("expected a frame with 'Processing...' status update")
	}
	if !foundFinalizing {
		t.Error("expected a frame with 'Finalizing...' status update")
	}
}

func TestObserveWithOpts_CancelledReturnsZeroValue(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()

	type res struct {
		val string
		err error
	}
	ch := make(chan res, 1)

	go func() {
		v, err := ObserveWithOpts("Working",
			func(update func(string)) (string, error) {
				time.Sleep(50 * time.Millisecond)
				return "done", nil
			},
			tea.WithInput(inR))
		ch <- res{v, err}
	}()

	time.Sleep(5 * time.Millisecond)
	inW.Write([]byte("q"))
	inW.Close()
	r := <-ch

	// action runs to completion even after cancel, so result is still returned
	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}
	if r.val != "done" {
		t.Errorf("expected 'done', got %q", r.val)
	}
}
