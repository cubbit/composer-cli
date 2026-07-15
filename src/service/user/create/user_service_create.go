package create

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cubbit/composer-cli/constants"
	"github.com/cubbit/composer-cli/src/api"
	"github.com/cubbit/composer-cli/src/configuration/configuration_models"
	"github.com/cubbit/composer-cli/utils"
	"github.com/cubbit/composer-cli/utils/printer"
	"github.com/cubbit/composer-cli/utils/printer/table"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type challengeAPI interface {
	GenerateChallenge(
		endpoints configuration_models.EndpointsV2,
		email *string,
		username *string,
		organizationName *string,
	) (*api.ChallengeResponseModel, error)
}

type Dependencies struct {
	UserAPI api.UserAPIInterface
	AuthAPI challengeAPI
}

type bulkCreateUsersFile struct {
	Users []bulkCreateUserFileItem `json:"users"`
}

type bulkCreateUserFileItem struct {
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	FirstName        *string  `json:"first_name,omitempty"`
	LastName         *string  `json:"last_name,omitempty"`
	Email            *string  `json:"email,omitempty"`
	AttachedPolicies []string `json:"attached_policies,omitempty"`
}

var (
	usernamePattern = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-_]+[a-zA-Z0-9]$`)
)

func ImportUsers(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, organizationName string) error {
	sampleFormat, err := cmd.Flags().GetString("sample")
	if err != nil {
		return fmt.Errorf("%s sample: %w", constants.ErrorRetrievingField, err)
	}
	if sampleFormat != "" {
		return printImportSample(cmd, sampleFormat)
	}

	filePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return fmt.Errorf("%s file: %w", constants.ErrorRetrievingField, err)
	}
	if filePath == "" {
		return fmt.Errorf("file is required")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %q: %w", filePath, err)
	}

	usersFile, err := parseBulkCreateUsersFile(filePath, data)
	if err != nil {
		return err
	}

	return createUsers(deps, cmd, profile, organizationName, usersFile)
}

func CreateUser(deps Dependencies, cmd *cobra.Command, profile configuration_models.ProfileV2, organizationName string) error {
	username, err := cmd.Flags().GetString("username")
	if err != nil {
		return fmt.Errorf("%s username: %w", constants.ErrorRetrievingField, err)
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return fmt.Errorf("%s password: %w", constants.ErrorRetrievingField, err)
	}
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return fmt.Errorf("%s email: %w", constants.ErrorRetrievingField, err)
	}
	policies, err := cmd.Flags().GetStringArray("policy")
	if err != nil {
		return fmt.Errorf("%s policy: %w", constants.ErrorRetrievingField, err)
	}

	var emailPtr *string
	if email != "" {
		emailPtr = &email
	}

	return createUsers(deps, cmd, profile, organizationName, bulkCreateUsersFile{
		Users: []bulkCreateUserFileItem{
			{
				Username:         username,
				Password:         password,
				Email:            emailPtr,
				AttachedPolicies: policies,
			},
		},
	})
}

func createUsers(
	deps Dependencies,
	cmd *cobra.Command,
	profile configuration_models.ProfileV2,
	organizationName string,
	usersFile bulkCreateUsersFile,
) error {
	request, err := createBulkCreateUsersRequest(deps, profile, organizationName, usersFile)
	if err != nil {
		return err
	}
	response, err := deps.UserAPI.BulkCreateIAMUsers(
		profile.Endpoints,
		profile.APIKey,
		profile.OrganizationID,
		request,
	)
	if err != nil {
		return fmt.Errorf("failed to bulk create users: %w", err)
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return fmt.Errorf("failed to get output flag: %w", err)
	}
	quiet, err := cmd.Flags().GetBool("quiet")
	if err != nil {
		return fmt.Errorf("failed to get quiet flag: %w", err)
	}
	if profile.Output != "" &&
		!cmd.Flags().Changed("output") &&
		!cmd.Flags().Changed("quiet") {
		output = string(profile.Output)
	}
	if output != "human" && !quiet {
		utils.PrintFormattedData(cmd.OutOrStdout(), response, output)
		return nil
	}

	if quiet {
		for _, item := range response.Data {
			email := ""
			if item.Email != nil {
				email = *item.Email
			}
			utils.PrintQuiet(cmd.OutOrStdout(), item.Username, item.ID, email, fmt.Sprintf("%t", item.Created), item.Status)
		}
		return nil
	}

	noHeaders, err := cmd.Flags().GetBool("no-headers")
	if err != nil {
		return fmt.Errorf("failed to get no-headers flag: %w", err)
	}

	tableColumns := []table.Column[api.BulkCreateIAMUserResponseItem]{
		{Title: "Username"},
		{Title: "ID"},
		{Title: "Email"},
		{Title: "Created"},
		{Title: "Status"},
	}

	rowMapper := func(item api.BulkCreateIAMUserResponseItem) []string {
		email := ""
		if item.Email != nil {
			email = *item.Email
		}

		return []string{item.Username, item.ID, email, fmt.Sprintf("%t", item.Created), item.Status}
	}

	return printer.CreateTable(
		cmd,
		response.Data,
		table.WithColumns(tableColumns),
		table.WithRowMapper(rowMapper),
		table.WithShowHeader[api.BulkCreateIAMUserResponseItem](!noHeaders),
		table.WithSuffix[api.BulkCreateIAMUserResponseItem]("\n"),
	)
}

func printImportSample(cmd *cobra.Command, sampleFormat string) error {
	switch sampleFormat {
	case "json":
		cmd.Println(`{
  "users": [
    {
      "username": "user-1",
      "password": "password-1",
      "first_name": "first-name-1",
      "last_name": "last-name-1",
      "email": "email-1@example.com",
      "attached_policies": []
    },
    {
      "username": "user-2",
      "password": "password-2",
      "first_name": "first-name-2",
      "last_name": "last-name-2"
    }
  ]
}`)
		return nil
	case "csv":
		cmd.Print(`username,password,first_name,last_name,email,attached_policies
