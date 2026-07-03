package status

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type progressModel struct {
	progress   progress.Model
	title      string
	statusText string
	percent    float64
	done       bool
}

func (m progressModel) Init() tea.Cmd { return nil }

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progressMsg:
		m.percent = msg.percent
		if msg.statusText != "" {
			m.statusText = msg.statusText
		}
		if m.percent >= 1.0 {
			m.done = true
			return m, tea.Quit
		}
		return m, nil
	case tea.KeyMsg:
		m.done = true
		return m, tea.Quit
	}
	return m, nil
}

func (m progressModel) View() string {
	if m.done {
		return ""
	}
	bar := m.progress.ViewAs(m.percent)
	if m.statusText != "" {
		return fmt.Sprintf(
			"%s %s  %s",
			titleStyle.Render(m.title),
			bar,
			m.statusText,
		)
	}
	return fmt.Sprintf(
		"%s %s",
		titleStyle.Render(m.title),
		bar,
	)
}

type progressMsg struct {
	percent    float64
	statusText string
}

// ProgressHandle controls a live progress bar. Create one with StartProgress,
// call Set to update the percentage and optional status text, and Stop to dismiss.
type ProgressHandle struct {
	program *tea.Program
	done    chan struct{}
}

// Set updates the bar to the given percentage (0.0–1.0) and optionally sets
// live status text. Pass an empty string to keep the previous text.
func (h *ProgressHandle) Set(pct float64, statusText string) {
	h.program.Send(progressMsg{percent: clamp(pct, 0.0, 1.0), statusText: statusText})
}

// Stop completes the bar and blocks until the UI is fully gone.
func (h *ProgressHandle) Stop() {
	h.program.Send(progressMsg{percent: 1.0})
	<-h.done
}

// StartProgress begins a progress bar immediately and returns a handle.
// Call Set(pct, text) to advance, then Stop() to complete.
//
// Usage:
//
//	h := StartProgress("Deploying")
//	h.Set(0.3, "Creating resources...")
//	// ... do work
//	h.Set(1.0, "Done")
//	h.Stop()
func StartProgress(title string, progOpts ...tea.ProgramOption) *ProgressHandle {
	p := progress.New(progress.WithSolidFill("69"))
	m := progressModel{progress: p, title: title, percent: 0.0}
	program := tea.NewProgram(m, progOpts...)
	done := make(chan struct{})
	go func() {
		program.Run()
		close(done)
	}()
	return &ProgressHandle{program: program, done: done}
}

// ProgressBar runs a worker function that sends progress updates (float64,
// 0.0 to 1.0) on the returned channel. The function blocks until the worker
// completes or the bar is closed. It returns any error from the worker.
//
// Usage:
//
//	err := ProgressBar("Uploading ...", func(ch chan<- float64) error {
//		for i := 0; i <= steps; i++ {
//			ch <- float64(i) / float64(steps)
//			time.Sleep(someWork)
//		}
//		return nil
//	})
func ProgressBar(title string, worker func(ch chan<- float64) error) error {
	return ProgressBarWithOpts(title, worker)
}

// ProgressBarWithOpts is like ProgressBar but accepts tea.ProgramOption
// values for testing (e.g. tea.WithInput, tea.WithOutput).
func ProgressBarWithOpts(title string, worker func(ch chan<- float64) error, progOpts ...tea.ProgramOption) error {
	p := progress.New(progress.WithSolidFill("69"))
	m := progressModel{progress: p, title: title, percent: 0.0}
	program := tea.NewProgram(m, progOpts...)

	errCh := make(chan error, 1)
	go func() {
		progCh := make(chan float64)
		go func() {
			errCh <- worker(progCh)
			close(progCh)
		}()

		for pct := range progCh {
			program.Send(progressMsg{percent: clamp(pct, 0.0, 1.0)})
		}
		program.Send(progressMsg{percent: 1.0})
	}()

	_, err := program.Run()
	if err != nil {
		return err
	}

	return <-errCh
}

// ObserveProgress runs a worker that sends progress + live status text.
// The worker receives a send function that accepts (percent, statusText).
//
// Usage:
//
//	err := ObserveProgress("Deploying", func(send func(float64, string)) error {
//		send(0.1, "Creating resources...")
//		time.Sleep(1 * time.Second)
//		send(0.5, "Configuring network...")
//		time.Sleep(1 * time.Second)
//		send(1.0, "Done")
//		return nil
//	})
func ObserveProgress(title string, worker func(func(float64, string)) error) error {
	return ObserveProgressWithOpts(title, worker)
}

// ObserveProgressWithOpts is like ObserveProgress but accepts tea.ProgramOption
// values for testing.
func ObserveProgressWithOpts(title string, worker func(func(float64, string)) error, progOpts ...tea.ProgramOption) error {
	p := progress.New(progress.WithSolidFill("69"))
	m := progressModel{progress: p, title: title, percent: 0.0}
	program := tea.NewProgram(m, progOpts...)

	errCh := make(chan error, 1)
	go func() {
		done := make(chan struct{})
		go func() {
			errCh <- worker(func(pct float64, text string) {
				program.Send(progressMsg{percent: clamp(pct, 0.0, 1.0), statusText: text})
			})
			close(done)
		}()
		<-done
		program.Send(progressMsg{percent: 1.0})
	}()

	_, err := program.Run()
	if err != nil {
		return err
	}

	return <-errCh
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
