// Package cmd provides CLI commands for managing IAM.
package cmd

import (
	"github.com/spf13/cobra"
)

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Execute commands in IAM sections",
}

func init() {
	rootCmd.AddCommand(iamCmd)
}
