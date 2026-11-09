# Cubbit Composer CLI v2 - Product Description & Implementation Guide

## Part 1: Improved Product Description

### Overview

Cubbit Composer is a platform for deploying and managing DS3, a geo-distributed object storage solution. Users deploy a Kubernetes operator on their clusters, which Composer orchestrates to provide S3-compatible storage across multiple locations.

### Core Concepts

#### Organization & Users

- **Organization**: The top-level container for all resources. Each user belongs to exactly one organization.
- **Root User**: The user who creates the organization. Has full administrative access.
- **IAM User**: Additional users invited to the organization. Each IAM user can create their own API keys for CLI/API access. IAM users can be used as service accounts (similar to AWS model).

#### Infrastructure Layer

- **Location**: A unified concept representing where infrastructure runs. Two types:
  - **Cluster Location**: A Kubernetes cluster with the Cubbit operator installed
  - **Virtual Location**: A logical grouping for hybrid/cloud resources

- **Node**: A compute unit within a location:
  - **Cluster Node**: A Kubernetes node (physical/virtual server)
  - **Hybrid Node**: A third-party cloud resource (e.g., S3 bucket acting as storage backend)

- **Agent**: A storage agent pod running on a node, one per disk. Managed automatically by the operator.

#### Resource Layer

- **Swarm**: A geo-distributed storage cluster spanning multiple locations. Configuration:
  - **Name and description**: Identifies the swarm
  - **Locations**: Selected locations to include in the swarm
  - **Nodes**: Specific nodes (servers) to use per location (with disk counts)
  - **Redundancy Classes**: Policies defining data distribution
  - The operator automatically deploys DS3 agents on selected nodes

- **Gateway**: An edge deployment exposing storage services. Configuration:
  - **Name and description**: Identifies the gateway
  - **Type**: Manual or Automatic (determines DNS/SSL management)
  - **Connection type**: API Key or other authentication method
  - **Provider**: DNS provider for automatic mode (e.g., AWS Route 53)
  - **Locations**: Selected locations with specific nodes for gateway deployment
  - **Connected Swarms**: Swarms linked to this gateway with their redundancy classes
  - Exposes: S3 Gateway, Filesystem Gateway, Web Console
  - Contains: Storage Classes and Smart Data Placement

- **Redundancy Class (RC)**: Defines how data is distributed across locations within a swarm. Configuration parameters:
  - **Disks per location**: Number of disks to use at each location
  - **Location tolerance**: How many locations can fail without data loss
  - **Disk tolerance**: How many disks can fail per location without data loss

  The system calculates derived values:
  - **Redundancy level**: Overall data redundancy (e.g., 1x, 2x, 3x)
  - **Location level**: Location-level erasure coding (e.g., 3+1 Erasure)
  - **Ring level**: Disk-level erasure coding (e.g., 11+4)

- **Storage Class**: A label/tier defined at the gateway level (e.g., "Hot storage", "Cold storage", "Archive storage"). Configuration:
  - **Name**: Human-readable name used by tenants in Smart Data Placement rules
  - **Routing string**: Tag used to identify the storage class (e.g., "env=production", "tier=hot")
  - **Availability**: Can be made available to all connected tenants

- **Smart Data Placement**: Flow-based policies that direct traffic to appropriate redundancy classes. Each rule consists of:
  - **Input**: A Storage class for the gateway or an event (e.g. cronjob)
  - **Conditions**: Criteria that must be true (e.g., "If file name ends with .mp4")
  - **Action**: Where to route data (e.g., "Route to redundancy class with dynamic criteria")
  - **Fallback**: Alternative action if primary fails
  - **Schedule**: Optional time-based activation (e.g., "Sun, 10:00")
  - **Status**: On/Off toggle
  - Rules can be saved as templates for reuse

#### Tenant Layer

- **Tenant**: A logical isolation boundary for end-users. Each tenant has:
  - Isolated S3 buckets and filesystem folders
  - Connection to one or more gateways
  - Own DNS endpoint and console link
  - Configurable settings (name, slug, description, support link)
  - OAuth login options (Google, Microsoft)
  - Project limits and threshold settings

- **Console User**: End-users within a tenant (separate from IAM users). Attributes:
  - Email, name, storage usage, bandwidth usage
  - Verification status (email verified)
  - 2FA authentication status
  - Maximum project number limit
  - Status: Active, Banned, Removed
  - Can be frozen or have sessions disconnected

- **Project**: Organizational containers within a tenant. Attributes:
  - Owner, users count, uploaded files count
  - Storage usage with threshold management
  - Status: Operational, Limited storage, Gateway error, Frozen, Banned
  - Can be frozen, unfrozen, banned, unbanned, or deleted

- **Bucket**: S3-compatible storage containers within projects. Attributes:
  - Name, owner project ID
  - Storage usage, object counts (total, empty, non-empty)
  - Locking and versioning settings
  - State: None, Suspended

#### IAM Layer

- **IAM Policy**: Access control policies for the Composer platform. Can grant access to:
  - Organization-level resources (infrastructure, swarms, gateways)
  - Tenant-level resources (specific tenant management)

- **IAM Group**: Collections of IAM users with shared policies. Works at both:
  - Organization scope (platform access)
  - Tenant scope (tenant-specific access)

### Data Flow

```
Console User (S3 SDK/CLI)
       |
   Gateway (endpoint + console link)
       |
   Storage Class (labels data tier)
       |
   Smart Data Placement (evaluate conditions and route traffic)
       |
   Redundancy Class (data distribution policy)
       |
   Swarm (geo-distributed storage)
       |
   Locations/Nodes (physical storage)
```

All relationships are N:N and geo-distributed:

- Tenant -> multiple Gateways
- Gateway -> multiple Swarms -> multiple Redundancy Classes
- Redundancy Class -> multiple Locations (within a Swarm)

---

## Part 2: New CLI Command Structure

### Design Principles

1. **Immediate parent ID only** for list/create operations (IDs are globally unique)
2. **Direct --id access** for single-resource operations (describe/edit/remove)
3. **Clear resource grouping** by domain (infrastructure, resources, tenants, IAM)
4. **Consistent naming**: noun for resources, verb for actions
5. **Helpful descriptions** that explain purpose, not just "manage X"

### Command Tree

