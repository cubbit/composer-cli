package cmd_docs

import (
	"bytes"
	"testing"
)

func TestDocsSubCmd_Structure_Markdown(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"markdown",
		"/tmp/test-docs",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("Markdown documentation generated")) {
		t.Fatalf("Expected markdown generation message, got %q", commandOutputString)
	}
}

func TestDocsSubCmd_Structure_MarkdownWithDefaultOutput(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"markdown",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("Markdown documentation generated")) {
		t.Fatalf("Expected markdown generation message, got %q", commandOutputString)
	}
}
