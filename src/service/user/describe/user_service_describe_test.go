package describe

import (
	"testing"

	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/spf13/cobra"
)

func TestDescribeUserRejectsSelfWithExplicitTarget(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		userID   string
		username string
		expected string
	}{
		{
			name:     "positional user id",
			args:     []string{"user-001"},
			expected: "failed request to describe IAM user: specify exactly one of USER_ID, --user-id, --username or --self",
		},
		{
			name:     "user id flag",
			userID:   "user-001",
			expected: "failed request to describe IAM user: specify exactly one of USER_ID, --user-id, --username or --self",
		},
		{
			name:     "username flag",
			username: "frank",
			expected: "failed request to describe IAM user: specify exactly one of USER_ID, --user-id, --username or --self",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("user-id", tc.userID, "")
			cmd.Flags().String("username", tc.username, "")
			cmd.Flags().Bool("self", true, "")

			err := DescribeUser(
				Dependencies{UserAPI: &api.MockUserAPI{}},
				cmd,
				nil,
				configuration_models.ProfileV2{},
				tc.args,
			)
			if err == nil {
				t.Fatal("Expected ambiguous target error, got nil")
			}
			if err.Error() != tc.expected {
				t.Fatalf("Expected error %q, got %q", tc.expected, err.Error())
			}
		})
	}
}
