package main

import (
    "context"
    pb "github.com/arturfil/m_commons/api"
)

type OrdersService interface {
    CreateOrder(context.Context) error
    ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
    Create(context.Context) error
}

