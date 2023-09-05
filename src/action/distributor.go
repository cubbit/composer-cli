package action

import (
	"fmt"
	"net/url"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/api"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/utils"
	"github.com/spf13/cobra"
)

func CreateDistributor(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var name, description, imageUrl, email, firstName, lastName, configPath string
	var swarmIDs []string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if imageUrl, err = cmd.Flags().GetString("image-url"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if imageUrl != "" {
		if _, err := url.ParseRequestURI(imageUrl); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorImageURL, err)
		}
	}

	if email, err = cmd.Flags().GetString("owner"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if firstName, err = cmd.Flags().GetString("first-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if lastName, err = cmd.Flags().GetString("last-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if swarmIDs, err = cmd.Flags().GetStringSlice("swarms"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if response, err = api.CreateDistributor(conf.Urls, *accessToken, name, &description, &imageUrl, swarmIDs, email, firstName, lastName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributor, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributer: %s created successfully", response.ID))

	return nil
}

func ListDistributor(cmd *cobra.Command, args []string) error {
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

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintList("Your Distributors List")

	for _, distributor := range distributors.Distributors {
		if verbose {
			fmt.Printf(" • %s, %s, %s\n", distributor.ID, distributor.Name, *distributor.Description)
		} else {
			fmt.Printf(" • %s\n", distributor.Name)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func RemoveDistributor(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, email, password, code, configPath, deleteDistributorToken string
	var conf *configuration.Config
	var challenge *api.ChallengeResponseModel

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if email, err = cmd.Flags().GetString("email"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if password, err = cmd.Flags().GetString("password"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if code, err = cmd.Flags().GetString("code"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributors *api.DistributorList

		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		for _, distributor := range distributors.Distributors {
			if name == distributor.Name {
				id = distributor.ID
			}
		}

		if id == "" {
			utils.PrintNotFound(fmt.Sprintf("Distributor %s not found", name))
			return nil
		}
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
