package cmd_version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the CLI version",
		Long:  "Print the current version of the Cubbit CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("cubbit version %s\n", version)
		},
	}
}
