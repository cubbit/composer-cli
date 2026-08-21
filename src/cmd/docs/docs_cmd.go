package cmd_docs

import (
	"github.com/spf13/cobra"
)

func NewDocsCmd() *cobra.Command {
	var docsCmd = &cobra.Command{
		Use:   "docs",
		Short: "Generate documentation",
		Long:  "Generate documentation for all commands in various formats",
	}

	docsMarkdownSubCmd := NewDocsSubCmdMarkdown()
	docsCmd.AddCommand(docsMarkdownSubCmd)

	docsManSubCmd := NewDocsSubCmdMan()
	docsCmd.AddCommand(docsManSubCmd)

	docsRstSubCmd := NewDocsSubCmdRst()
	docsCmd.AddCommand(docsRstSubCmd)

	docsYamlSubCmd := NewDocsSubCmdYaml()
	docsCmd.AddCommand(docsYamlSubCmd)

	docsTreeSubCmd := NewDocsSubCmdTree()
	docsCmd.AddCommand(docsTreeSubCmd)

	return docsCmd
}
