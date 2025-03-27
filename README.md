# Sample API

[![CI](https://github.com/henok321/sample-api/actions/workflows/CI.yml/badge.svg)](https://github.com/henok321/sample-api/actions/workflows/CI.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=henok321_sample-api&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=henok321_sample-api)

## Prerequisites

Ensure the following dependencies are installed:

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [pre-commit](https://pre-commit.com/) (`pip install pre-commit`)
- [Goose](https://github.com/pressly/goose) (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

## Setup and Development

### Run Setup

Execute the following command to set up the project:

```sh
make setup
```

This command will:

- Install commit hooks.
- Start the local database.
- Run database migrations.
- Create a `.env` file with necessary environment variables.

Reset database:

```shell
make reset
```

### Start the Application

To run the application locally:

```shell
set -o allexport
source .env
set +o allexport
go run cmd/main.go
```

### Build and run binary

#### Build

```shell
make build
```

#### Run

```shell
set -o allexport
source .env
set +o allexport
./bin/sample-api
```

### Health Check

Verify the service is running:

```sh
curl http://localhost:8080/health
```

### Makefile targets

For more information on available Makefile targets, run:

```shell
make help
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
