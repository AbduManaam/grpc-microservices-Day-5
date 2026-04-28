package client

import (
	"context"
	"log"
	userpb "proto/gen/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) *UserClient {
	conn, err := grpc.NewClient(addr,  grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return &UserClient{client: userpb.NewUserServiceClient(conn)}
}

func(c *UserClient)ValidateUser(userId string)error{

	_,err:= c.client.GetUser(context.Background(),&userpb.GetUserRequest{
		UserId: userId,

	})
	return err
}