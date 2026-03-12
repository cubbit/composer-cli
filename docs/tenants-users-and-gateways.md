# Tenants, Users & Gateways

This guide covers tenant lifecycle management, project and user administration within a tenant, and S3-compatible gateway setup.

---

## Tenants

A tenant is an isolated organizational unit within a DS3 Composer environment. Each tenant can have its own projects, users, gateways, and access policies.

```bash
# List all tenants
cubbit tenant list

# JSON output for scripting
cubbit tenant list --output json

# Get details of a specific tenant
cubbit tenant describe --tenant-id <tenant-id>

# Create a new tenant
cubbit tenant create --distributor-code <distributor-code> --name <tenant-name>
```

The `--distributor-code` is provided to the operator when they are invited to the service.

---

## Projects

Projects are namespaces within a tenant, typically used to organize storage by team, application, or environment.

```bash
# List projects in a tenant
cubbit tenant project list --tenant-id <tenant-id>
```

---

## Tenant Users

Users are accounts associated with a specific tenant. They can be granted access to projects and resources through IAM policies.

```bash
# List users in a tenant
cubbit tenant user list --tenant-id <tenant-id>
```

For managing operator-level permissions across tenants and swarms, see the [IAM guide](iam.md).

---

## Gateways

Gateways are the S3-compatible entry points to your DS3 swarm. They allow tenants and applications to interact with Cubbit storage using standard S3 APIs and tools.

### Creating a Gateway

```bash
cubbit gateway create --tenant-id <tenant-id> --location <location> --name <gateway-name>
```

### Configuring DNS

Before installing the gateway, configure the DNS domain for the tenant:

```bash
cubbit tenant configure-dns --tenant-id <tenant-id> --domain <domain>
```

### Listing Gateways

```bash
cubbit gateway list --tenant-id <tenant-id>
```

### Installing a Gateway

Once the gateway is created and DNS is configured, run the installation:

```bash
cubbit gateway install --tenant-id <tenant-id> --gateway-id <gateway-id>
```

---

## Next Steps

- [IAM](iam.md) — manage operator access and policies
