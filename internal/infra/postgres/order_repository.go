package postgres

import (
	"order-service/internal/domain"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        int64         `gorm:"column:id;primary_key;autoIncrement"`
	UserId    int64         `gorm:"column:user_id"`
	Amount    int64         `gorm:"column:amount"`
	Status    domain.Status `gorm:"column:status"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
}

func (Order) TableName() string {
	return "orders"
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Create(Order *domain.Order) error {
	return repo.DB.Create(&Order).Error
}
