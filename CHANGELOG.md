# Changelog

## 2.0.0

### Major Changes

- bb2411f: Rename the release binary and archives from `cubbit` to `ds3-composer`, and align the build and documentation references with the new name. This is a breaking change for download scripts that depend on the previous `cubbit_*` asset names.
- e495809: Remove bazel

### Minor Changes

- 8c04b55: Add `gateway update` plus `gateway connection`, `gateway routing-rules` (alias `rr`) and `gateway storage-class` commands backed by the Gateway V5 API.
- b84fbde: Replace `--version` flag with a dedicated version command
- a613c3e: Introduce a new command to create a swarm in connected clusters, specifying redundancy classes and allocated resources. The command initiates a background process, resulting in a fully operational swarm upon completion.
- 9df6003: Add swarm list and swarm describe commands
- 8bf7b48: Add interactive option to create swarm
- d7ffadf: Create Gateway V5
- 50d218a: Implement crud operations over domains
- f72998e: Add describe gateway command
- a661120: Gateway list cmd
- 17d69c1: Add `--endpoints` flag to auth commands with configuration v2 — lazy config init, `LoadTypeForAuthCommands` so `login`, `signup` and `activate` work without a pre-existing config file, `ParseAndValidateEndpointsV2Inits` with strict TOML validation, `EndpointsV2Inits.ToEndpointsV2()` base-URL resolution, and `CreateConfigV2` plus `CreateProfile`, `GetEndpoints` and `GetActiveProfile` methods on `ConfigV2`
- bd8a968: Tenant create command
- fcb7530: List and describe commands for tenant v5
- f327cd9: Improve api-key commands to manage other users' keys; add commands to create, list, describe, update and delete IAM users; add password reset and user API key management; implement bulk salt generation; support operator authentication without email address
- 04a9a4f: Extend printer with support for different output modes
- f3891f1: Add `verify-connection` command to trigger live DNS and TLS verification for a tenant connection; enhance `list-connections` to live-verify each connection before refreshing the list; add `--skip-verify` flag to return cached data
- c66abae: Two-factor authentication can now be managed from the CLI with `cubbit iam user tfa enable` and `cubbit iam user tfa disable`. Enabling takes the account password and either generates a setup key and a set of one-time recovery codes or uses the one supplied with `--setup-key`; disabling takes the account password together with a valid verification or recovery code.
- c66abae: The `iam api-key describe` and `iam api-key edit` commands now show a redacted form of the key instead of hiding it, so a key can be recognized from its detail view. The full secret is still shown only once, when the key is created.
- c66abae: `cubbit auth login` accepts a new `--tfa-code` flag, so operators with two-factor authentication enabled can sign in non-interactively instead of being prompted for a verification code. This allows scripted and automated logins to work with two-factor authentication enabled.
- 1948266: Add the infrastructure location update and delete commands and expose location metadata in the location list, describe and update commands.

### Patch Changes

- 73f7a0a: Align get redundancy class endpoint to v5.
- 47069e0: Fix infra location describe cli command
- 90d7895: Fix small issue linked with disk selection in interactive swarm creation

### Features

- **`--endpoints` flag for auth commands**: The `auth login`, `auth signup` and `auth activate` commands accept an optional `--endpoints` flag pointing to a TOML file with custom API endpoint configuration.

  ```bash
  cubbit auth login --profile dev --endpoints ./endpoints.toml
  cubbit auth activate --token <token> --endpoints ./endpoints.toml
  ```

- **Endpoints TOML format**: Define custom endpoints in a TOML file with either a `base` URL or individual per-service URLs:

  ```toml
  base = "http://localhost"
  iam = "http://localhost:8181"
  dash = "http://localhost:3001"
  ch = "http://localhost:8480"
  ```

  When `base` is provided, per-service URLs are resolved as `{base}{path}` using known service paths. When `base` is omitted, all three (`iam`, `dash`, `ch`) must be explicitly set.

- **Endpoint resolution per auth command**: Endpoints specified via `--endpoints` are parsed during `configuration_handler.Load()` for auth commands via `LoadTypeForAuthCommands`, which resolves endpoints with the following priority:

  1. Endpoints from `--endpoints` flag
  2. Endpoints from the active profile configuration (if a config file exists)
  3. Default endpoints from constants (`constants.BaseAPIURL + service URI`)

- **Configuration v2.0**: New TOML-based configuration format with `version = "v2"` and strict validation. Each profile stores its own endpoints:

  ```toml
  version = "v2"

  [active]
  profile = "composer"

  [profile.composer]
  output = "human"
  api_key = "<key>"
  organization_id = "<org-id>"

  [profile.composer.endpoints]
  iam = "https://iam.example.com"
  dash = "https://dash.example.com"
  ch = "https://ch.example.com"
  ```

- **New types**:

  - `EndpointsV2Inits` — struct for parsing `--endpoints` TOML files (with `base`, `iam`, `dash`, `ch` fields)
  - `ParseAndValidateEndpointsV2Inits(path)` — parse and validate an endpoints TOML file
  - `EndpointsV2Inits.ToEndpointsV2()` — resolve base URL into per-service endpoint URLs
  - `EndpointsV2` — per-profile endpoints stored in config (`iam`, `dash`, `ch`)
  - `ConfigV2` — full configuration model with `CreateProfile`, `GetEndpoints`, etc.

## 1.6.1

### Patch Changes

- ad6b30d: Align infrastructure location commands output

## 1.6.0

### Minor Changes

- 1f07b3c: Add create virtual clusters and nodes command

### Patch Changes

- c2eb60: Fix import and bazel build

## 1.5.0

### Minor Changes

- 33d55cf: New printer with custom table, tree, text and composable prints.

  Infrastrucutre location describe command.

## 1.4.1

### Patch Changes

- a5a6c71: Remove command related to old logic, unify cmd structure

## 1.4.0

### Minor Changes

- 2ecefbf: Authentication activate e2e tests

### Patch Changes

- 79cb109:

## 1.3.1

### Patch Changes

- 58d999b: List location infrastructure command

## 1.3.0

### Minor Changes

- 429ae80: - Implemented new functionality to allow users to log in using their username and password
  - Added the ability to enter an existing API key for authentication purposes
- ea77b57: Implement a new command to connect a cluster to an organization

### Patch Changes

- 219d7e5: Introduce E2E test for signup command

## 1.2.0

### Minor Changes

- 56a1844: Activate operator command
- 6606b27: Add root operator signup command

## 1.1.0

### Minor Changes

- 8bdc75e: Introduce command structure tests

## 1.0.1

### Patch Changes

- 8c6a53c: Make Enter prompt optional in CLI browser login flow
