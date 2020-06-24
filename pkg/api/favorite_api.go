package api

type GetFavoriteRequest struct {
}

type GetFavoriteResponse struct {
	Favorite bool `json:"favorite"`
}

type PutFavoriteRequest struct {
}

type PutFavoriteResponse struct {
}

type DeleteFavoriteRequest struct {
}

type DeleteFavoriteResponse struct {
}

func FavoriteModels() []interface{} {
	return []interface{}{
		GetFavoriteRequest{},
		GetFavoriteResponse{},
		PutFavoriteRequest{},
		PutFavoriteResponse{},
		DeleteFavoriteRequest{},
		DeleteFavoriteResponse{},
	}
}
