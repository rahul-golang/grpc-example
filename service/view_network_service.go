package service

import (
	"context"
	"interview/downstream"
	"interview/grpc/proto"
	"interview/models"
	logUtils "interview/utils/log_utils"
)

type viewNetworkService struct {
	networkDownstream   downstream.NetworkDownStream
	contactDownStream   downstream.ContactDownStream
	interestsDownStream downstream.InterestsDownStream
}

func (service viewNetworkService) ViewNetworkMembers(ctx context.Context, req *proto.UserViewingNetwork) (*proto.NetworkMembersView, error) {
	logUtils.GetLogger(ctx).Infof("viewNetworkService.ViewNetworkMembers: Inside view network member service: %+v", req)
	request, err := service.networkDownstream.ExecuteNetworkRequest(ctx, req.Network.Key)
	if err != nil {
		logUtils.GetLogger(ctx).Errorf("viewNetworkService.ViewNetworkMembers: Error from network downstream service. Error: %+v", err)
		return nil, err
	}
	logUtils.GetLogger(ctx).Infof("viewNetworkService.ViewNetworkMembers: Network Service Resp: %+v", request)

	memberDetails:=[]models.MemberDetails{}
	for _, v :=range  request.Keys{

		memDetails:=models.MemberDetails{}
		memDetails.UserKey=models.UserKey{Key:v.Key }
		contact, err := service.contactDownStream.ExecuteContactRequest(ctx,v.Key, v.Key)
		if err != nil {
			logUtils.GetLogger(ctx).Errorf("viewNetworkService.ViewNetworkMembers: Error from contact downstream service. Error: %+v", err)
			return nil, err
		}
		logUtils.GetLogger(ctx).Infof("viewNetworkService.ViewNetworkMembers: Contact Service Resp: %+v", contact)

		memDetails.Contact=contact


		interests, err := service.interestsDownStream.ExecuteInterestsRequest(ctx,v.Key, v.Key)
		if err != nil {
			logUtils.GetLogger(ctx).Errorf("viewNetworkService.ViewNetworkMembers: Error from interests downstream service. Error: %+v", err)
			return nil, err
		}
		logUtils.GetLogger(ctx).Infof("viewNetworkService.ViewNetworkMembers: Interests Service Resp: %+v", interests)
		memDetails.Interest=interests
		memberDetails = append(memberDetails, memDetails)
	}


	resp := models.NetworkMembersView{
		MembersDetails: memberDetails,
	}
	logUtils.GetLogger(ctx).Infof("viewNetworkService.ViewNetworkMembers: view network service Resp: %+v", resp)


	return resp.MapToProtoPb(), nil

}

func NewViewNetworkService(networkDownstream downstream.NetworkDownStream, contactDownStream downstream.ContactDownStream, interestsDownStream downstream.InterestsDownStream) proto.ViewNetworkServiceServer {
	return &viewNetworkService{
		networkDownstream:   networkDownstream,
		contactDownStream:   contactDownStream,
		interestsDownStream: interestsDownStream,
	}

}
