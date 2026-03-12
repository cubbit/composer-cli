# Configuration

The Cubbit CLI uses a profile-based configuration system that allows you to manage multiple environments (production, staging, local) from a single installation.

---

## Config File Location

The config file is stored at:

```
$XDG_CONFIG_HOME/cubbit/config.yaml
```

Which resolves to `~/.config/cubbit/config.yaml` on most Linux and macOS systems.

---

## Initializing the Config

```bash
cubbit config init
```

This creates the file with defaults if it doesn't exist yet. You can also edit it manually at any time.

---

## Profiles

Profiles are named sets of configuration values. You can have as many as you need — one per environment, per customer, or per project.

### Example config with multiple profiles

```ini
[default]
endpoint = "https://api.eu00wi.cubbit.services"
output = "json"

[profile.prod]
inherits = "default"
type = "composer"
api_key = "<prod_api_key>"

[profile.staging]
inherits = "default"
type = "composer"
endpoint = "https://staging.api.cubbit.services"
api_key = "<staging_api_key>"

[profile.dev]
inherits = "default"
type = "composer"
endpoint = "localhost"
api_key = "<dev_api_key>"
```

### Configuration options

| Option | Description | Example |
|--------|-------------|---------|
| `endpoint` | API endpoint for your DS3 Composer instance | `https://api.eu00wi.cubbit.services` |
| `output` | Default output format for all commands | `json`, `yaml`, `xml`, `csv` |
| `type` | Profile type | `composer` |
| `api_key` | Authentication API key (generated on login) | — |
| `inherits` | Inherit settings from another profile | `default` |

The `inherits` field lets you define shared settings once (like `endpoint` and `output`) in a `[default]` block and reuse them across profiles, overriding only what differs.

---

## Managing Profiles

```bash
# List all configured profiles
cubbit config profiles

# Switch the active profile
cubbit config switch-profile staging

# View the active configuration
cubbit config view
```

---

## Output Formats

You can set a default output format in your profile, or override it per command with the `--output` flag:

```bash
cubbit tenant list --output json
cubbit swarm list --output yaml
cubbit node list --output csv
```

Supported formats: `json`, `yaml`, `xml`, `csv`. When no format is set, output is human-readable.

---

## Authentication

API keys are generated through a browser-based OAuth flow and stored in your profile automatically. To refresh or create a key for a profile:

```bash
cubbit auth login --profile <profile_name>
```

See the [Getting Started guide](getting-started.md) for the full login walkthrough.

---

## Next Steps

- [Getting Started](getting-started.md)
