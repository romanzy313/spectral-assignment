package sensor

import protov1 "spectral-assignment/gen/proto/v1"

type GetPageResponseDTO struct {
	NextCursor *int64          `json:"nextCursor"`
	Data       []SensorDataDTO `json:"data"`
}

func ToProtoGetPageResponse(dto GetPageResponseDTO) *protov1.GetPageResponse {
	data := make([]*protov1.SensorData, len(dto.Data))

	for i, d := range dto.Data {
		data[i] = ToProtoSensorData(d)
	}

	return &protov1.GetPageResponse{
		NextCursor: dto.NextCursor,
		Data:       data,
	}
}

func FromProtoGetPageResponse(proto *protov1.GetPageResponse) GetPageResponseDTO {
	data := make([]SensorDataDTO, len(proto.Data))

	for i, d := range proto.Data {
		data[i] = FromProtoSensorData(d)
	}

	return GetPageResponseDTO{
		NextCursor: proto.NextCursor,
		Data:       data,
	}
}
