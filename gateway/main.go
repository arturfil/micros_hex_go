package main

import (
	"context"
	"log"
	"net/http"
	"time"

	common "github.com/arturfil/m_commons"
	"github.com/arturfil/m_commons/discovery"
	"github.com/arturfil/m_commons/discovery/consul"
	"github.com/arturfil/m_gateway/gateway"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

var (
    serviceName = "gateway"
    httpAddr = common.EnvString("HTTP_ADDR", ":8080")
    consulAddr = common.EnvString("CONSUL_ADDR", "localhost:8500")
)

func main() {

    registry, err := consul.NewRegistry(consulAddr, serviceName)
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    instanceID := discovery.GenerateInstanceID(serviceName)
    if err := registry.Register(ctx, instanceID, serviceName, httpAddr); err != nil {
        panic(err)
    }

    go func() {
        for {
            if err := registry.HealthCheck(instanceID, serviceName); err != nil {
                log.Fatal("failed to health check")
            }
            time.Sleep(time.Second * 1)
        }
    }()

    defer registry.DeRegister(ctx, instanceID, serviceName)

    mux := chi.NewMux()

    ordersGateway := gateway.NewGRPCGateway(registry)

    handler := NewHandler(ordersGateway)
    handler.registerRoutes(mux)

    log.Printf("Starting HTTP server at %s", httpAddr)

    if err := http.ListenAndServe(httpAddr, mux); err != nil {
        log.Fatal("Failed to start http server")
    }
}
