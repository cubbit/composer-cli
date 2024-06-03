package action

import (
	"fmt"
	"net/url"
	"strconv"
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
	var swarms []*api.Swarm
	var choices []string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &name},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &description},
		tui.Input{Placeholder: "Image URL", IsPassword: false, Value: &imageUrl}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if _, err = tui.TextInputs("Fill in the form to invite the associate operator",
		false,
		tui.Input{Placeholder: "First Name", IsPassword: false, Value: &firstName},
		tui.Input{Placeholder: "Last Name", IsPassword: false, Value: &lastName},
		tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}); err != nil {

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
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	if len(swarms) == 0 {
		utils.PrintNotFound("No swarms found")
		return nil
	}

	for _, sw := range swarms {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", sw.ID, sw.Name, sw.Description))
	}

	if choices, err = tui.ChooseMany("Which swarm would you like to associate to the distributor?", true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
	}

	for _, swarm := range choices {
		_, withoutPrefix, _ := strings.Cut(swarm, " ")
		delId, _, _ := strings.Cut(withoutPrefix, ",")
		swarmIDs = append(swarmIDs, delId)
	}

	if response, err = api.CreateDistributor(conf.Urls, *accessToken, name, &description, &imageUrl, swarmIDs, email, firstName, lastName); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor %s created successfully", response.ID))

	return nil
}

func RemoveDistributorInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, email, password, code, configPath, choice, deleteDistributorToken string
	var conf *configuration.Config
	var distributors *api.DistributorList
	var challenge *api.ChallengeResponseModel
	var choices []string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
	}

	if len(distributors.Distributors) == 0 {
		utils.PrintNotFound("No distributors found")
		return nil
	}

	for _, distributor := range distributors.Distributors {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributors found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor would you like to remove?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributor, err)
	}

	splits := strings.Split(choice, ",")
	_, id, _ = strings.Cut(splits[0], " ")
	_, name, _ := strings.Cut(splits[1], " ")

	if _, err = tui.TextInputs(fmt.Sprintf("Confirm your login to remove the distributor %s - %s 🚮", utils.RedBg.Render(name), utils.RedBg.Render(id)), true, tui.Input{Placeholder: "Email*", IsPassword: false, Value: &email}, tui.Input{Placeholder: "Password*", IsPassword: true, Value: &password}, tui.Input{Placeholder: "Code", IsPassword: false, Value: &code}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if challenge, err = api.GenerateOperatorChallenge(conf.Urls, email); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingOperatorChallengeRequest, err)
	}

	if deleteDistributorToken, err = api.ForgeDistributorDeleteToken(conf.Urls, email, password, conf.RefreshToken, challenge, code, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorForgingDistributorDeleteTokenRequest, err)
	}

	if err = api.RemoveDistributor(conf.Urls, *accessToken, id, deleteDistributorToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributorRequest, err)
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

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, true); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
	}

	utils.PrintList("Your Distributors List")

	if len(distributors.Distributors) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	var list []string
	for _, distributor := range distributors.Distributors {
		list = append(list, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
	}

	tui.List(list)

	return nil
}

func CreateDistributorCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, couponName, description, redemptionCount, configPath, zone, externalID string
	var maxRedemptions int
	var swarmIDs []string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var operator *api.Operator
	var swarms []*api.Swarm
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributor *api.Distributor
	var zones *api.ZoneMap

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if operator, err = api.GetOperatorSelf(conf.Urls, *accessToken); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingOperatorRequest, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponRequest, err)
		}
		id = distributor.ID
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "Name*", IsPassword: false, Value: &couponName},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &description},
		tui.Input{Placeholder: "Redemption Count", IsPassword: false, Value: &redemptionCount},
		tui.Input{Placeholder: "External ID", IsPassword: false, Value: &externalID},
	); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if redemptionCount != "" {

		if maxRedemptions, err = strconv.Atoi(redemptionCount); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRedemptionValue, err)
		}
	} else {
		maxRedemptions = -1
	}

	choices = []string{}
	if zones, err = api.GetGatwayZones(conf.Urls); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
	}

	if len(zones.Zones) != 0 {
		choices = append(choices, fmt.Sprintf("• %s", "Default"))
		for _, zn := range zones.Zones {
			choices = append(choices, fmt.Sprintf("• %s", zn.Name))
		}

		if choice, err = tui.ChooseOne("Which zone would you like to create your distributor code?", true, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
		}

		if choice != "" {
			_, value, _ := strings.Cut(choice, " ")
			if value != "Default" {
				for _, zn := range zones.Zones {

					if value == zn.Name {
						zone = zn.Key
						break
					}
				}
			}
		}
	}

	if swarms, err = api.ListSwarms(conf.Urls, *accessToken, operator.ID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingSwarmsRequest, err)
	}

	if len(swarms) == 0 {
		utils.PrintNotFound("No swarms found")
		return nil
	}

	choices = []string{}
	for _, sw := range swarms {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", sw.ID, sw.Name, sw.Description))
	}

	if choices, err = tui.ChooseMany("Which swarm would you like to associate to the distributor code?", true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
	}

	for _, swarm := range choices {
		_, withoutPrefix, _ := strings.Cut(swarm, " ")
		delId, _, _ := strings.Cut(withoutPrefix, ",")
		swarmIDs = append(swarmIDs, delId)
	}

	if response, err = api.CreateDistributorCoupon(conf.Urls, *accessToken, id, couponName, &description, swarmIDs, maxRedemptions, zone, externalID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor code %s created successfully", response.ID))

	return nil
}

func ListDistributorCouponsInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var distributorCoupons *api.DistributorCouponList
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributor *api.Distributor
	var list []string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)
		}
		id = distributor.ID
	}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	utils.PrintList("Your Distributor Codes List")

	if len(distributorCoupons.Coupons) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, coupon := range distributorCoupons.Coupons {
		list = append(list, fmt.Sprintf("• %s, %s, %s, %s", coupon.ID, coupon.Name, coupon.Description, coupon.Zone))
	}

	tui.List(list)

	return nil
}

func DescribeDistributorCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributor *api.Distributor
	var distributorCoupons *api.DistributorCouponList
	var format string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to select?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)
		}
		id = distributor.ID
	}

	choices = []string{}
	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor code found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor code would you like to describe?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	id, _, _ = strings.Cut(withoutPrefix, ",")

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		if id == coupon.ID {
			if coupon.MaxRedemptions == -1 {
				utils.PrintFormattedData(*coupon.ToHumanReadableDistributorCode(), format)
				return nil
			}

			utils.PrintFormattedData(*coupon, format)
			return nil
		}
	}

	return nil
}

func EditDistributorCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, couponName, couponID, description, redemptionCount, configPath, externalID string
	var maxRedemptions int
	var response *api.GenericIDResponseModel
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributor *api.Distributor
	var distributorCoupons *api.DistributorCouponList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingSwarmList, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)
		}
		id = distributor.ID
	}

	choices = []string{}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor code found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor code would you like to edit?", false, false, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	couponID, _, _ = strings.Cut(withoutPrefix, ",")

	if _, err = tui.TextInputs(
		"Fill in the form below",
		true,
		tui.Input{Placeholder: "Name", IsPassword: false, Value: &couponName},
		tui.Input{Placeholder: "Description", IsPassword: false, Value: &description},
		tui.Input{Placeholder: "Redemption Count", IsPassword: false, Value: &redemptionCount},
		tui.Input{Placeholder: "External ID", IsPassword: false, Value: &externalID}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("%s: %w", constants.ErrorDescriptionSize, err)
	}

	if redemptionCount != "" {

		if maxRedemptions, err = strconv.Atoi(redemptionCount); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRedemptionValue, err)
		}
	}

	if response, err = api.UpdateDistributorCoupon(conf.Urls, *accessToken, id, couponID, &couponName, &description, &maxRedemptions, &externalID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor code %s updated successfully", response.ID))

	return nil
}

func RevokeDistributorCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, couponID, configPath string
	var response *api.DistributorCouponCodeResponseModel
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributorCoupons *api.DistributorCouponList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var distributor *api.Distributor
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)
		}
		id = distributor.ID
	}

	choices = []string{}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor code found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor code would you like to revoke?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	couponID, _, _ = strings.Cut(withoutPrefix, ",")

	if response, err = api.RevokeDistributorCoupon(conf.Urls, *accessToken, id, couponID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRevokingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("new distributor code %s has been revoked successfully", response.CouponCode))

	return nil
}

func RemoveDistributorCouponInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, couponID, configPath string
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributorCoupons *api.DistributorCouponList

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to select?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		var distributor *api.Distributor
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)
		}
		id = distributor.ID
	}

	choices = []string{}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Description))
	}

	if len(choices) == 0 {
		utils.PrintNotFound("No distributor code found")
		return nil
	}

	if choice, err = tui.ChooseOne("Which distributor code would you like to remove?", false, true, choices); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
	}

	_, withoutPrefix, _ := strings.Cut(choice, " ")
	couponID, _, _ = strings.Cut(withoutPrefix, ",")

	if err = api.RemoveDistributorCoupon(conf.Urls, *accessToken, id, couponID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingDistributorCouponRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("distributor code %s removed successfully", couponID))

	return nil
}

func GetDistributorReportInteractive(cmd *cobra.Command) error {
	var err error
	var accessToken *string
	var id, name, configPath, coupon, from, to, output string
	var conf *configuration.Config
	var choice string
	var choices []string
	var distributors *api.DistributorList
	var distributorCoupons *api.DistributorCouponList
	var distributor *api.Distributor
	var download string
	var distributorReport *api.DistributorReportResponseModel
	var format string

	if conf, configPath, err = configuration.ReadConfig(cmd, configuration.SessionTypeOperator, false); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if id == "" && name == "" {
		if distributors, err = api.ListDistributors(conf.Urls, *accessToken); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
		}

		for _, distributor := range distributors.Distributors {
			choices = append(choices, fmt.Sprintf("• %s, %s, %s", distributor.ID, distributor.Name, distributor.Description))
		}

		if len(choices) == 0 {
			utils.PrintNotFound("No distributor found")
			return nil
		}

		if choice, err = tui.ChooseOne("Which distributor would you like to choose?", false, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		id, _, _ = strings.Cut(withoutPrefix, ",")
	}

	if id == "" {
		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)
		}
		id = distributor.ID
	}

	if _, err = tui.TextInputs(
		"Fill in the form below",
		false,
		tui.Input{Placeholder: "From* (DD/MM/YYY+HH:mm:ss)", IsPassword: false, Value: &from},
		tui.Input{Placeholder: "To* (DD/MM/YYY+HH:mm:ss)", IsPassword: false, Value: &to}); err != nil {

		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	choices = []string{}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, coupon := range distributorCoupons.Coupons {
		choices = append(choices, fmt.Sprintf("• %s, %s, %s", coupon.ID, coupon.Name, coupon.Description))
	}

	if len(choices) != 0 {
		if choice, err = tui.ChooseOne("Choose your distributor code", true, false, choices); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponList, err)
		}

		_, withoutPrefix, _ := strings.Cut(choice, " ")
		coupon, _, _ = strings.Cut(withoutPrefix, ",")

	}

	if download, err = tui.ChooseOne("Do you want to download the report?", false, false, []string{"Yes", "No"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	if download == "Yes" {
		if _, err = tui.TextInputs("File name or directory to download the report", true, tui.Input{Placeholder: "File Name or Dir*", IsPassword: false, Value: &output}); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
		}
		if output == "" {
			output = "."
		}

		var downloadedFile *string
		if downloadedFile, err = api.DownloadDistributorReport(conf.Urls, *accessToken, id, coupon, from, to, output); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDownloadingDistributorReportRequest, err)
		}

		utils.PrintSuccess(fmt.Sprintf("report downloaded successfully to : %s", *downloadedFile))
		return nil
	}

	if distributorReport, err = api.GetDistributorReport(conf.Urls, *accessToken, id, coupon, from, to); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorReportRequest, err)
	}

	if format, err = tui.ChooseOne("Choose your output format", false, true, []string{"json", "semantic", "csv"}); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRunningField, err)
	}

	utils.PrintList("Your Distributor Report")

	if len(distributorReport.Report) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintFormattedData(distributorReport.Report, format)

	return nil
}
