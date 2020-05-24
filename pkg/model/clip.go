package model

import (
	"time"
)

type ClipStatus string

const (
	CLIP_PUBLIC  ClipStatus = "PUBLIC"
	CLIP_DELETED ClipStatus = "DELETED"
)

type Clip struct {
	ID          string
	UserID      string
	Title       string
	Description string
	VideoID     string
	BeginAt     int
	EndAt       int
	Status      ClipStatus
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
	status ClipStatus,
) *Clip {
	return &Clip{
		ID:          NewID(),
		UserID:      userID,
		Title:       title,
		Description: description,
		VideoID:     videoID,
		BeginAt:     beginAt,
		EndAt:       EndAt,
		Status:      status,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
}
