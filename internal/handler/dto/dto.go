package dto

import "time"

type CreateOrderRequest struct {
	UserID int     `json:"user_id"`
	Amount float32 `json:"amount"`
	Status string  `json:"status"`
}

type OrderResponse struct {
	UserID    int       `json:"user_id"`
	Amount    float32   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderStatusRequest struct {
	Status string `json:"status"`
}
