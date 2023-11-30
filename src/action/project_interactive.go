package action

import (
	"fmt"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateProjectInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, configPath string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeAccount, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &name},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &description},
		tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if response, err = api.CreateProject(conf.Urls, *accessToken, name, &description, &imageUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingProjectRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("project  %s created successfully", response.ID))

	return nil
}
