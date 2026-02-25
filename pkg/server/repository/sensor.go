package repository

import (
	"context"
	"time"

	"spectral-assignment/pkg/server/model"
)

type SensorRepository interface {
	GetSensorReadings(ctx context.Context, cursor time.Time, limit int) (*model.SensorPage, error)
}
