# Swarms & Infrastructure

This guide covers the core infrastructure primitives of a Cubbit DS3 deployment: swarms, nexuses, nodes, agents, and redundancy classes.

---

## Concepts

| Concept | Description |
|---------|-------------|
| **Swarm** | The top-level distributed storage cluster |
| **Nexus** | A logical grouping of nodes within a swarm, typically representing a datacenter or availability zone |
| **Node** | A physical or virtual machine participating in a Nexus |
| **Agent** | A disk-level storage process running on a node |
| **Redundancy Class (RC)** | A policy that defines how data is distributed and protected across nodes and Nexuses |

A typical deployment hierarchy looks like:

```
Swarm
└── Nexus A (datacenter-1)
│   ├── Node A1
│   │   ├── Agent (disk-1)
│   │   └── Agent (disk-2)
│   └── Node A2
│       └── Agent (disk-1)
└── Nexus B (datacenter-2)
    └── Node B1
        └── Agent (disk-1)
```

---

## Swarms

```bash
# List all swarms
cubbit swarm list

# Get details of a specific swarm
cubbit swarm describe --swarm-id <swarm-id>

# Create a new swarm
cubbit swarm create --name my-swarm
```

---

## Nexuses

```bash
# List nexuses in a swarm
cubbit nexus list --swarm-id <swarm-id>

# Create a new nexus
cubbit nexus create --swarm-id <swarm-id> --name <name> --location <location> --provider-id <id>
```

---

## Nodes & Agents

Nodes and agents are created together in bulk using a JSON batch file. Each node entry automatically provisions its associated agents.

```bash
# Batch create nodes (and their agents) in a nexus
cubbit node create --batch --swarm-id <swarm-id> --nexus-id <nexus-id> --file nodes.json
```

**Batch file format (`nodes.json`):**

```json
{
  "nodes": [
    {
      "name": "node-1",
      "label": "node-1",
      "private_ip": "10.0.0.1",
      "public_ip": "203.0.113.1",
      "config": {},
      "agents": [
        {
          "port": 9000,
          "features": {},
          "volume": {
            "disk": "/dev/sdb",
            "mount_point": "/mnt/data"
          }
        }
      ]
    },
    {
      "name": "node-2",
      "label": "node-2",
      "private_ip": "10.0.0.2",
      "public_ip": "203.0.113.2",
      "config": {},
      "agents": [
        {
          "port": 9000,
          "features": {},
          "volume": {
            "disk": "/dev/sdb",
            "mount_point": "/mnt/data"
          }
        }
      ]
    }
  ]
}
```

Once all nodes are defined, deploy the nexus to generate the installation files needed to set up each server:

```bash
cubbit nexus deploy --swarm-id <swarm-id> --nexus-id <nexus-id> --output-dir <dir>
```

The output directory will contain the configuration and installation artifacts to be applied on each node.

```bash
# List nodes in a nexus
cubbit node list --swarm-id <swarm-id> --nexus-id <nexus-id>

# List agents on a node
cubbit agent list --swarm-id <swarm-id> --nexus-id <nexus-id> --node-id <node-id>
```

---

## Redundancy Classes

A Redundancy Class (RC) defines the erasure coding parameters used to distribute and protect data. It works at two levels:

- **Outer (geo) layer**: distributes data across Nexuses (availability zones)
- **Inner (local) layer**: distributes data across nodes within a single Nexus

### Parameters

| Flag | Description |
|------|-------------|
| `--name` | Human-readable name for the RC |
| `--outer-n` | Number of data shards at the geo level |
| `--outer-k` | Number of parity shards at the geo level |
| `--inner-n` | Number of data shards at the local level |
| `--inner-k` | Number of parity shards at the local level |
| `--anti-affinity-group` | Grouping strategy for node placement (e.g. `rack`) |
| `--nexuses` | Comma-separated list of Nexus IDs to include (must equal `outer-n + outer-k`) |

### Creating a redundancy class

The example below creates a geo-distributed RC spanning 5 Nexuses (3 data + 2 parity), with local 2+1 redundancy within each Nexus:

```bash
cubbit rc create \
  --name "geo-rc" \
  --swarm-id <swarm-id> \
  --outer-n 3 \
  --outer-k 2 \
  --inner-n 2 \
  --inner-k 1 \
  --anti-affinity-group 4 \
  --nexuses <nexus-id-1> <nexus-id-2> <nexus-id-3> <nexus-id-4> <nexus-id-5>
```

> The number of `--nexuses` values must equal `outer-n + outer-k`.

### Inspecting a redundancy class

```bash
# Get full details
cubbit rc describe --swarm-id <swarm-id> <rc-id>

# Check health status
cubbit rc status --swarm-id <swarm-id> <rc-id>

# List agents on a redundancy-class
cubbit agent list --swarm-id <swarm-id> --redundancy-class-id <rc-id>
```

### Expanding a redundancy class

```bash
cubbit rc expand --swarm-id <swarm-id> <rc-id>
```

---

## Next Steps

- [Tenants](tenants-users-and-gateways.md) — provision tenants on top of your swarm
