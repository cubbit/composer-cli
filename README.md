# Cubbit cli

## Usage

### Signup an operator

```bash
cubbit-operator-cli operator signup --api-server-url <api-server-url> --email <email> --password <password> --first-name <first-name> --last-name <last-name> --secret <secret>
```

> **_NOTE:_** The --api-server-url can either be the backend url or a keyword that identifies the file containing the service endpoint.

Keyword file example

./.env.local

```bash
iam: http://localhost:8181
hive: http://localhost:9151
dash:  http://localhost:3000
```

Interactively

```bash
cubbit-operator-cli operator signup --interactive
```

### Signin an operator

```bash
cubbit-operator-cli signin/login --interactive
```

Interactively

```bash
cubbit-operator-cli signin/login --api-server-url <api-server-url> --email <email> --password <password> --code <code> --profile <profile> --config <config>
```

### sign out an operator

```bash
cubbit-operator-cli logout --profile default
```

### Create a tenant

```bash
cubbit-operator-cli tenant create --name <name> --description <description> --image-url <image-url> --settings <settings>
```

Interactively

```bash
cubbit-operator-cli tenant create --interactive
```

### Describe a tenant

```bash
cubbit-operator-cli tenant describe --name <name> --id <id>  --format <json|csv|semantic>
```

Interactively

```bash
cubbit-operator-cli tenant describe --interactive
```

### List all of the available tenants

```bash
cubbit-operator-cli tenant list --verbose --line --query <query> --sort <sort_key> --page <page> --page-size <page-size>
```

### Delete a tenant

```bash
cubbit-operator-cli tenant remove --name <name> --id <id>  --email <email> --password <password> --code <code>
```

Interactively

```bash
cubbit-operator-cli tenant remove --interactive
```

### Edit the description of a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-description <description>
```

Interactively

```bash
cubbit-operator-cli tenant edit-description --interactive
```

### Edit the image of a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-image <image-url>
```

Interactively

```bash
cubbit-operator-cli tenant edit-image --interactive
```

### Invite an operator to a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id> add-operator --email <email> --first-name <first-name> --last-name <last-name> --role <role>
```

Interactively

```bash
cubbit-operator-cli tenant add-operator --interactive
```

### List operators of a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id>  list-operators --sort <sort_key> --page <page> --page-size <page-size>
```

### Remove an operator of a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id>  remove-operator <operator-id|operator-email>
```

Interactively

```bash
cubbit-operator-cli tenant remove-operator --interactive
```

### Create a swarm

```bash
cubbit-operator-cli swarm create --name <name> --description <description> --configuration <configuration>
```

Interactively

```bash
cubbit-operator-cli swarm create --interactive
```

### Describe a swarm

```bash
cubbit-operator-cli swarm describe --name <name> --id <id> --format <json|csv|semantic>
```

Interactively

```bash
cubbit-operator-cli swarm describe --interactive
```

### List all of the available swarms

```bash
cubbit-operator-cli swarm list --verbose --line --sort <sort_key> --page <page> --page-size <page-size>
```

### Delete a swarm

```bash
cubbit-operator-cli swarm remove --name <name> --id <id>  --email <email> --password <password> --code <code>
```

Interactively

```bash
cubbit-operator-cli swarm remove --interactive
```

### Edit the description of a swarm

```bash
cubbit-operator-cli swarm --name <name> --id <id>  edit-description <description>
```

Interactively

```bash
cubbit-operator-cli swarm edit-description --interactive
```

### Edit the name of a swarm

```bash
cubbit-operator-cli swarm --name <name> --id <id>  edit-name <name>
```

Interactively

```bash
cubbit-operator-cli swarm edit-name --interactive
```

### Invite an operator to a swarm

```bash
cubbit-operator-cli swarm --name <name> --id <id> add-operator --email <email> --first-name <first-name> --last-name <last-name> --role <role>
```

Interactively

```bash
cubbit-operator-cli swarm add-operator --interactive
```

### List operators of a swarm

```bash
cubbit-operator-cli swarm --name <name> --id <id> list-operators --sort <sort_key> --page <page> --page-size <page-size>
```

### Remove an operator of a swarm

```bash
cubbit-operator-cli swarm --name <name> --id <id>  remove-operator <operator-id|operator-email>
```

Interactively

```bash
cubbit-operator-cli swarm remove-operator --interactive
```

### Connect a swarm to a tenant

```bash
cubbit-operator-cli tenant --name <name> --id <id>  connect-swarm <swarm-name|swarm-id>
```

Interactively

```bash
cubbit-operator-cli tenant connect-swarm --interactive
```

### List tenant available swarms

```bash
cubbit-operator-cli tenant --name <name> --id <id> list-available-swarms --sort <sort_key> --page <page> --page-size <page-size>
```

Interactively

```bash
cubbit-operator-cli tenant list-available-swarms --interactive
```

### Create a distributor

```bash
cubbit-operator-cli distributor create/new --name <name> --description <description> --first-name <first-name> --image-url <image-url> --owner <owner> --swarms <swarm-ids> 
```

Interactively

```bash
cubbit-operator-cli distributor create/new --interactive
```

### List distributors

```bash
cubbit-operator-cli distributor list --sort <sort_key> --page <page> --page-size <page-size>
```

Interactively

```bash
cubbit-operator-cli distributor list --interactive
```

### Remove distributor

```bash
cubbit-operator-cli distributor remove --name <name> --id <id>  --email <email> --password <password> --code <code>
```

Interactively

```bash
cubbit-operator-cli distributor remove --interactive
```

### Generate distributors report

```bash
cubbit-operator-cli distributor --name <name> --id <id>  report --coupon <coupon> --format <json|csv|semantic> --from <from> --to <to> --output <output>
```

Interactively

```bash
cubbit-operator-cli distributor report --interactive
```

### Create a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id> create-coupon/new-coupon --coupon-name <coupon-name> --description <description> --redemption-count <redemption-count> --swarms <swarm-ids> --zone <zone>
```

