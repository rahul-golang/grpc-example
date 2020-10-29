package service

import (
	"context"
	"interview/grpc/proto"
)

type contactService struct {
}

func (c contactService) GetCommonContacts(ctx context.Context, keys *proto.TwoUserKeys) (*proto.Contacts, error) {
	return &proto.Contacts{
		Contacts: []string{"contact1", "contact2"},
	}, nil
}

func NewContactService() proto.ContactServiceServer {
	return &contactService{}

}
