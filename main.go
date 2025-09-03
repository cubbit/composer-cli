package main

import (
	"log"

	_ "embed"

	"github.com/cubbit/cubbit/client/cli/src/cmd"
)

//go:embed package.json
var packageJSON []byte

func main() {
	if len(packageJSON) == 0 {
		log.Fatalf("Fatal error: failed to embed package.json")
	}

	cmd.Execute(packageJSON)
}
