package main

import (
	"context"
	pb "github.com/arturfil/m_commons/api"
)

type OrdersService interface {
	CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
	ValidateOrder(context.Context, *pb.CreateOrderRequest) ([]*pb.Item, error)
}

type OrderStore interface {
	Create(context.Context) error
}
