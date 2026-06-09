package cmd_version

import (
	"bytes"
	"testing"
)

func TestVersionCmd_Structure_PrintsVersion(t *testing.T) {
	versionCmd := NewVersionCmd("1.2.3")

	commandOutput := new(bytes.Buffer)
	versionCmd.SetOut(commandOutput)
	versionCmd.SetErr(commandOutput)
	versionCmd.SetArgs([]string{})

	err := versionCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedOutput := "version version 1.2.3\n"
	if commandOutput.String() != expectedOutput {
		t.Fatalf("Expected output %q, got %q", expectedOutput, commandOutput.String())
	}
}

func TestVersionCmd_Structure_RejectsArgs(t *testing.T) {
	versionCmd := NewVersionCmd("1.2.3")

	commandOutput := new(bytes.Buffer)
	versionCmd.SetOut(commandOutput)
	versionCmd.SetErr(commandOutput)
	versionCmd.SetArgs([]string{"unexpected"})

	err := versionCmd.Execute()
	if err == nil {
		t.Fatal("Expected error when passing arguments, got nil")
	}
}
