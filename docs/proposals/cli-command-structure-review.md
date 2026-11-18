# CLI Command Structure Review - Cubbit Composer CLI

## Executive Summary

I've analyzed our current CLI command structure against 6 UX guidelines. Our CLI manages DS3 Composer infrastructure (geo-distributed object storage) with two main resource hierarchies: **Swarm** (infrastructure) and **Tenant** (customer/user management).

---

## Current Command Tree

```
cubbit
├── auth
│   ├── login (signin)
│   └── logout
├── config
│   ├── init
│   ├── view
│   ├── edit
│   ├── profiles
│   ├── switch-profile
│   └── validate
├── docs
│   ├── markdown, man, rst, yaml
│   └── tree
├── swarm
│   ├── create (new), list (ls), describe, edit, remove (rm), status
├── nexus [--swarm-id required]
│   ├── create, list, describe, edit, remove, deploy
├── node [--swarm-id, --nexus-id required]
│   ├── create, list, describe, edit, remove, deploy
├── agent [--swarm-id, --nexus-id, --node-id required]
│   ├── create, list, describe, edit, remove, status
├── rc [--swarm-id required]
│   ├── create, list, describe, status, expand
│   └── recovery
│       ├── start
│       └── status
├── tenant
│   ├── create, list (ls), describe, edit, remove (rm)
│   ├── report, configure-dns, verify-dns
├── gateway [--tenant-id required]
│   ├── create, list, describe, edit, remove, install
├── project [--tenant-id required]
│   ├── list, describe, edit, remove, restore, freeze, unfreeze
├── user [--tenant-id required]
│   ├── create, list, describe, edit, remove, restore
│   ├── freeze, unfreeze, force-logout
└── iam
    ├── user
    │   ├── create, list, describe, remove
    └── policy
        ├── promote, edit
```

---

## Detailed Review by Guideline

### 1. Clear and Concise Command Names

#### ✅ What Works Well
- Resource names are clear nouns: `swarm`, `tenant`, `gateway`, `node`
- Action verbs are consistent: `create`, `list`, `describe`, `edit`, `remove`
- Aliases improve usability: `ls` for `list`, `rm` for `remove`

#### ❌ Issues Found

| Current | Problem | Recommendation |
|---------|---------|----------------|
| `rc` | Abbreviation unclear to new users | Rename to `redundancy-class` (keep `rc` as alias) |
| `auth login` | OK but `signin` alias is redundant | Keep `login` only |
| `tenant configure-dns` | Verb-noun inconsistent with other commands | Consider `tenant dns configure` or keep as-is |
| `tenant verify-dns` | Same pattern issue | Consider `tenant dns verify` |
| `user force-logout` | Hyphenated action | Consider `user logout --force` or keep as-is |
| Root description: "for operators" | Vague | "Manage DS3 Composer infrastructure and tenants" |

#### ❌ Help Messages Need Improvement

Many descriptions are generic:
- `"Execute commands in swarm sections"` → Should be: `"Manage distributed storage swarms"`
- `"Execute commands in nexus sections"` → Should be: `"Manage nexus nodes within a swarm"`
- `"Execute commands in agent sections"` → Should be: `"Manage storage agents on nodes"`

---

### 2. Logical Command Grouping

#### ❌ Major Issue: Flat Hierarchy Requires Too Many Parent IDs

Current structure requires all ancestor IDs as flags:
```bash
# Current - requires full hierarchy
cubbit node list --swarm-id X --nexus-id Y
cubbit agent create --swarm-id X --nexus-id Y --node-id Z --agent-port 8080
```

**Problem**: Since all IDs are globally unique in the platform, requiring the full ancestor chain is unnecessary and verbose.

#### Recommendation

Since IDs are globally unique, we only need the immediate parent for `list`/`create` operations:
```bash
cubbit nexus list --swarm-id X
cubbit node list --nexus-id Y
cubbit agent list --node-id Z
cubbit agent create --node-id Z --port 8080
```

For single-resource operations (`describe`, `edit`, `remove`), use direct `--id` access:
```bash
cubbit agent describe --id jkl012
cubbit node edit --id ghi789
```

This dramatically simplifies commands while maintaining clear parent-child relationships.

#### ❌ Inconsistent Grouping

| Resource | Parent Command | Notes |
|----------|---------------|-------|
| `gateway` | Top-level | Requires `--tenant-id` |
| `project` | Top-level | Requires `--tenant-id` |
| `user` | Top-level | Requires `--tenant-id` |
| `iam user` | Under `iam` | Also requires `--tenant-id` or `--swarm-id` |

