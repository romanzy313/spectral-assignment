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

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM) // start shutdown process on signal
	defer cancel()

	client.Run(ctx, config)
}
