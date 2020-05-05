package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToVideo(video *model.Video) *api.Video {
	return &api.Video{
		ModelBase: api.ModelBase{
			Type: "Video",
			ID:   video.ID,
		},
		ChannelID:          video.ChannelID,
		Title:              video.Title,
		Description:        video.Description,
		Duration:           video.Duration,
		SmallThumbnailURL:  video.SmallThumbnailURL,
		MediumThumbnailURL: video.MediumThumbnailURL,
		LargeThumbnailURL:  video.LargeThumbnailURL,
		PublishedAt:        video.PublishedAt,
	}
}
