package utils

import (
	"fmt"
	"os"
)

func PrintSuccess(s string) {
	fmt.Printf("✨🐝 %s", s)
}

func PrintError(err error) {
	redBold := "\033[1;31m"
	reset := "\033[0m"

	formattedError := fmt.Sprintf("%sERR:%s %s%s", redBold, reset, err, reset)
	fmt.Fprintf(os.Stderr, formattedError+"\n")
}
