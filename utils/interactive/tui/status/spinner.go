package status

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type spinModel struct {
	spin  spinner.Model
	title string
	done  bool
}

func (m spinModel) Init() tea.Cmd { return m.spin.Tick }

func (m spinModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd
	case statusMsg:
		m.title = msg.text
		return m, nil
	case doneMsg:
		m.done = true
		return m, tea.Quit
	case tea.KeyMsg:
		m.done = true
		return m, tea.Quit
	}
	return m, nil
}

func (m spinModel) View() string {
	if m.done {
		return ""
	}
	return fmt.Sprintf("%s %s", m.spin.View(), m.title)
}

type statusMsg struct {
	text string
}

type doneMsg struct{}

func newSpin(title string) spinModel {
	s := spinner.New()
	s.Style = spinnerColor
	s.Spinner = spinner.MiniDot

	return spinModel{spin: s, title: title}
}

// SpinnerHandle controls a live animated spinner. Create one with StartSpinner,
// call SetText to update the displayed text, and Stop to dismiss.
type SpinnerHandle struct {
	program *tea.Program
	done    chan struct{}
}

// SetText updates the spinner's displayed status text in real time.
func (h *SpinnerHandle) SetText(text string) {
	h.program.Send(statusMsg{text: text})
}

// Stop dismisses the spinner and blocks until the UI is fully gone.
func (h *SpinnerHandle) Stop() {
	h.program.Send(doneMsg{})
	<-h.done
}

// StartSpinner begins an animated spinner immediately and returns a handle.
// Call SetText(...) to update the label, then Stop() to dismiss.
//
// Usage:
//
//	h := StartSpinner("Deploying")
//	h.SetText("Creating resources...")
//	// ... do work
//	h.SetText("Done")
//	h.Stop()
func StartSpinner(title string, progOpts ...tea.ProgramOption) *SpinnerHandle {
	m := newSpin(title)
	p := tea.NewProgram(m, progOpts...)
	done := make(chan struct{})
	go func() {
		p.Run()
		close(done)
	}()
	return &SpinnerHandle{program: p, done: done}
}

// Spin runs an action while showing an animated spinner with the given title.
// It returns the error from the action, or nil on success.
func Spin(title string, action func() error) error {
	return SpinWithOpts(title, action)
}

// SpinWithOpts is like Spin but accepts tea.ProgramOption values for testing
// (e.g. tea.WithInput, tea.WithOutput).
func SpinWithOpts(title string, action func() error, progOpts ...tea.ProgramOption) error {
	m := newSpin(title)
	p := tea.NewProgram(m, progOpts...)

	errCh := make(chan error, 1)
	go func() {
		errCh <- action()
		p.Send(doneMsg{})
	}()

	_, err := p.Run()
	if err != nil {
		return err
	}

	return <-errCh
}

// Observe runs an action while showing an animated spinner. The action receives
// a callback to update the displayed status text in real time.
// Returns the result value and any error from the action.
//
// Usage:
//
//	status, err := Observe("Deploying", func(update func(string)) (string, error) {
//		update("Creating resources...")
//		resp, err := api.Create(...)
//		if err != nil { return "", err }
//		update("Waiting for completion...")
//		err = api.Poll(...)
//		return resp.ID, err
//	})
func Observe[T any](title string, action func(func(string)) (T, error)) (T, error) {
	return ObserveWithOpts(title, action)
}

// ObserveWithOpts is like Observe but accepts tea.ProgramOption values for testing.
func ObserveWithOpts[T any](title string, action func(func(string)) (T, error), progOpts ...tea.ProgramOption) (T, error) {
	m := newSpin(title)
	p := tea.NewProgram(m, progOpts...)

	type result struct {
		val T
		err error
	}
	resCh := make(chan result, 1)
	go func() {
		v, err := action(func(s string) {
			p.Send(statusMsg{text: s})
		})
		resCh <- result{v, err}
		p.Send(doneMsg{})
	}()

	_, err := p.Run()
	if err != nil {
		var zero T
		return zero, err
	}

	res := <-resCh
	return res.val, res.err
}
