package model

import (
	"time"

	"github.com/senseyeio/duration"
)

type Video struct {
	Kind    string
	ETag    string
	ID      string
	Snippet struct {
		PublishedAt time.Time
		ChannelID   string
		Title       string
		Description string
		Thumbnails  struct {
			Default struct {
				URL    string
				Width  int
				Height int
			}
			Medium struct {
				URL    string
				Width  int
				Height int
			}
			High struct {
				URL    string
				Width  int
				Height int
			}
		}
		ChannelTitle string
		Tags         []string
		CategoryID   string
	}
	ContentDetails struct {
		Duration          duration.Duration
		Dimension         string
		Definition        string
		Caption           string
		LicensedContent   bool
		RegionRestriction map[string]interface{}
		ContentRating     map[string]string
	}
}

type VideoListResponse struct {
	Kind          string
	ETag          string
	NextPageToken string
	PrevPageToken string
	PageInfo      struct {
		TotalResults  int
		ResultPerPage int
	}
	Items []Video
}
