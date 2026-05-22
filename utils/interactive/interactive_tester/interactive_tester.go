package interactive_tester

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/creack/pty"
	"github.com/spf13/cobra"
)

var ansiPattern = regexp.MustCompile(`\033\[[\d;]*[a-zA-Z]|\033\][^\033]*\033\\|\033.`)

const defaultExpectTimeout = 10 * time.Second

// ExpectOption configures Expect behavior.
type ExpectOption func(*expectConfig)

type expectConfig struct {
	timeout time.Duration
}

func WithTimeout(timeout time.Duration) ExpectOption {
	return func(c *expectConfig) {
		c.timeout = timeout
	}
}

// InteractiveTester wraps a cobra command with a real PTY pair for interactive
// testing. It starts the command in the background and provides Expect/Write
// helpers to drive terminal-based TUI workflows.
type InteractiveTester struct {
	cmd      *cobra.Command
	ptmx     *os.File
	pts      *os.File
	output   chan string
	execErr  chan error
	readDone chan struct{}
	done     bool

	// Convenience key sequences
	Enter []byte
	Up    []byte
	Down  []byte
	Space []byte
}

type Option func(*InteractiveTester)

func WithTermSize(rows, cols uint16) Option {
	return func(h *InteractiveTester) {
		pty.Setsize(h.ptmx, &pty.Winsize{Rows: rows, Cols: cols})
	}
}

func New(cmd *cobra.Command, args []string, opts ...Option) (*InteractiveTester, error) {
	ptmx, pts, err := pty.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open PTY: %w", err)
	}

	h := &InteractiveTester{
		cmd:      cmd,
		ptmx:     ptmx,
		pts:      pts,
		output:   make(chan string, 512),
		execErr:  make(chan error, 1),
		readDone: make(chan struct{}),
		Enter:    []byte("\r"),
		Up:       []byte("\033[A"),
		Down:     []byte("\033[B"),
		Space:    []byte(" "),
	}

	pty.Setsize(h.ptmx, &pty.Winsize{Rows: 32, Cols: 160})

	for _, opt := range opts {
		opt(h)
	}

	cmd.SetIn(pts)
	cmd.SetOut(pts)
	cmd.SetErr(pts)
	cmd.SetArgs(args)

	return h, nil
}

// Start kicks off the command execution and output reader in background
// goroutines. Must be called before Expect or Write.
func (h *InteractiveTester) Start() {
	go func() {
		if err := h.cmd.Execute(); err != nil {
			h.execErr <- err
		}
		close(h.execErr)
	}()

	go func() {
		defer close(h.readDone)
		buf := make([]byte, 8192)
		for {
			n, err := h.ptmx.Read(buf)
			if n > 0 {
				chunk := make([]byte, n)
				copy(chunk, buf[:n])
				h.output <- string(chunk)
			}
			if err != nil {
				return
			}
		}
	}()
}

// Expect blocks until the given substring appears in the PTY output or the
// timeout elapses (default 10s). The accumulated output is returned in
// both success and failure cases.
func (h *InteractiveTester) Expect(expected string, opts ...ExpectOption) (string, error) {
	cfg := expectConfig{timeout: defaultExpectTimeout}
	for _, o := range opts {
		o(&cfg)
	}

	var accumulated strings.Builder
	deadline := time.After(cfg.timeout)

	for {
		select {
		case <-deadline:
			return accumulated.String(), fmt.Errorf(
				"timeout waiting for %q (accumulated: %s)", expected, accumulated.String(),
			)
		case chunk, ok := <-h.output:
			if !ok {
				return accumulated.String(), fmt.Errorf(
					"PTY output closed while waiting for %q", expected,
				)
			}
			accumulated.WriteString(chunk)
			if strings.Contains(accumulated.String(), expected) {
				return accumulated.String(), nil
			}
		case err, ok := <-h.execErr:
			if ok && err != nil {
				return accumulated.String(), fmt.Errorf(
					"command exited before %q matched: %w", expected, err,
				)
			}
		}
	}
}

