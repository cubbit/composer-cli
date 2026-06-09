package cmd_swarm

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestSwarmSubCmd_Structure_Describe_WithPositionalID(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm described successfully")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"describe", "swarm-123"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarm described successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestSwarmSubCmd_Structure_Describe_WithSwarmName(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm described successfully by name")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"describe", "--swarm-name", "test-swarm"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarm described successfully by name\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestSwarmSubCmd_Structure_Describe_WithSwarmIDFlag(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm described successfully by ID flag")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"describe", "--swarm-id", "swarm-123"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarm described successfully by ID flag\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestSwarmSubCmd_Structure_Describe_AliasInfo(t *testing.T) {
	mockService := service.NewSwarmServiceMock()
	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Swarm described via alias")
		return nil
	}

	swarmCmd := NewSwarmCmd(mockService)

	commandOutput := new(bytes.Buffer)
	swarmCmd.SetOut(commandOutput)
	swarmCmd.SetErr(commandOutput)
	swarmCmd.SetArgs([]string{"info", "--swarm-id", "swarm-123"})

	err := swarmCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Swarm described via alias\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
