# E2E Tests for Composer CLI

## General
This directory contains end-to-end (E2E) tests for the Composer CLI. These tests validate the integration and functionality of the CLI against a running Cubbit infrastructure, ensuring real-world scenarios are covered.

## Requirements
- **Cubbit Infrastructure**

    Make sure that following services are up.
    - IAM
    - Composer-Hub

- **Composer CLI**: Built and available in your PATH or referenced directly from root command.
- **Test Environment**: Ensure all required environmental and configuration variables including development credentials are set accordignly to the running infrastructure.

## Command to Run E2E Tests
Directory for the command should be the `.../composer-cli/`

`bazel test --test_output=all --test_env=HOME --cache_test_results=no --test_tag_filters=e2e ...`
