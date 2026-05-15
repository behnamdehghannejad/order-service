package repository

import "order-service/internal/domain"

type OrderRepository interface {
	Create(order domain.Order) error
	GetById(id int) (domain.Order, error)
	GetByUserId(userId int) (domain.Order, error)
	UpdateStatus(id int, status domain.Status) (domain.Order, error)
	AllOrders() ([]domain.Order, error)
	Delete(id int) error
}
