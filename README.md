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

How to sign in securely

```
cubbit operator --interactive signin
```

alternative less secure

```
cubbit operator signin --api-server-url https://api.cubbit.eu --email test@cubbit.io --password 123456 --code 000-000 --name default
```

How to sign out

```
cubbit operator --interactive signout --name default
```

## Build

```
go build -o build/cubbit github.com/cubbit/cubbit/client/cli
```
