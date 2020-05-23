package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToClip(clip *model.Clip, video *model.Video, favorites []*model.Favorite) *api.Clip {
	return &api.Clip{
		ModelBase: api.ModelBase{
			Type: "Clip",
			ID:   clip.ID,
		},
		Title:         clip.Title,
		Description:   clip.Description,
		BeginAt:       clip.BeginAt,
		EndAt:         clip.EndAt,
		FavoriteCount: len(favorites),
		Video:         ConvertToVideo(video),
	}
}

func ConvertToClips(clips []*model.Clip) []*api.Clip {
	res := make([]*api.Clip, 0)
	for _, clip := range clips {
		if clip.Video == nil {
			panic("no video")
		}
		res = append(res, ConvertToClip(clip, clip.Video, clip.Favorites))
	}
	return res
}
