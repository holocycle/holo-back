package api

type Liver struct {
	ModelBase
	Name      string `json:"name"`
	ChannelID string `json:"channelId"`
	MainColor string `json:"mainColor"`
	SubColor  string `json:"subColor"`
}

type ListLiversRequest struct {
}

type ListLiversResponse struct {
	Livers []*Liver `json:"livers"`
}

type GetLiverRequest struct {
}

type GetLiverResponse struct {
	Liver *Liver `json:"liver"`
}

func LiverModels() []interface{} {
	return []interface{}{
		Liver{},
		ListLiversRequest{},
		ListLiversResponse{},
		GetLiverRequest{},
		GetLiverResponse{},
	}
}
