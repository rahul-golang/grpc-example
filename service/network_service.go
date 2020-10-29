package service

import (
	"context"
	"interview/grpc/proto"
	logUtils "interview/utils/log_utils"
)

type networkService struct {
}

func (service networkService) GetNetworkMembers(ctx context.Context, nKey *proto.NetworkKey) (*proto.UserKeys, error) {
	logUtils.GetLogger(ctx).Infof("networkService.GetNetworkMembers: Inside get network members. Request: %+v", nKey)
	return &proto.UserKeys{
		Users: []*proto.UserKey{
			{Key: 1},
			{Key: 2},
		},
	}, nil

}

func NewNetworkService() proto.NetworkServiceServer {
	return &networkService{}

}