// ExpectMultiple blocks until all given substrings appear (in any order) in
// the PTY output, or the timeout elapses. The accumulated output is returned.
func (h *InteractiveTester) ExpectMultiple(expected []string, opts ...ExpectOption) (string, error) {
	cfg := expectConfig{timeout: defaultExpectTimeout}
	for _, o := range opts {
		o(&cfg)
	}

	remaining := make([]string, len(expected))
	copy(remaining, expected)

	var accumulated strings.Builder
	deadline := time.After(cfg.timeout)

	for len(remaining) > 0 {
		select {
		case <-deadline:
			return accumulated.String(), fmt.Errorf(
				"timeout waiting for remaining %v (accumulated: %s)", remaining, accumulated.String(),
			)
		case chunk, ok := <-h.output:
			if !ok {
				return accumulated.String(), fmt.Errorf(
					"PTY output closed while waiting for %v", remaining,
				)
			}
			fmt.Fprintf(os.Stderr, "Received chunk: %q\n", chunk)
			accumulated.WriteString(chunk)
			current := accumulated.String()
			var unmatched []string
			for _, r := range remaining {
				if !strings.Contains(current, r) {
					unmatched = append(unmatched, r)
				}
			}
			remaining = unmatched
		case err, ok := <-h.execErr:
			if ok && err != nil {
				return accumulated.String(), fmt.Errorf(
					"command exited before %v matched: %w", remaining, err,
				)
			}
		}
	}

	return accumulated.String(), nil
}

// WriteData writes raw bytes to the PTY master.
func (h *InteractiveTester) WriteData(data []byte) (int, error) {
	return h.ptmx.Write(data)
}

// WriteLine writes the given string followed by a carriage return (Enter key).
func (h *InteractiveTester) WriteLine(s string) (int, error) {
	return h.ptmx.Write([]byte(s + "\r"))
}

// WriteKeys writes each byte sequence in order, pausing briefly between them.
// Useful for multi-step key sequences like down-arrow → space → enter.
func (h *InteractiveTester) WriteKeys(keys ...[]byte) error {
	for _, k := range keys {
		if _, err := h.ptmx.Write(k); err != nil {
			return err
		}
	}
	return nil
}

// ExpectT is like Expect but calls t.Fatal on error.
func (h *InteractiveTester) ExpectT(t *testing.T, expected string, opts ...ExpectOption) string {
	t.Helper()
	out, err := h.Expect(expected, opts...)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// ExpectMultipleT is like ExpectMultiple but calls t.Fatal on error.
func (h *InteractiveTester) ExpectMultipleT(t *testing.T, expected []string, opts ...ExpectOption) string {
	t.Helper()
	out, err := h.ExpectMultiple(expected, opts...)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// WriteDataT is like WriteData but calls t.Fatal on error.
func (h *InteractiveTester) WriteDataT(t *testing.T, data []byte) int {
	t.Helper()
	n, err := h.WriteData(data)
	if err != nil {
		t.Fatalf("WriteData: %v", err)
	}
	return n
}

// WriteLineT is like WriteLine but calls t.Fatal on error.
func (h *InteractiveTester) WriteLineT(t *testing.T, s string) int {
	t.Helper()
	n, err := h.WriteLine(s)
	if err != nil {
		t.Fatalf("WriteLine: %v", err)
	}
	return n
}

// WriteKeysT is like WriteKeys but calls t.Fatal on error.
func (h *InteractiveTester) WriteKeysT(t *testing.T, keys ...[]byte) {
	t.Helper()
	if err := h.WriteKeys(keys...); err != nil {
		t.Fatalf("WriteKeys: %v", err)
	}
}

// Close releases PTY resources. Safe to call multiple times.

func (h *InteractiveTester) Close() {
	if h.done {
		return
	}
	h.done = true
	h.ptmx.Close()
	h.pts.Close()
}
