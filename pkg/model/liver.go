package model

import (
	"time"
)

type Liver struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ChannelID string     `json:"channelId"`
	MainColor string     `json:"mainColor"`
	SubColor  string     `json:"subColor"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func NewLiver(id, name, channelID, mainColor, subColor string) *Liver {
	return &Liver{
		ID:        id,
		Name:      name,
		ChannelID: channelID,
		MainColor: mainColor,
		SubColor:  subColor,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}
