# Location Commands

The Cubbit CLI provides infrastructure location commands to manage clusters and virtual nodes across your DS3 Composer environment. This lets you organize and control your distributed storage infrastructure.

---

## Concepts

| Concept | Description |
|---------|-------------|
| **Location** | An infrastructure cluster that aggregates physical and virtual nodes |
| **Virtual Location** | A logical cluster without physical hardware, used for organizing virtual nodes |
| **Virtual Node** | A storage node within a virtual location, configured with specific storage backends |

Locations can be **physical clusters** (with actual hardware nodes) or **virtual clusters** (logical groupings for virtual nodes). Virtual locations provide flexibility for testing, development, and multi-cloud scenarios.

---

## Locations

```bash
# List all infrastructure locations
cubbit infrastructure location list

# List all infrastructure locations (alias)
cubbit infrastructure location ls
```

### Describing a Location

To view detailed information about locations, including physical and virtual nodes:

```bash
# Describe all aggregated locations
cubbit infrastructure location describe

# Describe a specific location by name
cubbit infrastructure location describe --cluster-name <cluster-name>

# Describe a specific location by ID
cubbit infrastructure location describe --cluster-id <cluster-id>
```

**Note**: The `--cluster-name` and `--cluster-id` flags are mutually exclusive. When used, the command displays detailed information including:
- Cluster configuration and status
- Physical nodes with hardware details (OS, CPU, RAM, network, disks)
- Virtual nodes with storage configuration

---

## Virtual Locations

### Creating a Virtual Location

To create a virtual location (virtual cluster):

```bash
cubbit infrastructure location create-virtual --name <name> [--description <description>]
```

**Aliases**: `cv`

**Flags**:

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--name` | string | **Yes** | Name of the virtual location |
| `--description` | string | No | Description of the virtual location |

**Example**:

```bash
# Create a virtual location for testing
cubbit infrastructure location create-virtual \
  --name test-environment \
  --description "Testing environment for development"
```

The virtual location will be created and ready to accept virtual nodes.

---

## Virtual Nodes

### Creating a Virtual Node

To create a virtual node within a virtual location:

```bash
cubbit infrastructure location create-virtual-node \
  --cluster-id <cluster-id> \
  --name <name> \
  --storage-type <type> \
  --configuration '<json-configuration>'
```

**Aliases**: `cvn`

**Flags**:

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--cluster-id` | string | **Yes** | ID of the virtual cluster |
| `--name` | string | **Yes** | Name of the virtual node |
| `--storage-type` | string | **Yes** | Storage type (e.g., `s3`) |
| `--configuration` | string | **Yes** | JSON configuration object |

**Example**:

```bash
# Create an S3 virtual node
cubbit infrastructure location create-virtual-node \
  --cluster-id abc-123-def-456 \
  --name my-s3-node \
  --storage-type s3 \
  --configuration '{
    "endpoint": "https://s3.example.com",
    "bucket": "my-storage-bucket",
    "region": "us-east-1",
    "access_key": "my_s3_compatible_access_key",
    "secret_key": "my_s3_compatible_secret_key"
  }'
```

---

## Output Formats

All location commands support multiple output formats via the `--output` flag:

```bash
# Human-readable format (default)
cubbit infrastructure location list

# JSON format for scripting
cubbit infrastructure location list --output json

# YAML format for configuration
cubbit infrastructure location list --output yaml

# XML format
cubbit infrastructure location list --output xml
```

### Suppressing Headers

For cleaner output in scripts or CI/CD pipelines:

```bash
cubbit infrastructure location list --no-headers
```

### Quiet and Silent Modes

```bash
# Quiet mode - minimize non-essential output
cubbit infrastructure location describe --quiet

# Silent mode - redirect all output
cubbit infrastructure location list --silent
```

---

## Examples

### Complete Workflow: Creating a Virtual Environment

```bash
# 1. Create a virtual location
cubbit infrastructure location create-virtual \
  --name dev-environment \
  --description "Development testing environment"

# 2. List locations to get the cluster ID
cubbit infrastructure location list --output json

# 3. Create virtual nodes for different storage backends
cubbit infrastructure location create-virtual-node \
  --cluster-id <cluster-id-from-list> \
  --name s3-node-1 \
  --storage-type s3 \
  --configuration '{"endpoint": "https://s3.example.com", "bucket": "dev-bucket-1", "region": "eu-west-1", "access_key": "abcde", "secret_key": "abcde"}'

cubbit infrastructure location create-virtual-node \
  --cluster-id <cluster-id-from-list> \
  --name s3-node-2 \
  --storage-type s3 \
  --configuration '{"endpoint": "https://s3.example.com", "bucket": "dev-bucket-2", "region": "eu-west-1", "access_key": "abcde", "secret_key": "abcde"}'

# 4. Describe the location to see all nodes
cubbit infrastructure location describe --cluster-name dev-environment
```

### Inspecting Physical Infrastructure

```bash
# List all physical clusters
cubbit infrastructure location list

# Get detailed view of a production cluster
cubbit infrastructure location describe --cluster-name production

# View cluster details in JSON for automation
cubbit infrastructure location describe --cluster-id prod-cluster-123 --output json
```
