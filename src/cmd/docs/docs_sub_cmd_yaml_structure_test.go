package cmd_docs

import (
	"bytes"
	"testing"
)

func TestDocsSubCmd_Structure_Yaml(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"yaml",
		"/tmp/test-commands.yaml",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("YAML documentation generated")) {
		t.Fatalf("Expected YAML documentation generation message, got %q", commandOutputString)
	}
}

func TestDocsSubCmd_Structure_YamlWithDefaultOutput(t *testing.T) {
	docsCmd := NewDocsCmd()

	commandOutput := new(bytes.Buffer)
	docsCmd.SetOut(commandOutput)
	docsCmd.SetErr(commandOutput)
	docsCmd.SetArgs([]string{
		"yaml",
	})

	err := docsCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	commandOutputString := commandOutput.String()

	if !bytes.Contains([]byte(commandOutputString), []byte("YAML documentation generated")) {
		t.Fatalf("Expected YAML documentation generation message, got %q", commandOutputString)
	}
}