```
cubbit
├── auth
│   ├── login                    # Authenticate with Composer
│   ├── logout                   # Clear authentication
│   └── whoami                   # Show current user and organization
│
├── config
│   ├── init                     # Initialize CLI configuration
│   ├── view                     # View current configuration
│   ├── edit                     # Edit configuration file
│   ├── validate                 # Validate configuration
│   └── profile
│       ├── list                 # List all profiles
│       ├── create               # Create a new profile
│       ├── switch               # Switch active profile
│       └── remove               # Remove a profile
│
├── organization (alias: org)
│   ├── describe                 # Show organization details
│   └── edit                     # Update organization (e.g., name)
│
├── operator
│   ├── install                  # Install operator on cluster (creates cluster location)
│   ├── upgrade                  # Upgrade operator version (--location-name)
│   ├── uninstall                # Remove operator from cluster (--location-name)
│   └── status                   # Check operator health (--location-name)
│
├── infrastructure (alias: infra)
│   ├── status                   # Overview of all infrastructure health (includes resource usage)
│   ├── topology                 # Show visual ASCII topology (locations -> nodes -> disks)
│   ├── location
│   │   ├── create               # Create a virtual location (for hybrid nodes)
│   │   ├── list (ls)            # List all locations (cluster + virtual)
│   │   ├── describe             # Show location details (--id|--name)
│   │   ├── edit                 # Update location (--id|--name)
│   │   ├── remove (rm)          # Remove location (--id|--name)
│   │   └── status               # Check location health (--id|--name)
│   ├── node
│   │   ├── list (ls)            # List nodes (--location-id|--location-name)
│   │   ├── describe             # Show node details (--id|--name) - includes agents and disks
│   │   └── status               # Check node health (--id|--name)
│   ├── hybrid-node
│   │   ├── create               # Add hybrid node to virtual location (--location-id|--location-name)
│   │   ├── list                 # List hybrid nodes (--location-id|--location-name)
│   │   ├── describe             # Show hybrid node details (--id|--name)
│   │   ├── edit                 # Update hybrid node (--id|--name)
│   │   └── remove               # Remove hybrid node (--id|--name)
│   └── disk
│       ├── list                 # List disks (--node-name or --location-name)
│       ├── describe             # Show disk details (--id) - includes capacity/usage
│       └── status               # Check disk health (--id)
│
├── swarm
│   ├── create                   # Create swarm (--name, --location, --node, --interactive)
│   ├── list (ls)                # List all swarms
│   ├── describe                 # Show swarm details (--id|--name)
│   ├── remove (rm)              # Remove swarm (--id|--name)
│   ├── status                   # Check swarm creation/health status (--id|--name)
│   ├── settings
│   │   ├── view                 # View swarm settings (--swarm-name)
│   │   └── edit                 # Edit swarm settings (--swarm-name, --file|inline flags)
│   ├── metrics
│   │   └── usage                # Show swarm usage metrics (--swarm-name, --from, --to)
│   ├── location
│   │   ├── add                  # Add location to swarm (--swarm-name, --location-name)
│   │   ├── list                 # List locations in swarm (--swarm-name)
│   │   └── remove               # Remove location from swarm (--swarm-name, --location-name)
│   ├── agent
│   │   ├── list                 # List storage agents (--swarm-name, optional --location-name)
│   │   ├── describe             # Show agent details (--id) - includes disk assignment
│   │   ├── status               # Check agent health (--id)
│   │   ├── logs                 # View agent logs (--id)
│   │   └── restart              # Restart agent (--id)
│   └── redundancy-class (alias: rc)
│       ├── create               # Create RC (--swarm-name, --disks-per-location, --location-tolerance, --disk-tolerance)
│       ├── list (ls)            # List redundancy classes (--swarm-name)
│       ├── describe             # Show RC details (--id|--name) - shows level, tolerance, locations
│       ├── edit                 # Update RC (--id|--name)
│       ├── remove (rm)          # Remove RC (--id|--name)
│       ├── detach               # Detach RC from swarm (--id|--name)
│       ├── status               # Check RC health (--id|--name)
│       ├── metrics
│       │   └── usage            # Show RC usage metrics (--name, --from, --to)
│       └── recovery
│           ├── start            # Start data recovery (--id|--name)
│           └── status           # Check recovery status (--id|--name)
│
├── gateway
│   ├── create                   # Create gateway (--name, --type manual|automatic, --location, --node)
│   ├── list (ls)                # List all gateways with status
│   ├── describe                 # Show gateway overview (--id|--name) - status, position, storage classes, bandwidth
│   ├── remove (rm)              # Remove gateway (--id|--name) - DANGER: irreversible
│   ├── status                   # Check gateway health (--id|--name)
│   ├── settings
│   │   ├── view                 # View gateway settings (--gateway-name)
│   │   └── edit                 # Edit gateway settings (--gateway-name) - name, description, connection
│   ├── metrics
│   │   └── usage                # Show gateway usage metrics (--gateway-name, --from, --to)
│   ├── tenant
│   │   ├── list                 # List tenants connected to gateway (--gateway-name)
│   │   └── describe             # Show tenant details within gateway context (--gateway-name, --tenant-name)
│   ├── swarm                    # Manage gateway-swarm connections (destinations)
│   │   ├── connect              # Connect gateway to swarm (--gateway-name, --swarm-name, --rc-name)
│   │   ├── list                 # List connected swarms with RC and storage class (--gateway-name)
│   │   ├── set-default          # Set default swarm destination (--gateway-name, --swarm-name)
│   │   └── disconnect           # Disconnect swarm from gateway (--gateway-name, --swarm-name)
│   ├── storage-class
│   │   ├── create               # Create storage class (--gateway-name, --name, --routing-string, --available-all)
│   │   ├── list                 # List storage classes (--gateway-name) - shows usage by tenants/rules
│   │   ├── describe             # Show storage class details (--gateway-name, --name)
│   │   ├── edit                 # Update storage class (--gateway-name, --name)
│   │   └── remove               # Remove storage class (--gateway-name, --name)
│   └── location
│       ├── list                 # List locations configured for gateway (--gateway-name)
│       └── describe             # Show location details with nodes (--gateway-name, --location-name)
│
├── sdp (alias: rule)   # Routing rules for directing traffic
│   ├── create                   # Create routing rule (--name, --file rule.json | --interactive)
│   ├── list (ls)                # List all routing rules (--status all|draft|active)
│   ├── describe                 # Show rule details (--id|--name) - shows flow: input->conditions->action->fallback
│   ├── edit                     # Update rule (--id|--name, --file rule.json)
│   ├── remove (rm)              # Delete rule (--id|--name)
│   ├── enable                   # Enable rule (--id|--name)
│   ├── disable                  # Disable rule (--id|--name)
│   ├── schedule
│   │   ├── set                  # Set schedule for rule (--id|--name, --cron | --time)
│   │   ├── view                 # View rule schedule (--id|--name)
│   │   └── clear                # Remove schedule (--id|--name)
│   ├── template
│   │   ├── list                 # List available templates (--category data-saving|data-optimization)
│   │   ├── save                 # Save rule as template (--id|--name, --template-name)
│   │   └── use                  # Create rule from template (--template-name, --rule-name)
│   └── test                     # Test rule against sample input (--id|--name, --input-file)
│
├── tenant
│   ├── create                   # Create a new tenant (--name, --gateway-name)
│   ├── list (ls)                # List all tenants
│   ├── describe                 # Show tenant overview (--id|--name) - bandwidth, storage forecast, watchlist
│   ├── remove (rm)              # Remove tenant (--id|--name)
│   ├── gateway
│   │   ├── connect              # Connect tenant to gateway (--tenant-name, --gateway-name)
│   │   ├── list                 # List connected gateways with endpoint and console link (--tenant-name)
│   │   └── disconnect           # Disconnect from gateway (--tenant-name, --gateway-name)
│   ├── user                     # Console users (tenant end-users)
│   │   ├── create               # Create console user (--tenant-name, --email, --name)
│   │   ├── list                 # List console users (--tenant-name) - shows storage, bandwidth, verification, status
│   │   ├── describe             # Show user details (--tenant-name, --email) - 2FA, projects, storage overview
│   │   ├── edit                 # Update user (--tenant-name, --email, --max-projects)
│   │   ├── freeze               # Freeze user (--tenant-name, --email)
│   │   ├── unfreeze             # Unfreeze user (--tenant-name, --email)
│   │   ├── disconnect           # Disconnect all user sessions (--tenant-name, --email)
│   │   ├── ban                  # Ban user (--tenant-name, --email)
│   │   ├── unban                # Unban user (--tenant-name, --email)
│   │   └── remove               # Remove user (--tenant-name, --email)
│   ├── project
│   │   ├── list                 # List projects (--tenant-name) - shows storage, bandwidth, users, buckets, status
│   │   ├── describe             # Show project details (--tenant-name, --project-name) - owner, files, storage, buckets
│   │   ├── threshold
│   │   │   ├── view             # View project thresholds (--tenant-name, --project-name)
│   │   │   └── set              # Set storage/bandwidth thresholds (--tenant-name, --project-name, --storage, --bandwidth)
│   │   ├── freeze               # Freeze project (--tenant-name, --project-name)
│   │   ├── unfreeze             # Unfreeze project (--tenant-name, --project-name)
│   │   ├── ban                  # Ban project (--tenant-name, --project-name)
│   │   ├── unban                # Unban project (--tenant-name, --project-name)
│   │   └── remove               # Delete project (--tenant-name, --project-name)
│   ├── bucket
│   │   ├── list                 # List buckets (--tenant-name) - shows storage, objects, state
│   │   └── describe             # Show bucket details (--tenant-name, --bucket-name) - locking, versioning, project
│   ├── settings
│   │   ├── view                 # View all tenant settings (--tenant-name)
│   │   ├── edit                 # Edit general settings (--tenant-name, --name, --slug, --description, --support-link)
│   │   ├── icon
│   │   │   └── upload           # Upload tenant icon (--tenant-name, --file) - PNG/JPEG/GIF under 10MB
│   │   ├── projects
│   │   │   └── set-defaults     # Set default project limits (--tenant-name, --max-projects, --storage-limit, --bandwidth-limit)
│   │   ├── oauth
│   │   │   ├── google
│   │   │   │   ├── enable       # Enable Google OAuth (--tenant-name, --client-id, --secret)
│   │   │   │   ├── disable      # Disable Google OAuth (--tenant-name)
│   │   │   │   └── view         # View Google OAuth settings (--tenant-name)
│   │   │   └── microsoft
│   │   │       ├── enable       # Enable Microsoft OAuth (--tenant-name, --app-id, --directory-id, --secret)
│   │   │       ├── disable      # Disable Microsoft OAuth (--tenant-name)
│   │   │       └── view         # View Microsoft OAuth settings (--tenant-name)
│   │   ├── white-label
│   │   │   ├── view             # View white label settings (--tenant-name)
│   │   │   ├── enable           # Enable personalized branding (--tenant-name)
│   │   │   ├── disable          # Disable personalized branding (--tenant-name)
│   │   │   ├── upload-terms     # Upload terms & conditions PDF (--tenant-name, --file)
│   │   │   ├── upload-privacy   # Upload privacy policy PDF (--tenant-name, --file)
│   │   │   ├── upload-favicon   # Upload favicon (--tenant-name, --file)
│   │   │   ├── upload-logo      # Upload logo (--tenant-name, --file)
│   │   │   ├── upload-extended-logo  # Upload extended logo (--tenant-name, --file)
│   │   │   └── set-texts        # Set sign-up page texts (--tenant-name, --title, --description, --form-title)
│   │   └── domain               # Custom domains (connections)
│   │       ├── add              # Add custom domain (--tenant-name, --domain, --gateway-name)
│   │       ├── list             # List custom domains (--tenant-name)
│   │       ├── verify           # Verify domain DNS (--tenant-name, --domain)
│   │       ├── certificate      # Upload SSL certificate (--tenant-name, --domain, --cert, --key)
│   │       └── remove           # Remove custom domain (--tenant-name, --domain)
│   ├── metrics
│   │   ├── bandwidth            # Show bandwidth metrics (--tenant-name, --period 7d|15d|30d)
│   │   ├── storage              # Show storage usage trend (--tenant-name, --period 7d|15d|30d)
│   │   └── forecast             # Show storage usage forecast (--tenant-name)
│   └── report                   # Generate tenant usage report (--tenant-name, --format csv|json, --from, --to)
│
├── iam
│   ├── user
│   │   ├── create               # Create IAM user (--username, --password, --email, --group, --policy)
│   │   ├── import               # Bulk import users from CSV/JSON (--file)
│   │   ├── list                 # List IAM users (--status all|active|disabled|pending)
│   │   ├── describe             # Show IAM user details (--username) - groups, policies, status
│   │   ├── edit                 # Update IAM user (--username, --email)
│   │   ├── enable               # Enable user access (--username)
│   │   ├── disable              # Disable user access (--username)
│   │   ├── remove               # Remove IAM user (--username)
│   │   ├── group
│   │   │   ├── add              # Add user to group (--username, --group-name)
│   │   │   ├── list             # List user's groups (--username)
│   │   │   └── remove           # Remove user from group (--username, --group-name)
│   │   ├── policy
│   │   │   ├── attach           # Attach policy to user (--username, --policy-name)
│   │   │   ├── list             # List user's policies (--username)
│   │   │   └── detach           # Detach policy from user (--username, --policy-name)
│   │   └── password
│   │       ├── reset            # Reset user password (--username)
│   │       └── send             # Send password to user via email (--username)
│   ├── group
│   │   ├── create               # Create IAM group (--name, --description)
│   │   ├── list                 # List IAM groups - shows users count, policies
│   │   ├── describe             # Show group details (--name) - users, policies, created on
│   │   ├── edit                 # Update group (--name, --description)
│   │   ├── remove               # Delete group (--name)
│   │   ├── user
│   │   │   ├── add              # Add user to group (--group-name, --username)
│   │   │   ├── list             # List group members (--group-name)
│   │   │   └── remove           # Remove user from group (--group-name, --username)
│   │   └── policy
│   │       ├── attach           # Attach policy to group (--group-name, --policy-name)
│   │       ├── list             # List group's policies (--group-name)
│   │       └── detach           # Detach policy from group (--group-name, --policy-name)
│   ├── policy
│   │   ├── create               # Create policy (--name, --description, --file | --interactive)
│   │   ├── list                 # List policies - shows statements, users, groups, status
│   │   ├── describe             # Show policy details (--name) - statements, users, groups
│   │   ├── edit                 # Update policy (--name, --description)
│   │   ├── enable               # Enable policy (--name)
│   │   ├── disable              # Disable policy (--name)
│   │   ├── remove               # Delete policy (--name)
│   │   ├── statement
│   │   │   ├── add              # Add statement to policy (--policy-name, --effect, --where, --action, --entity)
│   │   │   ├── list             # List policy statements (--policy-name)
│   │   │   └── remove           # Remove statement from policy (--policy-name, --statement-id)
│   │   ├── user
│   │   │   ├── attach           # Attach policy to user (--policy-name, --username)
│   │   │   ├── list             # List users with this policy (--policy-name)
│   │   │   └── detach           # Detach policy from user (--policy-name, --username)
│   │   └── group
│   │       ├── attach           # Attach policy to group (--policy-name, --group-name)
│   │       ├── list             # List groups with this policy (--policy-name)
│   │       └── detach           # Detach policy from group (--policy-name, --group-name)
│   ├── api-key
│   │   ├── create               # Create API key for current user
│   │   ├── list                 # List API keys
│   │   ├── describe             # Show API key details (--id)
│   │   ├── revoke               # Revoke API key (--id)
│   │   └── rotate               # Rotate API key (--id)
│   └── template
│       ├── list                 # List policy templates (CSV, JSON)
│       └── download             # Download template file (--format csv|json)
│
└── docs
    ├── markdown                 # Generate markdown docs
    ├── man                      # Generate man pages
    └── tree                     # Show command tree
```

