package postgres

import (
	"order-service/internal/domain"
	"time"
)

type Order struct {
	ID        int64         `gorm:"column:id;primary_key"`
	UserId    int64         `gorm:"column:user_id"`
	Amount    int64         `gorm:"column:amount"`
	Status    domain.Status `gorm:"column:status"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
}
