package main

import (
	"context"
	"google.golang.org/grpc"
	"interview/grpc/proto"
	"log"
	"time"
)

func main(){
	//The service that provide the data runs on a different port, as part of a different application
	grpcConn, err := grpc.Dial(":3030", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	//networkClient := proto.NewNetworkServiceClient(grpcConn)
	//userClient := proto.NewUserServiceClient(grpcConn)
	//contactClient := proto.NewContactServiceClient(grpcConn)
	//interestsClient := proto.NewInterestsServiceClient(grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
req:=&proto.UserViewingNetwork{
	Network: &proto.NetworkKey{
		Key: 1,
	},
	User: &proto.UserKey{
		Key: 1,
	},
}
	client := proto.NewViewNetworkServiceClient(grpcConn)
	members, err := client.ViewNetworkMembers(ctx, req)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", members)

}