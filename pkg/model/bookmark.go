package model

import "time"

type Bookmark struct {
	CliplistID string
	UserID     string
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
