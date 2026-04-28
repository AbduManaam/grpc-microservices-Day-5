package clients

import (
	"context"
	"log"

	userpb "proto/gen/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserClient userpb.UserServiceClient

func InitUserClient(addr string) {
    conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    UserClient = userpb.NewUserServiceClient(conn)
}

func Register(username, password string) (*userpb.RegisterResponse, error) {
    return UserClient.Register(context.Background(), &userpb.RegisterRequest{
        Username: username, Password: password,
    })
}

func Login(username, password string) (*userpb.LoginResponse, error) {
    return UserClient.Login(context.Background(), &userpb.LoginRequest{
        Username: username, Password: password,
    })
}

func GetUser(id string) (*userpb.GetUserResponse, error) {
    return UserClient.GetUser(context.Background(), &userpb.GetUserRequest{
        UserId: id,
    })
}