package gateway

import (
	"context"
	"log"

	pb "github.com/arturfil/m_commons/api"
	"github.com/arturfil/m_commons/discovery"
)

type gateway struct {
    registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
    return &gateway{registry}
}

func (g *gateway) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
    conn, err := discovery.ServiceConnectin(ctx, "orders", g.registry)
    if err != nil {
        log.Fatal("Failed to dial server: %v", err) 
    }

    c := pb.NewOrderServiceClient(conn)

    return c.CreateOrder(ctx, &pb.CreateOrderRequest{
        CustomerID: p.CustomerID,
        Items: p.Items,
    })
}
