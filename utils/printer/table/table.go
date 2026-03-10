package table

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/lipgloss"
	lipgloss_table "github.com/charmbracelet/lipgloss/table"

	common "github.com/cubbit/composer-cli/utils/printer/common"
)

type Column[T any] struct {
	Title string
}

type RowMapper[T any] func(T) []string

type TableConfig[T any] struct {
	common.CommonConfig

	Columns    []Column[T]
	Mapper     RowMapper[T]
	ShowHeader bool
	Style      lipgloss.Style
}

type Option[T any] func(*TableConfig[T])

func WithColumns[T any](cols []Column[T]) Option[T] {
	return func(c *TableConfig[T]) {
		c.Columns = cols
	}
}

func WithRowMapper[T any](mapper RowMapper[T]) Option[T] {
	return func(c *TableConfig[T]) {
		c.Mapper = mapper
	}
}

func WithShowHeader[T any](show bool) Option[T] {
	return func(c *TableConfig[T]) {
		c.ShowHeader = show
	}
}

func WithStyle[T any](style lipgloss.Style) Option[T] {
	return func(c *TableConfig[T]) {
		c.Style = style
	}
}

func WithPrefix[T any](prefix string) Option[T] {
	return func(c *TableConfig[T]) {
		c.Prefix = prefix
	}
}

func WithSuffix[T any](suffix string) Option[T] {
	return func(c *TableConfig[T]) {
		c.Suffix = suffix
	}
}

func CreateTable[T any](data []T, opts ...Option[T]) string {
	cfg := &TableConfig[T]{
		Columns:    nil,
		Mapper:     nil,
		ShowHeader: true,
		Style:      lipgloss.NewStyle(),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.Mapper == nil {
		cfg.Mapper = defaultRowMapper[T]()
	}

	if len(cfg.Columns) == 0 {
		cfg.Columns = inferColumns[T]()
	}

	var headers []string
	for _, col := range cfg.Columns {
		trimmed := strings.TrimSpace(col.Title)
		padded := " " + trimmed + " "
		headers = append(headers, padded)
	}

	var rows [][]string
	for _, item := range data {
		rowData := cfg.Mapper(item)
		formattedRow := make([]string, len(rowData))
		for i, cell := range rowData {
			trimmed := strings.TrimSpace(cell)
			padded := " " + trimmed + " "
			formattedRow[i] = padded
		}
		rows = append(rows, formattedRow)
	}

	t := lipgloss_table.New().
		Headers(headers...).
		Rows(rows...)

	if !cfg.ShowHeader {
		t.Headers()
	}

	tableStyle := cfg.Style.Render(t.Render())

	if cfg.Prefix != "" {
		tableStyle = cfg.Prefix + tableStyle
	}

	if cfg.Suffix != "" {
		tableStyle = tableStyle + cfg.Suffix
	}

	return tableStyle
}

func inferColumns[T any]() []Column[T] {
	var zero T
	t := reflect.TypeOf(zero)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil
	}

	columns := make([]Column[T], 0, t.NumField())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.IsExported() {
			columns = append(columns, Column[T]{
				Title: field.Name,
			})
		}
	}

	return columns
}

func defaultRowMapper[T any]() RowMapper[T] {
	return func(item T) []string {
		v := reflect.ValueOf(item)
		t := reflect.TypeOf(item)

		if t.Kind() == reflect.Ptr {
			v = v.Elem()
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			return []string{fmt.Sprintf("%v", item)}
		}

		values := make([]string, 0, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.IsExported() {
				fieldValue := v.Field(i)
				values = append(values, fmt.Sprintf("%v", fieldValue.Interface()))
			}
		}

		return values
	}
}
