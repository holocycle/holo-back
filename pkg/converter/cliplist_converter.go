package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToCliplist(cliplist *model.Cliplist) *api.Cliplist {
	var firstItem *api.CliplistItem = nil
	if len(cliplist.CliplistContains) > 0 {
		firstItem = ConvertToCliplistItem(cliplist.CliplistContains[0])
	}

	return &api.Cliplist{
		ModelBase: api.ModelBase{
			Type: "Cliplist",
			ID:   cliplist.ID,
		},
		Title:       cliplist.Title,
		Description: cliplist.Description,
		Length:      len(cliplist.CliplistContains),
		FirstItem:   firstItem,
	}
}

func ConvertToCliplists(cliplists []*model.Cliplist) []*api.Cliplist {
	res := make([]*api.Cliplist, 0)
	for _, cliplist := range cliplists {
		res = append(res, ConvertToCliplist(cliplist))
	}
	return res
}

func ConvertToCliplistItem(cliplistContain *model.CliplistContain) *api.CliplistItem {
	clip := cliplistContain.Clip

	if clip.Status == model.ClipStatusDeleted {
		return &api.CliplistItem{
			Clip: api.Clip{
				ModelBase: api.ModelBase{
					ID:   clip.ID,
					Type: "CliplistItem",
				},
			},
			Available: false,
		}
	}
	return &api.CliplistItem{
		Clip: api.Clip{
			ModelBase: api.ModelBase{
				ID:   clip.ID,
				Type: "CliplistItem",
			},
			Title:         clip.Title,
			Description:   clip.Description,
			BeginAt:       clip.BeginAt,
			EndAt:         clip.EndAt,
			FavoriteCount: len(clip.Favorites),
			Video:         ConvertToVideo(clip.Video),
		},
		Available: true,
	}
}

func ConvertToCliplistItems(cliplistContains []*model.CliplistContain) []*api.CliplistItem {
	res := make([]*api.CliplistItem, 0)
	for _, cliplistContain := range cliplistContains {
		res = append(res, ConvertToCliplistItem(cliplistContain))
	}
	return res
}
