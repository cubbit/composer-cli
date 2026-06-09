package common

import (
	"github.com/charmbracelet/lipgloss"
)

type CommonConfig struct {
	Prefix string
	Suffix string
}

type CommonOption func(*CommonConfig)

func WithPrefix(prefix string) CommonOption {
	return func(c *CommonConfig) {
		c.Prefix = prefix
	}
}

func WithSuffix(suffix string) CommonOption {
	return func(c *CommonConfig) {
		c.Suffix = suffix
	}
}

type TextConfig struct {
	Prefix string
	Suffix string
	Style  lipgloss.Style
}

type TextOption func(*TextConfig)

func WithTextStyle(style lipgloss.Style) TextOption {
	return func(c *TextConfig) {
		c.Style = style
	}
}

type TableConfig struct {
	Prefix     string
	Suffix     string
	Columns    []interface{}
	Mapper     interface{}
	ShowHeader bool
	Style      lipgloss.Style
}

type TableOption func(*TableConfig)

func WithTableColumns(cols []interface{}) TableOption {
	return func(c *TableConfig) {
		c.Columns = cols
	}
}

func WithTableShowHeader(show bool) TableOption {
	return func(c *TableConfig) {
		c.ShowHeader = show
	}
}

func WithTableStyle(style lipgloss.Style) TableOption {
	return func(c *TableConfig) {
		c.Style = style
	}
}

type TreeConfig struct {
	Prefix    string
	Suffix    string
	Formatter interface{}
	Style     lipgloss.Style
}

type TreeOption func(*TreeConfig)

func WithTreeFormatter(formatter interface{}) TreeOption {
	return func(c *TreeConfig) {
		c.Formatter = formatter
	}
}

func WithTreeStyle(style lipgloss.Style) TreeOption {
	return func(c *TreeConfig) {
		c.Style = style
	}
}
