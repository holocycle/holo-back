package api

type User struct {
	ModelBase
	Name    string `json:"name"`
	Email   string `json:"email"`
	IconURL string `json:"iconUrl"`
}

func UserModels() []interface{} {
	return []interface{}{
		User{},
	}
}
