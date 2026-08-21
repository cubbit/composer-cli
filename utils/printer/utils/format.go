package utils

import "strings"

func FormatPtrString(s *string) string {
	if s == nil {
		return "none"
	}
	return *s
}

func FormatBool(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}

func FormatPtrBool(b *bool) string {
	if b == nil || !*b {
		return "no"
	}
	return "yes"
}

func FormatDisabledStatus(b *bool) string {
	if b != nil && *b {
		return "disabled"
	}
	return "enabled"
}

func FormatSecret(s *string, showSecrets bool) string {
	if s == nil {
		return "none"
	}
	if showSecrets {
		return *s
	}
	return "************"
}

func FormatStringSlicePtr(s *[]string) string {
	if s == nil || len(*s) == 0 {
		return "none"
	}
	return strings.Join(*s, ", ")
}
