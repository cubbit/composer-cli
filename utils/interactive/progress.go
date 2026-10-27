package interactive

import (
	"fmt"
	"io"

	"github.com/charmbracelet/lipgloss"
)

var (
	stepDoneStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	stepActiveStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).SetString("▸")
	stepPendingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).SetString("·")
	stepHeaderStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("111"))
)

// StepTracker tracks progress across a sequence of named steps.
type StepTracker struct {
	steps   []string
	current int
	writer  io.Writer
}

func (s *StepTracker) Header() string {
	return s.stepHeader(s.current)
}

// NextStep advances to the next step.
func (s *StepTracker) NextStep() {
	if s.current < len(s.steps) {
		s.current++
	}
}

// Done returns true when all steps have been completed.
func (s *StepTracker) Done() bool {
	return s.current >= len(s.steps)
}

// Render prints all steps with their status.
func (s *StepTracker) Render() {
	for i := range s.steps {
		s.writer.Write([]byte(s.stepLine(i, i < s.current) + "\n"))
	}
}

func (s *StepTracker) stepHeader(idx int) string {
	return stepHeaderStyle.Render(fmt.Sprintf("%d/%d - %s", idx+1, len(s.steps), s.steps[idx]))
}

func (s *StepTracker) stepLine(idx int, done bool) string {
	if done {
		return "  " + stepDoneStyle.String() + " " + s.steps[idx]
	}
	if idx == s.current {
		return "  " + stepActiveStyle.String() + " " + s.steps[idx]
	}
	return "  " + stepPendingStyle.String() + " " + s.steps[idx]
}
