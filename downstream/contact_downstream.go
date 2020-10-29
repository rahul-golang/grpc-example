package downstream

import (
	"context"
	"google.golang.org/grpc"
	"interview/grpc/proto"
	"interview/models"
	logUtils "interview/utils/log_utils"
)

type ContactDownStream interface {
	ExecuteContactRequest(ctx context.Context, key1, key2 int64) (models.Contact, error)
}

type contactDownStream struct {
	client proto.ContactServiceClient
}

func (n contactDownStream) ExecuteContactRequest(ctx context.Context, key1, key2 int64) (models.Contact, error) {
	req := &proto.TwoUserKeys{
		User1: &proto.UserKey{
			Key: key1,
		},
		User2: &proto.UserKey{
			Key: key2,
		},
	}
	members, err := n.client.GetCommonContacts(ctx, req)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("NetworkDownStream.ExecuteNetworkRequest: Error in calling network service. Error: %+v", err)
		return models.Contact{}, err
	}
	return models.NewContact(members), nil
}

func NewContactDownStream(conn *grpc.ClientConn) ContactDownStream {
	return &contactDownStream{
		client: proto.NewContactServiceClient(conn),
	}
}
