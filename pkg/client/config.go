package client

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port              uint16
	GrpcServerAddress string
	FrontendOrigin    string
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

	grpcServerAddress := os.Getenv("SPECTRAL_GRPC_SERVER_ADDRESS")
	if grpcServerAddress == "" {
		log.Fatalf("missing SPECTRAL_GRPC_SERVER_ADDRESS env variable")
	}

	frontendOrigin := os.Getenv("SPECTRAL_FRONTEND_ORIGIN")
	if frontendOrigin == "" {
		log.Fatalf("missing SPECTRAL_FRONTEND_ORIGIN env variable")
	}

	return Config{
		Port:              uint16(port),
		GrpcServerAddress: grpcServerAddress,
		FrontendOrigin:    frontendOrigin,
	}
}
