# Cubbit cli


## Usage

How to create an operator

```
cubbit operator signup --api-server-url https://api.cubbit.eu --email test@cubbit.io --password 123456 --first-name gigi --last-name esposito
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
tenant remove --id a37b4 --name cubbit --email test@cubbit.io --password 123456 --code 000-000
```

How to get the description of a tenant

```
tenant describe --id a37b4 --name cubbit --format default
```

## Build

```
go build -o build/cubbit github.com/cubbit/cubbit/client/cli
```
