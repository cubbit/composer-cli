package cmd_iam

import (
	"bytes"
	"testing"

	"github.com/cubbit/composer-cli/src/service"
	"github.com/spf13/cobra"
)

func TestIAMUserSubCmd_Structure_Import_WithFile(t *testing.T) {
	mockService := service.NewUserServiceMock()
	mockService.ImportUsersFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Users imported successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"import",
		"--file", "users.json",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Users imported successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_Import_WithSample(t *testing.T) {
	mockService := service.NewUserServiceMock()
	importCalled := false
	mockService.ImportUsersFunc = func(cmd *cobra.Command, args []string) error {
		importCalled = true
		cmd.Println("Mock: Sample generated")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"user", "import", "--sample", "json"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !importCalled {
		t.Fatal("ImportUsers should be called for sample generation")
	}
	expectedOutput := "Mock: Sample generated\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}
func TestIAMUserSubCmd_Structure_List(t *testing.T) {
	mockService := service.NewUserServiceMock()
	mockService.ListUsersFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: Users listed successfully")
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"user", "list"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: Users listed successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMUserSubCmd_Structure_List_WithFilters(t *testing.T) {
	mockService := service.NewUserServiceMock()
	listCalled := false
	mockService.ListUsersFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true

		search, _ := cmd.Flags().GetString("search")
		if search != "alice" {
			t.Fatalf("Expected search flag to be 'alice', got %q", search)
		}

		enabled, _ := cmd.Flags().GetString("enabled")
		if enabled != "true" {
			t.Fatalf("Expected enabled flag to be 'true', got %q", enabled)
		}

		sortKey, _ := cmd.Flags().GetString("sort-key")
		if sortKey != "username" {
			t.Fatalf("Expected sort-key flag to be 'username', got %q", sortKey)
		}

		sortOrder, _ := cmd.Flags().GetString("sort-order")
		if sortOrder != "asc" {
			t.Fatalf("Expected sort-order flag to be 'asc', got %q", sortOrder)
		}

		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{
		"user",
		"list",
		"--search", "alice",
		"--enabled", "true",
		"--sort-key", "username",
		"--sort-order", "asc",
	})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !listCalled {
		t.Fatal("ListUsers should be called")
	}
}

func TestIAMUserSubCmd_Structure_List_Alias(t *testing.T) {
	mockService := service.NewUserServiceMock()
	listCalled := false
	mockService.ListUsersFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(mockService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"user", "ls"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !listCalled {
		t.Fatal("ListUsers should be called via 'ls' alias")
	}
}