**Problem**: `user` (tenant users) vs `iam user` (operators) are confusing.

**Recommendation**:
```
tenant
├── user (end-users)
├── project
└── gateway

iam
└── operator (renamed from user)
```

---

### 3. Consistent Naming Conventions

#### ❌ Flag Naming Inconsistencies

| Pattern | Examples | Issue |
|---------|----------|-------|
| `--{resource}-id` | `--swarm-id`, `--tenant-id`, `--nexus-id` | ✅ Consistent |
| `--agent-*` prefix | `--agent-port`, `--agent-disk`, `--agent-mount-point` | ❌ Unnecessary prefix when under agent command |
| `--rc-id` | Redundancy class ID | ❌ Should be `--redundancy-class-id` if we rename the command |

**Recommendation**: Remove redundant prefixes:
```bash
# Current
cubbit agent create --agent-port 8080 --agent-disk /dev/sda

# Proposed
cubbit agent create --port 8080 --disk /dev/sda
```

#### ❌ Alias Bugs

```go
// In user.go - WRONG aliases
restoreUserSubCmd.Aliases = []string{"rm"}  // Should NOT be rm
unfreezeUserSubCmd.Aliases = []string{"rm"} // Should NOT be rm
```

These are bugs - `restore` and `unfreeze` should not have `rm` as aliases.

#### ❌ Inconsistent Action Names

| Action | Commands Using It | Notes |
|--------|-------------------|-------|
| `freeze/unfreeze` | project, user | For suspending resources |
| `ban/unban` | None | Could be clearer |
| `remove` | All | For deletion |
| `restore` | project, user | For recovering deleted resources |

**Recommendation**: Standardize on `suspend`/`resume` or `freeze`/`unfreeze` consistently.

---

### 4. Descriptive Help Messages

#### ❌ Current Help Messages Are Too Generic

```
swarm       Execute commands in swarm sections
nexus       Execute commands in nexus sections
node        Execute commands in node sections
agent       Execute commands in agent sections
tenant      Execute commands in tenant sections
```

#### Recommendation: Domain-Specific Descriptions

```
swarm       Manage geo-distributed storage swarms
nexus       Manage data centers (nexuses) within a swarm
node        Manage physical/virtual servers within a nexus
agent       Manage storage agents running on nodes
tenant      Manage customer tenants and their resources
gateway     Manage S3-compatible gateway deployments
rc          Manage redundancy classes for data protection
```

#### Subcommand Descriptions Should Explain Purpose

```bash
# Current
cubbit rc create --help
# "Create a new redundancy class"

# Should explain what it does
# "Create a redundancy class defining how data is distributed across nexuses
#  for fault tolerance. Requires specifying inner/outer erasure coding parameters."
```

---

### 5. Avoid Deep Nesting of Resource IDs

#### ❌ Critical Issue: Agent Commands Require 4 IDs

```bash
# Current - worst case
cubbit agent describe \
  --swarm-id abc123 \
  --nexus-id def456 \
  --node-id ghi789 \
  --agent-id jkl012
```

This is extremely verbose and error-prone.

#### Recommendation

Since all IDs are globally unique in our platform, we can simplify dramatically:

**For `list`/`create` operations** - require only the immediate parent ID:
```bash
cubbit nexus list --swarm-id X
cubbit node list --nexus-id Y
cubbit agent list --node-id Z
cubbit agent create --node-id Z --port 8080
```

**For single-resource operations** - use direct `--id` access:
```bash
cubbit agent describe --id jkl012
cubbit agent edit --id jkl012
cubbit agent remove --id jkl012
cubbit node describe --id ghi789
cubbit nexus describe --id def456
```

#### Improvement Summary

| Command | Current | Proposed | Improvement |
|---------|---------|----------|-------------|
| `nexus list` | 1 (swarm-id) | 1 (swarm-id) | Same |
| `node list` | 2 (swarm-id, nexus-id) | 1 (nexus-id) | -1 ID |
| `agent list` | 3 (swarm-id, nexus-id, node-id) | 1 (node-id) | -2 IDs |
| `agent describe` | 4 (all + agent-id) | 1 (id) | -3 IDs |
| `gateway list` | 1 (tenant-id) | 1 (tenant-id) | Same |
| `project list` | 1 (tenant-id) | 1 (tenant-id) | Same |

---

### 6. Support LTS APIs While Providing New APIs

#### Current State

The codebase uses both v1 and v2 API endpoints:
- `constants.go`: Defines both `/v1/` and `/v2/` paths
- Different resources use different API versions

