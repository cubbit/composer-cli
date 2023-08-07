package utils

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/urfave/cli/v2"
)

func ValidateDescriptionSize(ctx *cli.Context) error {
	description := ctx.Args().First()
	if len(description) >= 200 {
		return fmt.Errorf(constants.ErrorDescriptionSize)
	}
	return nil
}

func ValidateIDorNameNotEmpty(ctx *cli.Context) error {

	id := ctx.String("id")
	name := ctx.String("name")

	if name == "" && id == "" {
		return fmt.Errorf(constants.ErrorIdAndNameEmpty)
	}
	return nil
}
