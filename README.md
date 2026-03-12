# Cubbit Composer CLI

<p align="center">
  <img src="assets/logo.png" alt="Cubbit Logo" width="200"/>
</p>

<p align="center">
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/Go-1.24+-blue.svg" alt="Go Version"/></a>
  <a href="https://github.com/cubbit/composer-cli/releases"><img src="https://img.shields.io/github/package-json/v/cubbit/composer-cli" alt="Version"/></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"/></a>
</p>

The official CLI for managing your **Cubbit DS3 Composer** infrastructure. Built with Go, it provides a comprehensive and automation-friendly interface to control your entire DS3 environment — from swarms and tenants to gateways and user accounts.

---

## Features

- **Swarm Management** — Create, configure, and manage distributed swarms
- **Tenant Operations** — Full tenant lifecycle management and configuration
- **Infrastructure Control** — Deploy Nexus, Redundancy Classes, nodes, and agents
- **User & Account Management** — Handle tenant accounts and user administration
- **Gateway Configuration** — Set up and manage S3-compatible gateway operations
- **IAM** — Fine-grained operator and policy management
- **Interactive Workflows** — Guided, step-by-step processes for complex tasks
- **Automation-friendly** — JSON/YAML/CSV output, quiet and silent modes

---

## Installation

### Pre-built Binaries (Recommended)

Download the latest release for your platform from the [releases page](https://github.com/cubbit/composer-cli/releases), then:

```bash
# Linux / macOS
chmod +x cubbit
sudo mv cubbit /usr/local/bin/

# Verify
cubbit --version
```

### Go Install

```bash
go install github.com/cubbit/composer-cli@latest
```

Make sure `$GOPATH/bin` is in your system `PATH`.

For more installation options (build from source, Bazel), see the [Installation Guide](docs/installation.md).

---

## Quick Start

```bash
# 1. Initialize configuration
cubbit config init

# 2. Log in
cubbit auth login --profile <profile_name>

# 3. Verify setup
cubbit config view

# 4. Explore available commands
cubbit --help
cubbit docs tree
```

For a full walkthrough, see the [Getting Started Guide](docs/getting-started.md).

---

## Documentation

| Guide | Description |
|-------|-------------|
| [Installation](docs/installation.md) | All installation options: binaries, Go install, source, Bazel |
| [Getting Started](docs/getting-started.md) | First login, config setup, and your first commands |
| [Configuration](docs/configuration.md) | Profiles, endpoints, output formats, config file reference |
| [Swarms & Infrastructure](docs/swarms-and-infrastructure.md) | Swarms, Nexus, nodes, agents, and redundancy classes |
| [Tenants, Users & Gateways](docs/tenants-users-and-gateways.md) | Tenant lifecycle, projects, user management, and S3 gateway setup |
| [IAM](docs/iam.md) | Operators, policies, and access control |

---

## Platform Support

| Platform | Supported |
|----------|-----------|
| Linux (amd64, arm64, armv6, armv7, 386) | ✅ |
| macOS (amd64, arm64) | ✅ |
| Windows (amd64, 386) | ✅ |

---

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, code style guidelines, and the pull request process.

When reporting a bug, please include the CLI version (`cubbit --version`), your OS, the full error output, and steps to reproduce. [Open an issue →](https://github.com/cubbit/composer-cli/issues)

---

## License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.
