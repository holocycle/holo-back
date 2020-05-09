package converter

import (
	"github.com/holocycle/holo-back/pkg/api"
	"github.com/holocycle/holo-back/pkg/model"
)

func ConvertToUser(user *model.User) *api.User {
	return &api.User{
		ModelBase: api.ModelBase{
			Type: "User",
			ID:   user.ID,
		},
		Name:  user.Name,
		Email: user.Email,
		// FIXME iconURLはDBに格納されている情報を取得してくる必要がある
		IconURL: "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo",
	}
}
