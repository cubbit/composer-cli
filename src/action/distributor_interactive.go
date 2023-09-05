package action

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateDistributorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, email, firstName, lastName, configPath string
	var swarmIDs []string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if _, err = tui.TextInputs("Fill in the form below", false, tui.Input{Placeholder: "Name", IsPassword: false, Value: &name}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}, tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs("Fill in the form to invite the associate the operator", false, tui.Input{Placeholder: "First Name", IsPassword: false, Value: &firstName}, tui.Input{Placeholder: "Last Name", IsPassword: false, Value: &lastName}, tui.Input{Placeholder: "email", IsPassword: false, Value: &email}); err != nil {
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

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	if len(swarms) == 0 {
		utils.PrintNotFound("No swarms found")
		return nil
	}

	var choices []string

	for _, sw := range swarms {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", sw.SwarmID, sw.Name, sw.Description))
	}

	if choices, err = tui.ChooseMany("Which swarm would you like to associate to the distributer?", true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributor, err)
	}

	for _, swarm := range choices {
		_, withoutPrefix, _ := strings.Cut(swarm, " ")
		delId, _, _ := strings.Cut(withoutPrefix, ",")
		swarmIDs = append(swarmIDs, delId)
	}

	if response, err = api.CreateDistributor(conf.Urls, *accessToken, name, &description, &imageUrl, swarmIDs, email, firstName, lastName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributor, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor  %s created successfully", response.ID))

	return nil
}

func RemoveDistributorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, email, password, code, configPath, choice, deleteDistributorToken string
	var conf *configuration.Config
	var distributors *api.DistributorList
	var challenge *api.ChallengeResponseModel

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
	}

	if len(distributors.Distributors) == 0 {
		utils.PrintNotFound("No distributors found")
		return nil
	}

	var choices []string

	for _, distributor := range distributors.Distributors {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, *distributor.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributors found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor would you like to delete?", false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributor, err)
	}
	_, withoutPrefix, _ := strings.Cut(choice, " ")
	id, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Confirm your login to delete the distributor", true, tui.Input{Placeholder: "Email", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if deleteDistributorToken, err = api.ForgeDistributorDeleteToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteToken, err)
	}

	if err = api.RemoveDistributor(conf.Urls, *accessToken, id, deleteDistributorToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributor, err)
	}

	utils.PrintDelete(fmt.Sprintf("distributor %s removed successfully", id))

	return nil
}

func ListDistributorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var configPath string
	var conf *configuration.Config
	var distributors *api.DistributorList

	if conf, configPath, err = configuration.ReadConfig(cmd, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
	}

	utils.PrintList("Your Distributors List")
	for _, distributor := range distributors.Distributors {
		fmt.Printf("• %s, %s, %s\n", distributor.ID, distributor.Name, *distributor.Description)
	}

	return nil
}
