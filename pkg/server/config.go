package server

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port uint16
}

func NewConfigFromEnv() Config {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "3000"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse PORT: %v", err)
	}

	return Config{
		Port: uint16(port),
	}
}
