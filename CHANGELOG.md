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
