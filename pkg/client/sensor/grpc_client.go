package sensor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"
)

type GrpcClient interface {
	Close() error

	GetPage(ctx context.Context, req *protov1.GetPageRequest) (*protov1.GetPageResponse, error)
}

type GrpcClientImpl struct {
	conn   *grpc.ClientConn
	client protov1.SensorServiceClient
}

func NewGrpcClient(addr string) (*GrpcClientImpl, error) {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := protov1.NewSensorServiceClient(conn)

	return &GrpcClientImpl{
		conn:   conn,
		client: client,
	}, nil
}

func (c *GrpcClientImpl) Close() error {
	return c.conn.Close()
}

func (c *GrpcClientImpl) GetPage(ctx context.Context, req *protov1.GetPageRequest) (*protov1.GetPageResponse, error) {
	return c.client.GetPage(ctx, req)
}
