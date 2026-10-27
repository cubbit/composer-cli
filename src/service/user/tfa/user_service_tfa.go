package tfa

import (
	"crypto/hmac"
	crand "crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_handler"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/src/tui"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/spf13/cobra"
)

const recoveryCodeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

type Dependencies struct {
	AuthAPI api.AuthAPIInterface
	UserAPI api.UserAPIInterface
}

func EnableUserTFA(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
) error {
	operator, err := getCurrentUser(deps, profile)
	if err != nil {
		return err
	}
	if operator.TwoFactorEnabled {
		return fmt.Errorf("IAM user TFA is already enabled")
	}

	setupKey, err := readSetupKey(cmd)
	if err != nil {
		return err
	}
	if setupKey == "" {
		setupKey, err = generateTOTPSecret()
		if err != nil {
			return fmt.Errorf("failed to generate setup key: %w", err)
		}
	}
	setupKey = strings.ToUpper(strings.ReplaceAll(setupKey, " ", ""))

	recoveryCodes, err := generateRecoveryCodes(10)
	if err != nil {
		return fmt.Errorf("failed to generate recovery codes: %w", err)
	}

	operatorPassword, err := readOperatorPassword(cmd)
	if err != nil {
		return err
	}

	enableTFAToken, err := deps.AuthAPI.ForgeToken(
		profile.Endpoints,
		operator.ID,
		operator.Username,
		*operator.OrganizationName,
		operatorPassword,
		"",
		"enable_tfa",
		"",
		"",
	)
	if err != nil {
		return fmt.Errorf("failed to forge enable TFA token: %w", err)
	}

	validationCode, err := generateTOTPCode(setupKey, time.Now())
	if err != nil {
		return fmt.Errorf("failed to generate TOTP validation code: %w", err)
	}

	if err := deps.UserAPI.EnableIAMUserTFA(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		operator.ID,
		enableTFAToken,
		&api.UpdateIAMUserTFARequestBody{
			Secret:         setupKey,
			ValidationCode: validationCode,
			Codes:          recoveryCodes,
		},
	); err != nil {
		return err
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if quiet {
		quietFields := append([]string{setupKey}, recoveryCodes...)
		utils.PrintQuiet(cmd.OutOrStdout(), quietFields...)
		return nil
	}

	var output strings.Builder
	output.WriteString("IAM user TFA enabled successfully\n")
	output.WriteString(fmt.Sprintf("Setup key: %s\n", setupKey))
	output.WriteString("Recovery codes:\n")
	for _, code := range recoveryCodes {
		output.WriteString(fmt.Sprintf("  %s\n", code))
	}

	return printer.PrintText(cmd, handler, output.String())
}

func readSetupKey(cmd *cobra.Command) (string, error) {
	setupKey, err := cmd.Flags().GetString("setup-key")
	if err != nil {
		return "", fmt.Errorf("%s setup-key: %w", constants.ErrorRetrievingField, err)
	}

	return setupKey, nil
}

func DisableUserTFA(
	deps Dependencies,
	cmd *cobra.Command,
	handler configuration_handler.ConfigurationHandlerInterface,
	profile configuration_models.ProfileV2,
) error {
	operator, err := getCurrentUser(deps, profile)
	if err != nil {
		return err
	}
	if !operator.TwoFactorEnabled {
		quiet, err := cmd.Flags().GetBool("quiet")
		if err != nil {
			return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
		}
		if !quiet {
			return printer.PrintText(cmd, handler, "IAM user TFA is already disabled\n")
		}
		return nil
	}

	operatorPassword, err := readOperatorPassword(cmd)
	if err != nil {
		return err
	}

	tfaCode, err := cmd.Flags().GetString("tfa-code")
	if err != nil {
		return fmt.Errorf("%s tfa-code: %w", constants.ErrorRetrievingField, err)
	}
	tfaCode = strings.TrimSpace(tfaCode)
	if tfaCode == "" {
		tfaCode, err = readTFAInput("Please provide your TFA code", "TFA code")
		if err != nil {
			return err
		}
	}

	deleteTFAToken, err := deps.AuthAPI.ForgeToken(
		profile.Endpoints,
		operator.ID,
		operator.Username,
		*operator.OrganizationName,
		operatorPassword,
		tfaCode,
		"delete_tfa",
		"",
		"",
	)
	if err != nil {
		return fmt.Errorf("failed to forge delete TFA token: %w", err)
	}

	if err := deps.UserAPI.DisableIAMUserTFA(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		operator.ID,
		deleteTFAToken,
	); err != nil {
		return err
	}

	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("%s quiet: %w", constants.ErrorRetrievingField, err)
	}
	if !quiet {
		return printer.PrintText(cmd, handler, "IAM user TFA disabled successfully\n")
	}

	return nil
}