Interactively

```bash
cubbit-operator-cli distributor create-coupon/new-coupon --interactive
```

### Describe a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id> describe-coupon <coupon-id|coupon-name>
```

Interactively

```bash
cubbit-operator-cli distributor describe-coupon --interactive
```

### Edit a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id>  edit-coupon <coupon-id|coupon-name> --coupon-name <new-coupon-name> --description <description> --redemption-count <redemption-count>
```

Interactively

```bash
cubbit-operator-cli distributor edit-coupon --interactive
```

### List distributor coupons

```bash
cubbit-operator-cli distributor --name <name> --id <id>  list-coupons --sort <sort_key> --page <page> --page-size <page-size>
```

Interactively

```bash
cubbit-operator-cli distributor list -coupons--interactive
```

### Remove a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id> remove-coupon <coupon-id|coupon-name>
```

Interactively

```bash
cubbit-operator-cli distributor remove-coupon --interactive
```

### Revoke a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id> revoke-coupon <coupon-id|coupon-name>
```

Interactively

```bash
cubbit-operator-cli distributor revoke-coupon --interactive
```

### Assign a tenant to a distributor coupon

```bash
cubbit-operator-cli distributor --name <name> --id <id>  assign-tenant --tenant-id <tenant-id> --tenant-name <tenant-name> --coupon-code <coupon-code>
```

Interactively

```bash
cubbit-operator-cli distributor assign-tenant --interactive
```

## Build

```bash
go build -o build/cubbit github.com/cubbit/cubbit/client/cli
```

## TODO

### Tenant Commands

Edit tenant settings:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-settings {settings} 
```

List users registered under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> list-users --query <query> --sort <sort_key> --page <page> --page-size <page-size>
```

Create multiple users under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  create-user --emails <email1,email2>
```

Describe a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id> describe-user <user-id>
```

Delete a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id> delete-user <user-id>
```

Edit a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-user <user-id> --first-name <first-name> --last-name <last-name> --internal <true|false> --max-allowed-projects <max-allowed-projects> --endpoint-gateway <endpoint-gateway>
```

Freeze a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id> freeze-user <user-id>
```

Unfreeze a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  unfreeze-user <user-id>
```

Restore a user:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  restore-user <user-id>
```

Invalidate a user sessions:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  delete-user-sessions <user-id>
```

Update tenant coupon:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-coupon --coupon-code <coupon-code>
```

Get the operator with at least get access to the tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> describe-operator <operator-id>
```

Edit the role of an operator inside a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> edit-operator <operator-id> --role <role>
```
  
List projects registered under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> list-projects --query <query> --sort <sort_key> --page <page> --page-size <page-size>
```

Describe a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> describe-project <project-id>
```

Delete a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> delete-project <project-id>
```

Edit a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> edit-project <project-id> --image-url <image-url> --description <description> --name <name>
```

Freeze a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> freeze-project <project-id>
```

Unfreeze a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> unfreeze-project <project-id>
```

Restore a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> restore-project <project-id>
```

List swarms associated with a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> list-swarms --sort <sort_key> --page <page> --page-size <page-size>
```

### Swarm Commands

Get the swarm operator:

```bash
cubbit-operator-cli swarm --name <name> --id <id> describe-operator <operator-id>
```

Edit the role of a swarm operator:

```bash
cubbit-operator-cli swarm --name <name> --id <id> edit-operator <operator-id> edit --role <role>
```

### Generate project report

```bash
cubbit-operator-cli project --name <name> --id <id>  report --format <json|csv|semantic> --from <from> --to <to> --output <output>
```
