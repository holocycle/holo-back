package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToClip(clip *model.Clip, video *model.Video) *api.Clip {
	return &api.Clip{
		ModelBase: api.ModelBase{
			Type: "Clip",
			ID:   clip.ID,
		},
		Title:       clip.Title,
		Description: clip.Description,
		VideoID:     clip.VideoID,
		BeginAt:     clip.BeginAt,
		EndAt:       clip.EndAt,
		Video:       ConvertToVideo(video),
	}
}
