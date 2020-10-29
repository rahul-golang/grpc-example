package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"interview/grpc/proto"
	"testing"
)

type networkServiceTestSuite struct {
	suite.Suite
	context        context.Context
	networkService proto.NetworkServiceServer
}

func TestnetworkServiceTestSuite(t *testing.T) {
	suite.Run(t, new(networkServiceTestSuite))
}

func (suite *networkServiceTestSuite) SetupTest() {
	suite.context = context.Background()
	suite.networkService = NewNetworkService()

}
func (suite *networkServiceTestSuite) TearDownTest() () {
}

func (suite networkServiceTestSuite) TestGetSharedInterests_shouldReturnContactDetails() {
	expectd := &proto.UserKeys{
		Users: []*proto.UserKey{
			{Key: 1},
			{Key: 2},
		},
	}

	request := &proto.NetworkKey{
		Key: 1,
	}
	contacts, err := suite.networkService.GetNetworkMembers(suite.context, request)
	suite.Nil(err)
	suite.Equal(expectd, contacts)
}
