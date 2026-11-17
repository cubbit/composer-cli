package cmd_gateway

import (
	"fmt"
	"unicode/utf8"

	servicegateway "github.com/cubbit/composer-cli/src/service/gateway"
	"github.com/cubbit/composer-cli/utils"
	"github.com/spf13/cobra"
)

const (
	gatewayNameMinLength        = 3
	gatewayNameMaxLength        = 63
	gatewayDescriptionMaxLength = 120
	gatewayLocationMaxLength    = 240
)

func NewGatewaySubCmdUpdate(
	gatewayService servicegateway.GatewayServiceInterface,
) *cobra.Command {
	var gatewayUpdateCmd = &cobra.Command{
		Use:   "update GATEWAY_ID",
		Short: "Update a gateway using the Gateway V5 API",
		Long: `Update the name, description or location of a gateway using the Gateway V5 API.

The gateway is identified by GATEWAY_ID.
Only the provided flags are updated; the other fields are left unchanged.
At least one of --name, --description or --location is required.
Pass an empty value to --description or --location to clear it.

Examples:
  # Update the name of a gateway
  cubbit gateway update GATEWAY_ID --name "renamed-gateway"

  # Update description and location
  cubbit gateway update GATEWAY_ID \
    --description "Production gateway" \
    --location "eu-west-1"

  # Clear the description and location of a gateway
  cubbit gateway update GATEWAY_ID --description "" --location ""`,
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateUpdateFlags(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return gatewayService.Update(cmd, args)
		},
	}

	gatewayUpdateCmd.Flags().String("name", "", "New name of the gateway (3-63 characters)")
	gatewayUpdateCmd.Flags().String("description", "", "New description of the gateway (max 120 characters); pass an empty value to clear it")
	gatewayUpdateCmd.Flags().String("location", "", "New location of the gateway (max 240 characters); pass an empty value to clear it")

	return gatewayUpdateCmd
}

func validateUpdateFlags(cmd *cobra.Command) error {
	name, err := utils.GetOptionalStringFlag(cmd, "name")
	if err != nil {
		return err
	}

	description, err := utils.GetOptionalStringFlag(cmd, "description")
	if err != nil {
		return err
	}

	location, err := utils.GetOptionalStringFlag(cmd, "location")
	if err != nil {
		return err
	}

	if name == nil && description == nil && location == nil {
		return fmt.Errorf("specify at least one of --name, --description or --location")
	}

	if name != nil && (utf8.RuneCountInString(*name) < gatewayNameMinLength || utf8.RuneCountInString(*name) > gatewayNameMaxLength) {
		return fmt.Errorf("--name must be between %d and %d characters", gatewayNameMinLength, gatewayNameMaxLength)
	}

	if description != nil && utf8.RuneCountInString(*description) > gatewayDescriptionMaxLength {
		return fmt.Errorf("--description must be at most %d characters", gatewayDescriptionMaxLength)
	}

	if location != nil && utf8.RuneCountInString(*location) > gatewayLocationMaxLength {
		return fmt.Errorf("--location must be at most %d characters", gatewayLocationMaxLength)
	}

	return nil
}
