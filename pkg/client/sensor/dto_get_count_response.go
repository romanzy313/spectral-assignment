package sensor

import protov1 "github.com/romanzy313/spectral-assignment/gen/proto/v1"

type GetCountResponseDTO struct {
	Count uint64 `json:"count"`
}

func ToProtoGetCountResponse(dto GetCountResponseDTO) *protov1.GetCountResponse {

	return &protov1.GetCountResponse{
		Count: dto.Count,
	}
}

func FromProtoGetCountResponse(proto *protov1.GetCountResponse) GetCountResponseDTO {
	return GetCountResponseDTO{
		Count: proto.Count,
	}
}
