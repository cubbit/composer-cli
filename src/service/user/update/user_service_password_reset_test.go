package update

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func TestResetUserPasswordRequiresExplicitTarget(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().String("user-id", "", "")
	cmd.Flags().String("username", "", "")
	cmd.Flags().String("password", "", "")
	cmd.Flags().String("new-password", "new-password", "")
	cmd.Flags().String("tfa-code", "", "")

	err := ResetUserPassword(
		Dependencies{
			UserAPI: &api.MockUserAPI{},
			AuthAPI: api.NewAuthAPI(),
		},
		cmd,
		configuration_models.ProfileV2{},
		nil,
	)
	if err == nil {
		t.Fatal("Expected missing target error, got nil")
	}

	expectedErr := "specify exactly one of USER_ID, --user-id or --username"
	if err.Error() != expectedErr {
		t.Fatalf("Expected error %q, got %q", expectedErr, err.Error())
	}
}
