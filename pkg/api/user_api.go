package api

type User struct {
	ModelBase
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

type LoginUser struct {
	User
	Email string `json:"email"`
}

type ListUserRequest struct {
	Limit   int    `json:"limit"   validate:"min=0,max=100"`
	OrderBy string `json:"orderBy" validate:"oneof=any latest"`
}

type ListUserResponse struct {
	Users []*User `json:"users"`
}

type GetUserRequest struct {
}

type GetUserResponse struct {
	User User `json:"user"`
}

type GetLoginUserRequest struct {
}

type GetLoginUserResponse struct {
	LoginUser LoginUser `json:"loginUser"`
}

type GetUserFavoritesRequest struct {
}

type GetUserFavoritesResponse struct {
	FavoriteClips []*Clip `json:"favoriteClips"`
}

func UserModels() []interface{} {
	return []interface{}{
		User{},
		LoginUser{},
		ListUserRequest{},
		ListUserResponse{},
		GetUserRequest{},
		GetUserResponse{},
		GetLoginUserRequest{},
		GetLoginUserResponse{},
		GetUserFavoritesRequest{},
		GetUserFavoritesResponse{},
	}
}
