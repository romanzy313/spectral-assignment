package sensor

import (
	protov1 "spectral-assignment/gen/proto/v1"
)

type GetPageRequestDTO struct {
	Cursor int64 `query:"cursor"`
	Limit  int32 `query:"limit"`
}

func ToProtoGetPageRequest(dto GetPageRequestDTO) *protov1.GetPageRequest {
	return &protov1.GetPageRequest{
		Cursor: dto.Cursor,
		Limit:  dto.Limit,
	}
}

func FromProtoGetPageRequest(proto *protov1.GetPageRequest) GetPageRequestDTO {
	return GetPageRequestDTO{
		Cursor: proto.Cursor,
		Limit:  proto.Limit,
	}
}
