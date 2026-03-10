package tree

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	tree "github.com/charmbracelet/lipgloss/tree"

	common "github.com/cubbit/composer-cli/utils/printer/common"
)

type TreeNode struct {
	Value     any
	Children  []TreeNode
	Formatter TreeFormatter
}

type TreeFormatter func(any) string

type TreeConfig struct {
	common.CommonConfig

	Formatter TreeFormatter
	Style     lipgloss.Style
}

type Option func(*TreeConfig)

func WithFormatter(formatter TreeFormatter) Option {
	return func(c *TreeConfig) {
		c.Formatter = formatter
	}
}

func WithStyle(style lipgloss.Style) Option {
	return func(c *TreeConfig) {
		c.Style = style
	}
}

func WithPrefix(prefix string) Option {
	return func(c *TreeConfig) {
		c.Prefix = prefix
	}
}

func WithSuffix(suffix string) Option {
	return func(c *TreeConfig) {
		c.Suffix = suffix
	}
}

func CreateTree(nodes []TreeNode, opts ...Option) string {
	cfg := &TreeConfig{
		Formatter: defaultFormatter(),
		Style:     lipgloss.NewStyle(),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	if len(nodes) == 0 {
		return "No data"
	}

	trees := make([]*tree.Tree, len(nodes))
	for i, node := range nodes {
		trees[i] = buildTreeFromNode(node, cfg)
	}

	var combined strings.Builder
	for i, t := range trees {
		if i > 0 {
			combined.WriteString("\n")
		}
		combined.WriteString(t.String())
	}

	treeStyle := cfg.Style.Render(combined.String())

	if cfg.Prefix != "" {
		treeStyle = cfg.Prefix + treeStyle
	}

	if cfg.Suffix != "" {
		treeStyle = treeStyle + cfg.Suffix
	}

	return treeStyle
}

func defaultFormatter() TreeFormatter {
	return func(v any) string {
		switch val := v.(type) {
		case fmt.Stringer:
			return val.String()
		case string:
			return val
		default:
			return fmt.Sprintf("%v", val)
		}
	}
}

func buildTreeFromNode(node TreeNode, cfg *TreeConfig) *tree.Tree {
	var title string
	if node.Formatter != nil {
		title = node.Formatter(node.Value)
	} else if cfg.Formatter != nil {
		title = cfg.Formatter(node.Value)
	} else {
		title = defaultFormatter()(node.Value)
	}

	root := tree.Root(title)

	for _, child := range node.Children {
		root.Child(buildTreeFromNode(child, cfg))
	}

	return root
}
