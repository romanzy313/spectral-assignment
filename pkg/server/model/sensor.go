package model

type SensorData struct {
	Timestamp int64
	Value     float64
}

type SensorPage struct {
	Data   []SensorData
	Cursor *int64
}

type SensorCount struct {
	Count uint64
}
