package action

import (
	"encoding/json"
	"fmt"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
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

	swarmID := ctx.String("id")

	if config, configPath, err = configuration.ReadConfig(ctx); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}
	if accessToken, err = rehydrateTokenConfig(configPath, config); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(config.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperator, err)
	}
	if swarm, err = api.GetSwarm(config.Urls, *accessToken, operator.ID, swarmID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarm, err)
	}

	fmt.Printf("%v\n", swarm)

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
			fmt.Printf("%v\n", swarm)
		} else {
			fmt.Printf("%s\n", swarm.Name)
		}
	}

	return nil
}

func EditSwarm(ctx *cli.Context) {
	// TODO
}

func DeleteSwarm(clx *cli.Context) {
	// TODO
}

func AddProviderToSwarm(ctx *cli.Context) {
	// TODO
}

func ListSwarmProviders(ctx *cli.Context) {
	// TODO
}

func AddSecretToSwarm(ctx *cli.Context) {
	// TODO
}

func ListSwarmSecrets(ctx *cli.Context) {
	// TODO
}

func ListSwarmErrors(ctx *cli.Context) {
	// TODO
}