### Global Flags

```
--profile, -p      Configuration profile to use
--output, -o       Output format: human, json, yaml, xml (default: human)
--quiet, -q        Minimal output (IDs only)
--silent           Suppress all output
--no-headers       Omit table headers
--api-version      Force specific API version (v1, v2)
```

---

## Part 3: Command Examples

### Infrastructure Setup

```bash
# === CLUSTER LOCATIONS ===
# Cluster locations are auto-registered when you install the operator
# The operator install command handles everything in one step

cubbit operator install \
  --kubeconfig ~/.kube/config \
  --name "EU West Production"  # This becomes the location name

# After installation, the cluster appears as a location
cubbit infra location list
# Output shows: "EU West Production" (cluster) - connected

# Check cluster location status
cubbit infra location status --name "EU West Production"

# === VIRTUAL LOCATIONS ===
# Virtual locations are created via API for hybrid storage (e.g., S3)

cubbit infra location create \
  --name "AWS S3 Archive" \
  --type virtual

# Add S3 as a hybrid node in the virtual location
cubbit infra hybrid-node create \
  --location-name "AWS S3 Archive" \
  --name "S3 Archive Bucket" \
  --provider s3 \
  --bucket my-archive-bucket \
  --region eu-west-1 \
  --access-key AKIA... \
  --secret-key ...

# === LISTING INFRASTRUCTURE ===
# List all locations (both cluster and virtual)
cubbit infra location list

# List nodes in a cluster location (auto-discovered from k8s)
cubbit infra node list --location-name "EU West Production"

# List hybrid nodes in a virtual location
cubbit infra hybrid-node list --location-name "AWS S3 Archive"
```

### Infrastructure Management

