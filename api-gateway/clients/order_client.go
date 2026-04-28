package clients

import (
	"context"
	"log"

	orderpb "proto/gen/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var OrderClient orderpb.OrderServiceClient

func InitOrderClient(addr string) {
    conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal(err)
    }
    OrderClient = orderpb.NewOrderServiceClient(conn)
}

func CreateOrder(userID string) (*orderpb.CreateOrderResponse, error) {
    return OrderClient.CreateOrder(context.Background(),
        &orderpb.CreateOrderRequest{UserId: userID})
}