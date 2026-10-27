package cmd_docs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsSubCmdMan() *cobra.Command {
	var manCmd = &cobra.Command{
		Use:   "man [output-dir]",
		Short: "Generate man pages",
		Long:  "Generate man pages for all commands",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			outputDir := "./man"
			if len(args) > 0 {
				outputDir = args[0]
			}

			err := os.MkdirAll(outputDir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			header := &doc.GenManHeader{
				Title:   "cubbit",
				Section: "1",
			}

			err = doc.GenManTree(cmd.Root(), header, outputDir)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Man pages generated in %s\n", outputDir)
		},
	}

	return manCmd
}
