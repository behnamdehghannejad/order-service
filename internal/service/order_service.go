package service

import (
	"order-service/internal/domain"
	"order-service/internal/handler/dto"
)

type OrderService interface {
	Create(req dto.CreateOrderRequest) (err error)
	GetByID(id int) (order domain.Order, err error)
	UpdateOrderStatus(id int, status domain.Status) (err error)
	GetByUserId(id int) (order domain.Order, err error)
	ListAll() ([]domain.Order, error)
	Delete(id int) (err error)
}
