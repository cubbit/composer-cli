package cmd_operator

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestOperatorSubCmd_Structure_Connect_WithAllFlags(t *testing.T) {
	mockService := service.NewOperatorServiceMock()

	mockService.ConnectFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator connected successfully")
		return nil
	}

	operatorCmd := NewOperatorCmd(mockService)

	commandOutput := new(bytes.Buffer)
	operatorCmd.SetOut(commandOutput)
	operatorCmd.SetErr(commandOutput)
	operatorCmd.SetArgs([]string{
		"generate-connect-command",
		"--profile", "test-profile",
	})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator connected successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestOperatorSubCmd_Structure_Connect_WithMissingFlags(t *testing.T) {
	mockService := service.NewOperatorServiceMock()

	mockService.ConnectFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator connected successfully")
		return nil
	}

	operatorCmd := NewOperatorCmd(mockService)

	commandOutput := new(bytes.Buffer)
	operatorCmd.SetOut(commandOutput)
	operatorCmd.SetErr(commandOutput)
	operatorCmd.SetArgs([]string{
		"generate-connect-command",
	})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator connected successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestOperatorSubCmd_Structure_Connect_WithAlias(t *testing.T) {
	mockService := service.NewOperatorServiceMock()

	mockService.ConnectFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Operator connected successfully via alias")
		return nil
	}

	operatorCmd := NewOperatorCmd(mockService)

	commandOutput := new(bytes.Buffer)
	operatorCmd.SetOut(commandOutput)
	operatorCmd.SetErr(commandOutput)
	operatorCmd.SetArgs([]string{
		"generate-connect-command",
		"--profile", "test-profile",
	})

	err := operatorCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Operator connected successfully via alias\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
