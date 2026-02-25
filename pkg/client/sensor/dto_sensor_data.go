package sensor

import (
	pb "spectral-assignment/gen/proto"
)

type SensorDataDTO struct {
	Timestamp int64   `json:"t"`
	Value     float64 `json:"v"`
}

func ToProtoSensorData(d SensorDataDTO) *pb.SensorData {
	return &pb.SensorData{
		Timestamp: d.Timestamp,
		Value:     d.Value,
	}
}

func FromProtoSensorData(p *pb.SensorData) SensorDataDTO {
	return SensorDataDTO{
		Timestamp: p.Timestamp,
		Value:     p.Value,
	}
}
