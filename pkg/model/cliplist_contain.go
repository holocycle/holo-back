package model

import "time"

type CliplistContain struct {
	CliplistID string `gorm:"primary_key"`
	Index      int    `gorm:"primary_key"`
	ClipID     string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time

	Clip *Clip
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
