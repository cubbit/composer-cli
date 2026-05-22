package input

import (
	"errors"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func TestFormModel_Init(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "name"}},
	}
	cmd := m.Init()
	if cmd == nil {
		t.Error("expected Init to return a command, got nil")
	}
}

func TestFormModel_Typing(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "name"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("hello")})
	updatedModel := updated.(formModel)
	if got := updatedModel.inputs[0].Value(); got != "hello" {
		t.Errorf("expected input value 'hello', got %q", got)
	}
}

func TestFormModel_EscapeCancels(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "name"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEsc})
	updatedModel := updated.(formModel)

	if !updatedModel.quitting {
		t.Error("expected quitting to be true after escape")
	}
}

func TestFormModel_CtrlCCancels(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "name"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	updatedModel := updated.(formModel)

	if !updatedModel.quitting {
		t.Error("expected quitting to be true after ctrl+c")
	}
}

func TestFormModel_Navigation(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "a"}, {Name: "b"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "a"}, {Name: "b"}})

	if m.focus != 0 {
		t.Errorf("expected initial focus 0, got %d", m.focus)
	}

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
	m1 := updated.(formModel)
	if m1.focus != 1 {
		t.Errorf("expected focus 1 after tab, got %d", m1.focus)
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	m2 := updated.(formModel)
	if m2.focus != 0 {
		t.Errorf("expected focus 0 after shift+tab, got %d", m2.focus)
	}

	// Tab wraps from last field back to first
	updated, _ = m2.Update(tea.KeyMsg{Type: tea.KeyTab})
	updated, _ = updated.(formModel).Update(tea.KeyMsg{Type: tea.KeyTab})
	m4 := updated.(formModel)
	if m4.focus != 0 {
		t.Errorf("expected focus 0 after wrapping from last field, got %d", m4.focus)
	}

	// Shift+tab from first wraps to last
	updated, _ = m4.Update(tea.KeyMsg{Type: tea.KeyShiftTab})
	m5 := updated.(formModel)
	if m5.focus != len(m5.inputs)-1 {
		t.Errorf("expected focus %d after shift+tab from first, got %d", len(m5.inputs)-1, m5.focus)
	}
}

func TestFormModel_SubmitBlockedOnValidationError(t *testing.T) {
	validateErr := errors.New("too short")
	m := formModel{
		title: "Test",
		fields: []Field{{Name: "name", Validate: func(s string) error {
			if len(s) < 3 {
				return validateErr
			}
			return nil
		}}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("ab")})
	m1 := updated.(formModel)

	m1.validate()
	if m1.errs[0] != validateErr {
		t.Errorf("expected validation error, got %v", m1.errs[0])
	}

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m3 := updated.(formModel)

	if m3.quitting {
		t.Error("expected submit to be blocked when validation errors exist")
	}
}

func TestFormModel_SubmitSucceedsWithoutErrors(t *testing.T) {
	m := formModel{
		title: "Test",
		fields: []Field{{Name: "name", Validate: func(s string) error {
			if len(s) < 1 {
				return errors.New("empty")
			}
			return nil
		}}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("hello")})
	m1 := updated.(formModel)

	updated, cmd := m1.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m3 := updated.(formModel)

	if !m3.quitting {
		t.Error("expected submit to succeed after validation passes")
	}
	if cmd == nil {
		t.Error("expected a quit command on successful submit")
	}
}

func TestFormModel_ViewShowsTitleAndInputs(t *testing.T) {
	m := formModel{
		title:  "My Title",
		fields: []Field{{Name: "name"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	view := m.View()
	if !strings.Contains(view, "My Title") {
		t.Error("expected view to contain title")
	}
	if strings.Contains(view, "Submit") {
		t.Error("expected view to NOT contain Submit button")
	}
}

func TestFormModel_ViewShowsValidationError(t *testing.T) {
	m := formModel{
		title: "Test",
		fields: []Field{{Name: "name", Validate: func(s string) error {
			return errors.New("invalid input")
		}}},
	}
	m.inputs = newTextInputs([]Field{{Name: "name"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("x")})
	m1 := updated.(formModel)
	m1.validate()

	view := m1.View()
	if !strings.Contains(view, "invalid input") {
		t.Error("expected view to contain validation error text")
	}
}

func TestFormModel_ViewEmptyWhenQuitting(t *testing.T) {
	m := formModel{
		title:    "Test",
		fields:   []Field{{Name: "name"}},
		quitting: true,
	}
	view := m.View()
	if view != "" {
		t.Error("expected empty view when quitting")
	}
}

func TestFormModel_MultipleFieldsNavigation(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "a"}, {Name: "b"}, {Name: "c"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "a"}, {Name: "b"}, {Name: "c"}})

	expectedFocus := []int{0, 1, 2, 0}
	for i, exp := range expectedFocus {
		if m.focus != exp {
			t.Errorf("step %d: expected focus %d, got %d", i, exp, m.focus)
		}
		next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
		m = next.(formModel)
	}
}

func TestFormModel_SubmitFromAnyField(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "a"}, {Name: "b"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "a"}, {Name: "b"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("value_a")})
	m1 := updated.(formModel)

	updated, cmd := m1.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m5 := updated.(formModel)

	if !m5.quitting {
		t.Error("expected form to quit after submit")
	}
	if !m5.submitted {
		t.Error("expected submitted to be true")
	}
	if cmd == nil {
		t.Error("expected a quit command on successful submit")
	}
	if m5.inputs[0].Value() != "value_a" {
		t.Errorf("expected field a='value_a', got %q", m5.inputs[0].Value())
	}
}

func TestFormModel_SubmitButtonReturnsValues(t *testing.T) {
	m := formModel{
		title:  "Test",
		fields: []Field{{Name: "a"}, {Name: "b"}},
	}
	m.inputs = newTextInputs([]Field{{Name: "a"}, {Name: "b"}})

	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("value_a")})
	m1 := updated.(formModel)

	updated, _ = m1.Update(tea.KeyMsg{Type: tea.KeyTab})
	m2 := updated.(formModel)

	updated, _ = m2.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("value_b")})
	m3 := updated.(formModel)

	// Enter submits from any field
	updated, _ = m3.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m5 := updated.(formModel)

	if !m5.quitting {
		t.Error("expected form to quit after submit")
	}
	if !m5.submitted {
		t.Error("expected submitted to be true")
	}
	if m5.inputs[0].Value() != "value_a" {
		t.Errorf("expected field a='value_a', got %q", m5.inputs[0].Value())
	}
	if m5.inputs[1].Value() != "value_b" {
		t.Errorf("expected field b='value_b', got %q", m5.inputs[1].Value())
	}
}

func newTextInputs(fields []Field) []textinput.Model {
	inputs := make([]textinput.Model, len(fields))
	for i, f := range fields {
		t := textinput.New()
		t.Placeholder = f.Placeholder
		t.PlaceholderStyle = placeholderStyle
		if f.Password {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		if i == 0 {
			t.Focus()
		}
		inputs[i] = t
	}
	return inputs
}
