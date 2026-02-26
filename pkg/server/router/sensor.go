package router

import (
	"context"

	protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"
	"github.com/romanzy313/spectral-assignment/pkg/server/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SensorService interface {
	GetPage(ctx context.Context, cursor int64, limit int32) (*model.SensorPage, error)
	GetCount(ctx context.Context) (*model.SensorCount, error)
}

type SensorRouterV1 struct {
	protov1.UnimplementedSensorServiceServer

	sensorService SensorService
}

func NewSensorRouter(sensorService SensorService) *SensorRouterV1 {
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
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &protov1.GetCountResponse{
		Count: data.Count,
	}

	return resp, nil
}
