package router

import (
	"context"
	"log"

	protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"
	"github.com/romanzy313/spectral-assignment/pkg/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SensorRouterV1 struct {
	protov1.UnimplementedSensorServiceServer

	sensorService *service.SensorService
}

func NewSensorRouter(sensorService *service.SensorService) *SensorRouterV1 {
	return &SensorRouterV1{
		sensorService: sensorService,
	}
}

func (s *SensorRouterV1) Register(gs *grpc.Server) {
	protov1.RegisterSensorServiceServer(gs, s)
}

func (s *SensorRouterV1) GetPage(ctx context.Context, req *protov1.GetPageRequest) (*protov1.GetPageResponse, error) {
	data, err := s.sensorService.GetPage(ctx, req.Cursor, req.Limit)

	if err != nil {
		log.Printf("failed to get page data: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get page data: %s", err.Error())
	}

	resp := &protov1.GetPageResponse{
		NextCursor: nil,
		Data:       []*protov1.SensorData{},
	}

	if data.Cursor != nil {
		resp.NextCursor = data.Cursor
	}

	for _, v := range data.Data {
		resp.Data = append(resp.Data, &protov1.SensorData{
			Timestamp: v.Timestamp,
			Value:     v.Value,
		})
	}

	return resp, nil
}

func (s *SensorRouterV1) GetCount(ctx context.Context, _ *protov1.GetCountRequest) (*protov1.GetCountResponse, error) {
	data, err := s.sensorService.GetCount(ctx)

	if err != nil {
		log.Printf("failed to get count data: %s", err.Error())
		return nil, status.Errorf(codes.Internal, "failed to get count data: %s", err.Error())
	}

	resp := &protov1.GetCountResponse{
		Count: data.Count,
	}

	return resp, nil
}
