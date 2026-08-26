---
id: client-cli-swarm
title: "Swarm Commands"
sidebar_label: "Swarm"
slug: /client/cli/swarm
---

# Swarm Commands

The Cubbit CLI provides swarm management commands to create and manage distributed storage swarms across your DS3 Composer environment. Swarms are the foundation of your geo-distributed cloud storage infrastructure.

---

## Swarm Describe

Use `describe` to inspect a single swarm in detail.

```bash
cubbit swarm describe [SWARM_ID]
```

Aliases:

```bash
cubbit swarm info [SWARM_ID]
cubbit swarm show [SWARM_ID]
```

You can identify the swarm in exactly one of these ways:

- positional `SWARM_ID`
- `--swarm-id <swarm-id>`
- `--swarm-name <swarm-name>`

### Describe Flags

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--swarm-id` | string | No | Swarm ID, alternative to positional `SWARM_ID` |
| `--swarm-name` | string | No | Swarm name, alternative to `SWARM_ID` |
| `--output` | string | No | Output format: `human` (default), `json`, `yaml` |

### Describe Output

Human-readable output is organized in labeled sections:

- **Basic Info**: `id`, `name`, `description`, `organization_id`, `owner_id`
- **Storage**: `total_storage`, `used_storage`, `available_storage`
- **Metadata**: `created_at`, `creation_status`
- **Composition**: `nexus_count`, `redundancy_class_count`
- **Configuration**: key-value pairs from the swarm configuration
- **Status**: `evaluated_status`, `evaluated_status_last_updated_at`

When a nullable field is missing, the CLI prints `N/A`.

### Describe Examples

```bash
# Describe by positional swarm ID
cubbit swarm describe 8d54c31e-7d9f-4fcb-9d48-9d9d7d0e4c11

# Describe by explicit swarm ID flag
cubbit swarm describe --swarm-id 8d54c31e-7d9f-4fcb-9d48-9d9d7d0e4c11

# Describe by swarm name
cubbit swarm describe --swarm-name production-swarm

# Use an alias
cubbit swarm info 8d54c31e-7d9f-4fcb-9d48-9d9d7d0e4c11

# Programmatic output
cubbit swarm describe --swarm-name production-swarm --output json
cubbit swarm describe 8d54c31e-7d9f-4fcb-9d48-9d9d7d0e4c11 --output yaml
```

---

## Concepts

| Concept | Description |
|---------|-------------|
| **Swarm** | A geo-distributed cluster of storage nodes that provides high availability and data resilience |
| **Nexus** | A grouping of nodes from one or more clusters that contribute storage to the swarm |
| **Redundancy Class** | Configuration defining erasure coding parameters (local/geographic K/N) for data protection |
| **Node** | A physical or virtual machine contributing storage volumes to the swarm |
| **Volume** | A disk or storage unit on a node that stores data shards |

Swarms use advanced erasure coding to distribute data across multiple nodes and clusters, ensuring high availability even when individual nodes or entire clusters fail.

---

## Swarm List

Use `list` to view all swarms in your organization.

```bash
cubbit swarm list
```

Aliases:

```bash
cubbit swarm ls
```

### List Flags

| Flag | Type | Description |
|------|------|-------------|
| `--output` | string | Output format: `human` (default), `json`, `yaml` |
| `--quiet` | bool | Minimize stdout for CI/CD workflows |
| `--no-headers` | bool | Suppress table headers in human output |

### List Output (human-readable)

Human-readable output shows a table with the following columns:

| Column | Description |
|--------|-------------|
| **ID** | Swarm UUID |
| **Name** | Swarm name |
| **Total Storage** | Total storage capacity (human-readable) |
| **Used Storage** | Used storage (human-readable) |
| **Created At** | Creation timestamp |
| **Nexus Count** | Number of nexuses |
| **RC Count** | Number of redundancy classes |
| **Status** | Evaluated status (`online`, `offline`, `warning`, or `N/A`) |

### List Examples

```bash
# List all swarms (human-readable table)
cubbit swarm list

# List all swarms (JSON output)
cubbit swarm list --output json

