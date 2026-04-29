package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    userpb "proto/gen/user"

    "user-service/internal/db"
    "user-service/internal/interceptor"
    "user-service/internal/repository"
    "user-service/internal/service"
)

func main(){

	lis,err:= net.Listen("tcp",":50051")
	if err!=nil{
		log.Fatal(err)
	}

	s:=grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor),
	)

	database:=db.NewDB()
	repo:= &repository.Repo{DB:database}
	svc:=&service.Service{Repo:repo}

	//When a request comes for UserService, use svc to handle it
	userpb.RegisterUserServiceServer(s, svc)

    log.Println("User Service running on :50051")
    s.Serve(lis)
}