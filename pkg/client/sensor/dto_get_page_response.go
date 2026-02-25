package sensor

import (
	pb "spectral-assignment/gen/proto"
)

type GetPageResponseDTO struct {
	NextCursor *int64          `json:"nextCursor"`
	Data       []SensorDataDTO `json:"data"`
}

func ToProtoGetPageResponse(dto GetPageResponseDTO) *pb.GetPageResponse {
	data := make([]*pb.SensorData, len(dto.Data))

	for i, d := range dto.Data {
		data[i] = ToProtoSensorData(d)
	}

	return &pb.GetPageResponse{
		NextCursor: dto.NextCursor,
		Data:       data,
	}
}

func FromProtoGetPageResponse(proto *pb.GetPageResponse) GetPageResponseDTO {
	data := make([]SensorDataDTO, len(proto.Data))

	for i, d := range proto.Data {
		data[i] = FromProtoSensorData(d)
	}

	return GetPageResponseDTO{
		NextCursor: proto.NextCursor,
		Data:       data,
	}
}