# List all swarms (YAML output)
cubbit swarm list --output yaml

# Use the ls alias
cubbit swarm ls

# Suppress table headers for scripting
cubbit swarm list --no-headers

# Quiet mode (only error output)
cubbit swarm list --quiet
```

### Empty Results

When no swarms exist, the CLI prints:

```
No swarms found.
```

---

## Swarm Creation

The swarm creation process is **asynchronous** - when you submit a creation request, it returns immediately with a process ID. The actual swarm creation happens in the background.

### Creating a Swarm

```bash
cubbit swarm create \
  --name <swarm-name> \
  --nexus <cluster-id>:<node-id1>,<node-id2>,<node-id3> \
  --redundancy-class '<json-configuration>'
```

**Required Flags**:

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--name` | string | **Yes** | Name of the swarm (3-63 characters) |
| `--nexus` | string array | **Yes** | Nexus specification in format `cluster-id:node-id1,node-id2`. Can be specified multiple times |

**Optional Flags**:

| Flag | Type | Required | Description |
|------|------|----------|-------------|
| `--description` | string | No | Description of the swarm |
| `--redundancy-class` | string array | No | Redundancy class configuration in JSON format. Can be specified multiple times |
| `--redundancy-class-file` | string | No | Path to a JSON file containing an array of redundancy class configurations |

---

## Nexus Format

Each `--nexus` flag specifies which nodes from which clusters will participate in the swarm:

```
cluster-id:node-id1,node-id2,node-id3
```

**How it works**:

- The CLI automatically fetches cluster details from the backend
- All **volumes** (disks), in status ok, from each specified node are automatically included
- For **physical clusters**: specify physical node IDs
- For **virtual clusters**: specify virtual node IDs
- You can specify multiple `--nexus` flags to include nodes from different clusters

**Example Nexus Specifications**:

```bash
# Single node from a cluster
--nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba"

# Multiple nodes from the same cluster
--nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba,5dd6322d-e898-451e-8cbd-d7f6c7dd613b"

# Multiple clusters (use multiple --nexus flags)
--nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba,5dd6322d-e898-451e-8cbd-d7f6c7dd613b" \
--nexus "3002286e-7114-4357-8721-3e8f570b13ef:061e46e5-7026-4d28-bb18-4bbde944344f,cef1fad0-c4c6-4f02-bbe8-4201b7ae2ed0"
```

---

## Redundancy Class Configuration

Redundancy classes define how data is protected using erasure coding. Each redundancy class requires a JSON configuration:

```json
{
  "name": "rc-production",
  "outer_n": 1,
  "outer_k": 1,
  "inner_n": 4,
  "inner_k": 2,
  "cluster_ids": ["85e7e270-b827-4897-b506-91cfcf66f950", "12fc6e97-0120-4b9d-8f2a-46fcbafb3169"]
}
```

**Configuration Fields**:

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | **Yes** | Name of the redundancy class |
| `outer_n` | int | **Yes** | Number of data shards at the geographical level |
| `outer_k` | int | **Yes** | Number of parity shards at the geographical level. Customer will have `outer_n + outer_k` datacenters total and can safely lose up to `outer_k` datacenters before losing data |
| `inner_n` | int | **Yes** | Number of data shards at the machine/volume level |
| `inner_k` | int | **Yes** | Number of parity shards at the machine/volume level. Each datacenter will have `inner_n + inner_k` disks used and can safely lose up to `inner_k` disks before losing data |
| `cluster_ids` | string array | **Yes** | List of cluster IDs this redundancy class spans |
| `anti_affinity_group` | int | No | Anti-affinity group number represent how many volumes could be used from the same physical machine, multiple volumes sharing the same physical machine may face a common destiny if machine fall apart |

**Erasure Coding Explained**:

- **Geographical level (outer)**: Data is distributed across `outer_n + outer_k` datacenters. System tolerates failure of any `outer_k` datacenters
- **Machine/volume level (inner)**: Each datacenter uses `inner_n + inner_k` volumes. System tolerates failure of any `inner_k` volumes within each datacenter
- **Example**: `outer_n=2, outer_k=2` means 4 datacenters total, can lose any 2. `inner_n=2, inner_k=2` means 4 volumes per datacenter, can lose any 2

