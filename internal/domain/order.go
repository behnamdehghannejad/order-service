package domain

import (
	"time"
)

type Order struct {
	ID        int
	UserId    int
	Amount    float64
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
