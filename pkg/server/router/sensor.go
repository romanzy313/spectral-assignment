package router

import (
	"context"
	"fmt"
	"log"

	pb "spectral-assignment/gen/proto"
	"spectral-assignment/pkg/server/service"

	"google.golang.org/grpc"
)

type SensorRouter struct {
	pb.UnimplementedHelloServiceServer

	sensorService *service.SensorService
}

func NewSensorRouter(sensorService *service.SensorService) *SensorRouter {
	return &SensorRouter{
		sensorService: sensorService,
	}
}

func (s *SensorRouter) Register(gs *grpc.Server) {
	pb.RegisterHelloServiceServer(gs, s)
}

func (s *SensorRouter) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
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