**Anti affinity group**:

The value of the recommended antiy affinity group is determined by the number of volumes one is willing to lose divided by the number of machines a user is willing to lose.
A physical redundancy of `inner_n=4, inner_k=4` means there are 8 volumes per data center. Typically, we anticipate losing up to `5` volumes before risking data loss. However, setting anti-affinity to `2` and experiencing the simultaneous failure of `3` physical machines could result in the loss of up to `6` volumes. This is due to the fact that multiple volumes from the same machine may have been utilized, creating interdependent resources that should theoretically be independent. This functionality is particularly valuable in real-world scenarios where powerful machines are significantly underutilized and could potentially be connected to multiple volumes.

---

## Examples

### Basic Swarm Creation

Create a simple swarm with one physical cluster:

```bash
cubbit swarm create \
  --name "production-swarm" \
  --description "Production storage swarm" \
  --nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba" \
  --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["bc7450ef-b858-4741-a044-871ffc39c343"]}'
```

### Multi-Node Swarm

Create a swarm with multiple nodes from the same cluster:

```bash
cubbit swarm create \
  --name "multi-node-swarm" \
  --nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba,a3284691-f191-404f-9b5e-c2130d1b73ba,4302d028-520d-490d-949e-a9079a29722c" \
  --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":2,"inner_k":1,"cluster_ids":["bc7450ef-b858-4741-a044-871ffc39c343"]}'
```

### Multi-Cluster Swarm

Create a swarm spanning multiple clusters for geo-distribution:

```bash
cubbit swarm create \
  --name "geo-distributed-swarm" \
  --nexus "bc7450ef-b858-4741-a044-871ffc39c343:bfb3e9d4-6a35-4c76-bf4e-07ac50f109ba,a3284691-f191-404f-9b5e-c2130d1b73ba" \
  --nexus "4043dd06-eb16-481e-a5bc-5ebc7d3de41e:ec3efde9-7c4e-4039-b6c6-739718baab8d,39c003ba-5772-4593-b269-6f9164484d69" \
  --redundancy-class '{"name":"rc-geo","outer_n":1,"outer_k":1,"inner_n":1,"inner_k":1,"cluster_ids":["cluster-eu","cluster-us"]}'
```

### Using Redundancy Class File

For complex configurations, use a JSON file:

```bash
# Create redundancy-classes.json
cat > redundancy-classes.json <<EOF
[
  {
    "name": "rc-standard",
    "outer_n": 1,
    "outer_k": 1,
    "inner_n": 4,
    "inner_k": 2,
    "cluster_ids": ["b251b9b8-b8c8-461d-8e25-23bfb07ad03f","59b9b831-6ed9-4620-9d2d-75d407c6256f"]
  },
  {
    "name": "rc-premium",
    "outer_n": 2,
    "outer_k": 1,
    "inner_n": 6,
    "inner_k": 3,
    "cluster_ids": ["35c47c86-e5e0-4bd0-be87-b9e48435570c","cc86bf24-2cbc-48df-b5d1-dfbc9c3c9d02","858b5e94-edba-48f1-bfab-0c866d247c09"]
  }
]
EOF

# Use the file
cubbit swarm create \
  --name "multi-rc-swarm" \
  --nexus "47deed01-c5aa-4a61-a155-73de8d284af8:bf9d6ca8-fa10-4eba-adb3-56e901cddc80,dbfa59c6-b21b-4600-84f3-2cc1e34fb407" \
  --nexus "d777f072-5eeb-4f35-960e-5826d818c2af:91f9a8b6-c7fa-41a3-ba1c-9df557e361d3,efee5105-6562-4fb8-9b67-0e38d0d77420" \
  --redundancy-class-file redundancy-classes.json
```

### Mixed: Inline and File Configuration

You can combine both approaches:

