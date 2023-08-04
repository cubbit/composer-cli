package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cubbit/cubbit/client/cli/constants"
	"github.com/cubbit/cubbit/client/cli/src/request_utils"
	"github.com/cubbit/cubbit/client/cli/utils"
)

func extractChallengeResponseModel(response *ChallengeResponseModel) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractRefreshCookie(response *string) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		if res == nil {
			return nil
		}

		refreshTokenCookie, found := utils.Find(res.Cookies(), func(c *http.Cookie) bool {
			return c.Name == constants.RefreshTokenName
		})
		if !found {
			return fmt.Errorf(constants.ErrorTokenNotFound)
		}

		if refreshTokenCookie.Value == "" {
			return fmt.Errorf(constants.ErrorEmptyToken)
		}

		*response = refreshTokenCookie.Value
		return nil
	}
}

func extractTokenExpirationModel(response *TokenAndExpirationResponseModel) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractGenericIDResponseModel(response *GenericIDResponseModel) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractTenantListModel(response *TenantList) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractOperatorResponseModel(response *Operator) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractSwarmListModel(response *SwarmList) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {

		var body []byte
		var err error

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}
		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractSwarmResponseModel(response *Swarm) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func extractSwarmListResponseModel(response *[]Swarm) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {

		var body []byte
		var err error

		if res == nil {
			return nil
		}

		if body, err = ioutil.ReadAll(res.Body); err != nil {
			return err
		}
		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}
