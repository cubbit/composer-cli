package cmd_domain

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestDomainSubCmd_Structure_Create_WithAllFlags(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain created successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"create",
		"--domain-name", "example.com",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain created successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestDomainSubCmd_Structure_Create_MissingDomainName(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.CreateFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Create should not be called when required flags are missing")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"create",
	})

	err := domainCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --domain-name is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"domain-name\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestDomainSubCmd_Structure_Describe_WithAllFlags(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain described successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"describe",
		"--domain-id", "test-domain-id",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain described successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestDomainSubCmd_Structure_Describe_MissingDomainID(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.DescribeFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Describe should not be called when required flags are missing")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"describe",
	})

	err := domainCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --domain-id is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"domain-id\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestDomainSubCmd_Structure_List(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain list retrieved successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"list",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain list retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestDomainSubCmd_Structure_Delete_WithAllFlags(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.DeleteFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain deleted successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"delete",
		"--domain-id", "test-domain-id",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain deleted successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestDomainSubCmd_Structure_Delete_MissingDomainID(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.DeleteFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Delete should not be called when required flags are missing")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"delete",
	})

	err := domainCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --domain-id is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"domain-id\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestDomainSubCmd_Structure_Verify_WithAllFlags(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.VerifyFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain verified successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
		"--domain-id", "test-domain-id",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain verified successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}

func TestDomainSubCmd_Structure_Verify_MissingDomainID(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.VerifyFunc = func(cmd *cobra.Command, args []string) error {
		t.Fatal("Verify should not be called when required flags are missing")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"verify",
	})

	err := domainCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error when --domain-id is missing, got nil")
	}

	if !strings.Contains(err.Error(), "required flag(s) \"domain-id\" not set") {
		t.Fatalf("Expected error message to indicate missing required flags, got: %v", err)
	}
}

func TestDomainSubCmd_Structure_List_AliasLS(t *testing.T) {
	mockService := service.NewDomainServiceMock()

	mockService.ListFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Domain list retrieved successfully")
		return nil
	}

	domainCmd := NewDomainCmd(mockService)

	commandOutput := new(bytes.Buffer)
	domainCmd.SetOut(commandOutput)
	domainCmd.SetErr(commandOutput)
	domainCmd.SetArgs([]string{
		"ls",
	})

	err := domainCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	expectedOutput := "Mock: Domain list retrieved successfully\n"
	if commandOutputString != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutputString)
	}
}
