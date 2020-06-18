package model

import (
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	IconURL   string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUser(name, email, iconURL string) *User {
	return &User{
		ID:        NewID(),
		Name:      name,
		Email:     email,
		IconURL:   iconURL,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
