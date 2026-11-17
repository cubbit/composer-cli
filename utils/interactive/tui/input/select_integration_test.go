package input

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSelectWithOpts_SelectsFirstItem(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("\r")

	result, err := SelectWithOpts("Pick", []string{"alpha", "beta", "gamma"},
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "alpha" {
		t.Errorf("expected 'alpha', got %q", result)
	}
}

func TestSelectWithOpts_NavigateAndSelect(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("j\r")

	result, err := SelectWithOpts("Pick", []string{"alpha", "beta", "gamma"},
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "beta" {
		t.Errorf("expected 'beta', got %q", result)
	}
}

func TestSelectWithOpts_Cancel(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("\x1b")

	_, err := SelectWithOpts("Pick", []string{"alpha", "beta", "gamma"},
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	)
	if err != ErrCancelled {
		t.Fatalf("expected ErrCancelled, got %v", err)
	}
}

func TestMultiSelectWithOpts_SelectItems(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString(" j \r")

	result, err := MultiSelectWithOpts("Pick", []string{"a", "b", "c"},
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 selected items, got %d: %v", len(result), result)
	}
}

func TestConfirmWithOpts_Yes(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("\r")

	confirmed, err := ConfirmWithOpts("Continue?", tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !confirmed {
		t.Error("expected confirmed to be true")
	}
}

func TestConfirmWithOpts_No(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("j\r")

	confirmed, err := ConfirmWithOpts("Continue?", tea.WithInput(&inBuf), tea.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if confirmed {
		t.Error("expected confirmed to be false")
	}
}
