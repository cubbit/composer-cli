package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_LoginEndpoints_FlagRegistered(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("ok")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test",
		"--endpoints", "/tmp/endpoints.yaml",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if commandOutput.String() != "ok\n" {
		t.Fatalf("Expected output %q, got %q", "ok\n", commandOutput.String())
	}
}

func TestAuthSubCmd_LoginEndpoints_FlagValueCaptured(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	var capturedValue string
	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		val, err := cmd.Flags().GetString("endpoints")
		if err != nil {
			return err
		}
		capturedValue = val
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test",
		"--endpoints", "/path/to/custom.yaml",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if capturedValue != "/path/to/custom.yaml" {
		t.Fatalf("Expected captured endpoints value '/path/to/custom.yaml', got %q", capturedValue)
	}
}

func TestAuthSubCmd_LoginEndpoints_FlagOptional(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("ok")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error (flag should be optional), got %v", err)
	}
}

func TestAuthSubCmd_LoginEndpoints_FlagDefaultEmpty(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	var capturedValue string
	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		val, err := cmd.Flags().GetString("endpoints")
		if err != nil {
			return err
		}
		capturedValue = val
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test",
	})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if capturedValue != "" {
		t.Fatalf("Expected empty string default for endpoints flag, got %q", capturedValue)
	}
}
