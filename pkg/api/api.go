package api

type ModelBase struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func Models() []interface{} {
	models := []interface{}{
		ModelBase{},
	}
	models = append(models, VideoModels()...)
	models = append(models, ClipModels()...)
	return models
}
