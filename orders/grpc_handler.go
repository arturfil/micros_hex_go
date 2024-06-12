package main

import (
	"context"
	"log"

	pb "github.com/arturfil/m_commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
    pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
    handler := &grpcHandler{}
    pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
    log.Println("New order recieved! Order %v", p)
    o := &pb.Order{
        ID: "42",
    }
    return o, nil
}

