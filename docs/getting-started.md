# Getting Started

This guide walks you through your first login, initial configuration, and running your first commands against a Cubbit DS3 Composer instance.

---

## Prerequisites

Before you begin, make sure you have:

- The `cubbit` CLI installed — see the [Installation Guide](installation.md)
- Access to a Cubbit DS3 Composer instance (endpoint URL + credentials)

---

## Step 1: Initialize Configuration

The CLI stores its configuration at `$XDG_CONFIG/cubbit/config.yaml` (typically `~/.config/cubbit/config.yaml` on Linux/macOS).

Run the init command to create the file with sensible defaults:

```bash
cubbit config init
```

---

## Step 2: Log In

Authenticate with your DS3 Composer instance. Specify a profile name — this lets you manage multiple environments side by side.

```bash
cubbit auth login --profile prod
```

The login flow:

1. The CLI opens your browser automatically
2. A unique 8-digit verification code appears in your terminal
3. Enter the code in your browser to authorize
4. The CLI saves the generated API key to your profile

Once complete, your config will contain an entry like:

```ini
[profile.prod]
type = "composer"
endpoint = "https://api.eu00wi.cubbit.services"
api_key = "<generated_api_key>"
```

---

## Step 3: Verify Your Setup

Confirm the active profile is configured correctly:

```bash
cubbit config view
```

---

## Step 4: Explore Commands

```bash
# Top-level help
cubbit --help

# Full command tree
cubbit docs tree

# Help for a specific command
cubbit tenant --help
cubbit rc --help
```

---

## Your First Commands

```bash
# List tenants
cubbit tenant list

# List swarms
cubbit swarm list

# List IAM operators
cubbit iam user list --tenant-id <tenant-id>
```

---

## Typical Infrastructure Setup Flow

When setting up a new DS3 environment from scratch, the usual order is:

1. **Create a swarm** — the top-level distributed storage unit
2. **For each Nexus** — define the Nexus, register its nodes, deploy agents on each node, then move on to the next Nexus
3. **Create a redundancy class** — define how data is protected across the deployed nodes
4. **Create a tenant** — provision an isolated organizational unit on top of the swarm
5. **Connect the tenant to the redundancy class** — assign the RC to the tenant to define how its data is stored
6. **Create a gateway** — register an S3-compatible gateway for the tenant
7. **Configure DNS** — assign a domain to the tenant
8. **Install the gateway** — run the installation on the target server using the tenant and gateway IDs

Each step has its own guide in the [docs folder](.).

---

## Next Steps

- [Configuration](configuration.md) — profiles, output formats, multi-environment setup
- [Swarms & Infrastructure](swarms-and-infrastructure.md) — set up your first swarm
- [Tenants, Users & Gateways](tenants-users-and-gateways.md) — manage tenants and expose S3 endpoints
- [IAM](iam.md) — control operator permissions
