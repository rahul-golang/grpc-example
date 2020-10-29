package models

import "interview/grpc/proto"

type Interests struct {
	Interests []string
}

func NewInterests(interests *proto.Interests) Interests {
	return Interests{
		Interests: interests.Interests,
	}
}

func (in Interests) MapToProto() *proto.Interests {
	return &proto.Interests{
		Interests: in.Interests,
	}
}
