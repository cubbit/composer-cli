package utils

import "time"

func StringOrEmpty(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

func TimeOrEmpty(v *time.Time) string {
	if v == nil {
		return ""
	}

	return v.String()
}
