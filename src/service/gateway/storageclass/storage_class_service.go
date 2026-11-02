package storageclass

import (
	"fmt"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	api "github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	GatewayAPI api.GatewayAPIInterface
}

// gatewayIDFromFlags reads and validates the required --gateway-id flag.
func gatewayIDFromFlags(cmd *cobra.Command) (string, error) {
	gatewayID, err := cmd.Flags().GetString("gateway-id")
	if err != nil {
		return "", fmt.Errorf("%s gateway-id: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(gatewayID) == "" {
		return "", fmt.Errorf("--gateway-id must not be empty")
	}

	return strings.TrimSpace(gatewayID), nil
}

// List prints the default and custom storage classes configured on a gateway.
func List(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	response, err := deps.GatewayAPI.ListGatewayStorageClassesV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
	)
	if err != nil {
		return fmt.Errorf("failed to list storage classes: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, response.Data,
		func() error { return PrintStorageClassList(cmd, handler, response.Data) },
	)
}

// Create adds a custom storage class to a gateway.
func Create(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return fmt.Errorf("%s name: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(name) == "" {
		return fmt.Errorf("--name must not be empty")
	}

	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	storageClass, err := deps.GatewayAPI.CreateGatewayStorageClassV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		&api.CreateGatewayStorageClassV5Request{Name: name},
	)
	if err != nil {
		return fmt.Errorf("failed to create storage class: %w", err)
	}

	return printer.ComposeStructured(cmd, handler, storageClass,
		func() error { return PrintStorageClassCreated(cmd, handler, *storageClass) },
	)
}

// Delete removes a custom storage class from a gateway.
func Delete(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	storageClassID, err := cmd.Flags().GetString("id")
	if err != nil {
		return fmt.Errorf("%s id: %w", constants.ErrorRetrievingField, err)
	}

	if strings.TrimSpace(storageClassID) == "" {
		return fmt.Errorf("--id must not be empty")
	}

	gatewayID, err := gatewayIDFromFlags(cmd)
	if err != nil {
		return err
	}

	if err := deps.GatewayAPI.DeleteGatewayStorageClassV5(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		gatewayID,
		storageClassID,
	); err != nil {
		return fmt.Errorf("failed to delete storage class: %w", err)
	}

	return printer.PrintText(cmd, handler, buildStorageClassDeletedOutput(storageClassID))
}