user-1,password-1,first-name-1,last-name-1,email-1@example.com,
user-2,password-2,first-name-2,last-name-2,,
`)
		return nil
	default:
		return fmt.Errorf("invalid sample format %q: expected json or csv", sampleFormat)
	}
}

func parseBulkCreateUsersFile(filePath string, data []byte) (bulkCreateUsersFile, error) {
	if strings.EqualFold(filepath.Ext(filePath), ".csv") {
		usersFile, err := parseBulkCreateUsersCSV(data)
		if err != nil {
			return bulkCreateUsersFile{}, fmt.Errorf("error parsing CSV from file %q: %w", filePath, err)
		}
		return usersFile, nil
	}

	var usersFile bulkCreateUsersFile
	if err := json.Unmarshal(data, &usersFile); err != nil {
		return bulkCreateUsersFile{}, fmt.Errorf("error parsing JSON from file %q: %w", filePath, err)
	}
	return usersFile, nil
}

func parseBulkCreateUsersCSV(data []byte) (bulkCreateUsersFile, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return bulkCreateUsersFile{}, err
	}
	if len(records) == 0 {
		return bulkCreateUsersFile{}, fmt.Errorf("empty CSV")
	}

	header := make(map[string]int, len(records[0]))
	for index, column := range records[0] {
		header[strings.TrimSpace(column)] = index
	}
	if _, ok := header["username"]; !ok {
		return bulkCreateUsersFile{}, fmt.Errorf("missing username column")
	}
	if _, ok := header["password"]; !ok {
		return bulkCreateUsersFile{}, fmt.Errorf("missing password column")
	}

	usersFile := bulkCreateUsersFile{
		Users: make([]bulkCreateUserFileItem, 0, len(records)-1),
	}
	for _, record := range records[1:] {
		usersFile.Users = append(usersFile.Users, bulkCreateUserFileItem{
			Username:         csvValue(record, header, "username"),
			Password:         csvValue(record, header, "password"),
			FirstName:        optionalCSVValue(record, header, "first_name"),
			LastName:         optionalCSVValue(record, header, "last_name"),
			Email:            optionalCSVValue(record, header, "email"),
			AttachedPolicies: csvListValue(record, header, "attached_policies"),
		})
	}

	return usersFile, nil
}

func csvValue(record []string, header map[string]int, column string) string {
	index, ok := header[column]
	if !ok || index >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[index])
}

func optionalCSVValue(record []string, header map[string]int, column string) *string {
	value := csvValue(record, header, column)
	if value == "" {
		return nil
	}
	return &value
}

func csvListValue(record []string, header map[string]int, column string) []string {
	value := csvValue(record, header, column)
	if value == "" {
		return nil
	}

	rawItems := strings.Split(value, ";")
	items := make([]string, 0, len(rawItems))
	for _, rawItem := range rawItems {
		item := strings.TrimSpace(rawItem)
		if item != "" {
			items = append(items, item)
		}
	}
	return items
}

func createBulkCreateUsersRequest(
	deps Dependencies,
	profile configuration_models.ProfileV2,
	organizationName string,
	usersFile bulkCreateUsersFile,
) (*api.BulkCreateIAMUsersRequestBody, error) {
	if len(usersFile.Users) == 0 {
		return nil, fmt.Errorf("no users provided")
	}

	request := &api.BulkCreateIAMUsersRequestBody{
		Users: make([]api.BulkCreateIAMUserRequestBody, 0, len(usersFile.Users)),
	}

	for _, user := range usersFile.Users {
		if user.Username == "" {
			return nil, fmt.Errorf("username is required")
		}
		if !usernamePattern.MatchString(user.Username) {
			return nil, fmt.Errorf("invalid username %q: use only letters, numbers, hyphen or underscore, and start and end with a letter or number", user.Username)
		}
		if user.Password == "" {
			return nil, fmt.Errorf("password is required for user %q", user.Username)
		}
		for _, policyID := range user.AttachedPolicies {
			if _, err := uuid.Parse(policyID); err != nil {
				return nil, fmt.Errorf("invalid attached policy %q for user %q: expected UUID", policyID, user.Username)
			}
		}

		challenge, err := deps.AuthAPI.GenerateChallenge(
			profile.Endpoints,
			nil,
			&user.Username,
			&organizationName,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to generate challenge for user %q: %w", user.Username, err)
		}

		authenticationPublicKey, err := utils.AuthenticationPublicKeyFromPassword(
			user.Password,
			challenge.Salt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to generate authentication public key for user %q: %w", user.Username, err)
		}

		request.Users = append(request.Users, api.BulkCreateIAMUserRequestBody{
			Username:                user.Username,
			AuthenticationPublicKey: authenticationPublicKey,
			FirstName:               user.FirstName,
			LastName:                user.LastName,
			Email:                   user.Email,
			AttachedPolicies:        user.AttachedPolicies,
		})
	}

	return request, nil
}
