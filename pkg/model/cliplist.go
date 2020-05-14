package model

import "time"

type Cliplist struct {
	ID           string
	CreateUserID string
	Title        string
	Description  string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time

	Clip []*Clip
}

func Newcliplist(
	createUserID,
	title,
	description string,
) *Cliplist {
	return &Cliplist{
		ID:           NewID(),
		CreateUserID: createUserID,
		Title:        title,
		Description:  description,
		CreatedAt:    nil,
		UpdatedAt:    nil,
	}
}
