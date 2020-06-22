package api

import "time"

type Video struct {
	ModelBase
	ChannelID          string     `json:"channelId"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	Duration           int        `json:"duration"`
	SmallThumbnailURL  string     `json:"smallThumnailUrl"`
	MediumThumbnailURL string     `json:"mediumThumnailUrl"`
	LargeThumbnailURL  string     `json:"largeThumnailUrl"`
	PublishedAt        *time.Time `json:"publishedAt" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}

func NewVideo(
	id,
	channelID,
	title,
	description string,
	duration int,
	smallThumnailURL,
	mediumThumnailURL,
	largeThumnailURL string,
	publishedAt *time.Time,
) *Video {
	return &Video{
		ModelBase: ModelBase{
			Type: "Video",
			ID:   id,
		},
		ChannelID:          channelID,
		Title:              title,
		Description:        description,
		Duration:           duration,
		SmallThumbnailURL:  smallThumnailURL,
		MediumThumbnailURL: mediumThumnailURL,
		LargeThumbnailURL:  largeThumnailURL,
		PublishedAt:        publishedAt,
	}
}

func VideoModels() []interface{} {
	return []interface{}{
		Video{},
	}
}
