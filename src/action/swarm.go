package action

import (
	"encoding/json"
	"fmt"
	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/urfave/cli/v2"
)

func CreateSwarm(ctx *cli.Context) error {
	var config *configuration.Config
	var err error
	var configPath string
	var accessToken *string
	var operator *api.Operator
	var response *api.GenericIDResponseModel
	var swarmConfig map[string]interface{}

	name := ctx.String("name")
	description := ctx.String("description")
	swarmConfig1 := ctx.String("configuration")

	if err = json.Unmarshal([]byte(swarmConfig1), &swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorParsingJsonSettings, err)
	}

	if config, configPath, err = configuration.ReadConfig(ctx); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	if response, err = api.CreateSwarm(config.Urls, *accessToken, operator.ID, name, description, swarmConfig); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingSwarm, err)
	}

	fmt.Printf("Swarm %s created successfully\n", response.ID)

	return nil
}

func GetSwarm(ctx *cli.Context) error {

	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator
	var swarm *api.Swarm

	format := ctx.String("format")
	swarmID := ctx.String("id")
	swarmName := ctx.String("name")

	if config, configPath, err = configuration.ReadConfig(ctx); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}

	if swarmName != "" {
		var swarms []api.Swarm

		if swarms, err = api.ListSwarms(config.Urls, *accessToken, operator.ID); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
		}

		for _, sw := range swarms {
			if sw.Name == swarmName {
				utils.PrintFormattedData(sw, format)
				return nil
			}
		}

	}
	if swarm, err = api.GetSwarm(config.Urls, *accessToken, operator.ID, swarmID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}

	utils.PrintFormattedData(*swarm, format)

	return nil
}

func ListSwarms(ctx *cli.Context) error {
	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator
	var swarms []api.Swarm

	if config, configPath, err = configuration.ReadConfig(ctx); err != nil {
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

	verbose := ctx.Bool("verbose")

	for _, swarm := range swarms {
		if verbose {
			fmt.Printf("%s %s %s\n", swarm.SwarmID, swarm.Name, swarm.Description)
		} else {
			fmt.Printf("%s\n", swarm.Name)
		}
	}

	return nil
}
