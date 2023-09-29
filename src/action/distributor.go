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
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributorRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor %s created successfully", response.ID))

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
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
	}

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintList("Your Distributors List")

	if len(distributors.Distributors) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, distributor := range distributors.Distributors {
		if verbose {
			fmt.Printf(" • %s, %s, %s\n", distributor.ID, distributor.Name, distributor.Description)
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
			return fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
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

func CreateDistributorCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, couponName, description, configPath, zone string
	var maxRedemptions int
	var swarmIDs []string
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
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

	if couponName, err = cmd.Flags().GetString("coupon-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if maxRedemptions, err = cmd.Flags().GetInt("redemption-count"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if swarmIDs, err = cmd.Flags().GetStringSlice("swarms"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if zone, err = cmd.Flags().GetString("zone"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if zone != "" {
		var zones *api.ZoneMap
		if zones, err = api.GetGatwayZones(conf.Urls); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingZonesRequest, err)
		}

		if _, ok := zones.Zones[zone]; !ok {
			return fmt.Errorf(constants.ErrorInvalidZone)
		}
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)

		}
		id = distributor.ID
	}

	if response, err = api.CreateDistributorCoupon(conf.Urls, *accessToken, id, couponName, &description, swarmIDs, maxRedemptions, zone); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorCreatingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor coupon %s created successfully", response.ID))

	return nil
}

func ListDistributorCoupons(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config
	var distributorCoupons *api.DistributorCouponList

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)

		}
		id = distributor.ID
	}

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, *accessToken, id); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	var verbose, l bool

	if verbose, err = cmd.Flags().GetBool("verbose"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if l, err = cmd.Flags().GetBool("line"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintList("Your Distributor Coupons List")

	if len(distributorCoupons.Coupons) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	for _, coupon := range distributorCoupons.Coupons {
		if verbose {
			fmt.Printf("• %s, %s, %s, %s\n", coupon.ID, coupon.Name, coupon.Description, coupon.Zone)
		} else {
			fmt.Printf(" • %s\n", coupon.Name)
		}
		if l {
			fmt.Println()
		}
	}

	return nil
}

func DescribeDistributorCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, format, configPath string
	var conf *configuration.Config
	var distributorCoupon *api.DistributorCoupon

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)

		}
		id = distributor.ID
	}

	if distributorCoupon, err = getDistributorCouponByNameOrId(conf, *accessToken, id, args[0]); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponRequest, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	utils.PrintFormattedData(*distributorCoupon, format)

	return nil
}

func EditDistributorCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, couponName, description, configPath string
	var maxRedemptions int
	var response *api.GenericIDResponseModel
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if couponName, err = cmd.Flags().GetString("coupon-name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if description, err = cmd.Flags().GetString("description"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if len(description) > 200 {
		return fmt.Errorf("t%s: %w", constants.ErrorDescriptionSize, err)
	}

	if maxRedemptions, err = cmd.Flags().GetInt("redemption-count"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)

		}
		id = distributor.ID
	}

	var distributorCoupon *api.DistributorCoupon
	if distributorCoupon, err = getDistributorCouponByNameOrId(conf, *accessToken, id, args[0]); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponRequest, err)

	}
	couponID := distributorCoupon.ID

	if response, err = api.UpdateDistributorCoupon(conf.Urls, *accessToken, id, couponID, &couponName, &description, &maxRedemptions); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorEditingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("distributor coupon %s updated successfully", response.ID))

	return nil
}

func RevokeDistributorCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var response *api.DistributorCouponCodeResponseModel
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)

		}
		id = distributor.ID
	}

	var distributorCoupon *api.DistributorCoupon

	if distributorCoupon, err = getDistributorCouponByNameOrId(conf, *accessToken, id, args[0]); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponRequest, err)

	}
	couponID := distributorCoupon.ID

	if response, err = api.RevokeDistributorCoupon(conf.Urls, *accessToken, id, couponID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRevokingDistributorCouponRequest, err)
	}

	utils.PrintSuccess(fmt.Sprintf("new distributor coupon  code %s has been revoked successfully", response.CouponCode))

	return nil
}

