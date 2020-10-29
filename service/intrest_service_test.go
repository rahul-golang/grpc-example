package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"interview/grpc/proto"
	"testing"
)

type interestServiceTestSuite struct {
	suite.Suite
	context context.Context
	interestService proto.InterestsServiceServer
}

func TestInterestServiceTestSuite(t *testing.T){
	suite.Run(t,new(interestServiceTestSuite))
}

func(suite *interestServiceTestSuite)SetupTest(){
	suite.context=context.Background()
	suite.interestService=NewInterestService()

}
func(suite *interestServiceTestSuite) TearDownTest()(){
}

func (suite interestServiceTestSuite)TestGetSharedInterests_shouldReturnContactDetails(){
	expectd:=&proto.Interests{Interests:[]string{"intrest1", "intrest2"}}
	request:=&proto.TwoUserKeys{
		User1: &proto.UserKey{
			Key: 1,
		},
		User2: &proto.UserKey{
			Key: 1,
		},
	}
	contacts, err := suite.interestService.GetSharedInterests(suite.context,request)
	suite.Nil(err)
	suite.Equal(expectd,contacts)
}