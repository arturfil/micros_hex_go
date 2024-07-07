package main

import (
	"context"
	pb "github.com/arturfil/m_commons/api"
)

type OrdersService interface {
    GetOrder(context.Context, *pb.GetOrderRequest) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
	CreateOrder(context.Context, *pb.CreateOrderRequest, []*pb.Item) (*pb.Order, error)
}

type OrderStore interface {
	Create(context.Context, *pb.CreateOrderRequest, []*pb.Item) (string, error) 
    Get(ctx context.Context, id, customerID string) (*pb.Order, error)
}
