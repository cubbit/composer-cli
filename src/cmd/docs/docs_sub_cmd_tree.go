package cmd_docs

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewDocsSubCmdTree() *cobra.Command {
	var treeCmd = &cobra.Command{
		Use:   "tree",
		Short: "Show command tree",
		Long:  "Display a tree view of all available commands",
		Run: func(cmd *cobra.Command, args []string) {
			printTree(cmd.Root(), cmd.OutOrStdout(), "", true)
		},
	}

	return treeCmd
}

func printTree(cmd *cobra.Command, out interface{ Write([]byte) (int, error) }, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	fmt.Fprintf(out, "%s%s%s", prefix, connector, cmd.Name())
	if cmd.Short != "" {
		fmt.Fprintf(out, " - %s", cmd.Short)
	}
	fmt.Fprintln(out)

	printFlags(cmd, out, prefix, isLast)

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
		printTree(subcmd, out, childPrefix, isLastChild)
	}
}

func printFlags(cmd *cobra.Command, out interface{ Write([]byte) (int, error) }, prefix string, isParentLast bool) {
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

		fmt.Fprintf(out, "%s%s %s\n", flagPrefix, flagConnector, flag)
	}
}
