package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmInteractive(cmd *cobra.Command) error {
	var config *configuration.Config
	var err error
	var name, description, swarmConfig1, configPath string
	var accessToken *string
	var operator *api.Operator
	var response *api.GenericIDResponseModel
	var swarmConfig map[string]interface{}

	if config, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Fill in the form below", false, tui.Input{Placeholder: "Name", IsPassword: false, Value: &name}, tui.Input{Placeholder: "Description", IsPassword: false, Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if _, err = tui.TextAreas("Fill in the swarm configuration", true, tui.TextArea{Placeholder: "{}", Value: &swarmConfig1}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if err = json.Unmarshal([]byte(swarmConfig1), &swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	if response, err = api.CreateSwarm(config.Urls, *accessToken, operator.ID, name, description, swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarm, err)
	}

	utils.PrintSuccess(fmt.Sprintf("swarm %s created successfully", response.ID))
	return nil
}

func DescribeSwarmInteractive(cmd *cobra.Command) error {

	var err error
	var nameOrId, format, configPath string
	var accessToken *string
	var conf *configuration.Config
	var operator *api.Operator

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if _, err = tui.TextInputs("Enter your swarm ID or Name", false, tui.Input{Placeholder: "Swarm ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if format, err = tui.ChooseOne("Choose your output format", true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	var swarms []api.Swarm
	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	for _, swarm := range swarms {
		if swarm.Name == nameOrId || swarm.SwarmID == nameOrId {
			utils.PrintFormattedData(swarm, format)
			return nil
		}
	}

	return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)

}

func ListSwarmsInteractive(cmd *cobra.Command) error {
	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm

	if config, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if swarms, err = api.ListSwarms(config.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	utils.PrintList("Your Swarms List")
	for _, swarm := range swarms {
		fmt.Printf("• %s, %s, %s\n", swarm.SwarmID, swarm.Name, swarm.Description)

	}

	return nil
}

func RemoveSwarmInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, email, password, code, configPath, deleteSwarmToken, choice string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel
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

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	if len(swarms) == 0 {
		utils.PrintNotFound("No swarms found")
		return nil
	}

	var choices []string

	for _, swarm := range swarms {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.SwarmID, swarm.Name, swarm.Description))
	}

	if choice, err = tui.ChooseOne("Which swarm would you like to delete?", false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingSwarm, err)
	}
	_, withoutPrefix, _ := strings.Cut(choice, " ")
	id, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs("Confirm your login to delete the swarm", true, tui.Input{Placeholder: "Email", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallenge, err)
	}

	if deleteSwarmToken, err = api.ForgeOperatorDeleteSwarmToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingOperatorDeleteToken, err)
	}

	if err = api.RemoveSwarm(conf.Urls, *accessToken, id, deleteSwarmToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingSwarm, err)
	}

	utils.PrintDelete(fmt.Sprintf("swarm %s removed successfully", id))

	return nil
}

func EditSwarmDescriptionInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, configPath, description string
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm
	var found bool

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if _, err = tui.TextInputs("Enter your swarm ID or Name", false, tui.Input{Placeholder: "Swarm ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	for _, swarm := range swarms {
		if swarm.SwarmID == nameOrId || swarm.Name == nameOrId {
			nameOrId = swarm.SwarmID
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf(constants.ErrorRetrievingSwarm)
	}

	if _, err = tui.TextInputs("Enter your new swarm description", true, tui.Input{Placeholder: "New Swarm Description", Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if err = api.EditSwarmDescription(conf.Urls, *accessToken, nameOrId, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmDescription, err)
	}

	utils.PrintSuccess(fmt.Sprintf("swarm %s description updated successfully", nameOrId))

	return nil
}

func EditSwarmNameInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var nameOrId, configPath, description string
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm
	var found bool

	if conf, configPath, err = configuration.ReadConfig(cmd, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if _, err = tui.TextInputs("Enter your swarm ID or Name", false, tui.Input{Placeholder: "Swarm ID or Name", Value: &nameOrId}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	for _, swarm := range swarms {
		if swarm.SwarmID == nameOrId || swarm.Name == nameOrId {
			nameOrId = swarm.SwarmID
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf(constants.ErrorRetrievingSwarm)
	}

	if _, err = tui.TextInputs("Enter your new swarm name", true, tui.Input{Placeholder: "New Swarm Name", Value: &description}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if err = api.EditSwarmName(conf.Urls, *accessToken, nameOrId, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarmName, err)
	}

	utils.PrintSuccess(fmt.Sprintf("swarm %s name updated successfully", nameOrId))
	
	return nil
}
