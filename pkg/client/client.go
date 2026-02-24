package client

import (
	"context"
	"log"

	pb "spectral-assignment/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() {
	conn, err := grpc.NewClient("localhost:12000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	name := "Spectral"

	resp, err := client.Echo(context.Background(), &pb.EchoRequest{
		Name: &name,
	})
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	log.Printf("response: %s", resp.Message)

}
