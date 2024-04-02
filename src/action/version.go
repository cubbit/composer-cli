package action

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type commitizen struct {
	Version string `json:"version"`
}

type cz struct {
	Commitizen commitizen `json:"commitizen"`
}

func GetCliVersion(cmd *cobra.Command, args []string) error {
	var err error

	yamlFile, err := os.ReadFile(".cz.yaml")
	if err != nil {
		return err
	}

	var cz cz
	err = yaml.Unmarshal(yamlFile, &cz)
	if err != nil {
		return err
	}

	fmt.Println(cz.Commitizen.Version)

	return nil
}
