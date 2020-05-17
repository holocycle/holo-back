package model

import "time"

type Favorite struct {
	ClipID    string `gorm:"primary_key"`
	UserID    string `gorm:"primary_key"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewFavorite(
	clipID,
	userID string,
) *Favorite {
	return &Favorite{
		ClipID:    clipID,
		UserID:    userID,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
