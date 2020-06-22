package api

type ModelBase struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type PageInfo struct {
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"currentPage"`
	ItemPerPage int `json:"itemPerPage"`
}

func Models() []interface{} {
	models := []interface{}{
		ModelBase{},
		PageInfo{},
	}
	models = append(models, LiverModels()...)
	models = append(models, ClipModels()...)
	models = append(models, CommentModels()...)
	models = append(models, TagModels()...)
	models = append(models, UserModels()...)
	models = append(models, VideoModels()...)
	models = append(models, FavoriteModels()...)
	models = append(models, CliplistModels()...)
	return models
}
