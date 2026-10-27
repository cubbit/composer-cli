package cmd_docs

import (
	"bytes"
	"testing"
)

func TestDocsSubCmd_Structure_Tree(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"tree",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("docs")) {
		t.Fatalf("Expected command tree output with 'docs', got %q", commandOutputString)
	}
}
