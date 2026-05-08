package domain

import (
	"time"
)

type Order struct {
	ID        int64
	UserId    int64
	Amount    int64
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
