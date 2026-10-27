package status

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func TestProgressBarWithOpts_Success(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer
	inBuf.WriteString(" ")

	err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
		for i := 0; i <= 10; i++ {
			ch <- float64(i) / 10.0
		}
		return nil
	}, tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProgressBarWithOpts_ReturnsWorkerError(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer
	inBuf.WriteString(" ")
	want := errors.New("upload failed")

	err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
		ch <- 0.3
		return want
	}, tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if !errors.Is(err, want) {
		t.Fatalf("expected error %v, got %v", want, err)
	}
}

func TestProgressBarWithOpts_EmptyWorker(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer
	inBuf.WriteString(" ")

	err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
		return nil
	}, tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProgressBarWithOpts_ClampsOverflowValues(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer
	inBuf.WriteString(" ")

	err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
		ch <- 2.0
		ch <- 0.5
		return nil
	}, tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProgressBarWithOpts_OutputContainsTitle(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := ProgressBarWithOpts("Uploading files", func(ch chan<- float64) error {
			time.Sleep(15 * time.Millisecond)
			ch <- 0.0
			time.Sleep(15 * time.Millisecond)
			ch <- 1.0
			return nil
		}, tea.WithInput(inR), tea.WithOutput(&term))
		ch <- res{err}
	}()

	time.Sleep(50 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
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
		if strings.Contains(frame, "Uploading files") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a frame with %q, got %v", "Uploading files", frames)
	}
}

func TestProgressBarWithOpts_OutputContainsPercentage(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
			time.Sleep(15 * time.Millisecond)
			ch <- 0.3
			time.Sleep(30 * time.Millisecond)
			ch <- 1.0
			return nil
		}, tea.WithInput(inR), tea.WithOutput(&term))
		ch <- res{err}
	}()

	time.Sleep(80 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}

	frames := term.renderedFrames()
	if len(frames) == 0 {
		t.Fatal("expected rendered frames, got none")
	}

	found30 := false
	for _, frame := range frames {
		if strings.Contains(frame, "30%") {
			found30 = true
		}
	}
	if !found30 {
		t.Errorf("expected a frame with '30%%', got frames: %v", frames)
	}
}

func TestProgressBarWithOpts_OutputContainsBarCharacters(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := ProgressBarWithOpts("Uploading", func(ch chan<- float64) error {
			time.Sleep(15 * time.Millisecond)
			ch <- 0.5
			time.Sleep(30 * time.Millisecond)
			ch <- 1.0
			return nil
		}, tea.WithInput(inR), tea.WithOutput(&term))
		ch <- res{err}
	}()

	time.Sleep(80 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
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
		if strings.Contains(frame, "█") ||
			strings.Contains(frame, "▓") ||
			strings.Contains(frame, "░") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected a frame with block chars, got %v", frames)
	}
}

func TestProgressBarWithOpts_ProgressUpdatesOverTime(t *testing.T) {
	inR, inW := io.Pipe()
	defer inW.Close()
	var term terminalWriter

	type res struct{ err error }
	ch := make(chan res, 1)

	go func() {
		err := ProgressBarWithOpts("Processing", func(ch chan<- float64) error {
			time.Sleep(15 * time.Millisecond)
			ch <- 0.0
			time.Sleep(30 * time.Millisecond)
			ch <- 0.5
			time.Sleep(30 * time.Millisecond)
			ch <- 1.0
			return nil
		}, tea.WithInput(inR), tea.WithOutput(&term))
		ch <- res{err}
	}()

	time.Sleep(120 * time.Millisecond)
	inW.Write([]byte(" "))
	inW.Close()
	r := <-ch

	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}

	frames := term.renderedFrames()
	if len(frames) == 0 {
		t.Fatal("expected rendered frames, got none")
	}

	found0, found50 := false, false
	for _, frame := range frames {
		if strings.Contains(frame, "0%") {
			found0 = true
		}
		if strings.Contains(frame, "50%") {
			found50 = true
		}
	}
	if !found0 {
		t.Errorf("expected a frame with '0%%', got frames: %v", frames)
	}
	if !found50 {
		t.Errorf("expected a frame with '50%%', got frames: %v", frames)
	}
}
