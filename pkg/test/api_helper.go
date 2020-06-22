package test

import (
	"fmt"
	"time"

	"github.com/holocycle/holo-back/pkg/api"
)

func APIVideo(id int) *api.Video {
	return api.NewVideo(
		fmt.Sprintf("video%02d", id),
		fmt.Sprintf("channel%02d", id),
		fmt.Sprintf("video %02d", id),
		fmt.Sprintf("video %02d", id),
		100,
		"http://small",
		"http://medium",
		"http://large",
		tm("2020/1/1 00:00:00"),
	)
}

func APICliplist(id, length int, firstItem *api.CliplistItem) *api.Cliplist {
	return api.NewCliplist(
		fmt.Sprintf("cliplist%02d", id),
		fmt.Sprintf("cliplist %02d", id),
		fmt.Sprintf("cliplist %02d", id),
		length,
		firstItem,
	)
}

func APICliplistItem(id, favoriteCount int, video *api.Video, available bool) *api.CliplistItem {
	if !available {
		return api.NewCliplistItem(
			fmt.Sprintf("clip%02d", id),
			"",
			"",
			0,
			0,
			0,
			nil,
			false,
		)
	}
	return api.NewCliplistItem(
		fmt.Sprintf("clip%02d", id),
		fmt.Sprintf("clip %02d", id),
		fmt.Sprintf("clip %02d", id),
		1,
		10,
		0,
		video,
		true,
	)
}

func tm(value string) *time.Time {
	tm, err := time.Parse("2006/1/2 15:04:05", value)
	if err != nil {
		panic(err)
	}
	return &tm
}
