package cmd_iam

import (
	"bytes"
	"strings"
	"testing"

	apikey "github.com/cubbit/composer-cli/src/service/api_key"
	"github.com/cubbit/composer-cli/src/service/user"
	"github.com/spf13/cobra"
)

func TestIAMAPIKeySubCmd_Structure_Create(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	apiKeyService.CreateAPIKeyFunc = func(cmd *cobra.Command, args []string) error {
		cmd.Println("Mock: API key created successfully")
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"api-key", "create", "--name", "automation"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "Mock: API key created successfully\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestIAMAPIKeySubCmd_Structure_Create_MissingName(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	createCalled := false
	apiKeyService.CreateAPIKeyFunc = func(cmd *cobra.Command, args []string) error {
		createCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)
	iamCmd.SetOut(new(bytes.Buffer))
	iamCmd.SetErr(new(bytes.Buffer))
	iamCmd.SetArgs([]string{"api-key", "create"})

	err := iamCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when --name is missing, got nil")
	}
	if !strings.Contains(err.Error(), `required flag(s) "name" not set`) {
		t.Fatalf("Expected missing name flag error, got %v", err)
	}
	if createCalled {
		t.Fatal("CreateAPIKey should not be called when required flags are missing")
	}
}

func TestIAMAPIKeySubCmd_Structure_Revoke(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	revokeCalled := false
	apiKeyService.RevokeAPIKeyFunc = func(cmd *cobra.Command, args []string) error {
		revokeCalled = true
		cmd.Println("Mock: API key revoked successfully")
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"api-key", "revoke", "--id", "key-001"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !revokeCalled {
		t.Fatal("RevokeAPIKey should be called")
	}
}

func TestIAMAPIKeySubCmd_Structure_Edit(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	editCalled := false
	apiKeyService.EditAPIKeyFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		cmd.Println("Mock: API key edited successfully")
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)

	commandOutput := new(bytes.Buffer)
	iamCmd.SetOut(commandOutput)
	iamCmd.SetErr(commandOutput)
	iamCmd.SetArgs([]string{"api-key", "edit", "--id", "key-001", "--name", "automation"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditAPIKey should be called")
	}
}

func TestIAMAPIKeySubCmd_Structure_EditAlias(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	editCalled := false
	apiKeyService.EditAPIKeyFunc = func(cmd *cobra.Command, args []string) error {
		editCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)
	iamCmd.SetOut(new(bytes.Buffer))
	iamCmd.SetErr(new(bytes.Buffer))
	iamCmd.SetArgs([]string{"api-key", "update", "--id", "key-001", "--enabled=false"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !editCalled {
		t.Fatal("EditAPIKey should be called via 'update' alias")
	}
}

func TestIAMAPIKeySubCmd_Structure_ListAlias(t *testing.T) {
	apiKeyService := apikey.NewAPIKeyServiceMock()
	listCalled := false
	apiKeyService.ListAPIKeysFunc = func(cmd *cobra.Command, args []string) error {
		listCalled = true
		return nil
	}

	iamCmd := NewIAMCmd(user.NewUserServiceMock(), apiKeyService)
	iamCmd.SetOut(new(bytes.Buffer))
	iamCmd.SetErr(new(bytes.Buffer))
	iamCmd.SetArgs([]string{"api-key", "ls"})

	err := iamCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !listCalled {
		t.Fatal("ListAPIKeys should be called via 'ls' alias")
	}
}
