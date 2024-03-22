package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/tui"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, redundancyClassName, redundancyClassDescription, configPath string
	var innerK, innerN, outerK, outerN, antiAffinityGroup string
	var innerKInt, innerNInt, outerKInt, outerNInt, antiAffinityGroupInt int
	var conf *configuration.Config
	var redundancyClass *api.RedundancyClass
	var operator *api.Operator
	var choice string
	var choices []string
	var swarms []*api.Swarm

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
		}

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if choice, err = tui.ChooseOne("Which swarm would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		true,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &redundancyClassName},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &redundancyClassDescription},
		tui.Input{Placeholder: "Inner K*", IsPassword: false, Value: &innerK},
		tui.Input{Placeholder: "Inner N*", IsPassword: false, Value: &innerN},
		tui.Input{Placeholder: "Outer K*", IsPassword: false, Value: &outerK},
		tui.Input{Placeholder: "Outer N*", IsPassword: false, Value: &outerN},
		tui.Input{Placeholder: "Anti Affinity Group*", IsPassword: false, Value: &antiAffinityGroup}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(redundancyClassDescription) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if innerKInt, err = strconv.Atoi(innerK); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if innerNInt, err = strconv.Atoi(innerN); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if outerKInt, err = strconv.Atoi(outerK); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if outerNInt, err = strconv.Atoi(outerN); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if antiAffinityGroupInt, err = strconv.Atoi(antiAffinityGroup); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	bodyRequest := api.CreateRedundancyClassRequestBody{
		Name:              redundancyClassName,
		Description:       redundancyClassDescription,
		InnerK:            innerKInt,
		InnerN:            innerNInt,
		OuterK:            outerKInt,
		OuterN:            outerNInt,
		AntiAffinityGroup: antiAffinityGroupInt,
	}

	if redundancyClass, err = api.CreateRedundancyClass(conf.Urls, *accessToken, id, bodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRedundancyClassRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("redundancy class %s created successfully", redundancyClass.ID))

	return nil
}

func ListSwarmRedundancyClassesInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath string
	var conf *configuration.Config
	var RedundancyClassList *api.RedundancyClassList
	var operator *api.Operator
	var choice string
	var choices []string
	var swarms []*api.Swarm

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
		}

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if choice, err = tui.ChooseOne("Which swarm would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	utils.PrintList("Your Redundancy Classes")

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string

	for _, rc := range RedundancyClassList.Data {
		list = append(list, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	tui.List(list)

	return nil
}

func DescribeSwarmRedundancyClassInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath, format string
	var conf *configuration.Config
	var RedundancyClassList *api.RedundancyClassList
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var swarms []*api.Swarm

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("error while generating access and refresh tokens: %w", err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
		}

		if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
		}

		for _, swarm := range swarms {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", swarm.ID, swarm.Name, swarm.Description))
		}

		if choice, err = tui.ChooseOne("Which swarm would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDeletingTenant, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if RedundancyClassList, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(RedundancyClassList.Data) == 0 {
		utils.PrintNotFound("redundancy classes not found")
		return nil
	}

	choices = []string{}
	for _, rc := range RedundancyClassList.Data {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s, %s", rc.ID, rc.Name, rc.Description, rc.SwarmID))
	}

	if choice, err = tui.ChooseOne("Which redundancy class would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, rc := range RedundancyClassList.Data {
		if rc.ID == redundancyClassID {
			utils.PrintFormattedData(rc, format)
			return nil
		}
	}

	return nil
}
