package app

import (
	"order-service/internal/infra/repository/impl"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return err
	}

	err = postgres.AutoMigrate(&impl.OrderEntity{})

	return err
}
