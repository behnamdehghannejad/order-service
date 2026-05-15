package app

import (
	"fmt"
	"order-service/internal/infra/repository"
	"order-service/internal/infra/repository/impl"
	"order-service/internal/service"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&impl.OrderEntity{})

	repo := repository.NewTaskRepository(db)
	service := service.NewTaskService(repo)
	taskHandler := handler.NewTaskHandler(service)

	return err
}

func runServer(cfg *Config, taskHandler *handler2.TaskHandler) {
	grpcAddress := fmt.Sprintf("%s:%d", cfg.GRPC.Host, cfg.GRPC.Port)
	// Start HTTP Gateway in goroutine
	go server.RunHTTPGateway(grpcAddress, cfg.App.Port)

	// Start gRPC server in main thread (blocking)
	server.RunGrpcServer(cfg, grpcAddress, taskHandler)
}