#### ❌ No Version Flag or Version Switching

Users cannot:
- Choose which API version to use
- See which API version is being used
- Gradually migrate from v1 to v2

#### Recommendation: API Version Awareness

```bash
# Global flag to force API version
cubbit --api-version v1 tenant list
cubbit --api-version v2 tenant list

# Or per-profile setting in config
[profile.legacy]
api_version = "v1"

[profile.modern]
api_version = "v2"
```

#### Recommendation: Deprecation Warnings

When v1 endpoints are used, show warnings:
```
Warning: Using deprecated v1 API. Consider upgrading to v2.
See: https://docs.cubbit.io/migration-guide
```

---

## Priority Recommendations

### High Priority (UX Blockers)

1. **Fix alias bugs** - `restore` and `unfreeze` have wrong `rm` alias
2. **Reduce agent command complexity** - Allow direct `--id` access
3. **Improve help descriptions** - Replace generic "Execute commands in X sections"

### Medium Priority (Usability)

4. **Rename `rc` to `redundancy-class`** - Keep `rc` as alias
5. **Remove `--agent-*` prefixes** - Simplify to `--port`, `--disk`, etc.
6. **Rename `iam user` to `iam operator`** - Distinguish from tenant users
7. **Move gateway/project/user under tenant** - Better logical grouping

### Low Priority (Polish)

8. **Add `--api-version` flag** - For API version control
9. **Group DNS commands** - `tenant dns configure/verify` instead of `configure-dns`
10. **Consistent freeze/unfreeze** - Ensure all suspendable resources use same pattern

---

## Proposed Improved Command Tree

```
cubbit
├── auth
│   ├── login
│   └── logout
├── config
│   ├── init, view, edit, validate
│   └── profile list|switch|create|remove
├── docs [...]
│
├── swarm
│   ├── create, list, describe --id, edit --id, remove --id, status --id
│
├── nexus [--swarm-id for list/create, --id for single operations]
│   ├── create --swarm-id X
│   ├── list --swarm-id X
│   ├── describe --id Y
│   ├── edit --id Y
│   ├── remove --id Y
│   └── deploy --id Y
│
├── node [--nexus-id for list/create, --id for single operations]
│   ├── create --nexus-id Y
│   ├── list --nexus-id Y
│   ├── describe --id Z
│   ├── edit --id Z
│   ├── remove --id Z
│   └── deploy --id Z
│
├── agent [--node-id for list/create, --id for single operations]
│   ├── create --node-id Z
│   ├── list --node-id Z
│   ├── describe --id W
│   ├── edit --id W
│   ├── remove --id W
│   └── status --id W
│
├── redundancy-class (alias: rc) [--swarm-id for list/create]
│   ├── create, list, describe, status, expand
│   └── recovery start|status
│
├── tenant
│   ├── create, list, describe, edit, remove
│   ├── report
│   ├── dns configure|verify            # Grouped DNS commands
│   ├── user [...]                      # Moved under tenant
│   ├── project [...]                   # Moved under tenant
│   └── gateway [...]                   # Moved under tenant
│
└── iam
    ├── operator [...]                  # Renamed from user
    └── policy [...]
```

---

## Summary

I've identified **10 key improvements** across the 6 UX guidelines:

| Priority | Issue | Recommendation |
|----------|-------|----------------|
| 🔴 High | Alias bugs (`rm` on restore/unfreeze) | Fix immediately |
| 🔴 High | Commands require full ancestor chain | Immediate parent ID only + direct `--id` for single ops |
| 🔴 High | Generic help messages | Add domain-specific descriptions |
| 🟡 Medium | `rc` abbreviation unclear | Rename to `redundancy-class` |
| 🟡 Medium | `--agent-*` redundant prefixes | Simplify to `--port`, `--disk` |
| 🟡 Medium | `iam user` vs `user` confusion | Rename to `iam operator` |
| 🟡 Medium | Flat structure for nested resources | Move gateway/project/user under tenant |
| 🟢 Low | No API version control | Add `--api-version` flag |
| 🟢 Low | DNS commands naming | Consider `tenant dns configure/verify` |
| 🟢 Low | Inconsistent freeze/unfreeze | Standardize across resources |

**My Preferences for Implementation:**
- Breaking changes are acceptable for significant UX improvements
- Immediate parent ID for `list`/`create` operations (e.g., `cubbit agent list --node-id Y`)
- Direct `--id` access for single-resource operations (e.g., `cubbit agent describe --id Z`)
