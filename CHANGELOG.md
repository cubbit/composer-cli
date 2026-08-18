# Changelog

## 2.7.0

### Minor Changes

- f327cd9: Improve api-key commands to manage other users' keys
  Add CLI commands to create, list, describe, update, and delete IAM users
  Add CLI commands for password reset and user API key management
  Implement bulk salt generation
  Support operator authentication without email address

## 2.6.1

### Patch Changes

- 47069e0: Fix infra location describe cli command
- 90d7895: Fix small issue linked with disk selection in interactive swarm creation

## 2.6.0

### Minor Changes

- fcb5030: List and describe commands for tenant v5
- bd8a968: Tenant create command

## 2.5.0

### Minor Changes

- 17d69c1: Add `--endpoints` flag to auth commands with configuration v2.0

  ### Minor Changes

  - Lazy config init — move LoadConfig out of package init into PersistentPreRunE
  - Add `LoadTypeForAuthCommands` to allow auth commands (`login`, `signup`, `activate`) to work without a pre-existing config file
  - Wire `--endpoints` flag on `auth login`, `auth signup`, and `auth activate` commands
  - Implement `ParseAndValidateEndpointsV2Inits` with strict TOML validation
  - Implement `EndpointsV2Inits.ToEndpointsV2()` to resolve base URL into per-service endpoints
  - Add `ConfigurationV2VersionValue` constant (`"v2"`) and `CreateConfigV2` with endpoints support
  - Add `CreateProfile`, `GetEndpoints`, `GetActiveProfile` methods to `ConfigV2`

  ### Features

  - **`--endpoints` flag for auth commands**: The `auth login`, `auth signup`, and `auth activate` commands now accept an optional `--endpoints` flag pointing to a TOML file with custom API endpoint configuration.

    ```bash
    cubbit auth login --profile dev --endpoints ./endpoints.toml
    ```

    ```bash
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

## 2.4.0

### Minor Changes

- a661120: Gateway list cmd
- f72998e: Add describe gateway command
- 50d218a: Implement crud operations over domains
- d7ffadf: Create Gateway V5

### Patch Changes

- 73f7a0a: Align get redundancy class endpoint to v5.

## 2.3.0

### Minor Changes

- 8bf7b48: Add interactive option to create swarm

## 2.2.0

### Minor Changes

- 9df6003: Add new swarm describe command
- 9df6003: Add swarm list command

## 2.1.0

### Minor Changes

- a613c3e: Introduce a new command to create a swarm in connected clusters, specifying redundancy classes and allocated resources.
  The command initiates a background process, resulting in a fully operational swarm upon completion.

## 2.0.0

### Major Changes

- e495809: Remove bazel

### Minor Changes

- b84fbde: Rreplace --version flag with dedicated version command
- 4f691a4: Lazy config init — move LoadConfig out of package init into PersistentPreRunE
- 90e08ca: Wire ResolveURLsWithPriority into Login and persist endpoints via CreateProfileWithEndpoints
- 49bd9cb: Wire --endpoints flag on auth login and signup commands
- 3452af1: Implement ParseEndpointsFile with strict YAML validation
- 09f5c32: Update config_service.go to use URLs.BaseURL
- 5bcc0b7: Add error constants and ConfigVersion constant for configuration v2.0
- dbf631e: Add CreateProfileWithEndpoints, UpdateProfileEndpoints, ResolveURLsWithPriority and expand ConfigInterface

### Features

- **`--endpoints` flag for auth commands**: The `auth login` and `auth signup` commands now accept an optional `--endpoints` flag that specifies a path to a YAML file containing custom API endpoint configuration.

  ```bash
  cubbit auth login --profile dev --endpoints ./endpoints.yaml
  ```

- **Endpoints YAML format**: Define custom endpoints in a YAML file:

  ```yaml
  base: https://api.example.com
  iam: https://iam.example.com
  ```

- **Endpoint resolution priority**: The CLI resolves endpoints using the following priority chain:

  1. Endpoints specified via `--endpoints` flag
  2. Endpoints from the active profile configuration
  3. Default endpoints from the configuration

- **Configuration v2.0**: New configuration format with versioning and strict validation for configuration files.

- **New functions**:
  - `ParseEndpointsFile()` - Parse and validate endpoints YAML
  - `ResolveURLsWithPriority()` - Resolve endpoints with priority chain
  - `CreateProfileWithEndpoints()` - Create profile with endpoints
  - `UpdateProfileEndpoints()` - Update profile endpoints

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
