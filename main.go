package main

import (
	_ "embed"

	"github.com/cubbit/cubbit/client/cli/src/cmd"
)

//go:embed package.json
var packageJSON []byte

func main() {
	cmd.Execute(packageJSON)
}
