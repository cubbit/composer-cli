package interactive

import (
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/cubbit/composer-cli/utils/interactive/tui/input"
	"github.com/cubbit/composer-cli/utils/interactive/tui/status"
	"github.com/cubbit/composer-cli/utils/printer/text"
)

// ProgOpts is an alias for a slice of tea.ProgramOption, useful for config.
type ProgOpts []tea.ProgramOption

type Config struct {
	Stdout   io.Writer
	Stdin    io.Reader
	Steps    []string
	ProgOpts ProgOpts
}

type Interactive struct {
	config   Config
	progOpts []tea.ProgramOption
	tracker  *StepTracker
}

func New(cfg Config) *Interactive {
	if cfg.Stdout == nil {
		cfg.Stdout = os.Stdout
	}
	ic := &Interactive{
		config:   cfg,
		progOpts: []tea.ProgramOption{tea.WithOutput(cfg.Stdout)},
	}
	if cfg.Stdin != nil {
		ic.progOpts = append(ic.progOpts, tea.WithInput(cfg.Stdin))
	}
	ic.progOpts = append(ic.progOpts, cfg.ProgOpts...)
	if len(cfg.Steps) > 0 {
		ic.tracker = &StepTracker{
			steps:  cfg.Steps,
			writer: cfg.Stdout,
		}
	}
	return ic
}

// NextStep advances to the next step and writes the step header.
// Has no effect when Steps were not configured.
func (ic *Interactive) NextStep() {
	if ic.tracker != nil {
		ic.tracker.NextStep()
	}
}

func (ic *Interactive) Input(title string, opts ...input.FieldOption) (string, error) {
	return input.InputWithOpts(ic.prefix(title), ic.progOpts, opts...)
}

func (ic *Interactive) Secret(title string, opts ...input.FieldOption) (string, error) {
	return input.SecretWithOpts(ic.prefix(title), ic.progOpts, opts...)
}

func (ic *Interactive) MultiInput(title string, fields ...input.Field) (map[string]string, error) {
	return input.MultiInputWithOpts(ic.prefix(title), fields, ic.progOpts...)
}

func (ic *Interactive) Select(title string, options []string) (string, error) {
	return input.SelectWithOpts(ic.prefix(title), options, ic.progOpts...)
}

func (ic *Interactive) MultiSelect(title string, options []string) ([]string, error) {
	return input.MultiSelectWithOpts(ic.prefix(title), options, ic.progOpts...)
}

func (ic *Interactive) Confirm(title string) (bool, error) {
	return input.ConfirmWithOpts(ic.prefix(title), ic.progOpts...)
}

func (ic *Interactive) Spin(title string, action func() error) error {
	return status.SpinWithOpts(ic.prefix(title), action, ic.progOpts...)
}

// StartSpinner begins a spinner immediately and returns a handle.
// Call SetText on the handle to update the label, then Stop to dismiss.
func (ic *Interactive) StartSpinner(title string) *status.SpinnerHandle {
	return status.StartSpinner(ic.prefix(title), ic.progOpts...)
}

// StartProgress begins a progress bar immediately and returns a handle.
// Call Set on the handle to advance, then Stop to complete.
func (ic *Interactive) StartProgress(title string) *status.ProgressHandle {
	return status.StartProgress(ic.prefix(title), ic.progOpts...)
}

func (ic *Interactive) ProgressBar(title string, worker func(chan<- float64) error) error {
	return status.ProgressBarWithOpts(ic.prefix(title), worker, ic.progOpts...)
}

func (ic *Interactive) ProgressWithStatus(title string, worker func(func(float64, string)) error) error {
	return status.ObserveProgressWithOpts(ic.prefix(title), worker, ic.progOpts...)
}

func Observe[T any](ic *Interactive, title string, action func(func(string)) (T, error)) (T, error) {
	return status.ObserveWithOpts(ic.prefix(title), action, ic.progOpts...)
}

func (ic *Interactive) prefix(title string) string {
	if ic.tracker == nil {
		return title
	}
	return ic.tracker.Header() + "\n" + title
}

func (ic *Interactive) Success(msg string) string {
	s := successMark.String() + " " + text.CreateText(msg)
	ic.print(s)
	return s
}

func (ic *Interactive) Error(msg string) string {
	s := errorMark.String() + " " + text.CreateText(msg)
	ic.print(s)
	return s
}

func (ic *Interactive) Info(msg string) string {
	s := infoMark.String() + " " + text.CreateText(msg)
	ic.print(s)
	return s
}

func (ic *Interactive) print(s string) {
	ic.config.Stdout.Write([]byte(s + "\n"))
}

func (ic *Interactive) NewStepTracker(steps []string) *StepTracker {
	return &StepTracker{
		steps:  steps,
		writer: ic.config.Stdout,
	}
}

var ErrCancelled = input.ErrCancelled
