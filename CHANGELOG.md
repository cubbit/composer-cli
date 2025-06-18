# Changelog

## 0.20.0

### Minor Changes

- e495bbf: Introduced an interactive command for agent creation.

  Added a command to generate deployment files for agents, targeting a specific node or the Nexus when no node is specified.

  Enhanced the create node batch command with support for a deployment file generation flag.

- 408aa9f: used cobra's built-in tooling to generate documentation in multiple formats, including YAML, man pages, Markdown, and RST.
  introduced a command to print the full command tree along with their associated flags.

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.19.4](https://github.com/cubbit/cubbit/compare/cli-v0.19.0...cli-v0.19.4) (2025-06-11)

### Bug Fixes

- **cli:** fix the version bump ([#14023](https://github.com/cubbit/cubbit/issues/14023)) ([15034c3](https://github.com/cubbit/cubbit/commit/15034c338bd66968a3b273d0667b5b0411944f89))
- **cli:** update changelog generation ([#14017](https://github.com/cubbit/cubbit/issues/14017)) ([d187175](https://github.com/cubbit/cubbit/commit/d187175c8347c936581c186ebc578b9a625ff897))

### [0.19.3](https://github.com/cubbit/cubbit/compare/cli-v0.19.0...cli-v0.19.3) (2025-05-30)

### Features

- **composer-hub:** add more logs to ring creation ([#14316](https://github.com/cubbit/cubbit/issues/14316)) ([c7b7be7](https://github.com/cubbit/cubbit/commit/c7b7be7adc636da69dbbabf9ea829e1ea08476a8))
- **gods3:** make GODS3 fetch keys from Secret File and test the Diplomat flow in 'new_db_schema' e2e tests THOR-1754 ([#14465](https://github.com/cubbit/cubbit/issues/14465)) ([cea6386](https://github.com/cubbit/cubbit/commit/cea638671ed31ca15dd5c2c350578abf5324ed8a))
- **repo:** fix go mods ([c6cbc9d](https://github.com/cubbit/cubbit/commit/c6cbc9da1c59de982794a2859e8c7e541323e6fa))

### Bug Fixes

- **cli:** fix the version bump ([#14023](https://github.com/cubbit/cubbit/issues/14023)) ([15034c3](https://github.com/cubbit/cubbit/commit/15034c338bd66968a3b273d0667b5b0411944f89))
- **cli:** update changelog generation ([#14017](https://github.com/cubbit/cubbit/issues/14017)) ([d187175](https://github.com/cubbit/cubbit/commit/d187175c8347c936581c186ebc578b9a625ff897))
- **composer-hub:** enforce n+k nexus over ring creation ([#14323](https://github.com/cubbit/cubbit/issues/14323)) ([2ae82c5](https://github.com/cubbit/cubbit/commit/2ae82c51a559fab4d675eed1de37544c53517271))
- **network:** e2e tests offloader ([#15167](https://github.com/cubbit/cubbit/issues/15167)) ([d422370](https://github.com/cubbit/cubbit/commit/d422370e7106bb52df4ceb6616e802943569f207))
- **repo:** fix repo problems ([#14447](https://github.com/cubbit/cubbit/issues/14447)) ([c5f723e](https://github.com/cubbit/cubbit/commit/c5f723ec358f23c779e57f81bf756ec19a58eb5e))

### [0.19.2](https://github.com/cubbit/cubbit/compare/cli-v0.19.0...cli-v0.19.2) (2024-12-19)

### Bug Fixes

- **cli:** fix the version bump ([#14023](https://github.com/cubbit/cubbit/issues/14023)) ([15034c3](https://github.com/cubbit/cubbit/commit/15034c338bd66968a3b273d0667b5b0411944f89))
- **cli:** update changelog generation ([#14017](https://github.com/cubbit/cubbit/issues/14017)) ([d187175](https://github.com/cubbit/cubbit/commit/d187175c8347c936581c186ebc578b9a625ff897))

### [0.19.1](https://github.com/cubbit/cubbit/compare/cli-v0.19.0...cli-v0.19.1) (2024-11-14)

### Bug Fixes

- **cli:** fix the version bump ([#14023](https://github.com/cubbit/cubbit/issues/14023)) ([15034c3](https://github.com/cubbit/cubbit/commit/15034c338bd66968a3b273d0667b5b0411944f89))
- **cli:** update changelog generation ([#14017](https://github.com/cubbit/cubbit/issues/14017)) ([d187175](https://github.com/cubbit/cubbit/commit/d187175c8347c936581c186ebc578b9a625ff897))

## 0.19.0 (2024-03-22)

### Feat

- **cli**: manage redundancy classes and rings

## 0.18.0 (2024-03-15)

### Feat

- **cli**: manage swarm nodes

## 0.17.0 (2024-03-15)

### Feat

- **cli**: manage nexuses
- **cli**: add a custom server to handle token requests

## 0.16.1 (2024-03-11)

### Fix

- **cli**: update regex for the query filter

## 0.16.0 (2024-02-22)

### Feat

- **cli**: describe/edit a swarm collaborator

## 0.15.0 (2024-02-21)

### Feat

- **cli**: manage tenant projects

## 0.14.0 (2024-02-20)

### Feat

- **cli**: create and edit accounts in a tenant

### Fix

- **cli**: build

## 0.13.0 (2024-02-09)

### Feat

- **cli**: add tenant account cmds (#81)
- **cli**: configure bazel build and update workflows
- **project**: create projects
- **token**: move access token under operator
- **session**: avoid using wrong session for request
- **session**: improve session type
- **account**: introduce account signup, signin, signout

### Fix

- **cli**: update action permissions (#82)
- **account**: disable account and project section
- **account**: wrong endpoint called

## 0.12.0 (2023-10-06)

### Feat

- **cli**: update tenant coupon

### Fix

- **cli**: fix pagination

## 0.11.1 (2023-10-02)

### Fix

- **cli**: fix zone selection logic
- fix zone checking values

## 0.11.0 (2023-09-27)

### Feat

- **cli**: add access token generation cmd

## 0.10.0 (2023-09-22)

### Feat

- **cli**: add distributor report cmd

### Fix

- fix returned value when seletion is optional
- fix returned errors
- fix the cursor in multiple choice display

## 0.9.0 (2023-09-12)

### Feat

- **cli**: edit, revoke and delete coupon codes

### Fix

- fix nil pointer in distributor list

## 0.8.0 (2023-09-07)

### Feat

- **cli**: create, describe and list distributor coupons

### Fix

- unmark unrequired flags and fix redemption count default value

## 0.7.0 (2023-09-06)

### Feat

- **cli**: add aliases for distributor cmds
- **cli**: create, list and remove distributors

### Fix

- **bug**: fix bug in listing swarm operators and description

## 0.6.0 (2023-08-28)

### Feat

- **tenant**: connect a swarm to a tenant

## 0.5.0 (2023-08-28)

### Feat

- **operator**: add, remove and list operators for a swarm/tenant
- **tenant**: add the operator invite

## 0.4.0 (2023-08-24)

### Feat

- **swarm**: add edit-description, edit-name and remove swarm cmds

## 0.3.0 (2023-08-22)

### Feat

- **cli**: update ui components and add interactive swarm commands
- **cli**: add interative tenant deletion
- **cli**: add interative tenant creation & update tenant list
- **cli**: update interactive login
- **cli**: update interactive operator signup
- **cli**: create customized inputs
- **cli**: update string prints
- **cli**: integrate BubblTea & test animation for requests

### Fix

- **urls**: update default uri values for api endpoint

## 0.2.0 (2023-08-16)

### Feat

- **ci**: update the action to be triggered manually and on release
- **ci**: add upload artifacts to s3 action
- **ci**: add upload artifacts to s3 action
- **ci**: update build ci
- **ci**: add build and upload yaml

## 0.1.0 (2023-08-09)

### Feat

- **ci**: add bump version github action
- **ci**: update bump version github action
- **ci**: add bump version github action
- **operator**: add secret to create operator endpoint
- **tenant**: list available swarms
- **tenant**: list available swarms
- **.gitignore**: implemented .gitignore
- **url**: url struct
- **tenant**: edit image and description functions
- **tenant**: edit description and image commands implemented
- **readme**: tenant remove and tenant description implemented
- **readme**: updated description
- **tenant**: remove tenant function
- **tenant**: tenant remove logic
- **tenant**: csv formatting
- **tenant**: gives information abpout tenant
- **tenant**: remove tenants command
- **tenant**: actions.go
- **tenant**: list and describe commands with flags
- **tenant**: formatting tenant description
- **tenant**: gives a list of available tenants with their infomation
- **tenant**: lists all the tenants
- **tenant**: add list flag
- **tenant**: handling invalid image url
- **tenant**: errors are handled
- **tenant**: creates new tenants
- **tenant**: perform api request for creating tenant
- **tenant**: interface implementation
- initialize cli to current state o the art

### Fix

- **tenant**: wrong paarmeters
- **tenant**: deleted usless print
- **tenant**: return error in get tenant by name function
- **tenant**: get tenant by name function fixed
- **urls**: fixed return values
- **urls**: added function in create operator
- **urls**: readme and refactored urls function
- **tenant**: tenant action
- **tenant**: solved conflicts
- **readme**: added description to tenant list
- **tenant**: access and refresh tokens generation
- **tenant**: csv formatting
- **tenant**: format tenant function
- **tenant**: tenant describe alias and errors
- **operator**: corrected error
- **operator**: deleted print password
- **input**: deleted spaces before and after inputs
- **tenant**: fixed the accessToken print
- **main**: remove unused code

### Refactor

- add a formatter function and refactor prints
- centralize repeated error strings in a single file
- update naming and move shared constants
- **actions**: split actions.go into separate files
- **tenant**: list availabe swarms is more readable
- **urls**: implemented api server url configuration function
- **urls**: implemented multiple urls in a function
- **api-urls**: refactor for iam url
- **tenant**: specific iam url
- **config**: code is more readable and errors are shown
- **tenant**: moved logic
- **main**: improve code readability
