package status

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSpinWithOpts_Success(t *testing.T) {
	var outBuf bytes.Buffer

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := SpinWithOpts("Working", func() error {
			return nil
		}, tea.WithInput(nil), tea.WithOutput(&outBuf))
		ch <- res{err}
	}()

	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}
}

func TestSpinWithOpts_ReturnsActionError(t *testing.T) {
	var outBuf bytes.Buffer
	want := errors.New("something went wrong")

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := SpinWithOpts("Working", func() error {
			return want
		}, tea.WithInput(nil), tea.WithOutput(&outBuf))
		ch <- res{err}
	}()

	r := <-ch

	if !errors.Is(r.err, want) {
		t.Fatalf("expected error %v, got %v", want, r.err)
	}
}

func TestSpinWithOpts_OutputContainsTitle(t *testing.T) {
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := SpinWithOpts("Connecting to server", func() error {
			time.Sleep(30 * time.Millisecond)
			return nil
		}, tea.WithInput(nil), tea.WithOutput(&term))
		ch <- res{err}
	}()

	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}

	frames := term.renderedFrames()
	if len(frames) == 0 {
		t.Fatal("expected rendered frames, got none")
	}

	found := false
	for _, frame := range frames {
		if strings.Contains(frame, "Connecting to server") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a frame with %q, got %v", "Connecting to server", frames)
	}
}

func TestSpinWithOpts_OutputContainsSpinnerCharacter(t *testing.T) {
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := SpinWithOpts("Working", func() error {
			time.Sleep(30 * time.Millisecond)
			return nil
		}, tea.WithInput(nil), tea.WithOutput(&term))
		ch <- res{err}
	}()

	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}

	frames := term.renderedFrames()
	if len(frames) == 0 {
		t.Fatal("expected rendered frames, got none")
	}

	brailleChars := []string{
		"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
		"⣷", "⣯", "⣟", "⡿", "⢿", "⣻", "⣽", "⣾",
	}
	found := false
	for _, frame := range frames {
		for _, c := range brailleChars {
			if strings.Contains(frame, c) {
				found = true
				break
			}
		}
	}
	if !found {
		t.Errorf("expected a frame with a braille spinner char, got %v", frames)
	}
}
