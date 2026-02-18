package tests_setup

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/cmd"
	"github.com/cubbit/composer-cli/src/configuration"
	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func GetTestRunner(opts ...RunnerOption) (CLIRunner, error) {
	runnerBase := &RunnerBase{}
	for _, opt := range opts {
		if err := opt(runnerBase); err != nil {
			return nil, err
		}
	}

	binPath := os.Getenv("CLI_BIN_PATH")
	if binPath != "" {
		return &BinaryRunner{
			RunnerBase: *runnerBase,
			BinPath:    binPath,
		}, nil
	}

	configuration, err := configuration.LoadConfig()
	if err != nil {
		return nil, err
	}

	authAPI := api.NewAuthAPI(configuration)
	operatorAPI := api.NewOperatorAPI(configuration)
	locationAPI := api.NewLocationAPI()
	userAPI := api.NewUserAPI()

	agentService := service.NewAgentService(configuration)
	authService := service.NewAuthService(configuration, authAPI)
	operatorService := service.NewOperatorService(configuration, operatorAPI, userAPI)
	locationService := service.NewLocationService(configuration, locationAPI, userAPI)

	return &InProcessRunner{
		RunnerBase: *runnerBase,
		Root: cmd.NewRootCommand(
			agentService,
			authService,
			operatorService,
			locationService,
		),
	}, nil
}

func NewInProcessRunner(rootCmd *cobra.Command) *InProcessRunner {
	return &InProcessRunner{
		Root: rootCmd,
	}
}

func (r *InProcessRunner) Run(args ...string) (string, string, error) {
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	r.Root.SetOut(outBuf)
	r.Root.SetErr(errBuf)
	r.Root.SetArgs(args)

	err := r.Root.Execute()
	return outBuf.String(), errBuf.String(), err
}

func NewBinaryRunner(binPath string) *BinaryRunner {
	return &BinaryRunner{
		BinPath: binPath,
	}
}

func (r *BinaryRunner) Run(args ...string) (string, string, error) {
	cmd := exec.Command(r.BinPath, args...)
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	cmd.Stdout = outBuf
	cmd.Stderr = errBuf

	err := cmd.Run()
	return outBuf.String(), errBuf.String(), err
}
