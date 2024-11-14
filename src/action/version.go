package action

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type data struct {
	Version string `json:"version"`
}

func GetCliVersion(cmd *cobra.Command, args []string) error {
	var err error

	jsonFile, err := os.ReadFile("package.json")
	if err != nil {
		return err
	}

	var dv data
	err = json.Unmarshal(jsonFile, &dv)
	if err != nil {
		return err
	}

	fmt.Println(dv.Version)

	return nil
}
