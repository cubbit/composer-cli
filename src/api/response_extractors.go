package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/request_utils"
	"github.com/cubbit/composer-cli/utils"
)

func extractReport(response any) request_utils.RequestModifier {
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

func DownloadReport(output string, downloadedFile *string) request_utils.RequestModifier {
	return func(opt *request_utils.RequestOptions, res *http.Response) error {
		var err error
		var filename string
		var fileInfo fs.FileInfo

		if res == nil {
			return nil
		}

		filename = output
		if fileInfo, err = os.Stat(output); err == nil {

			if fileInfo.IsDir() {
				contentDisposition := res.Header.Get("Content-Disposition")
				if contentDisposition != "" {
					_, params, _ := mime.ParseMediaType(contentDisposition)
					filename = filepath.Join(output, params["filename"])
				}
			}
		}

		attachmentContent, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if filename != "" {
			err := ioutil.WriteFile(filename, attachmentContent, 0644)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("no filename found in Content-Disposition header")
		}

		*downloadedFile = filename
		return nil
	}
}

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
