package main

import "github.com/romanzy313/spectral-assignment/pkg/server"

func main() {
	config := server.NewConfigFromEnv()

	server.Run(config)
}
