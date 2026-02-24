package database

import (
	"context"
	"time"
)

type SensorReading struct {
	Timestamp  time.Time
	MeterUsage float64 // FIXME: should I be using decimal here?
}

type ReadingRequest struct {
	SensorId string
	Cursor   *time.Time
	Limit    int
}

type ReadingResponse struct {
	SensorId string
	Data     []*SensorReading
	Cursor   *time.Time
}

type Database interface {
	GetSensorReadings(context.Context, ReadingRequest) (*ReadingResponse, error)
}
