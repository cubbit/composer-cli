package text

import (
	"github.com/charmbracelet/lipgloss"

	common "github.com/cubbit/composer-cli/utils/printer/common"
)

type TextConfig struct {
	common.CommonConfig
	Style lipgloss.Style
}

type Option func(*TextConfig)

func WithStyle(style lipgloss.Style) Option {
	return func(c *TextConfig) {
		c.Style = style
	}
}

func CreateText(s string, opts ...Option) string {
	cfg := &TextConfig{
		Style: lipgloss.NewStyle(),
	}

	for _, opt := range opts {
		opt(cfg)
	}

	result := cfg.Style.Render(s)

	if cfg.Prefix != "" {
		result = cfg.Prefix + result
	}

	if cfg.Suffix != "" {
		result = result + cfg.Suffix
	}

	return result
}
