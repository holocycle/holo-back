package model

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUser(name, email string) *User {
	return &User{
		ID:        NewID(),
		Name:      name,
		Email:     email,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
