package main

import "github.com/romanzy313/spectral-assignment/pkg/client"

func main() {
	config := client.NewConfigFromEnv()

	client.Run(config)
}
