package table

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

type TestItem struct {
	Name  string
	Age   int
	Email string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type SimpleItem struct {
	Name string
}

func TestCreateTable_WithValidData(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
		{"Charlie", 35, "charlie@example.com"},
	}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
			{"Age"},
			{"Email"},
		}),
		WithShowHeader[TestItem](true),
	)

	expectedResult := strings.TrimSpace(`
╭─────────┬─────┬─────────────────────╮
│ Name    │ Age │ Email               │
├─────────┼─────┼─────────────────────┤
│ Alice   │ 30  │ alice@example.com   │
│ Bob     │ 25  │ bob@example.com     │
│ Charlie │ 35  │ charlie@example.com │
╰─────────┴─────┴─────────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_EmptyData(t *testing.T) {
	items := []TestItem{}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
		}),
	)

	expectedResult := strings.TrimSpace(`
╭──────╮
│ Name │
├──────┤
╰──────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_MissingMapper(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
	}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
		}),
		WithShowHeader[TestItem](true),
	)

	expectedResult := strings.TrimSpace(`
╭───────┬────┬───────────────────╮
│ Name  │    │                   │
├───────┼────┼───────────────────┤
│ Alice │ 30 │ alice@example.com │
╰───────┴────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_MissingColumns(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
	}

	result := CreateTable[TestItem](items)

	expectedResult := strings.TrimSpace(`
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ Alice │ 30  │ alice@example.com │
╰───────┴─────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_EmptyColumns(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
	}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{}),
	)

	expectedResult := strings.TrimSpace(`
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ Alice │ 30  │ alice@example.com │
╰───────┴─────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_NoHeader(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
	}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
			{"Age"},
		}),
		WithShowHeader[TestItem](false),
	)

	expectedResult := strings.TrimSpace(`
╭───────┬────┬───────────────────╮
│ Alice │ 30 │ alice@example.com │
│ Bob   │ 25 │ bob@example.com   │
╰───────┴────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_WithCustomStyle(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
	}

	customStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Padding(1)

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
		}),
		WithStyle[TestItem](customStyle),
	)

	exprectedResult := strings.TrimSpace(`
╔════════════════════════════════════╗
║                                    ║
║ ╭───────┬────┬───────────────────╮ ║
║ │ Name  │    │                   │ ║
║ ├───────┼────┼───────────────────┤ ║
║ │ Alice │ 30 │ alice@example.com │ ║
║ ╰───────┴────┴───────────────────╯ ║
║                                    ║
╚════════════════════════════════════╝
`)

	if result != exprectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + exprectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_SingleItem(t *testing.T) {
	items := []TestItem{{"Alice", 30, "alice@example.com"}}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
			{"Age"},
			{"Email"},
		}),
	)

	expectedResult := strings.TrimSpace(`
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ Alice │ 30  │ alice@example.com │
╰───────┴─────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_MultipleItems(t *testing.T) {
	items := make([]TestItem, 10)
	for i := 0; i < 10; i++ {
		items[i] = TestItem{
			Name:  "User" + string(rune('0'+i)),
			Age:   20 + i,
			Email: "user" + string(rune('0'+i)) + "@example.com",
		}
	}

	result := CreateTable[TestItem](items,
		WithColumns[TestItem]([]Column[TestItem]{
			{"Name"},
			{"Age"},
			{"Email"},
		}),
	)

	expectedResult := strings.TrimSpace(`
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ User0 │ 20  │ user0@example.com │
│ User1 │ 21  │ user1@example.com │
│ User2 │ 22  │ user2@example.com │
│ User3 │ 23  │ user3@example.com │
│ User4 │ 24  │ user4@example.com │
│ User5 │ 25  │ user5@example.com │
│ User6 │ 26  │ user6@example.com │
│ User7 │ 27  │ user7@example.com │
│ User8 │ 28  │ user8@example.com │
│ User9 │ 29  │ user9@example.com │
╰───────┴─────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_WithNumericData(t *testing.T) {
	products := []Product{
		{1, "Laptop", 999.99},
		{2, "Phone", 599.99},
	}

	result := CreateTable[Product](products,
		WithColumns[Product]([]Column[Product]{
			{"ID"},
			{"Name"},
			{"Price"},
		}),
	)

	expectedResult := strings.TrimSpace(`
╭────┬────────┬────────╮
│ ID │ Name   │ Price  │
├────┼────────┼────────┤
│ 1  │ Laptop │ 999.99 │
│ 2  │ Phone  │ 599.99 │
╰────┴────────┴────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_HeaderOnly(t *testing.T) {
	items := []SimpleItem{{"Test"}}

	result := CreateTable[SimpleItem](items,
		WithColumns[SimpleItem]([]Column[SimpleItem]{
			{"Name"},
		}),
	)

	expectedResult := strings.TrimSpace(`
╭──────╮
│ Name │
├──────┤
│ Test │
╰──────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_WithReflectionInferredColumns(t *testing.T) {
	items := []TestItem{
		{"Alice", 30, "alice@example.com"},
		{"Bob", 25, "bob@example.com"},
	}

	result := CreateTable[TestItem](items)

	expectedResult := strings.TrimSpace(`
╭───────┬─────┬───────────────────╮
│ Name  │ Age │ Email             │
├───────┼─────┼───────────────────┤
│ Alice │ 30  │ alice@example.com │
│ Bob   │ 25  │ bob@example.com   │
╰───────┴─────┴───────────────────╯
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}

func TestCreateTable_WithReflectionAndCustomStyle(t *testing.T) {
	items := []SimpleItem{
		{"Item1"},
		{"Item2"},
	}

	customStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		Padding(1)

	result := CreateTable[SimpleItem](items,
		WithStyle[SimpleItem](customStyle),
	)

	expectedResult := strings.TrimSpace(`
╔═══════════╗
║           ║
║ ╭───────╮ ║
║ │ Name  │ ║
║ ├───────┤ ║
║ │ Item1 │ ║
║ │ Item2 │ ║
║ ╰───────╯ ║
║           ║
╚═══════════╝
`)

	if result != expectedResult {
		t.Error("Expected table output does not match actual output expected:\n" + expectedResult + "\nactual:\n" + result)
	}
}
