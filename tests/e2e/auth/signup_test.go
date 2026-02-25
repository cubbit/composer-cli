package tests

import (
	"fmt"
	"testing"
	"time"

	setup "github.com/cubbit/composer-cli/tests/setup"
	. "github.com/cubbit/composer-cli/tests/utils"
)

func Test_E2E_Auth_SignUp_Success(t *testing.T) {
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

	email := Faker.Email()
	username := Faker.Username()
	organization := "org-" + Faker.NameUnique()
	firstName := Faker.Name()
	lastName := Faker.Name()
	settings := `{"settingKey":"settingValue"}`

	stdout, stderr, err := cli.Run(
		"auth", "signup",
		"--email", email,
		"--username", username,
		"--organization", organization,
		"--first-name", firstName,
		"--last-name", lastName,
		"--settings", settings,
	)
	if err != nil {
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	if stderr != "" {
		t.Fatalf("stderr is not empty: %s", stderr)
	}

	stdExpectedOutput := "Sign up completed successfully. Please check your email to verify your account.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}
}

func Test_E2E_Auth_SignUp_Success_WithPassword(t *testing.T) {
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

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
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	if stderr != "" {
		t.Fatalf("stderr is not empty: %s", stderr)
	}

	stdExpectedOutput := "Sign up completed successfully. Please check your email to verify your account.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}
}

func Test_E2E_Auth_SignUp_Fail_OrganizationNameConflict(t *testing.T) {
	t.Skip("flaky test - it's kafka dependent")
	cli, err := setup.GetTestRunner(setup.WithDevelopmentConfig(t))
	if err != nil {
		t.Fatalf("failed to get cli runner: %v", err)
	}

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
		t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
	}

	if stderr != "" {
		t.Fatalf("stderr is not empty: %s", stderr)
	}

	stdExpectedOutput := "Sign up completed successfully. Please check your email to verify your account.\n"
	if stdout != stdExpectedOutput {
		t.Fatalf("\n%-10s %s\n%-10s %s",
			"Expected:", stdExpectedOutput,
			"Got:", stdout)
	}

	err = KeepTesting(func() error {
		stdout, stderr, err = cli.Run(
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
			t.Fatalf("command failed: %v\nstdout: %s\nstderr: %s", err, stdout, stderr)
		}

		expectedOutput := `ERR failed during sign up request: failed to perform sign up request: code status expected 201, but received 409 instead
INF Conflict occurred
INF param organization.name` + "\n"

		if stderr != expectedOutput {
			return fmt.Errorf("\n%-10s %s\n%-10s %s",
				"Expected:", expectedOutput,
				"Got:", stderr,
			)
		}

		return nil
	}, 30*time.Second)

	if err != nil {
		t.Fatalf("failed to get expected error output after retries: %v", err)
	}
}
