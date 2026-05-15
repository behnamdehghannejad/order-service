package repository

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

type OrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{DB: db}
}

func (repo *OrderRepositoryImpl) Create(order domain.Order) error {
	entity := toEntity(order)
	entity.CreatedAt = time.Now()

	return repo.DB.Create(entity).Error
}

func (repo *OrderRepositoryImpl) GetById(id int) (domain.Order, error) {
	var order OrderEntity
	if err := repo.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return toDomain(order), nil
}

func (repo *OrderRepositoryImpl) GetByUserId(userId int) (domain.Order, error) {
	var order OrderEntity
	if err := repo.DB.Where("user_id = ?", userId).First(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return toDomain(order), nil
}

func (repo *OrderRepositoryImpl) UpdateStatus(id int, status domain.Status) (domain.Order, error) {
	var order OrderEntity

	if err := repo.DB.Model(&order).
		Where("id = ?", id).
		Update("status", status).
		Error; err != nil {
		return domain.Order{}, err
	}

	return toDomain(order), nil
}

func (repo *OrderRepositoryImpl) AllOrders() ([]domain.Order, error) {
	entities := make([]OrderEntity, 0, 50)

	if err := repo.DB.Find(&entities).Error; err != nil {
		return nil, err
	}

	orders := make([]domain.Order, 0, len(entities))
	for _, entity := range entities {
		orders = append(orders, toDomain(entity))
	}

	return orders, nil
}

func (repo *OrderRepositoryImpl) Delete(id int) error {
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

func toDomain(order OrderEntity) domain.Order {
	return domain.Order{
		ID:        order.ID,
		UserID:    order.UserId,
		Amount:    order.Amount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}
}
