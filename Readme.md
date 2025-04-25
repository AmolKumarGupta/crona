# Crona

Crona is an experimental job scheduler written in golang

## Getting Started

To start the server
```sh
./bin/crona
```

To start server with specified config
```sh
./bin/crona -c="path/to/config"
```

## How to build

> [!IMPORTANT]
> Golang 1.23 or above is required to be installed

Clone the git repository
```sh
git clone https://github.com/AmolKumarGupta/crona.git
```

Run build command
```sh
make build
```

To start the server
```sh
make run
```

## Testing

```sh
go test ./...
```