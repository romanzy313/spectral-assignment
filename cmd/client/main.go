package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/romanzy313/spectral-assignment/pkg/client"
)

func main() {
	config := client.NewConfigFromEnv()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	client.Run(ctx, config)
}
