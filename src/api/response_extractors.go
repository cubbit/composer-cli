package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/request_utils"
	"github.com/cubbit/composer-cli/utils"
)

func ExtractGenericModel(response any) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var body []byte

		if res == nil || response == nil {
			return nil
		}

		if body, err = io.ReadAll(res.Body); err != nil {
			return err
		}

		if err = json.Unmarshal(body, &response); err != nil {
			return err
		}

		return nil
	}
}

func ExtractRefreshCookie(response *string) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		if res == nil {
			return nil
		}

		refreshTokenCookie, found := utils.Find(res.Cookies(), func(c *http.Cookie) bool {
			return c.Name == constants.RefreshCookie
		})
		if !found {
			return fmt.Errorf("refresh token not found in cookies")
		}

		if refreshTokenCookie.Value == "" {
			return errors.New("refresh token cannot be empty")
		}

		if response == nil {
			return errors.New("response cannot be nil")
		}

		*response = refreshTokenCookie.Value
		return nil
	}
}
