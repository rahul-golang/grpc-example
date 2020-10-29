package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"interview/grpc/proto"
	"interview/mocks"
	"interview/models"
	"testing"
)

type viewNetworkServiceTestSuite struct {
	suite.Suite
	context             context.Context
	mockCtrl            *gomock.Controller
	contactDownstream   *mocks.MockContactDownStream
	networkDownstream   *mocks.MockNetworkDownStream
	interestsDownstream *mocks.MockInterestsDownStream
	viewNetworkService  proto.ViewNetworkServiceServer
}

func TestViewNetworkServiceTestSuite(t *testing.T) {
	suite.Run(t, new(viewNetworkServiceTestSuite))
}

func (suite *viewNetworkServiceTestSuite) SetupTest() {
	suite.context = context.Background()
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.contactDownstream = mocks.NewMockContactDownStream(suite.mockCtrl)
	suite.networkDownstream = mocks.NewMockNetworkDownStream(suite.mockCtrl)
	suite.interestsDownstream = mocks.NewMockInterestsDownStream(suite.mockCtrl)
	suite.viewNetworkService = NewViewNetworkService(suite.networkDownstream, suite.contactDownstream, suite.interestsDownstream)

}
func (suite *viewNetworkServiceTestSuite) TearDownTest() () {
}

func (suite viewNetworkServiceTestSuite) TestViewNetworkMembers_ShouldReturnAnErrorIfExecuteNetworkDownstreamReturnsAnError() {
	req := &proto.UserViewingNetwork{
		User: &proto.UserKey{
			Key: 1,
		},
		Network: &proto.NetworkKey{
			Key: 1,
		},
	}
	expectedError := errors.New("something-went-wrong")
	suite.networkDownstream.EXPECT().ExecuteNetworkRequest(suite.context, req.Network.Key).Return(models.UserKeys{}, expectedError)
	members, err := suite.viewNetworkService.ViewNetworkMembers(suite.context, req)
	suite.Equal(expectedError, err)
	suite.Nil(members)

}

func (suite viewNetworkServiceTestSuite) TestViewNetworkMembers_ShouldReturnAnErrorIfExecuteContactDownstreamReturnsAnError() {
	req := &proto.UserViewingNetwork{
		User: &proto.UserKey{
			Key: 1,
		},
		Network: &proto.NetworkKey{
			Key: 1,
		},
	}
	netwrokExpectedOutput := models.UserKeys{
		Keys: []models.UserKey{
			{Key: 1},
			{Key: 2},
		},
	}
	expectedError := errors.New("something-went-wrong")
	suite.networkDownstream.EXPECT().ExecuteNetworkRequest(suite.context, req.Network.Key).Return(netwrokExpectedOutput, nil)
	suite.contactDownstream.EXPECT().ExecuteContactRequest(suite.context, netwrokExpectedOutput.Keys[0].Key, netwrokExpectedOutput.Keys[0].Key).Return(models.Contact{}, expectedError)
	members, err := suite.viewNetworkService.ViewNetworkMembers(suite.context, req)
	suite.Equal(expectedError, err)
	suite.Nil(members)

}

func (suite viewNetworkServiceTestSuite) TestViewNetworkMembers_ShouldReturnAnErrorIfExecuteInterestsDownstreamReturnsAnError() {
	req := &proto.UserViewingNetwork{
		User: &proto.UserKey{
			Key: 1,
		},
		Network: &proto.NetworkKey{
			Key: 1,
		},
	}
	netwrokExpectedOutput := models.UserKeys{
		Keys: []models.UserKey{
			{Key: 1},
		},
	}
	contactServiceExpectedOutput := models.Contact{
		Contact: []string{"c1"},
	}
	expectedError := errors.New("something-went-wrong")
	suite.networkDownstream.EXPECT().ExecuteNetworkRequest(suite.context, req.Network.Key).Return(netwrokExpectedOutput, nil)
	suite.contactDownstream.EXPECT().ExecuteContactRequest(suite.context, netwrokExpectedOutput.Keys[0].Key, netwrokExpectedOutput.Keys[0].Key).Return(contactServiceExpectedOutput, nil)
	suite.interestsDownstream.EXPECT().ExecuteInterestsRequest(suite.context, netwrokExpectedOutput.Keys[0].Key, netwrokExpectedOutput.Keys[0].Key).Return(models.Interests{}, expectedError)
	members, err := suite.viewNetworkService.ViewNetworkMembers(suite.context, req)
	suite.Equal(expectedError, err)
	suite.Nil(members)

}


func (suite viewNetworkServiceTestSuite) TestViewNetworkMembers_ShouldSuccess() {
	req := &proto.UserViewingNetwork{
		User: &proto.UserKey{
			Key: 1,
		},
		Network: &proto.NetworkKey{
			Key: 1,
		},
	}
	netwrokExpectedOutput := models.UserKeys{
		Keys: []models.UserKey{
			{Key: 1},
		},
	}
	contactServiceExpectedOutput := models.Contact{
		Contact: []string{"c1"},
	}
	interestServiceExpectedOutput:=models.Interests{
		Interests: []string{"i1"},
	}

	memberDetails:=[]models.MemberDetails{
		{Contact: contactServiceExpectedOutput,UserKey: netwrokExpectedOutput.Keys[0],Interest: interestServiceExpectedOutput},
	}
	resp := models.NetworkMembersView{
		MembersDetails: memberDetails,
	}

	suite.networkDownstream.EXPECT().ExecuteNetworkRequest(suite.context, req.Network.Key).Return(netwrokExpectedOutput, nil)
	suite.contactDownstream.EXPECT().ExecuteContactRequest(suite.context, netwrokExpectedOutput.Keys[0].Key, netwrokExpectedOutput.Keys[0].Key).Return(contactServiceExpectedOutput, nil)
	suite.interestsDownstream.EXPECT().ExecuteInterestsRequest(suite.context, netwrokExpectedOutput.Keys[0].Key, netwrokExpectedOutput.Keys[0].Key).Return(interestServiceExpectedOutput, nil)

	members, err := suite.viewNetworkService.ViewNetworkMembers(suite.context, req)
	suite.Equal(resp.MapToProtoPb(), members)
	suite.Nil(err)

}