func RemoveDistributorCoupon(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)

		}
		id = distributor.ID
	}

	var distributorCoupon *api.DistributorCoupon

	if distributorCoupon, err = getDistributorCouponByNameOrId(conf, *accessToken, id, args[0]); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributor, err)

	}
	couponID := distributorCoupon.ID

	if err = api.RemoveDistributorCoupon(conf.Urls, *accessToken, id, couponID); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRemovingDistributorCouponRequest, err)
	}

	utils.PrintDelete(fmt.Sprintf("distributor coupon %s removed successfully", couponID))

	return nil
}

func GetDistributorReport(cmd *cobra.Command, args []string) error {
	var err error
	var accessToken *string
	var id, name, configPath, coupon, format, from, to, output string
	var conf *configuration.Config

	if id, err = cmd.Flags().GetString("id"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if name, err = cmd.Flags().GetString("name"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if from, err = cmd.Flags().GetString("from"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if to, err = cmd.Flags().GetString("to"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if format, err = cmd.Flags().GetString("format"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if output, err = cmd.Flags().GetString("output"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if coupon, err = cmd.Flags().GetString("coupon"); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingField, err)
	}

	if conf, configPath, err = configuration.ReadConfig(cmd); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorLoadingConfig, err)
	}

	if accessToken, err = rehydrateTokenConfig(configPath, conf); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorGeneratingToken, err)
	}

	if id == "" {
		var distributor *api.Distributor

		if distributor, err = getDistributorByNameOrId(conf, *accessToken, name); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorRequest, err)

		}
		id = distributor.ID
	}

	if coupon != "" {
		var distributorCoupon *api.DistributorCoupon
		if distributorCoupon, err = getDistributorCouponByNameOrId(conf, *accessToken, id, args[0]); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorCouponRequest, err)
		}
		coupon = distributorCoupon.ID
	}

	if output != "" {
		var fileName *string
		if fileName, err = api.DownloadDistributorReport(conf.Urls, *accessToken, id, coupon, from, to, output); err != nil {
			return fmt.Errorf("%s: %w", constants.ErrorDownloadingDistributorReportRequest, err)
		}

		utils.PrintSuccess(fmt.Sprintf("report downloaded successfully to : %s", *fileName))
		return nil
	}

	var distributorReport *api.DistributorReportResponseModel
	if distributorReport, err = api.GetDistributorReport(conf.Urls, *accessToken, id, coupon, from, to); err != nil {
		return fmt.Errorf("%s: %w", constants.ErrorRetrievingDistributorReportRequest, err)
	}

	utils.PrintList("Your Distributor Report")

	if len(distributorReport.Report) == 0 {
		utils.PrintEmptyList()
		return nil
	}

	utils.PrintFormattedData(distributorReport.Report, format)

	return nil
}

func getDistributorByNameOrId(conf *configuration.Config, accessToken string, distributor string) (*api.Distributor, error) {
	var err error
	var distributors *api.DistributorList

	if distributors, err = api.ListDistributors(conf.Urls, accessToken); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
	}

	for _, ds := range distributors.Distributors {
		if distributor == ds.Name || distributor == ds.ID {
			return ds, nil

		}
	}

	return nil, fmt.Errorf("distributor %s not found", distributor)

}

func getDistributorCouponByNameOrId(conf *configuration.Config, accessToken string, distributorID string, coupon string) (*api.DistributorCoupon, error) {
	var err error
	var distributorCoupons *api.DistributorCouponList

	if distributorCoupons, err = api.ListDistributorCoupons(conf.Urls, accessToken, distributorID); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingDistributorCouponsRequest, err)
	}

	for _, dc := range distributorCoupons.Coupons {
		if coupon == dc.Name || coupon == dc.ID {
			return dc, nil

		}
	}

	return nil, fmt.Errorf("distributor coupon %s not found", coupon)

}
