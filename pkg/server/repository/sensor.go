package repository

import (
	"context"

	"github.com/romanzy313/spectral-assignment/pkg/server/model"
)

type SensorRepository interface {
	GetPage(ctx context.Context, cursor int64, limit int32) (*model.SensorPage, error)
}
