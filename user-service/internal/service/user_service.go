package service

import (
	"context"
	"errors"
	userpb "proto/gen/user"
	"user-service/internal/auth"
	"user-service/internal/repository"
)

type Service struct {
	userpb.UnimplementedUserServiceServer
	Repo *repository.Repo
}

func(c *Service)Register(ctx context.Context,req *userpb.RegisterRequest)(*userpb.RegisterResponse, error){

	hash,err:= auth.HashPassword(req.Password)
	if err!=nil{
		return nil,err
	}

	if err:= c.Repo.CreateUser(req.Username, hash);err!=nil{
		return nil,err
	}

	return &userpb.RegisterResponse{Status:"created"},nil

}

func(c *Service)Login(ctx context.Context,req *userpb.LoginRequest)(*userpb.LoginResponse, error){

	u,err:=c.Repo.GetByUserName(req.Username)
	if err!=nil{
		return nil,err
	}

	if !auth.CheckPassword(req.Password,u.Password){
		return nil,errors.New("Invalid Credintial")
	}
	token,err:=auth.GenerateToken(u.Id)
    if err != nil {
        return nil, err
    }
	
	return &userpb.LoginResponse{Token:token},nil
}

func (s *Service) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
    u, err := s.Repo.GetUser(req.UserId)
    if err != nil {
        return nil, err
    }
    return &userpb.GetUserResponse{UserId: u.Id, Name: u.Username}, nil
}