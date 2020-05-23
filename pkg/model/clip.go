package model

import (
	"time"
)

type Clip struct {
	ID          string
	UserID      string
	Title       string
	Description string
	VideoID     string
	BeginAt     int
	EndAt       int
	CreatedAt   *time.Time
	UpdatedAt   *time.Time

	Video     *Video
	Favorites []*Favorite
}

func NewClip(
	userID,
	title,
	description,
	videoID string,
	beginAt,
	EndAt int,
) *Clip {
	return &Clip{
		ID:          NewID(),
		UserID:      userID,
		Title:       title,
		Description: description,
		VideoID:     videoID,
		BeginAt:     beginAt,
		EndAt:       EndAt,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
}
