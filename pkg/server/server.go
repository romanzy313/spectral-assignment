package server

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "spectral-assignment/gen/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
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
	port := 12000

	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	lis, err := net.Listen("tcp", portStr)
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
