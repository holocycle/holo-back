package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToLiver(liver *model.Liver, channel *model.Channel) *api.Liver {
	return &api.Liver{
		ModelBase: api.ModelBase{
			Type: "Liver",
			ID:   liver.ID,
		},
		Name:      liver.Name,
		MainColor: liver.MainColor,
		SubColor:  liver.SubColor,
		Channel: api.Channel{
			ModelBase: api.ModelBase{
				Type: "Channel",
				ID:   channel.ID,
			},
			Title:              channel.ID,
			Description:        channel.Description,
			SmallThumbnailURL:  channel.SmallThumbnailURL,
			MediumThumbnailURL: channel.MediumThumbnailURL,
			LargeThumbnailURL:  channel.LargeThumbnailURL,
			SmallBannerURL:     channel.SmallBannerURL,
			MediumBannerURL:    channel.MediumBannerURL,
			LargeBannerURL:     channel.LargeBannerURL,
			ViewCount:          channel.VideoCount,
			CommentCount:       channel.CommentCount,
			SubscriberCount:    channel.SubscriberCount,
			VideoCount:         channel.VideoCount,
			PublishedAt:        channel.PublishedAt,
		},
	}
}

func ConvertToLivers(livers []*model.Liver) []*api.Liver {
	res := make([]*api.Liver, 0)
	for _, liver := range livers {
		if liver.Channel == nil {
			panic("no channel")
		}
		res = append(res, ConvertToLiver(liver, liver.Channel))
	}
	return res
}
