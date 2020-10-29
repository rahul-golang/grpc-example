package downstream

import (
	"context"
	"google.golang.org/grpc"
	"interview/grpc/proto"
	"interview/models"
	logUtils "interview/utils/log_utils"
)

type InterestsDownStream interface {
	ExecuteInterestsRequest(ctx context.Context, key1, key2 int64) (models.Interests, error)
}

type interestsDownStream struct {
	client proto.InterestsServiceClient
}

func (n interestsDownStream) ExecuteInterestsRequest(ctx context.Context, key1, key2 int64) (models.Interests, error) {
	req := &proto.TwoUserKeys{
		User1: &proto.UserKey{
			Key: key1,
		},
		User2: &proto.UserKey{
			Key: key2,
		},
	}
	interests, err := n.client.GetSharedInterests(ctx, req)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("NetworkDownStream.ExecuteNetworkRequest: Error in calling network service. Error: %+v", err)
		return models.Interests{}, err
	}
	return models.NewInterests(interests), nil
}

func NewInterestsDownStream(conn *grpc.ClientConn) InterestsDownStream {
	return &interestsDownStream{
		client: proto.NewInterestsServiceClient(conn),
	}
}
