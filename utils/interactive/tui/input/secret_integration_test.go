package input

import (
	"bytes"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func TestSecretWithOpts_ReturnsValue(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("p@ssw0rd\t\r")

	result, err := SecretWithOpts("Password", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "p@ssw0rd" {
		t.Errorf("expected 'p@ssw0rd', got %q", result)
	}
}

func TestSecretWithOpts_EscapeCancels(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("\x1b")

	_, err := SecretWithOpts("Password", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != ErrCancelled {
		t.Fatalf("expected ErrCancelled, got %v", err)
	}
}

func TestSecret_EchoMode(t *testing.T) {
	inputs := make([]textinput.Model, 2)
	for i, f := range []Field{
		{Name: "user"},
		{Name: "pass", Password: true},
	} {
		t := textinput.New()
		if f.Password {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		inputs[i] = t
	}

	if inputs[0].EchoMode != textinput.EchoNormal {
		t.Error("expected normal input to have EchoNormal")
	}
	if inputs[1].EchoMode != textinput.EchoPassword {
		t.Error("expected password field to have EchoPassword")
	}
	if inputs[1].EchoCharacter != '•' {
		t.Errorf("expected password field EchoCharacter '•', got %q", inputs[1].EchoCharacter)
	}
}

func TestSecret_HidesCharactersInView(t *testing.T) {
	model := formModel{
		title:  "Login",
		fields: []Field{{Name: "pass", Password: true}},
	}
	model.inputs = newTextInputs([]Field{{Name: "pass", Password: true}})

	updated, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("secret123")})
	m := updated.(formModel)

	if m.inputs[0].Value() != "secret123" {
		t.Errorf("expected input value 'secret123', got %q", m.inputs[0].Value())
	}

	view := m.View()
	if strings.Contains(view, "secret123") {
		t.Error("expected secret value NOT to appear as plaintext in view")
	}
	if !strings.Contains(view, "•••••••••") {
		t.Errorf("expected secret to be masked with bullet characters in view")
	}
}

func TestSecret_OutputHidesCharacters(t *testing.T) {
	var inBuf bytes.Buffer
	var outBuf bytes.Buffer

	inBuf.WriteString("secret\t\r")

	result, err := SecretWithOpts("Password", []tea.ProgramOption{
		tea.WithInput(&inBuf),
		tea.WithOutput(&outBuf),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "secret" {
		t.Errorf("expected returned value 'secret', got %q", result)
	}
	output := outBuf.String()
	if strings.Contains(output, "secret") {
		t.Error("expected secret value NOT to appear as plaintext in program output")
	}
}
