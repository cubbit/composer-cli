package tests

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	setup "github.com/cubbit/composer-cli/tests/setup"
	. "github.com/cubbit/composer-cli/tests/utils"
	"github.com/segmentio/kafka-go"
)

func Test_E2E_Auth_Activate_Success(t *testing.T) {
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

	sub := NewTestSubscriber(
		NewDefaultTestSubscriberConfig("email.fct.send_requested.v1"),
	)
	defer sub.Close()
	sub.Start()

	email := Faker.Email()
	username := Faker.Username()
	organization := "org-" + Faker.NameUnique()
	password := Faker.Password()
	firstName := Faker.Name()
	lastName := Faker.Name()
	settings := `{"settingKey":"settingValue"}`

	stdout, stderr, err := cli.Run(
		"auth", "signup",
		"--email", email,
		"--username", username,
		"--organization", organization,
		"--password", password,
		"--first-name", firstName,
		"--last-name", lastName,
		"--settings", settings,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s\n", err, stdout, stderr)
	}

	stdExpectedOutput := "Sign up completed successfully. Please check your email to verify your account.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}

	token, err := sub.WaitAndEvaluate(30*time.Second, func(msg kafka.Message) (any, error) {
		key := CreateEmailSendRequestedMessageKey(firstName, lastName, email)

		if string(msg.Key) != key {
			return "", ErrorNotTheMessageIWasLookingFor
		}

		var body map[string]any
		err := json.Unmarshal(msg.Value, &body)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal message value: %w", err)
		}

		object, ok := body["object"].(map[string]any)
		if !ok {
			return "", fmt.Errorf("object is not a map[string]any")
		}

		templateVars, ok := object["template_vars"].(map[string]any)
		if !ok {
			return "", fmt.Errorf("template_vars is not a map[string]string")
		}

		command, ok := templateVars["safe_activation_command"].(string)
		if !ok {
			return "", fmt.Errorf("safe_activation_command is not a string")
		}

		if command == "" {
			return "", fmt.Errorf("safe_activation_command is empty")
		}

		parts := strings.Split(command, "--token ")
		if len(parts) != 2 {
			return "", fmt.Errorf("unexpected command format: %s", command)
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			return "", fmt.Errorf("token is empty")
		}

		return token, nil
	})
	if err != nil {
		t.Fatalf("failed during kafka message fetch: %v", err)
	}

	tokenStr, ok := token.(string)
	if !ok {
		t.Fatalf("token is not a string")
	}

	stdout, stderr, err = cli.Run(
		"auth", "activate",
		"--token", tokenStr,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	stdExpectedOutput = "Activation completed successfully. You can now log in.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}

	if stderr != "" {
		t.Fatalf("stderr is not empty: %s", stderr)
	}
}

func Test_E2E_Auth_Activate_FailWithReusedToken(t *testing.T) {
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

	sub := NewTestSubscriber(
		NewDefaultTestSubscriberConfig("email.fct.send_requested.v1"),
	)
	defer sub.Close()
	sub.Start()

	email := Faker.Email()
	username := Faker.Username()
	organization := "org-" + Faker.NameUnique()
	password := Faker.Password()
	firstName := Faker.Name()
	lastName := Faker.Name()
	settings := `{"settingKey":"settingValue"}`

	stdout, stderr, err := cli.Run(
		"auth", "signup",
		"--email", email,
		"--username", username,
		"--organization", organization,
		"--password", password,
		"--first-name", firstName,
		"--last-name", lastName,
		"--settings", settings,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s\n", err, stdout, stderr)
	}

	stdExpectedOutput := "Sign up completed successfully. Please check your email to verify your account.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}

	token, err := sub.WaitAndEvaluate(30*time.Second, func(msg kafka.Message) (any, error) {
		key := CreateEmailSendRequestedMessageKey(firstName, lastName, email)

		if string(msg.Key) != key {
			return "", ErrorNotTheMessageIWasLookingFor
		}

		var body map[string]any
		err := json.Unmarshal(msg.Value, &body)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal message value: %w", err)
		}

		object, ok := body["object"].(map[string]any)
		if !ok {
			return "", fmt.Errorf("object is not a map[string]any")
		}

		templateVars, ok := object["template_vars"].(map[string]any)
		if !ok {
			return "", fmt.Errorf("template_vars is not a map[string]string")
		}

		command, ok := templateVars["safe_activation_command"].(string)
		if !ok {
			return "", fmt.Errorf("safe_activation_command is not a string")
		}

		if command == "" {
			return "", fmt.Errorf("safe_activation_command is empty")
		}

		parts := strings.Split(command, "--token ")
		if len(parts) != 2 {
			return "", fmt.Errorf("unexpected command format: %s", command)
		}

		token := strings.TrimSpace(parts[1])
		if token == "" {
			return "", fmt.Errorf("token is empty")
		}

		return token, nil
	})
	if err != nil {
		t.Fatalf("failed during kafka message fetch: %v", err)
	}

	tokenStr, ok := token.(string)
	if !ok {
		t.Fatalf("token is not a string")
	}

	stdout, stderr, err = cli.Run(
		"auth", "activate",
		"--token", tokenStr,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	stdExpectedOutput = "Activation completed successfully. You can now log in.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}

	if stderr != "" {
		t.Fatalf("stderr is not empty: %s", stderr)
	}

	stdout, stderr, err = cli.Run(
		"auth", "activate",
		"--token", tokenStr,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	if stdout != "" {
		t.Fatalf("stdout is not empty: %s", stdout)
	}

	stdErrExpectedOutput := `ERR failed during activation request: failed to perform activation request: code status expected 204, but received 401 instead
INF unauthorized` + "\n"
	if stderr != stdErrExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdErrExpectedOutput,
			"Got:", stderr)
	}
}

func Test_E2E_Auth_Activate_FailMalformedToken(t *testing.T) {
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

	malformedToken := "malformed-token"
	stdout, stderr, err := cli.Run(
		"auth", "activate",
		"--token", malformedToken,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	if stdout != "" {
		t.Fatalf("stdout is not empty: %s", stdout)
	}

	stdErrExpectedOutput := `ERR failed during activation request: failed to perform activation request: code status expected 204, but received 401 instead
INF unauthorized` + "\n"
	if stderr != stdErrExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdErrExpectedOutput,
			"Got:", stderr)
	}
}
