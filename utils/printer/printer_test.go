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

	err := PrintTree(cmd, nil, nodes, func(nodes []tree.TreeNode) []tree.TreeNode { return nodes })

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

	err := PrintTree(cmd, nil, nodes, func(nodes []tree.TreeNode) []tree.TreeNode { return nodes })

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

func TestPrintTree_JSON(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "json")

	var out bytes.Buffer
	cmd.SetOut(&out)

	nodes := []tree.TreeNode{
		{Value: "Root"},
	}
	data := TestItem{Name: "Alice", Age: 30, Email: "alice@test.com"}

	err := PrintTree(cmd, nil, data, func(_ TestItem) []tree.TreeNode { return nodes })

	if err != nil {
		t.Errorf("Expected no error in json mode, got: %v", err)
	}

	expected := `{
  "Name": "Alice",
  "Age": 30,
  "Email": "alice@test.com"
}
`
	if out.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestPrintTree_YAML(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "yaml")

	var out bytes.Buffer
	cmd.SetOut(&out)

	nodes := []tree.TreeNode{
		{Value: "Root"},
	}
	data := TestItem{Name: "Bob", Age: 25, Email: "bob@test.com"}

	err := PrintTree(cmd, nil, data, func(_ TestItem) []tree.TreeNode { return nodes })

	if err != nil {
		t.Errorf("Expected no error in yaml mode, got: %v", err)
	}

	expected := `name: Bob
age: 25
email: bob@test.com
`
	if out.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestCreateTable_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintTable(cmd, nil, []TestItem{{Name: "test"}})

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

	err := PrintTable(cmd, nil, []TestItem{{Name: "test", Age: 30, Email: "test@example.com"}},
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

func TestCreateTable_JSON(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "json")

	var out bytes.Buffer
	cmd.SetOut(&out)

	data := []TestItem{
		{Name: "Alice", Age: 30, Email: "alice@test.com"},
		{Name: "Bob", Age: 25, Email: "bob@test.com"},
	}

	err := PrintTable(cmd, nil, data)

	if err != nil {
		t.Errorf("Expected no error in json mode, got: %v", err)
	}

	expected := `[
  {
    "Name": "Alice",
    "Age": 30,
    "Email": "alice@test.com"
  },
  {
    "Name": "Bob",
    "Age": 25,
    "Email": "bob@test.com"
  }
]
`
	if out.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestCreateTable_YAML(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "yaml")

	var out bytes.Buffer
	cmd.SetOut(&out)

	data := []TestItem{
		{Name: "Charlie", Age: 35, Email: "charlie@test.com"},
	}

	err := PrintTable(cmd, nil, data)

	if err != nil {
		t.Errorf("Expected no error in yaml mode, got: %v", err)
	}

	expected := `- name: Charlie
  age: 35
  email: charlie@test.com
`
	if out.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestPrintText_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintText(cmd, nil, "Hello World")

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

	err := PrintText(cmd, nil, "Hello World")

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	expectedResult := strings.TrimSpace("Hello World")
	actualResult := strings.TrimSpace(out.String())
	if actualResult != expectedResult {
		t.Error("Expected text output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}

func TestPrintText_JSON(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "json")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintText(cmd, nil, "Hello World")

	if err != nil {
		t.Errorf("Expected no error in json mode, got: %v", err)
	}

	expected := `{
  "message": "Hello World"
}
`
	if out.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestPrintText_YAML(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "yaml")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := PrintText(cmd, nil, "Hello YAML")

	if err != nil {
		t.Errorf("Expected no error in yaml mode, got: %v", err)
	}

	expected := `message: Hello YAML
`
	if out.String() != expected {
		t.Errorf("YAML snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestComposeStructured_Human(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "false")

	var out bytes.Buffer
	cmd.SetOut(&out)

	err := ComposeStructured(cmd, nil, TestItem{Name: "Alice", Age: 30},
		func() error { return PrintText(cmd, nil, "human output") },
	)

	if err != nil {
		t.Errorf("Expected no error in human mode, got: %v", err)
	}

	actualResult := strings.TrimSpace(out.String())
	if actualResult != "human output" {
		t.Errorf("Expected human text output, got %q", actualResult)
	}
}

func TestComposeStructured_JSON(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("output", "human", "output format")
	cmd.Flags().Set("output", "json")

	var out bytes.Buffer
	cmd.SetOut(&out)

	data := TestItem{Name: "Structured", Age: 42, Email: "s@test.com"}
	unreachable := func() error { t.Error("printFunc should not be called in json mode"); return nil }

	err := ComposeStructured(cmd, nil, data, unreachable)

	if err != nil {
		t.Errorf("Expected no error in json mode, got: %v", err)
	}

	expected := `{
  "Name": "Structured",
  "Age": 42,
  "Email": "s@test.com"
}
`
	if out.String() != expected {
		t.Errorf("JSON snapshot mismatch\nexpected:\n%s\nactual:\n%s", expected, out.String())
	}
}

func TestComposeStructured_Quiet(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("quiet", false, "quiet mode")
	cmd.Flags().Set("quiet", "true")

	var out bytes.Buffer
	cmd.SetOut(&out)

	unreachable := func() error { t.Error("printFunc should not be called in quiet mode"); return nil }

	err := ComposeStructured(cmd, nil, "data", unreachable)

	if err != nil {
		t.Errorf("Expected no error in quiet mode, got: %v", err)
	}

	if out.Len() > 0 {
		t.Errorf("Expected no output in quiet mode, got: %s", out.String())
	}
}
