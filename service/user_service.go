package service

import (
	"context"
	"interview/grpc/proto"
	staticDataPopulator"interview/static_data_populator"
)

type userService struct {
	userData staticDataPopulator.UserData
}

func (u userService) GetUser(ctx context.Context, key *proto.UserKey) (*proto.User, error) {
	data, err := u.userData.GetUseData(ctx, key.Key)
	if err!=nil{
		return nil, err
	}
	return data.MapToProto(),nil
}

func NewUserService(userData staticDataPopulator.UserData) proto.UserServiceServer {
	return &userService{
		userData: userData,
	}
}
