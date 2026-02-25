package service

import (
	"context"
	"spectral-assignment/pkg/server/model"
	"spectral-assignment/pkg/server/repository"
	"time"
)

type SensorService struct {
	repo repository.SensorRepo
}

func NewSensorService(repo repository.SensorRepo) *SensorService {
	return &SensorService{
		repo: repo,
	}
}

func (s *SensorService) GetSensorReadings(ctx context.Context, cursor time.Time, limit int) (*model.SensorPage, error) {
	page, err := s.repo.GetSensorReadings(ctx, cursor, limit)
	if err != nil {
		return nil, err
	}
	return page, nil
}
