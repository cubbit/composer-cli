package cmd_auth

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestAuthSubCmd_Structure_LoginWithAllRequiredFlags(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Login successful")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
	})
	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Login successful\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestAuthSubCmd_Structure_LoginWithEndpointsFlag(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	var capturedEndpoints string
	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		val, err := cmd.Flags().GetString("endpoints")
		if err != nil {
			return err
		}
		capturedEndpoints = val
		cmd.Println("Mock: Login successful")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{
		"login",
		"--profile", "test-profile",
		"--endpoints", "/tmp/fake.yaml",
	})
	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if capturedEndpoints != "/tmp/fake.yaml" {
		t.Fatalf("Expected endpoints flag value '/tmp/fake.yaml', got %q", capturedEndpoints)
	}
}

func TestAuthSubCmd_Structure_LoginWithoutProfileFlag(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Login successful")
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{"login"})

	err := authCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error (profile flag should be optional), got %v", err)
	}

	if commandOutput.String() != "Mock: Login successful\n" {
		t.Fatalf("Expected output %q, got %q", "Mock: Login successful\n", commandOutput.String())
	}
}

func TestAuthSubCmd_Structure_LoginProfileFlagIsOptional(t *testing.T) {
	mockAuthService := service.NewAuthServiceMock()

	mockAuthService.LoginFunc = func(cmd *cobra.Command, args []string) error {
		return nil
	}

	authCmd := NewAuthCmd(mockAuthService)

	commandOutput := new(bytes.Buffer)
	authCmd.SetOut(commandOutput)
	authCmd.SetErr(commandOutput)
	authCmd.SetArgs([]string{"login"})

	if err := authCmd.Execute(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	loginCmd, _, err := authCmd.Find([]string{"login"})
	if err != nil {
		t.Fatalf("Expected login command, got %v", err)
	}

	flag := loginCmd.Flags().Lookup("profile")
	if flag == nil {
		t.Fatal("expected profile flag to be defined")
	}

	if _, required := flag.Annotations[cobra.BashCompOneRequiredFlag]; required {
		t.Error("expected the profile flag to be optional")
	}
}
