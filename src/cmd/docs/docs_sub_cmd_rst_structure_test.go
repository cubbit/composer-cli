package cmd_docs

import (
	"bytes"
	"testing"
)

func TestDocsSubCmd_Structure_Rst(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"rst",
		"/tmp/test-rst",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("ReST documentation generated")) {
		t.Fatalf("Expected ReST documentation generation message, got %q", commandOutputString)
	}
}

func TestDocsSubCmd_Structure_RstWithDefaultOutput(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"rst",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("ReST documentation generated")) {
		t.Fatalf("Expected ReST documentation generation message, got %q", commandOutputString)
	}
}
