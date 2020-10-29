package downstream

import (
	"context"
	"google.golang.org/grpc"
	"interview/grpc/proto"
	"interview/models"
	logUtils "interview/utils/log_utils"
)

type NetworkDownStream interface {
	ExecuteNetworkRequest(ctx context.Context, networkKey int64) (models.UserKeys, error)
}

type networkDownStream struct {
	client proto.NetworkServiceClient
}

func (n networkDownStream) ExecuteNetworkRequest(ctx context.Context, networkKey int64) (models.UserKeys, error) {
	logUtils.GetLogger(ctx).Info("NetworkDownStream.ExecuteNetworkRequest: Creating downstream call.")

	req := &proto.NetworkKey{
		Key: networkKey,
	}
	members, err := n.client.GetNetworkMembers(ctx, req)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("NetworkDownStream.ExecuteNetworkRequest: Error in calling network service. Error: %+v", err)
		return models.UserKeys{}, err
	}
	return models.NewUserKeys(members), nil
}

func NewNetworkDownStream(conn *grpc.ClientConn) NetworkDownStream {
	return &networkDownStream{
		client: proto.NewNetworkServiceClient(conn),
	}
}