```bash
cubbit swarm create \
  --name "hybrid-swarm" \
  --nexus "1913d047-965c-44ff-b756-c345ef9206fa:65284758-3567-434d-85f6-421c46041dd3,ffd7f2c7-66aa-4268-8311-74e68303cbc4" \
  --redundancy-class '{"name":"rc-basic","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":1,"cluster_ids":["1913d047-965c-44ff-b756-c345ef9206fa"]}' \
  --redundancy-class-file additional-rcs.json
```

---

## Asynchronous Process

Swarm creation is **asynchronous**:

1. **Submit Request**: CLI sends creation request to backend
2. **Immediate Response**: Returns process ID immediately
3. **Background Processing**: Swarm creation happens asynchronously
4. **Monitor Progress**: Use interactive mode to track progress

```bash
# Non-interactive mode (returns immediately)
cubbit swarm create \
  --name "my-swarm" \
  --nexus "f4267663-ed58-49f5-af07-d51429a576f4:1ab945c7-af0c-446f-9fdd-731c571e152b,8385db64-0687-419a-823b-dc2bb41a8264" \
  --redundancy-class '{"name":"rc-1","outer_n":1,"outer_k":0,"inner_n":1,"inner_k":0,"cluster_ids":["f4267663-ed58-49f5-af07-d51429a576f4"]}'

# Output: Swarm creation started. Process ID: abc-123-def-456
```

### Race Condition Prevention

The CLI automatically checks for ongoing swarm creation processes:

- If a `swarm_creation` process is already **running**, the command fails with:

  ```
  Error: swarm creation already in progress.
  ```

- This prevents duplicate swarm creation requests
- Only one swarm creation process can run at a time per organization

---

## Complete Workflow Example

### Step 1: List Available Clusters

First, identify available clusters and nodes:

```bash
# List all infrastructure locations
cubbit infrastructure location list

# Get detailed cluster information
cubbit infrastructure location describe --cluster-name production
```

### Step 2: Plan Your Swarm

Determine:

- Which clusters to use
- Which nodes from each cluster
- Redundancy class parameters based on your availability requirements

### Step 3: Create the Swarm

```bash
cubbit swarm create \
  --name "production-swarm-2024" \
  --description "Production swarm for Q1 2024" \
  --nexus "2f98fe0b-dd03-43a6-bf05-c030171394e1:fe913ba4-88bc-4723-8591-21ddc6bfd372,7b6cb15b-c8c0-44af-a4a2-a23fb1153110" \
  --nexus "93523c74-cee6-4caf-8c9e-f68c9317c6d6:4f643534-e79f-4a29-84c3-6e0d61ea6b72,ed1e1d4a-c8a7-4016-9d59-7a77c44c5b67" \
  --redundancy-class '{
    "name": "rc-production",
    "outer_n": 1,
    "outer_k": 1,
    "inner_n": 1,
    "inner_k": 1,
    "cluster_ids": ["2f98fe0b-dd03-43a6-bf05-c030171394e1", "93523c74-cee6-4caf-8c9e-f68c9317c6d6"],
    "anti_affinity_group": 2
  }'
```

## Troubleshooting

### "Swarm creation already in progress"

**Cause**: Another swarm creation process is running

**Solution**:
Wait creation to complete.

### "Cluster not found"

**Cause**: Invalid cluster ID in nexus specification

**Solution**:

```bash
# List available clusters
cubbit infrastructure location list

# Verify cluster ID exists
cubbit infrastructure location describe --cluster-id <your-cluster-id>
```

### "Node not found in cluster"

**Cause**: Invalid node ID or node doesn't belong to specified cluster

**Solution**:

```bash
# Get detailed cluster information with all nodes
cubbit infrastructure location describe --cluster-id <cluster-id>
```

### "No usable volumes"

**Cause**: Specified node has no volumes in status ok

**Solution**:

- Ensure the node has configured and used storage volumes
- Check node configuration in the cluster details

---

## See Also

- `cubbit infrastructure location` - Manage infrastructure clusters
- `cubbit infrastructure location describe` - View cluster and node details
- Process monitoring commands (when available)
