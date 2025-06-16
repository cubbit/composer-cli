package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/pflag"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation",
	Long:  "Generate documentation for all commands in various formats",
}

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

		fmt.Printf("Markdown documentation generated in %s\n", outputDir)
	},
}

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
			Title:   "cubbit-operator-cli",
			Section: "1",
		}

		err = doc.GenManTree(cmd.Root(), header, outputDir)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Man pages generated in %s\n", outputDir)
	},
}

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

		fmt.Printf("ReST documentation generated in %s\n", outputDir)
	},
}

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

		fmt.Printf("YAML documentation generated: %s\n", outputFile)
	},
}

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Show command tree",
	Long:  "Display a tree view of all available commands",
	Run: func(cmd *cobra.Command, args []string) {
		printTree(cmd.Root(), "", true)
	},
}

func printTree(cmd *cobra.Command, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	fmt.Printf("%s%s%s", prefix, connector, cmd.Name())
	if cmd.Short != "" {
		fmt.Printf(" - %s", cmd.Short)
	}
	fmt.Println()

	printFlags(cmd, prefix, isLast)

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	commands := cmd.Commands()
	var visibleCommands []*cobra.Command
	for _, subcmd := range commands {
		if !subcmd.Hidden {
			visibleCommands = append(visibleCommands, subcmd)
		}
	}

	for i, subcmd := range visibleCommands {
		isLastChild := i == len(visibleCommands)-1
		printTree(subcmd, childPrefix, isLastChild)
	}
}

func printFlags(cmd *cobra.Command, prefix string, isParentLast bool) {
	var flags []string

	cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Hidden {
			flagStr := fmt.Sprintf("--%s", flag.Name)
			if flag.Shorthand != "" {
				flagStr = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
			}
			if flag.Value.Type() != "bool" {
				flagStr += fmt.Sprintf(" <%s>", flag.Value.Type())
			}
			if flag.Usage != "" {
				flagStr += fmt.Sprintf(" - %s", flag.Usage)
			}
			flags = append(flags, flagStr)
		}
	})

	var persistentFlags []string
	cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Hidden {
			flagStr := fmt.Sprintf("--%s", flag.Name)
			if flag.Shorthand != "" {
				flagStr = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
			}
			if flag.Value.Type() != "bool" {
				flagStr += fmt.Sprintf(" <%s>", flag.Value.Type())
			}
			if flag.Usage != "" {
				flagStr += fmt.Sprintf(" - %s", flag.Usage)
			}
			persistentFlags = append(persistentFlags, flagStr)
		}
	})

	for i, flag := range flags {
		flagPrefix := prefix
		if isParentLast {
			flagPrefix += "    "
		} else {
			flagPrefix += "│   "
		}

		isLastFlag := i == len(flags)-1 && len(persistentFlags) == 0
		flagConnector := "├── "
		if isLastFlag {
			flagConnector = "└── "
		}

		fmt.Printf("%s%s %s\n", flagPrefix, flagConnector, flag)
	}

	for i, flag := range persistentFlags {
		flagPrefix := prefix
		if isParentLast {
			flagPrefix += "    "
		} else {
			flagPrefix += "│   "
		}

		isLastFlag := i == len(persistentFlags)-1
		flagConnector := "├── "
		if isLastFlag {
			flagConnector = "└── "
		}

		fmt.Printf("%s%s🌍 %s (persistent)\n", flagPrefix, flagConnector, flag)
	}
}

func init() {
	docsCmd.AddCommand(markdownCmd)
	docsCmd.AddCommand(manCmd)
	docsCmd.AddCommand(rstCmd)
	docsCmd.AddCommand(yamlCmd)
	docsCmd.AddCommand(treeCmd)

	rootCmd.AddCommand(docsCmd)
}
