package utils

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

var byteUnits = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}

func FormatBytes(bytes int64) string {
	if bytes == 0 {
		return "0"
	}

	if bytes < KB {
		return fmt.Sprintf("%d B", bytes)
	}

	div := int64(KB)
	for i := 1; i < len(byteUnits)-1; i++ {
		if bytes < div*KB {
			return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), byteUnits[i])
		}
		div *= KB
	}

	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), byteUnits[len(byteUnits)-1])
}
