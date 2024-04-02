package api

import (
	"net/http"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/configuration"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
)

func CreateDistributor(urls configuration.Url, accessToken, name string, description *string, imageUrl *string, swarmIDs []string, email, firstName, lastName string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + constants.Distributors

	requestBody := map[string]interface{}{
		"name":   name,
		"swarms": swarmIDs,
		"invite": map[string]interface{}{
			"email":      email,
			"first_name": firstName,
			"last_name":  lastName,
		},
	}

	if description != nil {
		requestBody["description"] = *description
	}

	if imageUrl != nil {
		requestBody["image_url"] = *imageUrl
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListDistributors(urls configuration.Url, accessToken string) (*DistributorList, error) {
	var err error
	url := urls.IamUrl + constants.Distributors
	var response DistributorList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func RemoveDistributor(urls configuration.Url, accessToken, distributorId, deleteDistributorToken string) error {
	var err error
	url := urls.IamUrl + constants.Distributors + "/" + distributorId + "?token=" + deleteDistributorToken
	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusNoContent),
	); err != nil {
		return err
	}

	return nil
}

func GetDistributorReport(urls configuration.Url, accessToken, distributorId, couponId, from, to string) (*DistributorReportResponseModel, error) {
	var err error
	var url string
	var response DistributorReportResponseModel

	if couponId != "" {
		url = urls.DashUrl + constants.BaseDashURI + constants.Distributors + "/" + distributorId + "/coupons/" + couponId + "/report" + "?from_date=" + from + "&to_date=" + to

	} else {
		url = urls.DashUrl + constants.BaseDashURI + constants.Distributors + "/" + distributorId + "/report" + "?from_date=" + from + "&to_date=" + to

	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		extractReport(&response),
	); err != nil {
		return nil, err
	}

	return &response, err
}

func DownloadDistributorReport(urls configuration.Url, accessToken, distributorId, couponId, from, to, output string) (*string, error) {
	var err error
	var url string
	var response string
	if couponId != "" {
		url = urls.DashUrl + constants.BaseDashURI + constants.Distributors + "/" + distributorId + "/coupons/" + couponId + "/report" + "?from_date=" + from + "&to_date=" + to

	} else {
		url = urls.DashUrl + constants.BaseDashURI + constants.Distributors + "/" + distributorId + "/report" + "?from_date=" + from + "&to_date=" + to

	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodGet),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAttachement(),
		DownloadReport(output, &response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func CreateDistributorCoupon(urls configuration.Url, accessToken, distributorID, name string, description *string, swarmIDs []string, maxRedemptions int, zone string) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons"

	requestBody := map[string]interface{}{
		"name":            name,
		"swarms":          swarmIDs,
		"max_redemptions": maxRedemptions,
	}

	if description != nil {
		requestBody["description"] = description
	}

	if zone != "" {
		requestBody["zone"] = zone
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusCreated),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}

	return &response, nil
}

func ListDistributorCoupons(urls configuration.Url, accessToken, distributorID string) (*DistributorCouponList, error) {
	var err error
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons"
	var response DistributorCouponList

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func GetDistributorCoupon(urls configuration.Url, accessToken, distributorID, couponID string) (*DistributorCoupon, error) {
	var err error
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons/" + couponID
	var response DistributorCoupon

	if err = request_utils.DoRequest(
		url,
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func UpdateDistributorCoupon(urls configuration.Url, accessToken string, distributorID, couponID string, name, description *string, maxRedemption *int) (*GenericIDResponseModel, error) {
	var err error
	var response GenericIDResponseModel
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons/" + couponID

	requestBody := map[string]interface{}{}

	if name != nil && *name != "" {
		requestBody["name"] = *name
	}

	if description != nil && *description != "" {
		requestBody["description"] = *description
	}

	if maxRedemption != nil && *maxRedemption != 0 {
		requestBody["max_redemptions"] = *maxRedemption
	}

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPatch),
		request_utils.WithRequestBody(requestBody),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func RevokeDistributorCoupon(urls configuration.Url, accessToken string, distributorID, couponID string) (*DistributorCouponCodeResponseModel, error) {
	var err error
	var response DistributorCouponCodeResponseModel
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons/" + couponID + "/revoke"

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodPost),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		ExtractGenericModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func RemoveDistributorCoupon(urls configuration.Url, accessToken string, distributorID, couponID string) error {
	var err error
	url := urls.IamUrl + constants.Distributors + "/" + distributorID + "/coupons/" + couponID

	if err = request_utils.DoRequest(
		url,
		request_utils.WithRequestMethod(http.MethodDelete),
		request_utils.WithAccessToken(accessToken),
		request_utils.WithExpectedStatusCode(http.StatusOK),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return err
	}
	return nil
}
