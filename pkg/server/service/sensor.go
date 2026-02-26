package service

import (
	"context"
	"fmt"

	"github.com/romanzy313/spectral-assignment/pkg/logger"
	"github.com/romanzy313/spectral-assignment/pkg/server/model"
	"github.com/romanzy313/spectral-assignment/pkg/server/repository"
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
		logger.FromContext(ctx).Error("failed to get page data", "error", err)
		return nil, fmt.Errorf("failed to get page data: %w", err)
	}

	return page, nil
}
