package model

import (
	"time"
)

type Video struct {
	ID                 string
	ChannelID          string
	Title              string
	Description        string
	Duration           int
	SmallThumbnailURL  string
	MediumThumbnailURL string
	LargeThumbnailURL  string
	PublishedAt        *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

func NewVideo(
	id,
	channelID,
	title,
	description string,
	duration int,
	smallThumbnailURL,
	mediumThumbnailURL,
	largeThumbnailURL string,
	publishedAt *time.Time,
) *Video {
	return &Video{
		ID:                 id,
		ChannelID:          channelID,
		Title:              title,
		Description:        description,
		Duration:           duration,
		SmallThumbnailURL:  smallThumbnailURL,
		MediumThumbnailURL: mediumThumbnailURL,
		LargeThumbnailURL:  largeThumbnailURL,
		PublishedAt:        publishedAt,
		CreatedAt:          nil,
		UpdatedAt:          nil,
	}
}
