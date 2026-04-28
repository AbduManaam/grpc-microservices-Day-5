package main

import (
	"log"
	"net"
	"order-service/internal/client"
	"order-service/internal/db"
	"order-service/internal/repository"
	"order-service/internal/service"
	orderpb "proto/gen/order"

	"google.golang.org/grpc"
)

func main() {

	lis,err:= net.Listen("tcp",":50052")
	if err!=nil{
	   log.Fatal(err)
	}
	s:= grpc.NewServer()
	database:= db.NewDB()
	repo:= &repository.Repo{DB: database}
	userClient:=client.NewUserClient("localhost:50051")
	svc:= &service.Service{Repo: repo,UserClient: userClient}

	orderpb.RegisterOrderServiceServer(s,svc)
	log.Println("Order Service running on :50052")
	s.Serve(lis)

}