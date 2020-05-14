package api

type ModelBase struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func Models() []interface{} {
	models := []interface{}{
		ModelBase{},
	}
	models = append(models, LiverModels()...)
	models = append(models, ClipModels()...)
	models = append(models, CommentModels()...)
	models = append(models, TagModels()...)
	models = append(models, UserModels()...)
	models = append(models, VideoModels()...)
	models = append(models, FavoriteModels()...)
	return models
}
