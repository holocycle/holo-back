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
		Title:         clip.Title,
		Description:   clip.Description,
		BeginAt:       clip.BeginAt,
		EndAt:         clip.EndAt,
		FavoriteCount: 0,
		Video:         ConvertToVideo(video),
	}
}

func ConvertToClips(clips []*model.Clip) []*api.Clip {
	res := make([]*api.Clip, 0)
	for _, clip := range clips {
		if clip.Video == nil {
			panic("no video")
		}
		res = append(res, ConvertToClip(clip, clip.Video))
	}
	return res
}
