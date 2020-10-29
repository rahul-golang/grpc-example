package main

import (
	"interview/downstream"
	"interview/grpc/proto"
	"interview/service"
	dataPopulator "interview/static_data_populator"
	"net"

	"google.golang.org/grpc"
)

func main() {


	lis, err := net.Listen("tcp", ":3030")

	if err != nil {
		panic(err)
	}
	grpcConn, err := grpc.Dial(":3030", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	//popular := dataPopulator.NewNetworkDataPopular()
	networkDownstream := downstream.NewNetworkDownStream(grpcConn)
	contactDownstream:= downstream.NewContactDownStream(grpcConn)
	interestsDownstream := downstream.NewInterestsDownStream(grpcConn)
	viewNetworkService := service.NewViewNetworkService(networkDownstream,contactDownstream,interestsDownstream)
	proto.RegisterViewNetworkServiceServer(grpcServer,viewNetworkService)

	userData := dataPopulator.NewUserData()
	userService := service.NewUserService(userData)
	proto.RegisterUserServiceServer(grpcServer,userService)

	interestService := service.NewInterestService()
	proto.RegisterInterestsServiceServer(grpcServer,interestService)

	contactService := service.NewContactService()
	proto.RegisterContactServiceServer(grpcServer,contactService)

	networkService := service.NewNetworkService()
	proto.RegisterNetworkServiceServer(grpcServer, networkService)

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}

}
