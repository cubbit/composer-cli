package tests_setup

import "github.com/spf13/cobra"

type CLIRunner interface {
	Run(args ...string) (stdout string, stderr string, err error)
}

type RunnerBase struct {
	config *string
}

type RunnerOption = func(base *RunnerBase) error

type InProcessRunner struct {
	RunnerBase
	Root *cobra.Command
}

type BinaryRunner struct {
	RunnerBase
	BinPath string
}
