# spectral-assignment

An language-independent gRPC-based microservice.

First, I’d like to clarify the naming. I have shortened the names for clarity. What is referred to in the task document as `gRPC-server` is `server`, `gRPC-client` is `client`, and `Front-End` is `frontend`. I hope this is not too confusing.

## Development Setup

I am developing on Ubuntu 25.10. Some commands are for Ubuntu systems. macOS commands may differ a bit.

### Golang

- [Install Go](https://go.dev/doc/install). Verify with `go version`. This project uses `1.25.6`; anything higher is okay.
- [Install protobuf compiler](https://protobuf.dev/installation/). Verify version with `protoc --version`. Any version 3+ should be okay.
- Install make, if not already installed (e.g. `apt install make`).
- Install the development go binaries with `make install-tools`.
- Make sure that go binary paths are accessible by `protoc` (e.g. `export PATH="$PATH:$(go env GOPATH)/bin"`). Check with `protoc-gen-go --version`
- Install go packages `go mod download`

### Node

- [Install Node.js](https://nodejs.org/en/download). Use `nvm` install option.
- Run `nvm install` to install the correct version of Node.js.
- Run `nvm use` to enable the project version of Node.js. Check with `node -v`.
- Run `corepack enable pnpm` to activate package manager.
- Finally, run `pnpm install` to install dependencies.

### Other

- [Install Docker](https://docs.docker.com/engine/install/)
- Install playwright dependencies with `sudo pnpm exec playwright install-deps`

## How to Use

### Development

Run these 3 commands at the same time to run the applications with automatic reload on changes: `make dev-server`, `make dev-client`, `make dev-frontend`.

Server is available on port `12000`, client at `12001`, and frontend at `12002`. The frontend is served at [http://localhost:12002](http://localhost:12002).

If the protobuf schema changes, generate new files with `make proto`.

### Run Containerized ("Production")

Run everything with `docker compose up --build` or simply `make run-docker`. The frontend is served at [http://localhost:3000](http://localhost:3000).

### Testing

Run unit tests with `make test`. This will both run golang and JavaScript tests. Coverage is printed out to the console.

To run end-to-end tests, first start the Docker Compose with `make run-docker`. Then execute `make test-e2e` or `make test-e2e-ui` for headless and UI tests, respectively.

## Design Decisions

There were a couple TODO

### Data Modeling

I have chosen to implement cursor-based pagination for sensor data.

```proto
message GetPageResponse {
  optional int64 next_cursor = 1;
  repeated SensorData data = 2;
}
```

Cursor pagination is a perfect choice when time-series data is involved, because previous meter readings remain unchanged and new records are always added at the next unique timestamp. The data is retrieved in chunks of up-to 10000 datapoints, trading round-trip chattiness for smaller memory consumption on the backend services. This is a simple, robust, and well-known method of sending large amounts of data. The complexity of reconstructing the full data set is shifted to the API consumer.

As for the representation of sensor data, I went with the following schema:

```proto
message SensorData {
  int64 timestamp = 1;
  double value = 2;
}
```

I’ve decided to use Unix timestamps to minimize data transfer and reduce conversion overhead compared to a "native" protobuf datatype, `google.protobuf.Timestamp`. For sensor reading, I went with using a floating-point data type instead of the strings. Since this data is used purely for visualization, a number is better for performance and simplicity.

However, if the data from the `server` is used in financial calculations, it should be represented as a `string` rather than a `double`. Then it will be the consumer's task to perform appropriate decimal conversions and calculations.

### Project Structure and Architecture

The `gRPC-Server` in `./pkg/server` uses a separation-of-concerns architecture. Its sub-packages are `model`, `repository`, and `service`. This makes it easier to maintain and scale the application. For now, service just proxies the repository. In the future, services will be wider and can use multiple repositories and work with various models.

The `gRPC-Client` at `./pkg/client` uses vertical slice architecture. A go package exists per feature (sensor in this example). Inside each feature, a flat folder structure defines everything needed to implement it. `DTOs` with mapper functions are used to transform the data between JSON and protobuf encodings. This architecture best fits the `gRPC-Client`, as it simply proxies data. In the future, auth should be added as a middleware, outside each feature implementation.

The `Front-End` at `./apps/frontend` utilizes a typical React application structure. Components, hooks, and UI elements are split into separate folders. Vertical slice architecture is used again to separate different features, which I call modules. A special `_runtime.ts` file is used to initialize relevant modules as global dependencies. It is like a singleton, but on import level. The frontend application should be developed in terms of modules and reusable components.

### Technology Choices

Here is the list of key technologies used for this assignment, as well as a short explanation of why I chose them:

- Official protobuf compiler. It works for backend purposes and is well-known to developers.
- `Echo` as http server. It's simple to use and has lots of useful middleware out of the box
- `slog` for structured logging. It's a solid library with a good API.
- `React` with `Vite` bundler, `Vitest` testing framework, `tailwind` styling for frontend. Its a typical front-end stack I am familiar with, and its development speed is great!
- `Playwright` for end-to-end testing. Really good project, something I am familiar with

### Testing

Unit test coverage is currently below 80%. To compensate, I have implementedend-to-end testing (`frontend` → `client` → `server` → `mock db`). Integration testing TODO (`http request` → `client` → `server` → `mock db`).

### Other

- Protobuf schemas, gRPC, and http services are versioned for backwards compatibility.
- Backend services implement graceful shutdown.
- In development, everything is run though `make`. Watch mode is present for all dev use-cases.
- All components are Dockerized. They use mount caching for dependency downloads.


## What Can Be Improved

TODO

Testing, Docker builds where contexts dont collide with each other.
Proper e2e setup with real db.
Integration tests with test-containers
