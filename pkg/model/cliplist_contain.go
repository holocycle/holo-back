package model

import "time"

type CliplistContain struct {
	CliplistID string
	Index      int
	ClipID     string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func NewCliplistContain(
	cliplistID string,
	index int,
	clipID string,
) *CliplistContain {
	return &CliplistContain{
		CliplistID: cliplistID,
		Index:      index,
		ClipID:     clipID,
		CreatedAt:  nil,
		UpdatedAt:  nil,
	}
}
