package model

import (
	"time"
)

type HealthCheck struct {
	ID        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		ID:        NewID(),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
