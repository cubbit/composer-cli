package action

import (
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateSwarmRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, redundancyClassName, redundancyClassDescription, configPath string
	var innerK, innerN, outerK, outerN, antiAffinityGroup int
	var conf *configuration.Config
	var redundancyClass *api.RedundancyClass
	var nexuses []string

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassName, err = cmd.Flags().GetString("redundancy-class-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if redundancyClassDescription, err = cmd.Flags().GetString("redundancy-class-description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if innerK, err = cmd.Flags().GetInt("inner-k"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if innerN, err = cmd.Flags().GetInt("inner-n"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outerK, err = cmd.Flags().GetInt("outer-k"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if outerN, err = cmd.Flags().GetInt("outer-n"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if antiAffinityGroup, err = cmd.Flags().GetInt("anti-affinity-group"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if nexuses, err = cmd.Flags().GetStringSlice("nexuses"); err != nil {
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

	bodyRequest := api.CreateRedundancyClassRequestBody{
		Name:              redundancyClassName,
		Description:       redundancyClassDescription,
		InnerK:            innerK,
		InnerN:            innerN,
		OuterK:            outerK,
		OuterN:            outerN,
		AntiAffinityGroup: antiAffinityGroup,
		Nexuses:           nexuses,
	}

	if redundancyClass, err = api.CreateRedundancyClassV4(conf.Urls, *accessToken, id, bodyRequest); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingRedundancyClassRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("Redundancy class %s created successfully\n", redundancyClass.ID))

	return nil
}

func ListSwarmRedundancyClasses(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, id, configPath string
	var conf *configuration.Config
	var redundancyClasses *api.RedundancyClassList
	var verbose, l bool

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
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

	if redundancyClasses, err = api.ListRedundancyClasses(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingRedundancyClassesRequest, err)
	}

	if len(redundancyClasses.Data) == 0 {
		utils.PrintSuccess("No redundancy classes found")
		return nil
	}

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if verbose {
		utils.PrintVerbose(redundancyClasses.Data, l)
		return nil
	}

	for _, rc := range redundancyClasses.Data {
		fmt.Printf(" • %s\n", rc.Name)

		if l {
			fmt.Println()
		}
	}

	return nil
}

func DescribeSwarmRedundancyClass(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var redundancyClassID, format, configPath string
	var conf *configuration.Config
	var redundancyClass *api.RedundancyClass

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	redundancyClassID = args[0]

	if redundancyClass, err = api.GetRedundancyClass(conf.Urls, *accessToken, redundancyClassID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingRedundancyClassRequest, err)
	}

	utils.PrintFormattedData(*redundancyClass, format)

	return nil
}
