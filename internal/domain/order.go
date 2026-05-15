package domain

import (
	"time"
)

type Order struct {
	ID        int
	UserID    int
	Amount    float32
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status string

const (
	CREATED   = "CREATED"
	ACCEPTED  = "ACCEPTED"
	READY     = "READY"
	DELIVERED = "DELIVERED"
)
