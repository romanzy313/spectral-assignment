package model

import "time"

type SensorData struct {
	Timestamp time.Time
	Value     float64 // FIXME: should I be using decimal here?
}

type SensorPage struct {
	Data   []SensorData
	Cursor *time.Time
}
