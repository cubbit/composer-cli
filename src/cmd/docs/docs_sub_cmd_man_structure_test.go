package cmd_docs

import (
	"bytes"
	"testing"
)

func TestDocsSubCmd_Structure_Man(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"man",
		"/tmp/test-man",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("Man pages generated")) {
		t.Fatalf("Expected man pages generation message, got %q", commandOutputString)
	}
}

func TestDocsSubCmd_Structure_ManWithDefaultOutput(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"man",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("Man pages generated")) {
		t.Fatalf("Expected man pages generation message, got %q", commandOutputString)
	}
}
