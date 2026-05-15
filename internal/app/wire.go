package app

import (
	"fmt"
	handler "order-service/internal/handler/grpc"
	"order-service/internal/infra/repository"
	"order-service/internal/service"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&repository.OrderEntity{})

	orderRepository := repository.NewOrderRepository(postgres)
	orderService := service.NewOrderService(orderRepository)
	orderGrpcHandler := handler.NewOrderGrpcHandler(orderService)
	runServer(cfg, orderGrpcHandler)

	return err
}

func runServer(cfg *Config, taskHandler *handler.OrderGrpcHandler) {
	grpcAddress := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)

	// Start HTTP Gateway in goroutine
	go RunHTTPGateway(grpcAddress, cfg.App.Port)
	// Start gRPC server in main thread (blocking)
	RunGrpcServer(grpcAddress, taskHandler)
}
