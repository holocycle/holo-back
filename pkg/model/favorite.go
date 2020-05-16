package model

import "time"

type Favorite struct {
	ClipID    string
	UserID    string
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
