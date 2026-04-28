package service

import (
	"context"
	"log"
	"order-service/internal/client"
	"order-service/internal/repository"
	orderpb "proto/gen/order"
)

type Service struct {
	orderpb.UnimplementedOrderServiceServer
	Repo       *repository.Repo
	UserClient *client.UserClient
}

func (s *Service) CreateOrder(ctx context.Context, req *orderpb.CreateOrderRequest) (*orderpb.CreateOrderResponse, error) {

	if err := s.UserClient.ValidateUser(req.UserId); err != nil {
		log.Fatal(err)
	}
	status := s.Repo.Create(req.UserId)
	return &orderpb.CreateOrderResponse{Status: status}, nil
}
