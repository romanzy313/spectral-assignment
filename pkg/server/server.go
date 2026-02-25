package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"spectral-assignment/pkg/server/repository"
	"spectral-assignment/pkg/server/router"
	"spectral-assignment/pkg/server/service"
)

// basic implementation from
// https://grpc.io/docs/languages/go/basics/
func Run() {
	port := 12000
	address := fmt.Sprintf("0.0.0.0:%d", port)

	// dependencies are initialized here
	s := grpc.NewServer()

	mockData, err := repository.ReadCsvData("./meterusage.csv")
	if err != nil {
		log.Fatalf("failed to read csv data: %v", err)
	}

	sensorRepo := repository.NewSensorMemoryRepository(mockData)
	sensorService := service.NewSensorService(sensorRepo)
	sensorRouter := router.NewSensorRouter(sensorService)

	sensorRouter.Register(s)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening on port %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