```bash
# Get overall infrastructure status (like the dashboard)
cubbit infrastructure status

# Example output:
# ┌─────────────────────────────────────────────────────────────────┐
# │  INFRASTRUCTURE STATUS           Last sync: Oct 22, 2025        │
# ├─────────────────────────────────────────────────────────────────┤
# │  Nodes          Storage Usage      Avg. RAM Load                │
# │  ● 189 Healthy  72% (+3%)          57%                          │
# │  ⚠ 33 Warning   vs last month      RAM: 34%  CPU: 63%           │
# │  ✖ 12 Error                                                     │
# ├─────────────────────────────────────────────────────────────────┤
# │  Locations: 12                                                  │
# └─────────────────────────────────────────────────────────────────┘

# List all locations with summary
cubbit infra location list

# Example output (table view):
# NAME          PROVIDER    NODES   VOLUMES   HEALTHY   WARNINGS   ERRORS
# US East 1     AWS         24      400       380       12         8
# EU Central 1  Azure       24      400       395       3          2
# AP Southeast  GCP         24      400       400       0          0

# Show location details
cubbit infra location describe --name "US East 1"

# Example output:
# Location: US East 1
# Status: Healthy
# Provider: AWS
# Position: Virginia, USA
# Last Update: Oct 22, 2025, 11:25 AM
#
# Volumes: 422 Healthy, 14 Warnings, 5 Errors
# Storage Usage: 72% (+3% vs last month)
# Avg RAM Load: 57%

# View visual infrastructure topology (locations -> nodes -> disks)
cubbit infra topology

# Example output:
# ┌─────────────────────────────────────────────────────────────────┐
# │  US East 1 (AWS, Virginia) ✓         24 nodes, 400 volumes      │
# │  ●●●●●●●●●●●●●●●●●●●●○○○○                                       │
# │  ●●●●●●●●●●●●●●●●●●●●●●●●                                       │
# ├─────────────────────────────────────────────────────────────────┤
# │  EU Central 1 (Azure, Ireland) ✓     24 nodes, 400 volumes      │
# │  ●●●●●●●●●●●●●●●●●●●●●●●●                                       │
# │  ●●●●●●●●●●●●●●●●●●●●●●●●                                       │
# ├─────────────────────────────────────────────────────────────────┤
# │  AP Southeast (GCP, Singapore) ⚠     24 nodes, 400 volumes      │
# │  ●●●●●●●●●●●●●●●●⚠⚠⚠⚠○○○○                                       │
# │  ●●●●●●●●●●●●●●●●●●●●●●●●                                       │
# └─────────────────────────────────────────────────────────────────┘
# Legend: ● Healthy  ⚠ Warning  ○ Error

# === NODE MANAGEMENT ===

# List nodes in a location (table view like UI)
cubbit infra node list --location-name "US East 1"

# Example output:
# NODE                LOCATION    POSITION        DISKS   CAPACITY   WARNINGS   ERRORS   HEALTH
# udine-rack3-srv05   US East 1   Virginia, USA   400     32 TB      12         8        Warning
# paris-cache1-srv08  US East 1   Virginia, USA   250     20 TB      0          0        Healthy
# sofia-backups-srv04 US East 1   Virginia, USA   500     40 TB      15         10       Error

# Show node details
cubbit infra node describe --name "udine-rack3-srv05"

# Example output:
# Node: udine-rack3-srv05
# Status: ● Healthy
# ID: node-0-5
# OS: Ubuntu 22.04
#
# Network:
#   Public IP: 194.67.120.35
#   Local IP: 10.3.5.12
#   Ports: 4001, 4002
#
# Resources:
#   CPU: 8 vCPU (35% usage last hour)
#   RAM: 32GB (62% usage last hour)
#   Uptime: 6d 12h
#
# Storage:
#   Used: 542 TB / 960 TB (57%)
#   Disk Allocation: 120 DISKS × 8TB SATA = 960 TB RAW
#   Health: 60/70 healthy
#
# Tenants: Acme Enterprises
# Assigned Data: 120 TB / 250 TB

# Cordon a node (prevent new workloads)
cubbit infra node cordon --name "udine-rack3-srv05"

# Uncordon a node
cubbit infra node uncordon --name "udine-rack3-srv05"

# === DISK MANAGEMENT ===

# List disks on a node (shows visual health grid)
cubbit infra disk list --node-name "udine-rack3-srv05"

# List all disks in a location
cubbit infra disk list --location-name "US East 1"

# Filter to show only problems
cubbit infra disk list --node-name "udine-rack3-srv05" --problems-only

# Show disk details
cubbit infra disk describe --id "disk-104"

# Example output:
# Disk: disk-104
# Status: ○ Offline
# Mount Point: data/agent00
# OCCP Port: 4001
# Public Key: cs02-4b5c
# Last Sync: 2m 11s ago
# Disk Type: HDD
#
# Storage: 342 GB / 1 TB (34% used)
#
# Activity Log:
#   2025-06-17 10:42:19  Sync completata con ridondanza
#   2025-06-18 11:15:30  Backup eseguito con successo
#   2025-06-19 09:30:45  Aggiornamento sistema completato
#   2025-06-20 14:52:10  Verifica integrita file riuscita
#   2025-06-21 08:09:55  Sincronizzazione con nodo esterno

# Check disk health
cubbit infra disk status --id "disk-104"

# Mark disk error as resolved
cubbit infra disk resolve --id "disk-104"
```

### Swarm Agent Management

```bash
# List all storage agents in a swarm
cubbit swarm agent list --swarm-name "Global Storage"

# Filter agents by location
cubbit swarm agent list --swarm-name "Global Storage" --location-name "EU West Production"

# View agent details (includes disk assignment, resource usage)
cubbit swarm agent describe --id <agent-id>

# Check agent health
cubbit swarm agent status --id <agent-id>

# View agent logs (stream)
cubbit swarm agent logs --id <agent-id> --follow

# View last 100 lines of logs
cubbit swarm agent logs --id <agent-id> --tail 100

# Restart a problematic agent
cubbit swarm agent restart --id <agent-id>
```

### Swarm Management

