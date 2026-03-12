# IAM (Identity & Access Management)

The Cubbit CLI provides IAM commands to manage operators and their policies across tenants and swarms. This lets you control who can perform which actions on which resources.

---

## Concepts

| Concept | Description |
|---------|-------------|
| **Operator** | A user with management-level access to the DS3 Composer environment |
| **Policy** | A set of permissions attached to an operator, scoped to a tenant or swarm |

Policies can be scoped to a **tenant** (controls access to tenant-level resources like projects and users) or to a **swarm** (controls access to infrastructure resources like nodes and redundancy classes). The `--tenant-id` and `--swarm-id` flags are mutually exclusive.

---

## Operators

```bash
# List IAM operators for a tenant
cubbit iam user list --tenant-id <tenant-id>

# List IAM operators for a swarm
cubbit iam user list --swarm-id <swarm-id>
```

### Inviting a Collaborator

To invite a new operator and assign them a policy in one step:

```bash
# Invite IAM operators for a tenant
cubbit iam user create --email <email> --policy-id <policy-id> --tenant-id <tenant-id>

# Invite IAM operators for a swarm
cubbit iam user create --email <email> --policy-id <policy-id> --swarm-id <swarm-id>
```

The invited user will receive an email to complete their registration.

---

## Next Steps

- [Tenants & Users](tenants-users-and-gateways.md) — manage tenant-level users