func readTFAInput(title string, placeholder string) (string, error) {
	var tfaCode string
	if _, err := tui.TextInputs(
		title,
		false,
		tui.Input{Placeholder: placeholder, Value: &tfaCode},
	); err != nil {
		return "", fmt.Errorf("failed to read TFA code: %w", err)
	}

	tfaCode = strings.TrimSpace(tfaCode)
	if tfaCode == "" {
		return "", fmt.Errorf("TFA code is required")
	}

	return tfaCode, nil
}

func readOperatorPassword(cmd *cobra.Command) (string, error) {
	operatorPassword, err := cmd.Flags().GetString("password")
	if err != nil {
		return "", fmt.Errorf("%s password: %w", constants.ErrorRetrievingField, err)
	}
	if operatorPassword == "" {
		operatorPassword = os.Getenv("PASSWORD")
	}
	if operatorPassword == "" {
		if _, err = tui.TextInputs(
			"Please provide your login details",
			false,
			tui.Input{Placeholder: "password", IsPassword: true, Value: &operatorPassword},
		); err != nil {
			return "", fmt.Errorf("failed to read operator password: %w", err)
		}
	}

	return operatorPassword, nil
}

func getCurrentUser(
	deps Dependencies,
	profile configuration_models.ProfileV2,
) (*api.IAMUser, error) {
	operator, err := deps.UserAPI.GetIAMUserSelf(profile.Endpoints, "", profile.APIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve current IAM user: %w", err)
	}
	if operator.ID == "" {
		return nil, fmt.Errorf("current IAM user does not expose an id")
	}
	if operator.Username == "" {
		return nil, fmt.Errorf("current IAM user does not expose a username")
	}
	if operator.OrganizationName == nil || *operator.OrganizationName == "" {
		return nil, fmt.Errorf("current IAM user does not expose an organization name")
	}

	return operator, nil
}

func generateTOTPSecret() (string, error) {
	buffer := make([]byte, 20)
	if _, err := crand.Read(buffer); err != nil {
		return "", err
	}

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(buffer), nil
}

func generateTOTPCode(secret string, now time.Time) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		key, err = base32.StdEncoding.DecodeString(secret)
		if err != nil {
			return "", err
		}
	}

	counter := uint64(now.Unix() / 30)
	message := make([]byte, 8)
	binary.BigEndian.PutUint64(message, counter)

	hash := hmac.New(sha1.New, key)
	if _, err := hash.Write(message); err != nil {
		return "", err
	}

	sum := hash.Sum(nil)
	offset := sum[len(sum)-1] & 0x0f
	binaryCode := (uint32(sum[offset])&0x7f)<<24 |
		(uint32(sum[offset+1])&0xff)<<16 |
		(uint32(sum[offset+2])&0xff)<<8 |
		(uint32(sum[offset+3]) & 0xff)

	return fmt.Sprintf("%06d", binaryCode%1000000), nil
}

func generateRecoveryCodes(count int) ([]string, error) {
	codes := make([]string, 0, count)
	for range count {
		left, err := randomRecoveryCodePart()
		if err != nil {
			return nil, err
		}
		right, err := randomRecoveryCodePart()
		if err != nil {
			return nil, err
		}
		codes = append(codes, left+"-"+right)
	}

	return codes, nil
}

func randomRecoveryCodePart() (string, error) {
	var builder strings.Builder
	builder.Grow(5)

	for range 5 {
		index, err := crand.Int(crand.Reader, big.NewInt(int64(len(recoveryCodeAlphabet))))
		if err != nil {
			return "", err
		}
		builder.WriteByte(recoveryCodeAlphabet[index.Int64()])
	}

	return builder.String(), nil
}
