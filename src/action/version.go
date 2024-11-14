package action

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type commitizen struct {
	Version string `json:"version"`
}

type cz struct {
	Commitizen commitizen `json:"commitizen"`
}

func GetCliVersion(cmd *cobra.Command, args []string) error {
	var err error

	versionFile, err := os.ReadFile("VERSION")
	if err != nil {
		return err
	}

	fmt.Println(string(versionFile))

	return nil
}
