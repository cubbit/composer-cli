package cmd_docs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func NewDocsSubCmdYaml() *cobra.Command {
	var yamlCmd = &cobra.Command{
		Use:   "yaml [output-file]",
		Short: "Generate YAML documentation",
		Long:  "Generate YAML documentation for all commands",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			outputFile := "./commands.yaml"
			if len(args) > 0 {
				outputFile = args[0]
			}

			file, err := os.Create(outputFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			err = doc.GenYaml(cmd.Root(), file)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "YAML documentation generated: %s\n", outputFile)
		},
	}

	return yamlCmd
}
