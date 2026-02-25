package sensor

import protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"

type SensorDataDTO struct {
	Timestamp int64   `json:"t"`
	Value     float64 `json:"v"`
}

func ToProtoSensorData(d SensorDataDTO) *protov1.SensorData {
	return &protov1.SensorData{
		Timestamp: d.Timestamp,
		Value:     d.Value,
	}
}

func FromProtoSensorData(p *protov1.SensorData) SensorDataDTO {
	return SensorDataDTO{
		Timestamp: p.Timestamp,
		Value:     p.Value,
	}
}
