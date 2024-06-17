package gateway

import (
    "context"
    pb "github.com/arturfil/m_commons/api"
)

type OrdersGateway interface {
    CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error)
}
