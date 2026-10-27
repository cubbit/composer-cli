package tree

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestCreateTree_NoNodes(t *testing.T) {
	result := CreateTree([]TreeNode{})

	expectedResult := "No data"

	if result != expectedResult {
		t.Error("Expected 'No data' but got:", result)
	}
}

func TestCreateTree_SingleNode(t *testing.T) {
	nodes := []TreeNode{
		{Value: "Root"},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Root
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithChildren(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Children: []TreeNode{
				{Value: "Child1"},
				{Value: "Child2"},
			},
		},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Root
├── Child1
└── Child2
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_Nested(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Children: []TreeNode{
				{
					Value: "Child1",
					Children: []TreeNode{
						{Value: "Grandchild1"},
						{Value: "Grandchild2"},
					},
				},
				{Value: "Child2"},
			},
		},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Root
├── Child1
│   ├── Grandchild1
│   └── Grandchild2
└── Child2
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithGlobalFormatter(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "test",
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(s any) string {
			return s.(string) + " formatted"
		}),
	)

	expectedResult := strings.TrimSpace(`
test formatted
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithNodeFormatter(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "test",
			Formatter: func(v any) string {
				return v.(string) + " node-formatted"
			},
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(s any) string {
			return s.(string) + " global-formatted"
		}),
	)

	expectedResult := strings.TrimSpace(`
test node-formatted
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_NodeFormatterOverrideGlobal(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "test1",
		},
		{
			Value:     "test2",
			Formatter: func(v any) string { return v.(string) + " (custom)" },
		},
		{
			Value: "test3",
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(s any) string {
			return s.(string) + " (global)"
		}),
	)

	expectedResult := strings.TrimSpace(`
test1 (global)
test2 (custom)
test3 (global)
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithCustomStyle(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Children: []TreeNode{
				{Value: "Child"},
			},
		},
	}

	customStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Padding(1)

	result := CreateTree(nodes,
		WithStyle(customStyle),
	)

	if result == "" {
		t.Error("Expected non-empty tree string with custom style")
	}
}

func TestCreateTree_MultipleRoots(t *testing.T) {
	nodes := []TreeNode{
		{Value: "Root1"},
		{Value: "Root2"},
		{Value: "Root3"},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Root1
Root2
Root3
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

type TestStruct struct {
	Name  string
	Value int
}

func TestCreateTree_WithStruct(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: TestStruct{Name: "Item1", Value: 100},
		},
		{
			Value: TestStruct{Name: "Item2", Value: 200},
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(v any) string {
			s := v.(TestStruct)
			return s.Name + ": $" + string(rune('0'+s.Value/100))
		}),
	)

	expectedResult := strings.TrimSpace(`
Item1: $1
Item2: $2
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_MixedTypes(t *testing.T) {
	nodes := []TreeNode{
		{Value: "String node"},
		{Value: 123},
		{
			Value: TestStruct{Name: "Struct", Value: 456},
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(v any) string {
			switch val := v.(type) {
			case string:
				return val + " (str)"
			case int:
				return fmt.Sprintf("%d (int)", val)
			case TestStruct:
				return val.Name + " (struct)"
			default:
				return fmt.Sprintf("%v", val)
			}
		}),
	)

	expectedResult := strings.TrimSpace(`
String node (str)
123 (int)
Struct (struct)
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_MultilineValue(t *testing.T) {
	multilineValue := "Line 1\nLine 2\nLine 3"

	nodes := []TreeNode{
		{
			Value: multilineValue,
		},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Line 1
Line 2
Line 3
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_MultilineWithChildren(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root\nMulti",
			Children: []TreeNode{
				{Value: "Child1"},
				{
					// the x are the padding, default padding is a space in lipgloss tree
					// but they are not feasible for the test
					Value: "Child2\nMultix\nlinexx",
					Children: []TreeNode{
						{Value: "Grandchild1"},
						{Value: "Grandchild2"},
					},
				},
				{Value: "Child3"},
			},
		},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
Root
Multi
├── Child1
├── Child2
│   Multix
│   linexx
│   ├── Grandchild1
│   └── Grandchild2
└── Child3
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_CustomFormatterWithMultiline(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "test",
		},
	}

	customFormatter := func(v any) string {
		s := v.(string)
		return "Prefix:\n" + s + "\nSuffix"
	}

	result := CreateTree(nodes, WithFormatter(customFormatter))

	expectedResult := strings.TrimSpace(`
Prefix:
test
Suffix
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_NodeFormatterWithStyle(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Children: []TreeNode{
				{Value: "Child"},
			},
			Formatter: func(v any) string {
				return "CUSTOM: " + v.(string)
			},
		},
	}

	customStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("196"))

	result := CreateTree(nodes, WithStyle(customStyle))

	if result == "" {
		t.Error("Expected non-empty tree string with custom style")
	}

	if !strings.Contains(result, "CUSTOM: Root") {
		t.Error("Expected custom formatter output not found in result")
	}
}

func TestCreateTree_StructWithNodeFormatter(t *testing.T) {
	type Person struct {
		Name  string
		Age   int
		Email string
	}

	nodes := []TreeNode{
		{
			Value: Person{
				Name:  "John Doe",
				Age:   30,
				Email: "john@example.com",
			},
			Formatter: func(v any) string {
				switch val := v.(type) {
				case Person:
					return fmt.Sprintf("%s (%d years old)", val.Name, val.Age)
				default:
					return fmt.Sprintf("%v", val)
				}
			},
		},
	}

	result := CreateTree(nodes)

	expectedResult := strings.TrimSpace(`
John Doe (30 years old)
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_DeeplyNestedWithFormatters(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Formatter: func(v any) string {
				return "R: " + v.(string)
			},
			Children: []TreeNode{
				{
					Value: "Child",
					Children: []TreeNode{
						{
							Value: "Grandchild",
							Formatter: func(v any) string {
								return "G: " + v.(string)
							},
						},
					},
				},
			},
		},
	}

	result := CreateTree(nodes,
		WithFormatter(func(v any) string {
			return "GLOBAL: " + v.(string)
		}),
	)

	expectedResult := strings.TrimSpace(`
R: Root
└── GLOBAL: Child
    └── G: Grandchild
`)

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithPrefix(t *testing.T) {
	nodes := []TreeNode{
		{Value: "Root"},
	}

	result := CreateTree(nodes,
		WithPrefix("PREFIX: "),
	)

	expectedResult := "PREFIX: Root"

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithSuffix(t *testing.T) {
	nodes := []TreeNode{
		{Value: "Root"},
	}

	result := CreateTree(nodes,
		WithSuffix(" :SUFFIX"),
	)

	expectedResult := "Root :SUFFIX"

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_WithPrefixAndSuffix(t *testing.T) {
	nodes := []TreeNode{
		{Value: "Root"},
	}

	result := CreateTree(nodes,
		WithPrefix("START: "),
		WithSuffix(" :END"),
	)

	expectedResult := "START: Root :END"

	if result != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTree_NestedWithPrefixSuffixAndStyle(t *testing.T) {
	nodes := []TreeNode{
		{
			Value: "Root",
			Children: []TreeNode{
				{
					Value: "Child 1",
					Children: []TreeNode{
						{Value: "Grandchild 1.1"},
						{Value: "Grandchild 1.2"},
					},
				},
				{
					Value: "Child 2",
					Children: []TreeNode{
						{Value: "Grandchild 2.1"},
					},
				},
			},
		},
	}

	result := CreateTree(nodes,
		WithPrefix("[\n"),
		WithSuffix("\n]"),
	)

	expectedResult := strings.TrimSpace(`[
Root
├── Child 1
│   ├── Grandchild 1.1
│   └── Grandchild 1.2
└── Child 2
    └── Grandchild 2.1
]`)
	actualResult := strings.TrimSpace(result)

	if actualResult != expectedResult {
		t.Error("Expected tree output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + actualResult)
	}
}
