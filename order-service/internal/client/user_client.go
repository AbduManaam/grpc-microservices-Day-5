package client

import (
	"context"
	"log"
	userpb "proto/gen/user"

	"google.golang.org/grpc"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) *UserClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
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