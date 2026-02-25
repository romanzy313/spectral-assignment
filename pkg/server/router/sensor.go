package router

import (
	"context"
	"log"

	pb "spectral-assignment/gen/proto"
	"spectral-assignment/pkg/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	page, err := s.sensorService.GetPage(ctx, req.Cursor.AsTime(), int(req.Limit))

	if err != nil {
		log.Printf("failed to get page data: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get page data: %s", err.Error())
	}

	resp := &pb.GetPageResponse{
		NextCursor: nil,
		Data:       []*pb.SensorReading{},
	}

	if page.Cursor != nil {
		resp.NextCursor = timestamppb.New(*page.Cursor)
	}

	for _, reading := range page.Data {
		resp.Data = append(resp.Data, &pb.SensorReading{
			Timestamp: timestamppb.New(reading.Timestamp),
			Value:     reading.Value,
		})
	}

	return resp, nil
}
