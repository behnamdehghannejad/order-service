package app

import (
	"order-service/internal/utils"
)

func Start() error {
	cfg := LoadConfig()
	postgres, err := NewPostgres(cfg)
	if err != nil {
		return utils.AppError{
			Code:    404,
			Message: "user not found",
		}
	}

	err = postgres.AutoMigrate()

	return err
}
