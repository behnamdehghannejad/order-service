package impl

import (
	"order-service/internal/domain"
	"time"

	"gorm.io/gorm"
)

type OrderEntity struct {
	ID        int           `gorm:"column:id;primary_key;autoIncrement"`
	UserId    int           `gorm:"column:user_id"`
	Amount    float32       `gorm:"column:amount"`
	Status    domain.Status `gorm:"column:status"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}

func (OrderEntity) TableName() string {
	return "orders"
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (repo *Repository) Create(order domain.Order) error {
	entity := toEntity(order)
	entity.CreatedAt = time.Now()

	return repo.DB.Create(entity).Error
}

func (repo *Repository) GetById(id int64) (*domain.Order, error) {
	var order OrderEntity
	if err := repo.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return toDomain(order), nil
}

func (repo *Repository) GetByUserId(userId int64) (*domain.Order, error) {
	var order OrderEntity
	if err := repo.DB.Where("user_id = ?", userId).First(&order).Error; err != nil {
		return nil, err
	}
	return toDomain(order), nil
}

func (repo *Repository) updateOrderStatus(id int64, status domain.Status) (*domain.Order, error) {
	var order OrderEntity
	if err := repo.DB.Model(&order).
		Where("id = ?", id).
		Update("status", status).
		Error; err != nil {
		return nil, err
	}

	return toDomain(order), nil
}

func (repo *Repository) DeleteOrder(id int64) error {
	result := repo.DB.Delete(&OrderEntity{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func toEntity(order domain.Order) OrderEntity {
	return OrderEntity{
		UserId: order.UserID,
		Amount: order.Amount,
		Status: order.Status,
	}
}

func toDomain(order OrderEntity) *domain.Order {
	return &domain.Order{
		ID:        order.ID,
		UserID:    order.UserId,
		Amount:    order.Amount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
