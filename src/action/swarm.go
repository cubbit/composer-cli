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
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)

		}

	}
	if swarm, err = api.GetSwarm(config.Urls, *accessToken, operator.ID, swarmID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}

	utils.PrintFormattedData(*swarm, format)

	return nil
}
func EditSwarmDescription(ctx *cli.Context) error {

	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator

	swarmID := ctx.String("id")
	swarmName := ctx.String("name")
	description := ctx.Args().First()

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
				swarmID = sw.SwarmID
				break
			}
		}
	}
	if swarmID == "" {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}
	if err = api.EditSwarmDescription(config.Urls, *accessToken, operator.ID, swarmID, description); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarm, err)
	}

	return nil

}
func EditSwarmName(ctx *cli.Context) error {

	var err error
	var configPath string
	var accessToken *string
	var config *configuration.Config
	var operator *api.Operator

	swarmID := ctx.String("id")
	swarmName := ctx.String("name")
	newName := ctx.Args().First()

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
				swarmID = sw.SwarmID
				break
			}
		}
	}
	if swarmID == "" {
		return fmt.Errorf(constants.ErrorRetrievingSwarm)
	}

	if err = api.EditSwarmName(config.Urls, *accessToken, operator.ID, swarmID, newName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingSwarm, err)
	}

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

func ListSwarmOperators(ctx *cli.Context) error {
	var err error
	var accessToken *string
	var configPath string
	var config *configuration.Config
	var operator *api.Operator
	var operators *api.OperatorList

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
				swarmID = sw.SwarmID
				break
			}

		}

	}
	if swarmID == "" {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}
	if operators, err = api.ListSwarmOperators(config.Urls, *accessToken, swarmID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmOperators, err)
	}

	verbose := ctx.Bool("verbose")

	for _, operator := range operators.Operators {
		if verbose {
			fmt.Printf("%s %s %s\n", operator.ID, operator.FirstName, operator.LastName)
		} else {
			fmt.Printf("%s\n", operator.ID)
		}
	}

	return nil
}
