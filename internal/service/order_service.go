package service

import (
	"order-service/internal/domain"
	"order-service/internal/handler/dto"
)

type OrderService interface {
	Create(req dto.CreateOrderRequest) (err error)
	GetByID(id int) (domain domain.Order, err error)
}
