.PHONY: install-tools
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.11
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.1

.PHONY: proto
proto:
	@rm -rf gen/proto/v1
	@mkdir -p gen/proto/v1
	protoc \
		--proto_path=. \
		--go_out=gen --go_opt=paths=source_relative \
		--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
		proto/v1/schema.proto

.PHONY: run-client
run-client:
	go run cmd/client/main.go

.PHONY: run-server
run-server:
	go run cmd/server/main.go

.PHONY: run-frontend
run-frontend:
	pnpm --filter frontend build
	pnpm --filter frontend preview

.PHONY: go-test
go-test:
	go test -v ./...
