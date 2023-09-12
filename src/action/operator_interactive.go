package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateOperatorInteractive(cmd *cobra.Command) error {
	var urls *configuration.Url
	var apiServerUrl, firstName, lastName, email, password, secret string
	var err error

	if _, err = tui.TextInputs("Enter your API server URL", false, tui.Input{Placeholder: "API server url: (default https://api.cubbit.eu)", IsPassword: false, Value: &apiServerUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if urls, err = configuration.ConfigureAPIServerURL(apiServerUrl); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConfiguringAPIURL, err)
	}

	if _, err = tui.TextInputs("Fill in the form bellow", true, tui.Input{Placeholder: "First Name*", IsPassword: false, Value: &firstName}, tui.Input{Placeholder: "Last Name*", IsPassword: false, Value: &lastName}, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Secret", IsPassword: true, Value: &secret}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if err = api.CreateOperator(*urls, firstName, lastName, email, password, secret); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingOperatorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("operator %s created successfully", email))

	return nil
}
