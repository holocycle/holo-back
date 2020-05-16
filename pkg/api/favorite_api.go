package api

type Favorite struct {
	ModelBase
	ClipID string `json:"clipId"`
	UserID string `json:"userId"`
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
		Favorite{},
		PutFavoriteRequest{},
		PutFavoriteResponse{},
		DeleteFavoriteRequest{},
		DeleteFavoriteResponse{},
	}
}
