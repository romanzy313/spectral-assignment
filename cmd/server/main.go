package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/romanzy313/spectral-assignment/pkg/server"
)

func main() {
	config := server.NewConfigFromEnv()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	server.Run(ctx, config)
}
