# Cubbit cli


## Usage

How to create an operator

```
cubbit operator signup --api-server-url https://api.cubbit.eu --email test@cubbit.io --password 123456 --first-name gigi --last-name esposito
//--api-server-url can either be the backend url or a keyword that identifies the file containing the service endpoint 
```

Keyword file example

```
./.env.local
iam: http://localhost:8181
hive: http://localhost:9151
```

alternative

```
cubbit operator --interactive signup
```

How to log in securely

```
cubbit login --interactive
```

alternative less secure

```
cubbit login --api-server-url https://api.cubbit.eu --email test@cubbit.io --password 123456 --code 000-000 --profile default
```

How to sign out

```
cubbit logout --profile default
```

How to create a tenant

```
tenant create --name cubbit --description "The Cubbit tenant" --image-url https://image.png --settings "{\"test\": 42}"
```

How to list all of the available tenants

```
tenant list --verbose --line
```

How to delete a tenant

```
tenant --id a37b4 --name cubbit remove --email test@cubbit.io --password 123456 --code 000-000
```

How to get the description of a tenant

```
tenant --id a37b4 --name cubbit describe --format default
```

How to change the description of a tenant

```
tenant --id a37b4 --name cubbit edit-description
```

How to change the image of a tenant

```
tenant --id a37b4 --name cubbit edit-image
```
## Build

```
go build -o build/cubbit github.com/cubbit/cubbit/client/cli
```
