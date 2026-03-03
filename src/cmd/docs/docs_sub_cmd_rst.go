package cmd_docs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsSubCmdRst() *cobra.Command {
	var rstCmd = &cobra.Command{
		Use:   "rst [output-dir]",
		Short: "Generate ReST documentation",
		Long:  "Generate ReST documentation for all commands",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			outputDir := "./rst"
			if len(args) > 0 {
				outputDir = args[0]
			}

			err := os.MkdirAll(outputDir, 0755)
			if err != nil {
				log.Fatal(err)
			}

			err = doc.GenReSTTree(cmd.Root(), outputDir)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "ReST documentation generated in %s\n", outputDir)
		},
	}

	return rstCmd
}
