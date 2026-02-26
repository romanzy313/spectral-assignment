package server

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	IsDev bool
	Port  uint16
}

func NewConfigFromEnv() Config {
	isDev := os.Getenv("IS_DEV") == "true"

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "80"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse PORT: %v", err)
	}

	return Config{
		IsDev: isDev,
		Port:  uint16(port),
	}
}
