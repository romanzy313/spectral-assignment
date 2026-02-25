package sensor

import (
	pb "spectral-assignment/gen/proto"
)

type GetPageRequestDTO struct {
	Cursor int64 `query:"cursor"`
	Limit  int32 `query:"limit"`
}

func ToProtoGetPageRequest(dto GetPageRequestDTO) *pb.GetPageRequest {
	return &pb.GetPageRequest{
		Cursor: dto.Cursor,
		Limit:  dto.Limit,
	}
}

func FromProtoGetPageRequest(proto *pb.GetPageRequest) GetPageRequestDTO {
	return GetPageRequestDTO{
		Cursor: proto.Cursor,
		Limit:  proto.Limit,
	}
}