```bash
# === SWARM LIST ===
# List all swarms (table view like UI)
cubbit swarm list

# Example output:
# NAME                      USED CAPACITY   % USAGE   LOCATIONS   CREATED ON           LAST SYNC      STATUS
# Storage EU West           25/50 TB        50%       4           Jan 15, 2025, 14:30  Today, 11:25   Success
# Storage US East           30/75 TB        40%       3           Feb 20, 2025, 09:15  Today, 10:00   Error
# Backup Asia Pacific       20/40 TB        50%       5           Mar 10, 2025, 16:45  Today, 09:30   Success
# Database NA Central       15/30 TB        50%       2           Apr 05, 2025, 11:00  Today, 12:15   Success
# Production US West        40/60 TB        67%       4           Aug 15, 2025, 10:00  Today, 16:30   Success

# Filter by status
cubbit swarm list --status degraded
cubbit swarm list --status error

# Show draft swarms (not yet deployed)
cubbit swarm list --drafts

# === SWARM CREATION ===
# Interactive mode guides you through the wizard (like the UI)
cubbit swarm create --interactive

# Or create non-interactively with all parameters
cubbit swarm create \
  --name "Storage EU West" \
  --description "Multi-region AWS deployment" \
  --location "EU North 1" \
  --location "AP Southeast 1" \
  --location "EU North 2" \
  --location "AP Southeast 4"

# Select specific nodes per location
cubbit swarm create \
  --name "Storage EU West" \
  --description "Production geo-distributed storage" \
  --location "EU North 1" \
  --node "udine-rack3-srv05" \
  --node "udine-rack3-srv10" \
  --location "AP Southeast 1" \
  --node "singapore-srv01"

# Save as draft (don't deploy yet)
cubbit swarm create --name "Test Swarm" --location "US East 1" --draft

# === SWARM DETAILS ===
# Show swarm details (like the swarm detail page)
cubbit swarm describe --name "Storage EU West"

# Example output:
# Swarm: Storage EU West
# Status: ● Online
# Last Update: 04/06/2025, 11:25 AM
#
# ⚠ Alert: US East 1 - 1 ring offline, Redundancy temporarily reduced to 2x
#
# Storage Usage:
#   Total Used: 500 TB
#   Total Assigned: 147 TB
#   Total Unused: 12 TB
#
# Usage per Redundancy Class:
#   RC 2x: ████████░░ 130 TB
#   RC 3x: █████████████░░ 250 TB
#   RC 4x: ██████████░░ 200 TB
#   RC 6x: █████░░░░░ 114 TB
#
# Locations (6):
#   LOCATION      REDUNDANCY CLASS   USAGE %   HEALTH
#   US East 1     3x                 52%       ⚠ Warning
#   US East 1     3x                 52%       ● Healthy
#   US West 2     5x                 70%       ● Healthy
#   EU Central 1  2x                 35%       ● Healthy
#   AP Southeast  4x                 60%       ● Healthy

# Check swarm status
cubbit swarm status --name "Storage EU West"

# Add a new location to an existing swarm
cubbit swarm location add --swarm-name "Storage EU West" --location-name "AP Northeast 1"

# Remove a location from a swarm
cubbit swarm location remove --swarm-name "Storage EU West" --location-name "US East 1"

# === REDUNDANCY CLASS MANAGEMENT ===
# List all redundancy classes in a swarm
cubbit swarm rc list --swarm-name "Storage EU West"

# Example output:
# NAME            DISKS   TOLERANCE   LEVEL   LOCATION                    STATE
# Standard (3x)   120     5/22        3x      Milan, London, +3           ○ Offline
# Standard (3x)   120     5/22        3x      Milan, London, +3           ⚠ Degraded
# Premium (5x)    200     8/22        5x      Paris, Berlin, +5           ● Online
# Deluxe (7x)     300     12/22       7x      New York, Tokyo, +7         ● Online
# Basic (1x)      60      2/22        1x      Amsterdam, Madrid, +1       ● Online
# Business (4x)   150     6/22        4x      Toronto, Sydney, +4         ● Online
# Elite (6x)      250     10/22       6x      Dubai, Singapore, +6        ● Online
# Economy (2x)    80      4/22        2x      Lisbon, Copenhagen, +2      ● Online

# Filter by state
cubbit swarm rc list --swarm-name "Storage EU West" --state offline
cubbit swarm rc list --swarm-name "Storage EU West" --state degraded

# Download list as CSV
cubbit swarm rc list --swarm-name "Storage EU West" --output csv > redundancy-classes.csv

# Create a new redundancy class (interactive wizard with preview)
cubbit swarm rc create --swarm-name "Storage EU West" --interactive

# Create non-interactively with fault tolerance parameters
cubbit swarm rc create \
  --swarm-name "Storage EU West" \
  --name "Standard" \
  --disks-per-location 120 \
  --location-tolerance 5 \
  --disk-tolerance 22

# Show redundancy class details
cubbit swarm rc describe --name "Standard (3x)"

# Example output:
# Redundancy Class: Standard (3x)
# State: ● Online
#
# Configuration:
#   Disks: 120
#   Tolerance: 5/22
#   Level: 3x
#
# Locations:
#   Milan, London, Paris, Berlin, +3 more
#
# Health:
#   ●●●●●●●●●●●●●●●●●●●●○○○○  (20/24 healthy)

# Edit a redundancy class
cubbit swarm rc edit --name "Standard (3x)"

# Delete a redundancy class
cubbit swarm rc remove --name "Standard (3x)"

# Check RC status
cubbit swarm rc status --name "Standard (3x)"

# Start data recovery for an RC
cubbit swarm rc recovery start --name "Standard (3x)"

# Check recovery status
cubbit swarm rc recovery status --name "Standard (3x)"

# === GATEWAY LIST ===
# List all gateways (table view like UI)
cubbit gateway list

# Example output:
# NAME              LOCATION   TYPE        N. SWARM   TENANTS   INGRESS BW   EGRESS BW   STATUS
# Gateway Cubbit 1  Milano     Manual      3          15        3 GB         3 GB        ● Online
# Gateway Cubbit 2  Rome       Automatic   4          20        3 GB         3 GB        ○ Offline
# Gateway Cubbit 3  Florence   Manual      2          10        3 GB         3 GB        ● Online
# Gateway Cubbit 4  Naples     Automatic   5          25        3 GB         3 GB        ● Online

# Filter by status
cubbit gateway list --status online
cubbit gateway list --status offline

# === GATEWAY CREATION ===
# Create an automatic gateway (DNS managed by Composer via API Key)
cubbit gateway create \
  --name "Gateway 01" \
  --description "Production gateway for EU region" \
  --type automatic \
  --provider "AWS Route 53" \
  --api-key "AIzaSyDasxGmWKa4JsXZ-HjGw7ISLn_3nam" \
  --location "EU North 1" \
  --node "udine-rack3-srv05" \
  --node "udine-rack3-srv06" \
  --location "EU North 2" \
  --node "paris-srv01"

# Create a manual gateway (user manages DNS)
cubbit gateway create \
  --name "Custom Gateway" \
  --description "Manual DNS configuration" \
  --type manual \
  --location "US East 1"

# Interactive mode for gateway creation (wizard-style)
cubbit gateway create --interactive

# === GATEWAY DETAILS ===
# Show gateway overview (like the gateway detail page)
cubbit gateway describe --name "Gateway_01"

# Example output:
# Gateway: Gateway_01
# Status: ● Online
# Position: Bologna, IT
#
# Storage Classes: glacier, pd-standard, gp2, +2 more
#
# Average Bandwidth (last 7 days):
#   Ingress: 150 kB    Egress: 36 kB
#
# Traffic Distribution by SDP Rule:
#   Saving large files    90% (2.1 TB)
#   Large file saving     7% (161 GB)
#   Archive old files     2% (8 GB)
#   Other                 1% (3 GB)
#   Total: 2.3 TB (+12%)
#
# Active Smart Data Placement: 6
#   RULE NAME               STORAGE CLASS   STATUS
#   Saving large files      Glacier         On
#   Monthly data backup     SkyVault        Off
#   Project synchronization SyncPlus        Off

# === GATEWAY SETTINGS ===
# View gateway settings
cubbit gateway settings view --gateway-name "Gateway 01"

# Edit gateway settings (name, description)
cubbit gateway settings edit --gateway-name "Gateway 01" \
  --name "Gateway 01 Updated" \
  --description "Updated description for production gateway"

# === GATEWAY-SWARM CONNECTIONS (Destinations) ===
# Connect gateway to a swarm with redundancy class
cubbit gateway swarm connect \
  --gateway-name "Gateway 01" \
  --swarm-name "EU-West-1" \
  --rc-name "3+1 Erasure" \
  --storage-class "Hot"

# List connected swarms (destinations)
cubbit gateway swarm list --gateway-name "Gateway 01"

# Example output:
# SWARM           LOCATION   REDUNDANCY CLASS   STORAGE CLASS   STATUS
# EU-West-1       Pavia      3+1 Erasure        Hot             ○ Not active
# US-East-1       Milano     2+1 Erasure        Cold            ○ Not active
# AP-Southeast-1  Bologna    API Key            Hot             ● Active

# Add new destination
cubbit gateway swarm connect \
  --gateway-name "Gateway 01" \
  --swarm-name "AP-Northeast-1" \
  --rc-name "5+2 Erasure" \
  --default

# Disconnect swarm from gateway
cubbit gateway swarm disconnect --gateway-name "Gateway 01" --swarm-name "EU-West-1"

# === STORAGE CLASSES ===
# List storage classes for a gateway
cubbit gateway storage-class list --gateway-name "Gateway 01"

# Example output:
# NAME             ROUTING STRING     USED BY TENANTS   USED IN RULES   STATUS
# Hot storage      env=production     0                 5               ● Active
# Cold storage     env=development    1                 10              ○ Unused
# Backup storage   env=test           2                 15              ● Active
# Archive storage  env=staging        3                 20              ● Active

# Create a storage class
cubbit gateway storage-class create \
  --gateway-name "Gateway 01" \
  --name "Hot storage" \
  --routing-string "tier=hot" \
  --available-all  # Available to all connected tenants

# Edit storage class
cubbit gateway storage-class edit \
  --gateway-name "Gateway 01" \
  --name "Hot storage" \
  --routing-string "tier=hot,env=production"

# Remove storage class
cubbit gateway storage-class remove --gateway-name "Gateway 01" --name "Archive storage"

# === TENANTS CONNECTED TO GATEWAY ===
# List tenants connected to a gateway
cubbit gateway tenant list --gateway-name "Gateway_01"

# Example output:
# ID                LINK CONSOLE               ENDPOINT                    INGRESS BW   STATUS
# Tenant Cubbit 1   console.loremipsum.com     customdomain.loremipsum.com 3 GB         ● Healthy
# Tenant Cubbit 2   console.loremipsum.com     customdomain.loremipsum.com 4 GB         ⚠ Warning
# Tenant Cubbit 3   console.loremipsum.com     customdomain.loremipsum.com 4 GB         ⚠ Ready

# Show tenant details within gateway context
cubbit gateway tenant describe --gateway-name "Gateway_01" --tenant-name "Tenant Cubbit 1"

# === GATEWAY LOCATIONS ===
# List locations configured for gateway
cubbit gateway location list --gateway-name "Gateway 01"

# Example output:
# LOCATION      NODES
# EU North 1    udine-rack3-srv05, udine-rack3-srv06, udine-rack3-srv07, udine-rack3-srv08
# EU North 2    paris-srv01, paris-srv02
# EU West 1     london-srv01 (10 nodes)
# US East 3     nyc-srv01 (12 nodes)
# AP Southeast 5 singapore-srv01 (7 nodes)

# === DELETE GATEWAY (DANGER ZONE) ===
# Delete gateway permanently (requires confirmation)
cubbit gateway remove --name "Gateway 01" --confirm

# Warning message:
# ⚠ DANGER: Deleting a gateway permanently revokes its access, blocks all requests,
# and makes any cached but undistributed data unreachable.
# This action is irreversible; once confirmed, the information cannot be retrieved.
# Type the gateway name to confirm: Gateway 01
```

