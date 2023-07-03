# Public Dashboard CLI

All the commands has to work both in a interactive and programmatic way (passing all arguments).

## Authentication

Credentials passed through ENV variables overrides the one stored in the configuration file

- `cubbit login --email=<email> --password=<password>`: prompts for operator's email and password, asks for config, creates an API key and stores the authentication information in a file called `~/.config/cubbit/credentials`
- `cubbit logout`: deletes the file created with `cubbit login`
- `cubbit config`: prompts (and changes) configuration (endpoint, language, etc.)

## Tenant

Provides all the APIs useful to manage a tenant.

- `cubbit tenant list/ls`: lists all the available tenants for the current logged operator
- `cubbit tenant create/new --name=<name> --description=<description> --coupon=<coupon>`: creates a new tenant. It prompts for name and description
- `cubbit tenant remove/rm <id>/<name>`: removes the tenant with the specified id.
- `cubbit tenant report --format<format> <id>/<name>`: creates a report for the tenant
- `cubbit tenant describe/info <id>/<name>`: shows information about the tenant with the specified id
- `cubbit tenant --name=<name> --id=<id> list-available-swarms`: lists the swarms that can be connected
- `cubbit tenant --name=<name> --id=<id> connect-swarm --name=<name> --id=<id>`: connects the swarm with the specified id/name to the tenant with the specified id/name
- `cubbit tenant --name=<name> --id=<id> add-operator --email=<email> --permission=<permission>`: adds a new operator to the specified tenant
- `cubbit tenant --name=<name> --id=<id> remove-operator --email=<email> --id=<id>`: removes the specified operator
- `cubbit tenant --name=<name> --id=<id> list-operators`: lists all the operators for the tenant
- `cubbit tenant --name=<name> --id=<id> edit-description <description>`: changes the description
- `cubbit tenant --name=<name> --id=<id> edit-image <url>`: changes the image

## Gateway

Provides all the APIs to manage a S3 GW for a tenant.

- `cubbit gateway --tenant-name=<name> --tenant-id=<id> describe/info`: shows information about the tenant gateway

## Swarms

Provides all the APIs to create and manage swarms.

- `cubbit swarm list/ls`: lists all the available swarms
- `cubbit swarm create/new --name=<name> --description=<description>`: creates a new swarm. It prompts for name and description
- `cubbit swarm remove/rm <id>/<name>`: removes the swarm with the specified id.
- `cubbit swarm report --format<format> <id>/<name>`: creates a report for the swarm
- `cubbit swarm describe/info <id>/<name>`: shows information about the swarm with the specified id
- `cubbit swarm --name=<name> --id=<id> add-operator --email=<email> --permission=<permission>`: adds a new operator to the specified swarm
- `cubbit swarm --name=<name> --id=<id> remove-operator --email=<email> --id=<id>`: removes the specified operator
- `cubbit swarm --name=<name> --id=<id> list-operators`: lists all the operators
- _`cubbit swarm --name=<name> --id=<id> edit-name <name>`: changes the name_
- `cubbit swarm --name=<name> --id=<id> edit-description <description>`: changes the description

## Nodes

Provides all the APIs to create and manages nodes for a swarm.

- `cubbit node add --swarm-name=<name> --swarm-id=<id> -e <env> --data-volume=<data-volume> --config-volume=<config-volume> --name=<name> --description=<description> --provider-id=<provider-id> --method=docker`: creates a new node without running it, using the specified method. It prints the command to run it
- `cubbit node run --swarm-name=<name> --swarm-id=<id> -e <env> --data-volume=<data-volume> --config-volume=<config-volume> --name=<name> --description=<description> --provider-id=<provider-id> --method=docker`: creates a new node running it
- `cubbit node list/ls --swarm-name=<name> --swarm-id=<id>`: lists all nodes in the swarm
- `cubbit node remove --swarm-name=<name> --swarm-id=<id> <id>`: removes the node with the specified id (works only if the node is suspended)
- _`cubbit node suspend --swarm-name=<name> --swarm-id=<id> <id>`: suspends the node with the specified id_
- `cubbit node describe --swarm-name=<name> --swarm-id=<id> <id>`: describes the node with the specified id
- `cubbit node edit-name --swarm-name=<name> --swarm-id=<id> --id=<id> <name>`: edits the name of a node.
- `cubbit node edit-description --swarm-name=<name> --swarm-id=<id> --id=<id> <description>`: edits the description of a node.
- `cubbit node edit-provider --swarm-name=<swarm> --swarm-id=<id> --id=<id> <provider-id>`: edits the provider of a node

## Providers

Provides all the APIs to create and manage providers for a swarm.

- `cubbit provider create/new --swarm-name=<name> --swarm-id=<id> --name=<name> --email=<email>`: creates a new provider for the given swarm
- `cubbit provider list --swarm-name=<name> --swarm-id=<id>`: lists all the providers for the given swarm
- `cubbit provider remove <id>`: removes the provider with the given id
- `cubbit provider edit-name --name=<name> --id=<id> <name>`: changes the name for the given provider
- `cubbit provider edit-email --name=<name> --id=<id> <email>`: changes the email for the given provider

## Distributor

- `cubbit distributor report --coupon=<coupon> --format=<format> --output=<output> <distributor-id>`: downloads/prints a full report for the distributor
- `cubbit distributor create-coupon --description=<description> --redemption-count=<redemption-count>`: creates a new coupon for the distributor
- `cubbit distributor describe-coupon <coupon-id>`: shows information about the coupon
- `cubbit distributor remove-coupon --description=<description> --redemption-count=<redemption-count>`: removes a new coupon for the distributor
- `cubbit distributor list-coupons`: lists all coupons for the distributor
