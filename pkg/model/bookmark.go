package model

import "time"

type Bookmark struct {
	CliplistID string `gorm:"primary_key"`
	UserID     string `gorm:"primary_key"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewBookMark(
	cliplistID,
	userID string,
) *Bookmark {
	return &Bookmark{
		CliplistID: cliplistID,
		UserID:     userID,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}
}
