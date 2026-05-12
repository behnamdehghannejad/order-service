package repository

import "order-service/internal/domain"

type OrderRepository interface {
	Create(order domain.Order) error
	GetById(id int64) (*domain.Order, error)
	GetByUserId(userId int64) (*domain.Order, error)
	updateOrderStatus(id int64, status domain.Status) (*domain.Order, error)
	DeleteOrder(id int64) error
}
