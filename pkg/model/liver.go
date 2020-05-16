package model

import (
	"time"
)

type Liver struct {
	ID        string
	Name      string
	ChannelID string
	MainColor string
	SubColor  string
	CreatedAt *time.Time
	UpdatedAt *time.Time

	Channel *Channel
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
