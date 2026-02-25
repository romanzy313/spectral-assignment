package client

import (
	"context"
	"log"
	"time"

	pb "spectral-assignment/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Run() {
	conn, err := grpc.NewClient("localhost:12000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewSensorServiceClient(conn)

	resp, err := client.GetPage(context.Background(), &pb.GetPageRequest{
		Cursor: timestamppb.New(time.Time{}),
		Limit:  99999,
	})
	if err != nil {
		log.Fatalf("RPC failed: %v", err)
	}

	log.Printf("response: %v", resp)

}
