package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToTag(tag *model.Tag) *api.Tag {
	return &api.Tag{
		ModelBase: api.ModelBase{
			Type: "Tag",
			ID:   tag.ID,
		},
		Name:  tag.Name,
		Color: tag.Color,
	}
}

func ConvertToTags(tags []*model.Tag) []*api.Tag {
	res := make([]*api.Tag, 0)
	for _, tag := range tags {
		res = append(res, ConvertToTag(tag))
	}
	return res
}

func ConvertClipTaggedToTags(clipTaggedList []*model.ClipTagged) []*api.Tag {
	res := make([]*api.Tag, 0)
	for _, clipTagged := range clipTaggedList {
		res = append(res, ConvertToTag(clipTagged.Tag))
	}
	return res
}
