package service

import (
	"context"
	"interview/grpc/proto"
)

type interestService struct {
}

func (i interestService) GetSharedInterests(ctx context.Context, keys *proto.TwoUserKeys) (*proto.Interests, error) {
	return &proto.Interests{
		Interests: []string{"intrest1","intrest2"},
	},nil
}

func NewInterestService() proto.InterestsServiceServer {
	return &interestService{}
}
