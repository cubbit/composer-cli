package cmd_docs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsSubCmdMarkdown() *cobra.Command {
	var markdownCmd = &cobra.Command{
		Use:   "markdown [output-dir]",
		Short: "Generate Markdown documentation",
		Long:  "Generate Markdown documentation for all commands",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			outputDir := "./docs"
			if len(args) > 0 {
				outputDir = args[0]
			}

			err := os.MkdirAll(outputDir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			err = doc.GenMarkdownTree(cmd.Root(), outputDir)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Markdown documentation generated in %s\n", outputDir)
		},
	}

	return markdownCmd
}
