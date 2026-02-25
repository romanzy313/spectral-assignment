package service

import (
	"context"
	"spectral-assignment/pkg/server/model"
	"spectral-assignment/pkg/server/repository"
)

type SensorService struct {
	repo repository.SensorRepository
}

func NewSensorService(repo repository.SensorRepository) *SensorService {
	return &SensorService{
		repo: repo,
	}
}

func (s *SensorService) GetPage(ctx context.Context, cursor int64, limit int32) (*model.SensorPage, error) {
	page, err := s.repo.GetPage(ctx, cursor, limit)
	if err != nil {
		return nil, err
	}
	return page, nil
}
