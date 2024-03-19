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

func SetupSwarmRingInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath string
	var conf *configuration.Config
	var RedundancyClassList *api.RedundancyClassList
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var nexuses *api.NexusList
	var redundancyClass api.RedundancyClass
	var nexusIDs []string
	var ringList *api.RingList
	var ringsNumber int

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
		var swarms []api.Swarm

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

	if choice, err = tui.ChooseOne("Which redundancy class would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	for _, rc := range RedundancyClassList.Data {
		if rc.ID == redundancyClassID {
			redundancyClass = rc
			break
		}
	}

	if nexuses, err = api.ListNexuses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingNexusesRequest, err)
	}

	if len(nexuses.Nexuses) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	choices = []string{}

	for _, nx := range nexuses.Nexuses {
		var nodes *api.NodeList

		if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nx.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
		}

		if len(nodes.Nodes) < (redundancyClass.InnerN + redundancyClass.InnerK) {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s %s", nx.ID, nx.Name, nx.Description, "[Not enough nodes]"))
		} else {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", nx.ID, nx.Name, nx.Description))
		}
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No nexuses found")
		return nil
	}

	if nexusIDs, err = tui.ChooseMany(fmt.Sprintf("Which nexuses would you like to choose? (select at least %d nexus)", redundancyClass.OuterK+redundancyClass.OuterN), false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingTenant, err)
	}

	if len(nexusIDs) < (redundancyClass.OuterK + redundancyClass.OuterN) {
		return fmt.Errorf("%s: %s", constants.ErrorNotEnoughNexuses, "create or select more!")
	}

	for i, nexus := range nexusIDs {
		_, withoutPrefix, _ := strings.Cut(nexus, " ")
		nexusID, _, _ := strings.Cut(withoutPrefix, ",")
		nexusIDs[i] = nexusID
	}

	ringBulk := api.RingBulk{
		RedundancyClassID: redundancyClassID,
		Nexuses:           nexusIDs,
	}

	if ringList, err = api.CreateRing(conf.Urls, *accessToken, id, redundancyClassID, ringBulk, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRingRequest, err)
	}

	if ringList.Count == 0 {
		return fmt.Errorf("%s: %s", constants.ErrorCreatingRingRequest, fmt.Sprintf("you can create only %d rings", ringList.Count))
	}

	var ringsNumberString string

	if _, err = tui.TextInputs(fmt.Sprintf("Fill in the number of rings (we suggest %d)", ringList.Count), true, tui.Input{Placeholder: "Rings Number", IsPassword: false, Value: &ringsNumberString}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)

	}

	if ringsNumber, err = strconv.Atoi(ringsNumberString); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorConvertingField, err)
	}

	if ringsNumber > ringList.Count {
		return fmt.Errorf("%s: %s", constants.ErrorCreatingRingRequest, fmt.Sprintf("you can create only %d rings", ringList.Count))
	}

	if ringList, err = api.CreateRing(conf.Urls, *accessToken, id, redundancyClassID, ringBulk, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRingRequest, err)
	}

	utils.PrintSuccess("rings created successfully")

	for _, ring := range ringList.Data {
		fmt.Printf("• ring %s created \n", ring.ID)
	}

	return nil
}

func ListSwarmRingInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var name, id, configPath string
	var conf *configuration.Config
	var RedundancyClassList *api.RedundancyClassList
	var operator *api.Operator
	var choices []string
	var choice string
	var redundancyClassID string
	var ringList *api.RingList

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
		var swarms []api.Swarm

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

	if choice, err = tui.ChooseOne("Which redundancy class would you like to choose?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	redundancyClassID, _, _ = strings.Cut(withoutPrefix, ",")

	if ringList, err = api.ListRings(conf.Urls, *accessToken, id, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRingsRequest, err)
	}

	utils.PrintList("Your Rings List")

	if len(ringList.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string

	for _, ring := range ringList.Data {
		list = append(list, fmt.Sprintf("• %s, %s, %s", ring.ID, ring.SwarmID, ring.Status))
	}

	tui.List(list)

	return nil
}
