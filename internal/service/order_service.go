package service

import (
	"order-service/internal/domain"
)

type OrderService interface {
	Create(order domain.Order) (err error)
	GetByID(id int) (order domain.Order, err error)
	UpdateStatus(id int, status domain.Status) (err error)
	GetByUserId(id int) (order domain.Order, err error)
	ListAll() ([]domain.Order, error)
	Delete(id int) (err error)
}
