# Public Dashboard CLI

All the commands has to work both in a interactive and programmatic way (passing all arguments).

## Authentication

Credentials passed through ENV variables overrides the one stored in the configuration file

- `ds3-composer login --email=<email> --password=<password>`: prompts for operator's email and password, asks for config, creates an API key and stores the authentication information in a file called `~/.config/cubbit/credentials`
- `ds3-composer logout`: deletes the file created with `ds3-composer login`
- `ds3-composer config`: prompts (and changes) configuration (endpoint, language, etc.)

## Tenant

Provides all the APIs useful to manage a tenant.

- `ds3-composer tenant list/ls`: lists all the available tenants for the current logged operator
- `ds3-composer tenant create/new --name=<name> --description=<description> --coupon=<coupon>`: creates a new tenant. It prompts for name and description
- `ds3-composer tenant remove/rm <id>/<name>`: removes the tenant with the specified id.
- `ds3-composer tenant report --format<format> <id>/<name>`: creates a report for the tenant
- `ds3-composer tenant describe/info <id>/<name>`: shows information about the tenant with the specified id
- `ds3-composer tenant --name=<name> --id=<id> list-available-swarms`: lists the swarms that can be connected
- `ds3-composer tenant --name=<name> --id=<id> connect-swarm --name=<name> --id=<id>`: connects the swarm with the specified id/name to the tenant with the specified id/name
- `ds3-composer tenant --name=<name> --id=<id> add-operator --email=<email> --permission=<permission>`: adds a new operator to the specified tenant
- `ds3-composer tenant --name=<name> --id=<id> remove-operator --email=<email> --id=<id>`: removes the specified operator
- `ds3-composer tenant --name=<name> --id=<id> list-operators`: lists all the operators for the tenant
- `ds3-composer tenant --name=<name> --id=<id> edit-description <description>`: changes the description
- `ds3-composer tenant --name=<name> --id=<id> edit-image <url>`: changes the image

## Gateway

Provides all the APIs to manage a S3 GW for a tenant.

- `ds3-composer gateway --tenant-name=<name> --tenant-id=<id> describe/info`: shows information about the tenant gateway

## Swarms

Provides all the APIs to create and manage swarms.

- `ds3-composer swarm list/ls`: lists all the available swarms
- `ds3-composer swarm create/new --name=<name> --description=<description>`: creates a new swarm. It prompts for name and description
- `ds3-composer swarm remove/rm <id>/<name>`: removes the swarm with the specified id.
- `ds3-composer swarm report --format<format> <id>/<name>`: creates a report for the swarm
- `ds3-composer swarm describe/info/show <id>`: shows information about the swarm with the specified id, or via `--swarm-id=<id>` / `--swarm-name=<name>`
- `ds3-composer swarm --name=<name> --id=<id> add-operator --email=<email> --permission=<permission>`: adds a new operator to the specified swarm
- `ds3-composer swarm --name=<name> --id=<id> remove-operator --email=<email> --id=<id>`: removes the specified operator
- `ds3-composer swarm --name=<name> --id=<id> list-operators`: lists all the operators
- _`ds3-composer swarm --name=<name> --id=<id> edit-name <name>`: changes the name_
- `ds3-composer swarm --name=<name> --id=<id> edit-description <description>`: changes the description

## Nodes

Provides all the APIs to create and manages nodes for a swarm.

- `ds3-composer node add --swarm-name=<name> --swarm-id=<id> -e <env> --data-volume=<data-volume> --config-volume=<config-volume> --name=<name> --description=<description> --provider-id=<provider-id> --method=docker`: creates a new node without running it, using the specified method. It prints the command to run it
- `ds3-composer node run --swarm-name=<name> --swarm-id=<id> -e <env> --data-volume=<data-volume> --config-volume=<config-volume> --name=<name> --description=<description> --provider-id=<provider-id> --method=docker`: creates a new node running it
- `ds3-composer node list/ls --swarm-name=<name> --swarm-id=<id>`: lists all nodes in the swarm
- `ds3-composer node remove --swarm-name=<name> --swarm-id=<id> <id>`: removes the node with the specified id (works only if the node is suspended)
- _`ds3-composer node suspend --swarm-name=<name> --swarm-id=<id> <id>`: suspends the node with the specified id_
- `ds3-composer node describe --swarm-name=<name> --swarm-id=<id> <id>`: describes the node with the specified id
- `ds3-composer node edit-name --swarm-name=<name> --swarm-id=<id> --id=<id> <name>`: edits the name of a node.
- `ds3-composer node edit-description --swarm-name=<name> --swarm-id=<id> --id=<id> <description>`: edits the description of a node.
- `ds3-composer node edit-provider --swarm-name=<swarm> --swarm-id=<id> --id=<id> <provider-id>`: edits the provider of a node

## Providers

Provides all the APIs to create and manage providers for a swarm.

- `ds3-composer provider create/new --swarm-name=<name> --swarm-id=<id> --name=<name> --email=<email>`: creates a new provider for the given swarm
- `ds3-composer provider list --swarm-name=<name> --swarm-id=<id>`: lists all the providers for the given swarm
- `ds3-composer provider remove <id>`: removes the provider with the given id
- `ds3-composer provider edit-name --name=<name> --id=<id> <name>`: changes the name for the given provider
- `ds3-composer provider edit-email --name=<name> --id=<id> <email>`: changes the email for the given provider

## Distributor

- `ds3-composer distributor report --coupon=<coupon> --format=<format> --output=<output> <distributor-id>`: downloads/prints a full report for the distributor
- `ds3-composer distributor create-coupon --description=<description> --redemption-count=<redemption-count>`: creates a new coupon for the distributor
- `ds3-composer distributor describe-coupon <coupon-id>`: shows information about the coupon
- `ds3-composer distributor remove-coupon --description=<description> --redemption-count=<redemption-count>`: removes a new coupon for the distributor
- `ds3-composer distributor list-coupons`: lists all coupons for the distributor