### Smart Data Placement Management

```bash
# === ROUTING RULES LIST ===
# List all routing rules
cubbit sdp list

# Example output:
# RULE NAME                  STORAGE CLASS   GATEWAY   TENANTS   LAST ACTIVITY   STATUS   SCHEDULE
# Salvataggio file grandi    Glacier         12        8         2 hours ago     On       ---
# Salvataggio file grandi    Glacier         12        8         2 hours ago     On       Sun, 10:00
# Backup dati mensili        SkyVault        20        15        2 hours ago     On       ---
# Archiviazione video        CloudStore      30        22        30 minutes ago  Off      ---
# Sincronizzazione progetti  SyncPlus        25        18        10 minutes ago  On       ---

# Filter by status
cubbit sdp list --status draft
cubbit sdp list --status active

# List user's templates
cubbit sdp template list

# Example output:
# CATEGORY           TEMPLATE NAME
# Data Saving        Large file archival
# Data Saving        Backup policy
# Data Optimization  Video compression routing
# Data Optimization  Cold storage migration

# === CREATE ROUTING RULE ===
# Create from JSON file (defines the full flow)
cubbit sdp create --name "my-archive-rule" --file rule.json

# Example rule.json structure:
# {
#   "input": {
#     "event": "file_created",
#     "bucket": "acme-corp"
#   },
#   "conditions": [
#     { "field": "file.extension", "operator": "ends_with", "value": ".mp4" }
#   ],
#   "action": {
#     "type": "route_dynamic",
#     "redundancy_classes": ["rc-001", "rc-002", "rc-003"],
#     "distribution": "weighted"
#   },
#   "fallback": {
#     "type": "route",
#     "target": "default-rc"
#   }
# }

# Create from template
cubbit sdp template use \
  --template-name "Large file archival" \
  --rule-name "Archive MP4 files"

# Interactive mode (wizard)
cubbit sdp create --interactive

# === ROUTING RULE DETAILS ===
# Show rule details (flow visualization)
cubbit sdp describe --name "myrule"

# Example output:
# Rule: [myrule]
# Status: Saved (Draft)
# Last saved: 2 mins ago
# Enabled: No
#
# Flow:
#   START
#     └─► INPUT: When file created in acme-corp
#           └─► CONDITION: If file name ends with ".mp4"
#                 └─► ACTION: Route with dynamic criteria
#                       Distributed across 3 redundancy classes (weighted)
#                       └─► FALLBACK: Route to default RC
#                             └─► END

# === ENABLE/DISABLE RULES ===
cubbit sdp enable --name "myrule"
cubbit sdp disable --name "myrule"

# === SCHEDULE MANAGEMENT ===
# Set schedule for rule
cubbit sdp schedule set --name "Backup rule" --time "Sun 10:00"
cubbit sdp schedule set --name "Nightly sync" --cron "0 2 * * *"

# View schedule
cubbit sdp schedule view --name "Backup rule"

# Clear schedule
cubbit sdp schedule clear --name "Backup rule"

# === TEMPLATES ===
# Save rule as template
cubbit sdp template save \
  --name "myrule" \
  --template-name "My custom archival template"

# List templates by category
cubbit sdp template list --category "Data Saving"

# === DELETE RULE ===
cubbit sdp remove --name "myrule"

# === TEST RULE ===
# Test rule against sample input
cubbit sdp test --name "myrule" --input-file test-event.json
```

### Tenant Management

