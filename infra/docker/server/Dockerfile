FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

# Download modules with cache and build
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/server/main.go


FROM scratch AS production

WORKDIR /app

COPY --from=builder /app/bin/server ./main

# copy mock data
COPY --from=builder /app/meterusage.csv ./meterusage.csv

EXPOSE 80

CMD ["./main"]
