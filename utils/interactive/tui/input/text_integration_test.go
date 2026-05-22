package input

import (
	"bytes"
	"io"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInputWithOpts_EnterSubmitsDirectly(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	// Enter submits directly without needing to tab to submit button first
	inBuf.WriteString("hello\r")

	result, err := InputWithOpts("Name", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "hello" {
		t.Errorf("expected 'hello', got %q", result)
	}
}

func TestInputWithOpts_TabThenEnterStillWorks(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	// Tab to submit button then enter still works as before
	inBuf.WriteString("hello\t\r")

	result, err := InputWithOpts("Name", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "hello" {
		t.Errorf("expected 'hello', got %q", result)
	}
}

func TestInputWithOpts_CancelWithEscape(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("\x1b")

	_, err := InputWithOpts("Name", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != ErrCancelled {
		t.Fatalf("expected ErrCancelled, got %v", err)
	}
}

func TestInputWithOpts_ProducesOutput(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("hello\r")

	_, err := InputWithOpts("MyPrompt", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if outBuf.Len() == 0 {
		t.Error("expected non-empty output from program")
	}
}

func TestMultiInputWithOpts_MultipleFields(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("Alice\t30\r")

	result, err := MultiInputWithOpts("Details", []Field{
		{Name: "name"},
		{Name: "age"},
	}, tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["name"] != "Alice" {
		t.Errorf("expected name='Alice', got %q", result["name"])
	}
	if result["age"] != "30" {
		t.Errorf("expected age='30', got %q", result["age"])
	}
}

func TestInputWithOpts_PipeInput(t *testing.T) {
	var outBuf bytes.Buffer
	inR, inW := io.Pipe()

	type result struct {
		val string
		err error
	}
	resultCh := make(chan result, 1)

	go func() {
		val, err := InputWithOpts("Name", []tea.ProgramOption{
			tea.WithInput(inR),
			tea.WithOutput(&outBuf),
		})
		resultCh <- result{val, err}
	}()

	inW.Write([]byte("streamed\r"))
	inW.Close()

	r := <-resultCh
	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}
	if r.val != "streamed" {
		t.Errorf("expected 'streamed', got %q", r.val)
	}
}
