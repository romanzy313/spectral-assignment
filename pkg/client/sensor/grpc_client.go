package sensor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "spectral-assignment/gen/proto"
)

type GrpcClient interface {
	Close() error

	GetPage(ctx context.Context, req *pb.GetPageRequest) (*pb.GetPageResponse, error)
}

type GrpcClientImpl struct {
	conn   *grpc.ClientConn
	client pb.SensorServiceClient
}

func NewGrpcClient(addr string) (*GrpcClientImpl, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewSensorServiceClient(conn)

	return &GrpcClientImpl{
		conn:   conn,
		client: client,
	}, nil
}

func (c *GrpcClientImpl) Close() error {
	return c.conn.Close()
}

func (c *GrpcClientImpl) GetPage(ctx context.Context, req *pb.GetPageRequest) (*pb.GetPageResponse, error) {
	return c.client.GetPage(ctx, req)
}
