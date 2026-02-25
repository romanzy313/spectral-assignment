package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/romanzy313/spectral-assignment/pkg/server/repository"
	"github.com/romanzy313/spectral-assignment/pkg/server/router"
	"github.com/romanzy313/spectral-assignment/pkg/server/service"
)

// basic implementation from
// https://grpc.io/docs/languages/go/basics/
func Run(config Config) {

	s := grpc.NewServer()

	// dependencies are initialized here
	mockData, err := repository.ReadCsvData("./meterusage.csv")
	if err != nil {
		log.Fatalf("failed to read csv data: %v", err)
	}
	sensorRepo := repository.NewSensorMemoryRepository(mockData)
	sensorService := service.NewSensorService(sensorRepo)
	sensorRouter := router.NewSensorRouter(sensorService)
	sensorRouter.Register(s)

	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening on port %d", config.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
