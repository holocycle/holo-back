package model

import (
	"time"
)

type Channel struct {
	ID                 string
	Title              string
	Description        string
	SmallThumbnailURL  string
	MediumThumbnailURL string
	LargeThumbnailURL  string
	SmallBannerURL     string
	MediumBannerURL    string
	LargeBannerURL     string
	ViewCount          int64
	CommentCount       int64
	SubscriberCount    int64
	VideoCount         int64
	PublishedAt        *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

func NewChannel(
	id,
	title,
	description,
	smallThumbnailURL,
	mediumThumbnailURL,
	largeThumbnailURL,
	smallBannerURL,
	mediumBannerURL,
	largeBannerURL string,
	viewCount,
	commentCount,
	subscriberCount,
	videoCount int64,
	publishedAt *time.Time,
) *Channel {
	return &Channel{
		ID:                 id,
		Title:              title,
		Description:        description,
		SmallThumbnailURL:  smallThumbnailURL,
		MediumThumbnailURL: mediumThumbnailURL,
		LargeThumbnailURL:  largeThumbnailURL,
		SmallBannerURL:     smallBannerURL,
		MediumBannerURL:    mediumBannerURL,
		LargeBannerURL:     largeBannerURL,
		ViewCount:          viewCount,
		CommentCount:       commentCount,
		SubscriberCount:    subscriberCount,
		VideoCount:         videoCount,
		PublishedAt:        publishedAt,
		CreatedAt:          nil,
		UpdatedAt:          nil,
	}
}
