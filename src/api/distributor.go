package api

import (
	"fmt"
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
		extractGenericIDResponseModel(&response),
		request_utils.WithAccessToken(accessToken),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorCreatingDistributor, err)
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
		extractDistributorListModel(&response),
	); err != nil {
		return nil, fmt.Errorf("%s: %w", constants.ErrorListingDistributorsRequest, err)
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
		return fmt.Errorf("%s: %w", constants.ErrorDeletingDistributor, err)
	}

	return nil
}