```bash
# === TENANT LIST & CREATION ===
# List all tenants
cubbit tenant list

# Create a tenant
cubbit tenant create \
  --name "Tenant Cubbit 1" \
  --admin-email admin@acme.com

# Show tenant details
cubbit tenant describe --name "Tenant Cubbit 1"

# === GATEWAY CONNECTIONS ===
# List connected gateways
cubbit tenant gateway list --tenant-name "Tenant Cubbit 1"

# Example output:
# GATEWAY              POSITION              TYPE            CERTIFICATE        STATUS
# app.mycompany.com    Pavia                 CUSTOM DOMAIN   ⚠ Expiring 10.09   ○ Offline
# app.mycompany.com    Milano                API Key         ---                ○ Inactive
# cdn.mycompany.com    Bologna               API Key         ---                ● Online
# test-offloader-url   San Giovanni in Per.  CUSTOM DOMAIN   ○ Missing          ● Online
# cdn.mycompany.com    San Giovanni in Per.  CUSTOM DOMAIN   ● Valid            ● Online

# Connect tenant to gateway
cubbit tenant gateway connect \
  --tenant-name "Tenant Cubbit 1" \
  --gateway-name "gw-eu-west-1"

# Disconnect from gateway
cubbit tenant gateway disconnect \
  --tenant-name "Tenant Cubbit 1" \
  --gateway-name "gw-eu-west-1"

# === CUSTOM DOMAINS ===
# List custom domains
cubbit tenant domain list --tenant-name "Tenant Cubbit 1"

# Example output:
# DOMAIN                          GATEWAY        STATUS
# app.mycompany.com               gw-eu-west-1   ✓ Verified
# app.mycompany.com               gw-eu-west-2   ✓ Verified
# cdn.mycompany.com               gw-us-east-1   ✓ Verified
# custom.yourcompany.com          ---            ⟳ Verification
# app.mycompany.com               gw-us-east-1   ○ Unverified

# Add a custom domain
cubbit tenant domain add \
  --tenant-name "Tenant Cubbit 1" \
  --domain "app.mycompany.com" \
  --gateway-name "gw-eu-west-1"

# Verify domain DNS
cubbit tenant domain verify --tenant-name "Tenant Cubbit 1" --domain "app.mycompany.com"

# Upload SSL certificate for custom domain
cubbit tenant domain certificate \
  --tenant-name "Tenant Cubbit 1" \
  --domain "app.mycompany.com" \
  --certificate-file cert.pem \
  --key-file key.pem

# Remove custom domain
cubbit tenant domain remove --tenant-name "Tenant Cubbit 1" --domain "app.mycompany.com"

# === CONSOLE USERS (Tenant End Users) ===
# List console users
cubbit tenant user list --tenant-name "Tenant Cubbit 1"

# Example output:
# EMAIL                          USED STORAGE   BANDWIDTH   VERIFIED   PROJECTS   STATUS
# mario.rossi@example.com        22 GB          128 kB      Yes        12         ○ Removed
# lucy.smith@example.com         45 GB          256 kB      No         8          ● Active
# john.doe@example.com           30 GB          500 kB      Yes        5          ⊘ Banned
# emma.jones@example.com         60 GB          1 MB        ---        10         ● Active
# noah.brown@example.com         25 GB          750 kB      Yes        15         ● Active

# Filter by status
cubbit tenant user list --tenant-name "Tenant Cubbit 1" --status active
cubbit tenant user list --tenant-name "Tenant Cubbit 1" --status banned
cubbit tenant user list --tenant-name "Tenant Cubbit 1" --status removed

# Download as CSV
cubbit tenant user list --tenant-name "Tenant Cubbit 1" --output csv > users.csv

# Show user details
cubbit tenant user describe --tenant-name "Tenant Cubbit 1" --email "mario.rossi@example.com"

# Example output:
# User: Mario Rossi
# Email: mario.rossi@example.com
# 2FA Authentication: ○ Not active
# Projects: 1
# Maximum project number: 1
#
# Storage Overview:
#   Total Used: 22 GB
#
# Projects:
#   NAME            STORAGE    STATUS
#   Project name 1  120 MB     ⚠ Limited storage
#   Project name 2  120 MB     ● Operational
#   Project name 3  120 MB     ○ Gateway error
#   Project name 4  120 MB     ❄ Freezed

# Create tenant user
cubbit tenant user create \
  --tenant-name "Tenant Cubbit 1" \
  --email user@acme.com \
  --name "John Doe"

# Set max projects for user
cubbit tenant user edit \
  --tenant-name "Tenant Cubbit 1" \
  --email "mario.rossi@example.com" \
  --max-projects 5

# Ban/unban user
cubbit tenant user ban --tenant-name "Tenant Cubbit 1" --email "user@example.com"
cubbit tenant user unban --tenant-name "Tenant Cubbit 1" --email "user@example.com"

# Remove user
cubbit tenant user remove --tenant-name "Tenant Cubbit 1" --email "user@example.com"

# === WHITE LABEL SETTINGS ===
# View white label settings
cubbit tenant settings white-label view --tenant-name "Tenant Cubbit 1"

# Enable/disable personalized branding
cubbit tenant settings white-label enable --tenant-name "Tenant Cubbit 1"
cubbit tenant settings white-label disable --tenant-name "Tenant Cubbit 1"

# Upload legal documents
cubbit tenant settings white-label upload-terms \
  --tenant-name "Tenant Cubbit 1" \
  --file terms-and-conditions.pdf

cubbit tenant settings white-label upload-privacy \
  --tenant-name "Tenant Cubbit 1" \
  --file privacy-policy.pdf

# Upload branding assets
cubbit tenant settings white-label upload-favicon \
  --tenant-name "Tenant Cubbit 1" \
  --file favicon.png

cubbit tenant settings white-label upload-logo \
  --tenant-name "Tenant Cubbit 1" \
  --file logo.png

cubbit tenant settings white-label upload-extended-logo \
  --tenant-name "Tenant Cubbit 1" \
  --file extended-logo.png

# Set sign-up page texts
cubbit tenant settings white-label set-texts \
  --tenant-name "Tenant Cubbit 1" \
  --title "Welcome to Lore Tenant, the coolest cloud object storage" \
  --description "Join the test Tenant managed by Jurenz" \
  --form-title "Sign-up form below"

# === TENANT PROJECTS ===
# List projects in tenant
cubbit tenant project list --tenant-name "Tenant Cubbit 1"

# Show project details
cubbit tenant project describe --tenant-name "Tenant Cubbit 1" --project-name "Project name 1"

# Freeze/unfreeze project
cubbit tenant project freeze --tenant-name "Tenant Cubbit 1" --project-name "Project name 1"
cubbit tenant project unfreeze --tenant-name "Tenant Cubbit 1" --project-name "Project name 1"

# === TENANT BUCKETS ===
# List buckets in tenant
cubbit tenant bucket list --tenant-name "Tenant Cubbit 1"

# Filter by project
cubbit tenant bucket list --tenant-name "Tenant Cubbit 1" --project-name "Project name 1"

# Show bucket details
cubbit tenant bucket describe --tenant-name "Tenant Cubbit 1" --bucket-name "my-bucket"

# === USAGE REPORTS ===
# Generate usage report
cubbit tenant report --tenant-name "Tenant Cubbit 1" --format csv --from 2024-01-01 --to 2024-01-31
```

### IAM Management

```bash
# === IAM USERS ===
# List all IAM users
cubbit iam user list

# Example output:
# USERNAME              EMAIL                    GROUPS   POLICIES   STATUS      ACCESS   LAST ACTIVITY
# mario.rossi           john.doe@company.com     26       2          ● Active    On       Aug 12, 2025, 12:00
# jane.smith            -                        4        0          ○ Disabled  Off      Aug 12, 2025, 12:00
# mike.jones            mike.jones@company.com   6        1          ◐ Pending   Off      Aug 12, 2025, 12:00
# lucy.brown            lucy.brown@company.com   2345     1          ● Active    On       Aug 12, 2025, 12:00

# Filter by status
cubbit iam user list --status active
cubbit iam user list --status disabled
cubbit iam user list --status pending

# Download as CSV
cubbit iam user list --output csv > iam-users.csv

# Create a single IAM user
cubbit iam user create \
  --username "john.doe" \
  --password "SecurePass123!" \
  --email "john.doe@company.com" \
  --group "Admins" \
  --group "Developers" \
  --policy "S3FullAccess" \
  --send-password-email  # Send password to user via email

# Bulk import users from CSV/JSON
cubbit iam user import --file users.csv

# Download template for bulk import
cubbit iam template download --format csv > users-template.csv
cubbit iam template download --format json > users-template.json

# Show user details
cubbit iam user describe --username "john.doe"

# Enable/disable user access
cubbit iam user enable --username "john.doe"
cubbit iam user disable --username "john.doe"

# Add user to group
cubbit iam user group add --username "john.doe" --group-name "Data Analysts"

# Remove user from group
cubbit iam user group remove --username "john.doe" --group-name "Developers"

# Attach policy directly to user
cubbit iam user policy attach --username "john.doe" --policy-name "ReadOnlyAccess"

# Detach policy from user
cubbit iam user policy detach --username "john.doe" --policy-name "ReadOnlyAccess"

# Reset password
cubbit iam user password reset --username "john.doe"

# Send password via email
cubbit iam user password send --username "john.doe"

# Remove user
cubbit iam user remove --username "john.doe"

# === IAM GROUPS ===
# List all groups
cubbit iam group list

# Example output:
# NAME                 USERS   POLICIES                            CREATED ON
# Admin                12      Policy name 1, Policy 2, +3          Aug 10, 2025, 12:00
# Swarm Admin          26      Policy name 1, Policy 2, +3          Aug 10, 2025, 12:00
# Node Manager         18      Policy name 2, Policy 3, +5          Sep 15, 2025, 14:30
# Data Analyst         22      Policy name 3, Policy 4, +2          Oct 20, 2025, 09:00

# Download as CSV
cubbit iam group list --output csv > groups.csv

# Create a group
cubbit iam group create \
  --name "Infrastructure Admins" \
  --description "Full access to infrastructure resources"

# Show group details
cubbit iam group describe --name "Admin"

# Edit group
cubbit iam group edit --name "Admin" --description "Updated description"

# Add user to group
cubbit iam group user add --group-name "Admin" --username "new.user"

# Remove user from group
cubbit iam group user remove --group-name "Admin" --username "old.user"

# Attach policy to group
cubbit iam group policy attach --group-name "Admin" --policy-name "FullAccess"

# Detach policy from group
cubbit iam group policy detach --group-name "Admin" --policy-name "OldPolicy"

# Delete group
cubbit iam group remove --name "Deprecated Group"

# === IAM POLICIES ===
# List all policies
cubbit iam policy list

# Example output:
# NAME            STATEMENTS   USERS   GROUPS                    STATUS     CREATED ON
# Policy name 1   12           12      Admin, Customers, +3      ● Active   Aug 10, 2025, 12:00
# Policy name 2   20           15      Manager, Partners, +5     ● Active   Sep 15, 2025, 14:30
# Policy name 3   10           8       User, Investors, +2       ● Active   Oct 1, 2025, 09:00

# Download as CSV
cubbit iam policy list --output csv > policies.csv

# Create policy interactively (guided mode)
cubbit iam policy create --name "SwarmReadOnly" --interactive

# Create policy from JSON file
cubbit iam policy create --name "CustomPolicy" --file policy.json

# Show policy details
cubbit iam policy describe --name "Policy name 1"

# Add statement to policy (guided mode)
cubbit iam policy statement add \
  --policy-name "CustomPolicy" \
  --effect allow \
  --where swarm \
  --entity "Swarm 1,Swarm 2" \
  --action "Read,Write"

# List statements in policy
cubbit iam policy statement list --policy-name "CustomPolicy"

# Remove statement from policy
cubbit iam policy statement remove --policy-name "CustomPolicy" --statement-id 3

# Enable/disable policy
cubbit iam policy enable --name "CustomPolicy"
cubbit iam policy disable --name "CustomPolicy"

# Attach policy to user
cubbit iam policy user attach --policy-name "ReadOnly" --username "viewer"

# Attach policy to group
cubbit iam policy group attach --policy-name "FullAccess" --group-name "Admins"

# Delete policy
cubbit iam policy remove --name "OldPolicy"

# === API KEYS ===
# Create API key for current user
cubbit iam api-key create --name "CI/CD Pipeline" --expires-in 90d

# List API keys
cubbit iam api-key list

# Describe API key
cubbit iam api-key describe --id "key-abc123"

# Revoke API key
cubbit iam api-key revoke --id "key-abc123"

# Rotate API key (creates new key, revokes old)
cubbit iam api-key rotate --id "key-abc123"
```

