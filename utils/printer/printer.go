package printer

import (
	"fmt"

	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/text"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	"github.com/spf13/cobra"
)

func shouldPrint(cmd *cobra.Command) (bool, error) {
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return true, err
	}
	return !quiet, nil
}

func PrintTree(cmd *cobra.Command, nodes []tree.TreeNode, opts ...tree.Option) error {
	should, err := shouldPrint(cmd)
	if err != nil || !should {
		return err
	}

	result := tree.CreateTree(nodes, opts...)
	_, err = fmt.Fprint(cmd.OutOrStdout(), result)

	return err
}

func CreateTable[T any](cmd *cobra.Command, data []T, opts ...table.Option[T]) error {
	should, err := shouldPrint(cmd)
	if err != nil || !should {
		return err
	}

	result := table.CreateTable(data, opts...)
	_, err = fmt.Fprint(cmd.OutOrStdout(), result)

	return err
}

func PrintText(cmd *cobra.Command, s string, opts ...text.Option) error {
	should, err := shouldPrint(cmd)
	if err != nil || !should {
		return err
	}
	result := text.CreateText(s, opts...)
	_, err = fmt.Fprint(cmd.OutOrStdout(), result)
	return err
}

func Compose(cmd *cobra.Command, printFuncs ...func() error) error {
	should, err := shouldPrint(cmd)
	if err != nil || !should {
		return err
	}

	for _, fn := range printFuncs {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}
