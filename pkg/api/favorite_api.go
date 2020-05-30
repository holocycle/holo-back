package api

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
		PutFavoriteRequest{},
		PutFavoriteResponse{},
		DeleteFavoriteRequest{},
		DeleteFavoriteResponse{},
	}
}
