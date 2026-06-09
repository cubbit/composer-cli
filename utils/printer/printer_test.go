package printer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/cubbit/composer-cli/utils/printer/tree"
	"github.com/spf13/cobra"
)

type TestItem struct {
	Name  string
	Age   int
	Email string
}

func TestPrintTree_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	nodes := []tree.TreeNode{
		{Value: "Root"},
	}

	err := PrintTree(cmd, nodes)

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestPrintTree_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	nodes := []tree.TreeNode{
		{Value: "Root"},
	}

	err := PrintTree(cmd, nodes)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace(`
Root
`)

	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestCreateTable_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := CreateTable(cmd, []TestItem{{Name: "test"}})

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestCreateTable_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := CreateTable(cmd, []TestItem{{Name: "test", Age: 30, Email: "test@example.com"}},
		table.WithColumns[TestItem]([]table.Column[TestItem]{
			{Title: "Name"},
			{Title: "Age"},
			{Title: "Email"},
		}),
	)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace(`
╭──────┬─────┬──────────────────╮
│ Name │ Age │ Email            │
├──────┼─────┼──────────────────┤
│ test │ 30  │ test@example.com │
╰──────┴─────┴──────────────────╯
`)

	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestPrintText_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintText(cmd, "Hello World")

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestPrintText_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintText(cmd, "Hello World")

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace("Hello World")
	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected text output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestCompose_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := Compose(cmd,
		func() error { return PrintText(cmd, "test1") },
		func() error { return PrintText(cmd, "test2") },
	)

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestCompose_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := Compose(cmd,
		func() error { return PrintText(cmd, "Line 1") },
		func() error { return PrintText(cmd, "Line 2") },
	)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace("Line 1Line 2")
	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected compose output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestCompose_Heterogeneous_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := Compose(cmd,
		func() error { return PrintText(cmd, "text output") },
		func() error { return CreateTable(cmd, []TestItem{{Name: "test", Age: 30}}) },
		func() error { return PrintTree(cmd, []tree.TreeNode{{Value: "tree"}}) },
	)

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}

func TestCompose_Heterogeneous_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := Compose(cmd,
		func() error { return PrintText(cmd, "Text Section\n") },
		func() error {
			return CreateTable(cmd, []TestItem{
				{Name: "Alice", Age: 30, Email: "alice@example.com"},
			},
				table.WithSuffix[TestItem]("\n"))
		},
		func() error {
			return PrintTree(cmd, []tree.TreeNode{
				{
					Value: "Root",
					Children: []tree.TreeNode{
						{
							Value: "Child 1",
							Children: []tree.TreeNode{
								{Value: "Grandchild 1.1"},
								{Value: "Grandchild 1.2"},
							},
						},
						{
							Value: "Child 2",
							Children: []tree.TreeNode{
								{Value: "Grandchild 2.1"},
							},
						},
					},
				},
			})
		},
	)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace(`Text Section
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ Alice │ 30  │ alice@example.com │
╰───────┴─────┴───────────────────╯
Root
├── Child 1
│   ├── Grandchild 1.1
│   └── Grandchild 1.2
└── Child 2
    └── Grandchild 2.1`)
	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected compose output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}
