package action

import (
	"fmt"
	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func SetupSwarmRing(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, redundancyClassID, configPath string
	var conf *configuration.Config
	var redundancyClass *api.RedundancyClass
	var nexusIDs []string
	var ringsNumber int
	var nodes *api.NodeList
	var ringList *api.RingList

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("redundancy-class-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexusIDs, err = cmd.Flags().GetStringSlice("nexus-ids"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if ringsNumber, err = cmd.Flags().GetInt("rings-number"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if redundancyClass, err = api.GetRedundancyClass(conf.Urls, *accessToken, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	if len(nexusIDs) < (redundancyClass.OuterK + redundancyClass.OuterN) {
		return fmt.Errorf("%s: %s", constants.ErrorNotEnoughNexuses, " create more nexuses")
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	for _, nexus := range nexusIDs {
		if nodes, err = api.ListNodes(conf.Urls, *accessToken, id, nexus); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingNodesRequest, err)
		}

		if len(nodes.Nodes) < (redundancyClass.InnerN + redundancyClass.InnerK) {
			return fmt.Errorf("%s: %s", constants.ErrorNotEnoughNodes, " create more nodes for the nexus "+id)
		}
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

	changed := cmd.Flags().Changed("rings-number")

	if changed {
		if ringsNumber > ringList.Count {
			return fmt.Errorf("%s: %s", constants.ErrorCreatingRingRequest, fmt.Sprintf("you can create only %d rings", ringList.Count))
		} else {
			ringBulk.RingsNumber = &ringsNumber
		}
	} else {
		ringBulk.RingsNumber = &ringList.Count

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

func ListSwarmRings(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, redundancyClassID, configPath string
	var conf *configuration.Config
	var ringList *api.RingList
	var verbose, l bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassID, err = cmd.Flags().GetString("redundancy-class-id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		if id, err = getSwarmByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
		}
	}

	if ringList, err = api.ListRings(conf.Urls, *accessToken, id, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRingsRequest, err)
	}

	utils.PrintList("Your Rings List")

	if len(ringList.Data) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if verbose {
		utils.PrintVerbose(ringList.Data, l)
		return nil

	}

	for _, ring := range ringList.Data {
		fmt.Printf(" • %s\n", ring.ID)

		if l {
			fmt.Println("")
		}
	}

	return nil
}
