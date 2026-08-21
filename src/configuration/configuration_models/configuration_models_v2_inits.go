package configuration_models

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/cubbit/composer-cli/constants"
)

type EndpointsV2Inits struct {
	Base string `toml:"base,omitempty"`
	IAM  string `toml:"iam,omitempty"`
	Dash string `toml:"dash,omitempty"`
	CH   string `toml:"ch,omitempty"`
}

func (e EndpointsV2Inits) Validate() error {
	if e.Base == "" {
		if e.IAM == "" || e.Dash == "" || e.CH == "" {
			return fmt.Errorf("either base endpoint or all individual endpoints (IAM, Dash, CH) must be set")
		}

		return nil
	}

	return nil
}

func (e EndpointsV2Inits) ToEndpointsV2() *EndpointsV2 {
	iam := e.IAM
	if iam == "" {
		iam = e.Base + constants.BaseIamURI
	}

	dash := e.Dash
	if dash == "" {
		dash = e.Base + constants.BaseDashURI
	}

	ch := e.CH
	if ch == "" {
		ch = e.Base + constants.BaseChURI
	}

	return &EndpointsV2{
		IAM:  iam,
		Dash: dash,
		CH:   ch,
	}
}

func ParseAndValidateEndpointsV2Inits(path string) (*EndpointsV2Inits, error) {
	var endpoints EndpointsV2Inits
	if _, err := toml.DecodeFile(path, &endpoints); err != nil {
		return nil, fmt.Errorf("failed to decode configuration v2 inits file %q: %w", path, err)
	}

	if err := endpoints.Validate(); err != nil {
		return nil, fmt.Errorf("invalid endpoints configuration: %w", err)
	}

	return &endpoints, nil
}
