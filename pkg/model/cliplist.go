package model

import "time"

type Cliplist struct {
	ID          string
	UserID      string
	Title       string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time

	Clip []*Clip
}

func NewCliplist(
	userID,
	title,
	description string,
) *Cliplist {
	return &Cliplist{
		ID:          NewID(),
		UserID:      userID,
		Title:       title,
		Description: description,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
}