### Statement Effect and Where Options

When creating policy statements, use these options:

**Effect:**
- `allow` - Grant permission
- `deny` - Explicitly deny permission

**Where (Resource Scope):**
- `everywhere` - All resources
- `swarm` - Swarm resources
- `gateway` - Gateway resources
- `tenant` - Tenant resources
- `users` - User resources
- `infrastructure` - Infrastructure resources

**Actions (examples):**
- `Read`, `Write`, `Delete`, `Edit`, `All`, `None`, `Moderate`, `Selective`

**Entity:**
- Cubbit Resource Name (CRN) or `*` for all resources in scope

---

## Part 4: Legacy CLI Migration

### Binary Distribution

- **New CLI**: `cubbit` (this implementation)
- **Legacy CLI**: `cubbit-legacy` (separate binary, current implementation)

### Migration Path

1. Ship both binaries in the same package
2. Legacy commands show deprecation warnings pointing to new equivalents
3. Documentation maps old commands to new commands

### Command Mapping (Legacy -> New)

| Legacy Command | New Command | Notes |
|---------------|-------------|-------|
| `cubbit nexus create --swarm-id X` | `cubbit infra location create --type cluster` | Nexus = Location |
| `cubbit nexus list --swarm-id X` | `cubbit swarm location list --swarm-id X` | |
| `cubbit node create --swarm-id X --nexus-id Y` | Automatic (operator managed) | Nodes auto-discovered |
| `cubbit node list --swarm-id X --nexus-id Y` | `cubbit infra node list --location-id Y` | |
| `cubbit agent create ...` | Automatic (operator managed) | Agents auto-deployed |
| `cubbit agent list ...` | `cubbit swarm agent list --swarm-name X` | |
| `cubbit rc create --swarm-id X` | `cubbit swarm rc create --swarm-id X` | Nested under swarm |
| `cubbit gateway create --tenant-id X` | `cubbit gateway create` + `cubbit tenant gateway connect` | Decoupled |
| `cubbit user create --tenant-id X` | `cubbit tenant user create --tenant-id X` | Nested under tenant |
| `cubbit project list --tenant-id X` | `cubbit tenant project list --tenant-id X` | Nested under tenant |
| `cubbit iam user create` | `cubbit iam user create` | Same |

---

## Part 5: Implementation Plan

### Phase 1: Core Infrastructure (Week 1-2)

1. **Create infrastructure command group**
   - `src/cmd/infrastructure.go` - Parent command with status/topology
   - Move `location` under infra (was nexus)
   - Move `node` under infra
   - Add `hybrid-node` under infra
   - Add `agent` under infra
   - Add `disk` under infra

2. **Create organization and operator commands**
   - `src/cmd/organization.go` - Org describe/edit
   - `src/cmd/operator.go` - Install/upgrade/uninstall/status

3. **Update API layer**
   - `src/api/organization.go` - Organization endpoints
   - `src/api/operator.go` - Operator endpoints
   - `src/api/infrastructure.go` - Location, node, agent, disk endpoints

4. **Update action layer**
   - Mirror API layer structure

### Phase 2: Resource Layer (Week 3-4)

1. **Update Swarm commands**
   - Add `swarm location add/list/remove` subcommands
   - Move `redundancy-class` under swarm as `swarm rc`

2. **Create Gateway commands**
   - Add `gateway storage-class` subcommands
   - Add `gateway smart-data-placement` subcommands

3. **Deprecate legacy commands**
   - Hide `nexus`, `agent` top-level commands
   - Point to new locations in help text

### Phase 3: Tenant Layer (Week 5-6)

1. **Restructure Tenant commands**
   - Nest `user` under `tenant user`
   - Nest `project` under `tenant project`
   - Add `tenant gateway connect/disconnect`
   - Add `tenant dns configure/verify`

2. **Update API/Action layers**

### Phase 4: IAM Layer (Week 7-8)

1. **Expand IAM commands**
   - Add `iam group` with member management
   - Add `iam api-key` commands
   - Update `iam policy` commands

2. **Add auth enhancements**
   - Add `auth whoami` command

### Phase 5: Polish & Documentation (Week 9-10)

1. **Update all help descriptions**
2. **Fix alias bugs** (rm on restore/unfreeze)
3. **Add smart data placement validation**
4. **Generate documentation**
5. **Create migration guide**

---

## Part 6: Files to Create/Modify

### New Files

```
src/cmd/organization.go      # Organization commands
src/cmd/operator.go          # Operator install/upgrade/status
src/cmd/infrastructure.go    # Infrastructure parent + status/topology
src/api/organization.go
src/api/operator.go
src/api/infrastructure.go    # Combines location, node, agent, disk APIs
src/api/storage_class.go
src/api/smart_data_placement.go
src/action/organization.go
src/action/operator.go
src/action/infrastructure.go
src/action/storage_class.go
src/action/smart_data_placement.go
```

### Files to Modify

```
src/cmd/root.go              # Update description, add new commands
src/cmd/swarm.go             # Add location and redundancy-class subcommands
src/cmd/gateway.go           # Add storage-class, smart-data-placement subcommands
src/cmd/tenant.go            # Restructure: nest user, project, add gateway/dns
src/cmd/user.go              # Move under tenant
src/cmd/project.go           # Move under tenant
src/cmd/iam.go               # Add group, api-key subcommands
src/cmd/redundancy_class.go  # Move under swarm, keep rc alias
src/cmd/auth.go              # Add whoami subcommand
src/cmd/node.go              # Move under infrastructure
src/cmd/nexus.go             # Deprecate/remove (becomes infra location)
src/cmd/agent.go             # Move under infrastructure
constants/constants.go       # New API paths
constants/errors.go          # New error messages
```

---

## Part 7: Verification

### Testing Checklist

- [ ] All new commands have unit tests
- [ ] Integration tests for full workflows:
  - [ ] Infrastructure setup (location -> swarm -> rc)
  - [ ] Gateway setup (gateway -> storage-class -> smart-data-placement)
  - [ ] Tenant setup (tenant -> gateway connect -> user)
  - [ ] IAM setup (user -> group -> policy)
  - [ ] Infrastructure management (topology -> agent list -> agent status -> disk status)
- [ ] Legacy binary works independently
- [ ] Help text is accurate and helpful
- [ ] Output formats work (human, json, yaml, xml)
- [ ] Error messages are clear and actionable

### Manual Testing Scenarios

1. **Full onboarding flow**: New org -> install operator -> create swarm -> create gateway -> create tenant -> verify S3 access
2. **Hybrid setup**: Create virtual location -> add S3 hybrid node -> include in swarm
3. **Multi-tenant**: Create multiple tenants -> connect to same gateway -> verify isolation
4. **IAM flow**: Create IAM user -> create group -> attach policy -> verify access
5. **Infrastructure management**: View topology -> list agents -> check agent status -> view disk health
