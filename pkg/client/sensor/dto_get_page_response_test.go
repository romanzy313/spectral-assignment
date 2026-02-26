package sensor

import (
	"testing"

	protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"
	"github.com/stretchr/testify/assert"
)

func TestToProtoGetPageResponse(t *testing.T) {
	proto := ToProtoGetPageResponse(GetPageResponseDTO{
		NextCursor: nil,
		Data: []SensorDataDTO{
			{Timestamp: 10, Value: 1},
		},
	})

	assert.Equal(t, true, proto.NextCursor == nil)
	assert.Equal(t, []*protov1.SensorData{
		{Timestamp: 10, Value: 1},
	}, proto.Data)
}

func TestFromProtoGetPageResponse(t *testing.T) {
	cursor := int64(42)
	dto := FromProtoGetPageResponse(&protov1.GetPageResponse{
		NextCursor: &cursor,
		Data: []*protov1.SensorData{
			{Timestamp: 10, Value: 1},
		},
	})

	assert.Equal(t, true, *dto.NextCursor == 42)
	assert.Equal(t, []SensorDataDTO{
		{Timestamp: 10, Value: 1},
	}, dto.Data)
}

// func FromProtoGetPageResponse(proto *protov1.GetPageResponse) GetPageResponseDTO {
// 	data := make([]SensorDataDTO, len(proto.Data))

// 	for i, d := range proto.Data {
// 		data[i] = FromProtoSensorData(d)
// 	}

// 	return GetPageResponseDTO{
// 		NextCursor: proto.NextCursor,
// 		Data:       data,
// 	}
// }
