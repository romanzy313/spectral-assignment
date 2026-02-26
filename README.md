# spectral-assignment

Language-independent gRPC-based microservice.

## Dependency Setup

I am developing on Ubuntu 24.04 LTS. MacOS installation instructions may differ.

### Golang

- [Install Go](https://go.dev/doc/install). Verify with `go version`. This project uses `1.25.6`, anything higher is okay.
- [Install protobuf compiler](https://protobuf.dev/installation/). Verify version with `protoc --version`. Any version 3+ should be okay.
- Install make, if not already included (`apt install make`).
- Then install development go binaries with `make install-tools`.
- Make sure to add go binary paths to env `export PATH="$PATH:$(go env GOPATH)/bin"`. Check with `protoc-gen-go --version`
- Install go packages `go mod download`

### Node

- [Install Node.js](https://nodejs.org/en/download). I recommend using `nvm`.
- Run `nvm install` to install the correct version of Node.js.
- Run `nvm use` to enable the correct version of Node.js. Check with `node -v`.
- Run `corepack enable pnpm` to activate package manager
- Finally, run `pnpm install` to install dependencies.

## How to Use

### Development

Run these 3 commands in separate shell environments to serve applications with automatic reload on changes:

1. `make dev-server`
3. `make dev-client`
2. `make dev-frontend`

Server is available on port `12000`, client at `12001`, and frontend at `12002`. The frontend can be accessed on [http://localhost:12002](http://localhost:12002).

Generate protobuf files with `make proto`.

### Run Containerized ("Production")

Run everything with `docker compose up --build`. The frontend is available at [http://localhost:3000](http://localhost:3000).

## Design Decisions

First talk about naming. server is `gRPC server`, client is `gRPC client`

TODO: talk about echo, react, and grpc impl.

paste schema here?

For simplicity and performance, timestamps on the backend are integers, not native `time.Time` types.

In real system, the `gRPC Server` should return sensor reading values as `string` rather then `float`. The floating point operations on sensor reading data can introduce rounding errors and precision loss. In the case of this assignment, the values are only used for visualization. Therefore, the precision loss is neglegable. However, other services that need to do financial calcualations, should instead use a `decimal` data type.
