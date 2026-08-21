package text

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestCreateText_SimpleString(t *testing.T) {
	result := CreateText("hello")

	if result != "hello" {
		t.Errorf("Expected 'hello', got %q", result)
	}
}

func TestCreateText_WithWhitespace(t *testing.T) {
	result := CreateText("  test  ")

	if result != "  test  " {
		t.Errorf("Expected '  test  ', got %q", result)
	}
}

func TestCreateText_EmptyString(t *testing.T) {
	result := CreateText("")

	if result != "" {
		t.Errorf("Expected empty string, got %q", result)
	}
}

func TestCreateText_WithStyle(t *testing.T) {
	boldStyle := lipgloss.NewStyle().Bold(true)
	result := CreateText("hello", WithStyle(boldStyle))

	if result == "" {
		t.Errorf("Expected non-empty result")
	}

	if !strings.Contains(result, "hello") {
		t.Errorf("Expected result to contain 'hello', got %q", result)
	}
}
