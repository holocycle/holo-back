package model

import (
	"time"

	"github.com/google/uuid"
)

type HealthCheck struct {
	ID        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{
		ID:        "healthcheck-" + uuid.New().String(),
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
