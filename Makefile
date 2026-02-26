.PHONY: install-tools
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.1
	go install github.com/mitranim/gow@latest

.PHONY: proto
proto:
	@rm -rf gen/proto/v1
	@mkdir -p gen/proto/v1
	protoc \
		--proto_path=. \
		--go_out=gen --go_opt=paths=source_relative \
		--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
		proto/v1/schema.proto

.PHONY: dev-client
dev-client:
	PORT=12001 \
	SPECTRAL_FRONTEND_ORIGIN=http://localhost:12002 \
	SPECTRAL_GRPC_SERVER_ADDRESS=localhost:12000 \
	gow -c run cmd/client/main.go

.PHONY: dev-server
dev-server:
	PORT=12000 \
	gow run cmd/server/main.go

.PHONY: dev-frontend
dev-frontend:
	PUBLIC_SPECTRAL_GRPC_CLIENT_ORIGIN=http://localhost:12001 \
    pnpm --filter frontend dev --port 12002

.PHONY: go-test
go-test:
	go test -v ./...

.PHONY: dev-go-test
dev-go-test:
	gow -c test ./...



.PHONY: test
test: go-test
