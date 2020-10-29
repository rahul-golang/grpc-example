package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"interview/grpc/proto"
	"testing"
)

type contactServiceTestSuite struct {
	suite.Suite
	context context.Context
	contactService proto.ContactServiceServer
}

func TestContactServiceTestSuite(t *testing.T){
	suite.Run(t,new(contactServiceTestSuite))
}

func(suite *contactServiceTestSuite)SetupTest(){
	suite.context=context.Background()
	suite.contactService=NewContactService()

}
func(suite *contactServiceTestSuite) TearDownTest()(){
}

func (suite contactServiceTestSuite)TestGetCommonContacts_shouldReturnContactDetails(){
	expectd:=&proto.Contacts{Contacts:[]string{"contact1", "contact2"}}
	request:=&proto.TwoUserKeys{
		User1: &proto.UserKey{
			Key: 1,
		},
		User2: &proto.UserKey{
			Key: 1,
		},
	}
	contacts, err := suite.contactService.GetCommonContacts(suite.context,request)
	suite.Nil(err)
	suite.Equal(expectd,contacts)
}