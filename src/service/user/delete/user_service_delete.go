package delete

import (
	"fmt"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/service/user/shared"
	"github.com/spf13/cobra"
)

type Dependencies struct {
	UserAPI api.UserAPIInterface
}

func DeleteUser(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	args []string,
) error {
	userID, err := shared.ResolveUserID(deps.UserAPI, cmd, args, profile)
	if err != nil {
		return err
	}

	if err := deps.UserAPI.DeleteIAMUser(profile.Endpoints, profile.APIKey, profile.OrganizationID, userID, ""); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingIAMUserRequest, err)
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if !quiet {
		cmd.Printf("IAM user %s deleted successfully\n", userID)
	}
	return nil
}
