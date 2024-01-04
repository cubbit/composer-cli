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
cubbit-operator-cli tenant list --verbose --line
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
cubbit-operator-cli tenant --name <name> --id <id>  list-operators
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
cubbit-operator-cli swarm list --verbose --line
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
cubbit-operator-cli swarm --name <name> --id <id> list-operators
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
cubbit-operator-cli tenant --name <name> --id <id> list-available-swarms
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
cubbit-operator-cli distributor list
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
cubbit-operator-cli distributor --name <name> --id <id>  list-coupons
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

Update account max projects:

```bash
cubbit-operator-cli tenant --name <name> --id <id> account <account-id> edit-max-projects <max-projects>
```

Edit tenant settings:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-settings {settings} 
```

List accounts registered under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account list/ls
```

Create multiple accounts under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account create/new <email1,email2>
```

Describe an account:

```bash
cubbit-operator-cli tenant --name <name> --id <id> account describe <account-id>
```

Delete an account:

```bash
cubbit-operator-cli tenant --name <name> --id <id> account delete <account-id>
```

Edit an account:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account <account-id> edit --first-name <first-name> --last-name <last-name> --internal <true|false> --max-allowed-projects <max-allowed-projects> --endpoint-gateway <endpoint-gateway>
```

Ban an account:
  
```bash
cubbit-operator-cli tenant --name <name> --id <id> account ban <account-id>
```

Unban an account:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account unban <account-id>
```

Restore an account:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account restore <account-id>
```

Invalidate an account sessions:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  account delete-sessions <account-id>
```

Update tenant coupon:

```bash
cubbit-operator-cli tenant --name <name> --id <id>  edit-coupon --coupon-code <coupon-code>
```

Get the operator with at least get access to the tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> operator describe <operator-id>
```

Edit the role of an operator inside a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> operator <operator-id> edit --policy-id <policy-id>
```
  
List projects registered under a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project list/ls
```

Describe a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project describe <project-id>
```

Delete a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project delete <project-id>
```

Edit a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project <project-id> edit --image-url <image-url> --description <description> --name <name>
```

Ban a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project ban <project-id>
```

Unban a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project unban <project-id>
```

Restore a project:

```bash
cubbit-operator-cli tenant --name <name> --id <id> project restore <project-id>
```

List sign-in providers:

```bash
cubbit-operator-cli tenant --name <name> --id <id> sign-in-providers list/ls
```

List swarms associated with a tenant:

```bash
cubbit-operator-cli tenant --name <name> --id <id> swarm list/ls
```

### Swarm Commands

Get the swarm operator:

```bash
cubbit-operator-cli swarm --name <name> --id <id> operator describe <operator-id>
```

Edit the role of a swarm operator:

```bash
cubbit-operator-cli swarm --name <name> --id <id> operator <operator-id> edit --policy-id <policy-id>
```

### Operator Commands

List operators:

```bash
cubbit-operator-cli operator list/ls
```

Update operator email:

```bash
cubbit-operator-cli operator --id <operator-id> --email <email>
```

Get self operator:

```bash
cubbit-operator-cli operator describe 
```

Get other operator:

```bash
cubbit-operator-cli operator describe --id <operator-id>
```

Delete self operator:

```bash
cubbit-operator-cli operator delete
```

Delete other operator:

```bash
cubbit-operator-cli operator delete --id <operator-id>
```

Update self operator credentials:

```bash
cubbit-operator-cli operator edit --password <password>
```

Add TFA to self operator:

```bash
cubbit-operator-cli operator tfa add --codes <codes> --secret <secret> --validation-code <validation-code>
```

Add TFA to other operator:

```bash
cubbit-operator-cli operator --id <operator-id> tfa add --codes <codes> --secret <secret> --validation-code <validation-code>
```

Delete TFA from self operator:

```bash
cubbit-operator-cli operator tfa delete 
```

Delete TFA from other operator:

```bash
cubbit-operator-cli operator --id <operator-id> tfa delete 
```

List operator allowed actions:

```bash
cubbit-operator-cli operator --id <operator-id> list-allowed-actions
```
