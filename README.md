# Cubbit CLI

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org/)
[![Version](https://img.shields.io/github/package-json/v/cubbit/composer-cli)](https://github.com/cubbit/composer-cli/releases)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

The official Cubbit CLI (Command-Line Interface) for managing your DS3 composer infrastructure.

## Overview

The Cubbit CLI is a powerful command-line tool designed to provide comprehensive management capabilities for your DS3 composer environment. Built with Go and featuring an intuitive interface, it streamlines operations across your entire infrastructure stack.

### Key Capabilities

- **Swarm Management** - Create, configure, and manage distributed swarms
- **Tenant Operations** - Complete tenant lifecycle management and configuration
- **Infrastructure Control** - Deploy and manage Nexus, Redundancy Classes, nodes, and agents
- **User & Account Management** - Handle tenant accounts and user administration
- **Gateway Configuration** - Set up and manage gateway operations
- **Interactive Workflows** - User-friendly guided processes for complex tasks

The CLI leverages modern Go libraries including Cobra for command structure and Bubble Tea for terminal user interfaces, supporting both automated scripting and interactive modes.

## Installation

### Prerequisites

- Go 1.24.1 or higher
- Git (for source installation)

### Option 1: Pre-built Binaries (Recommended)

Download the latest release for your platform from our [releases page](https://github.com/cubbit/composer-cli/releases).

### Option 2: Go Install

Install directly using Go's package manager:

```bash
go install github.com/cubbit/composer-cli@latest
```

This will download, compile, and install the `composer-cli` binary to your `$GOPATH/bin` directory. Make sure `$GOPATH/bin` is in your system's PATH.

### Option 3: Build from Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/cubbit/composer-cli.git
   cd composer-cli
   ```

2. **Build the CLI:**
   ```bash
   # Build for current platform
   go build -o build/cubbit .

   # Cross-compile for specific platform (example for macOS)
   env GOOS=darwin GOARCH=amd64 go build -o build/cubbit .
   ```

3. **Install to your PATH:**
   ```bash
   # Linux/macOS
   sudo cp build/cubbit /usr/local/bin/

   # Windows: Copy build/cubbit.exe to a directory in your PATH
   ```

### Verify Installation

```bash
cubbit --version
```

## Quick Start

### Initial Setup

1. **Initialize configuration** (optional - creates config file at `$HOME/.config/cubbit/config.toml`):
   ```bash
   cubbit config init
   ```

2. **Authenticate with your composer account:**
   ```bash
   cubbit auth login --profile <profile_name>
   ```

   This will open your browser for secure authentication and automatically configure your API key.

   **Optional: Specify custom endpoints**

   You can provide custom API endpoints via a TOML file:

   ```bash
   cubbit auth login --profile <profile_name> --endpoints <path/to/endpoints.toml>
   ```

   Example `endpoints.toml`:

   ```toml
   base = "https://api.example.com"
   iam = "https://iam.example.com"
   dash = "https://dash.example.com"
   ch = "https://ch.example.com"
   ```

   All keys are optional. When `base` is provided, the CLI derives default URIs for `iam`, `dash`, and `ch` by appending path suffixes. Individual endpoints (`iam`, `dash`, `ch`) override the derived values.

3. **Verify your setup:**
   ```bash
   cubbit config view
   ```

### Basic Commands

```bash
# Display help and available commands
cubbit --help

# View command structure
cubbit docs tree

# Manage configuration profiles
cubbit config [command]

# Manage tenants
cubbit tenant [command]

# Manage swarms
cubbit swarm [command]
```

## Configuration

The CLI uses a profile-based configuration system stored in `$HOME/.config/cubbit/config.toml` (or `$XDG_CONFIG_HOME/cubbit/config.toml`). This allows you to manage multiple environments and accounts efficiently.

### Configuration Version

The CLI uses configuration version `v2` with strict validation for configuration files.

### Configuration Example

```toml
version = "v2"

[active]
profile = "composer"

[profile.composer]
output = "human"
api_key = "<your_api_key>"
organization_id = "<your_organization_id>"
updated_at = 2025-01-01T00:00:00Z

[profile.composer.endpoints]
iam = "https://iam.eu00wi.cubbit.services"
dash = "https://dashboard.cubbit.eu/api"
ch = "https://api.eu00wi.cubbit.services/composer-hub"

[profile.dev]
output = "human"
api_key = "<your_api_key>"
organization_id = "<your_organization_id>"
updated_at = 2025-01-01T00:00:00Z

[profile.dev.endpoints]
iam = "http://localhost:8181/iam"
dash = "http://localhost:8181/api"
ch = "http://localhost:8181/composer-hub"
```

### Endpoints Configuration

You can specify custom API endpoints using a TOML file with the `--endpoints` flag during `auth login`:

```toml
base = "https://api.example.com"
iam = "https://iam.example.com"
dash = "https://dash.example.com"
ch = "https://ch.example.com"
```

All keys are optional. When `base` is set without individual endpoints, the CLI derives defaults by appending path suffixes (`/iam`, `/api`, `/composer-hub`). Individual endpoints override derived values.

### Configuration Options

| Field | Profile Option | Description | Required |
|-------|---------------|-------------|----------|
| `output` | Output format | `human` (default) | Yes |
| `api_key` | Authentication API key | - | Yes |
| `organization_id` | Organization identifier | - | Yes |
| `updated_at` | Profile last-updated timestamp | ISO 8601 | Yes |
| `endpoints` | API endpoint set | `iam`, `dash`, `ch` | Yes |

### Profile Management

```bash
# Switch between profiles
cubbit config switch-profile <profile_name>

# List available profiles
cubbit config profiles

# View current configuration
cubbit config view

# Edit configuration in $EDITOR
cubbit config edit

# Validate configuration
cubbit config validate
```

## Authentication

Authentication is handled through API keys generated via secure browser-based login or inline password-based flow.

### Login Process

1. Run the login command:
   ```bash
   cubbit auth login --profile <profile_name>
   ```

2. Choose an authentication method when prompted:
   - **Browser-based**: Opens your browser for device-code authentication
   - **Inline**: Enter username, organization, and password directly
   - **API key**: Provide an existing API key

3. For browser-based login, copy the displayed one-time device code and enter it in the browser to complete authorization.

4. The CLI will automatically update your configuration with the new API key and organization ID.

### Non-interactive Login

For scripting and automation, you can provide credentials via flags:

```bash
# Using username/password
cubbit auth login --profile <profile_name> --username <user> --organization <org> --password <pass>

# Using an API key
cubbit auth login --profile <profile_name> --api-key <key>

# Using environment variables
API_KEY=<key> cubbit auth login --profile <profile_name>
```

**Warning**: Providing credentials via CLI flags may expose them in shell history and process lists. Prefer environment variables or interactive prompts.

### Logout

```bash
# Logout from specific profile
cubbit auth logout --profile <profile_name>

# Logout from all profiles
cubbit auth logout --all
```

## Interactive Mode

For complex operations and guided workflows, use interactive mode:

```bash
cubbit --interactive
```

Interactive mode provides step-by-step assistance for gateway configuration and installation processes.

## Advanced Usage

### Scripting and Automation

The CLI provides multiple modes for automation and scripting:

- **Quiet mode** - Suppresses non-essential output
- **Silent mode** - Redirects all output to /dev/null

```bash
# Quiet mode example
cubbit tenant list --quiet

# Silent mode example
cubbit swarm deploy --silent
```

## Features

- ✅ **Cross-platform Support** - Linux, macOS, Windows
- ✅ **Profile-based Configuration** - Manage multiple environments and accounts
- ✅ **Interactive Workflows** - Guided setup processes
- ✅ **Automation-friendly** - Scriptable with multiple output formats
- ✅ **Comprehensive Documentation** - Built-in help system
- ✅ **Secure Authentication** - Browser-based OAuth flow with device code

## Documentation

- **Cubbit Official Documentation**: [docs.cubbit.io](https://docs.cubbit.io)
- **Command Reference**: Run `cubbit docs` for detailed command documentation

## Support

### Getting Help

- **Issues**: [GitHub Issues](https://github.com/cubbit/composer-cli/issues)
- **Discussions**: [GitHub Discussions](https://github.com/cubbit/composer-cli/discussions)
- **Documentation**: [docs.cubbit.io](https://docs.cubbit.io)

### Reporting Bugs

Please include the following information when reporting issues:
- CLI version (`cubbit --version`)
- Operating system and version
- Complete error messages
- Steps to reproduce

## Contributing

We welcome contributions from the community! Please see our [Contributing Guide](CONTRIBUTING.md) for:

- Code style guidelines
- Development setup instructions
- Pull request process
- Issue reporting guidelines

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for complete details.

## Changelog

For a complete history of changes and new features, see [CHANGELOG.md](CHANGELOG.md).

---

**Need help?** Check our [documentation](https://docs.cubbit.io) or open an [issue](https://github.com/cubbit/composer-cli/issues).
