# spectral-assignment

## Dev Setup

### Golang stuff

- Install go. Verify with `go version`
- [Install protobuf compiler](https://protobuf.dev/installation/). I have  used `apt` for my machine. Verify version with `protoc --version`. It works on my machine with `3.21.12`, any version 3+ should be okay.
- Install makefile. Then run `make install-tools`
- Make sure to add go binary paths to env `export PATH="$PATH:$(go env GOPATH)/bin"`. Check with `protoc-gen-go --version`

### Node stuff

- `nvm use`, `nvm install`, `corepack enable pnpm`, `pnpm install`


# Dev notes:

- in dev backend is served on `12000`, frontend at `12001`, webapp at `12002`
