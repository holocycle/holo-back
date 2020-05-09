package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToLiver(liver *model.Liver) *api.Liver {
	return &api.Liver{
		ModelBase: api.ModelBase{
			Type: "Liver",
			ID:   liver.ID,
		},
		Name:      liver.Name,
		ChannelID: liver.ChannelID,
		MainColor: liver.MainColor,
		SubColor:  liver.SubColor,
	}
}

func ConvertToLivers(livers []*model.Liver) []*api.Liver {
	res := make([]*api.Liver, 0)
	for _, liver := range livers {
		res = append(res, ConvertToLiver(liver))
	}
	return res
}
