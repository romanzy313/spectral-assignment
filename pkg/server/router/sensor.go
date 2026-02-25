package router

import (
	"context"
	"log"

	pb "spectral-assignment/gen/proto"
	"spectral-assignment/pkg/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SensorRouter struct {
	pb.UnimplementedSensorServiceServer

	sensorService *service.SensorService
}

func NewSensorRouter(sensorService *service.SensorService) *SensorRouter {
	return &SensorRouter{
		sensorService: sensorService,
	}
}

func (s *SensorRouter) Register(gs *grpc.Server) {
	pb.RegisterSensorServiceServer(gs, s)
}

func (s *SensorRouter) GetPage(ctx context.Context, req *pb.GetPageRequest) (*pb.GetPageResponse, error) {
	page, err := s.sensorService.GetPage(ctx, req.Cursor, req.Limit)

	if err != nil {
		log.Printf("failed to get page data: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get page data: %s", err.Error())
	}

	resp := &pb.GetPageResponse{
		NextCursor: nil,
		Data:       []*pb.SensorData{},
	}

	if page.Cursor != nil {
		resp.NextCursor = page.Cursor
	}

	for _, v := range page.Data {
		resp.Data = append(resp.Data, &pb.SensorData{
			Timestamp: v.Timestamp,
			Value:     v.Value,
		})
	}

	return resp, nil
}
