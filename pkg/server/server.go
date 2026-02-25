package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "spectral-assignment/gen/proto"
	"spectral-assignment/pkg/server/repository"
	"spectral-assignment/pkg/server/service"
)

type server struct {
	pb.UnimplementedHelloServiceServer

	// inject the services here
}

func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	var name string
	if req.Name != nil {
		name = *req.Name
	} else {
		name = "World"
	}

	resp := &pb.EchoResponse{
		Message: fmt.Sprintf("Hello, %s", name),
	}

	log.Printf("responding with %v\n", resp)

	return resp, nil
}

// basic implementation from
// https://grpc.io/docs/languages/go/basics/
func Run() {
	// dependencies here
	mockData, err := repository.ReadCsvData("./meterusage.csv")
	if err != nil {
		log.Fatalf("failed to read csv data: %v", err)
	}

	sensorRepo := repository.NewSensorMemRepo(mockData)
	sensorService := service.NewSensorService(sensorRepo)

	res, err := sensorService.GetSensorReadings(context.Background(), time.Time{}, 9999)
	if err != nil {
		log.Fatalf("failed to get sensor readings: %v", err)
	}
	log.Printf("Sensor readings: %+v", res.Data[0])

	port := 12000
	address := fmt.Sprintf("0.0.0.0:%d", port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})

	log.Printf("server listening on port %d", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
