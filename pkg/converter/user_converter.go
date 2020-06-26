package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToUser(user *model.User) *api.User {
	return &api.User{
		ModelBase: api.ModelBase{
			Type: "LoginUser",
			ID:   user.ID,
		},
		Name:    user.Name,
		IconURL: user.IconURL,
	}
}

func ConvertToLoginUser(user *model.User) *api.LoginUser {
	convertedUser := ConvertToUser(user)
	return &api.LoginUser{
		User:  *convertedUser,
		Email: user.Email,
	}
}

func ConvertToFavoriteClips(favorites []*model.Favorite) []*api.Clip {
	res := make([]*api.Clip, 0)
	for _, favorite := range favorites {
		clip := favorite.Clip
		res = append(res, ConvertToClip(clip, clip.Video, nil))
	}
	return res
}
