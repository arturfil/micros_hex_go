package main

import (
	"context"
	"log"
	"net"
	"time"

	common "github.com/arturfil/m_commons"
	"github.com/arturfil/m_commons/discovery"
	"github.com/arturfil/m_commons/discovery/consul"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	serviceName = "orders"
	grpcAddr    = common.EnvString("GRPC_ADDR", "localhost:2000")
	consulAddr  = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main() {

	grpcServer := grpc.NewServer()

	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatalf("failed to health check %v", err)
			}
			time.Sleep(time.Second * 1)
		}
	}()

	defer registry.DeRegister(ctx, instanceID, serviceName)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())

	log.Printf("GRPC Server started at %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
