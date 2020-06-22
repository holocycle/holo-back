package model

import "time"

type CliplistStatus string

const (
	CliplistStatusPublic  CliplistStatus = "PUBLIC"
	CliplistStatusDeleted CliplistStatus = "DELETED"
)

type Cliplist struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Status      CliplistStatus
	CreatedAt   *time.Time
	UpdatedAt   *time.Time

	CliplistContains []*CliplistContain
}

func NewCliplist(
	userID,
	title,
	description string,
	status CliplistStatus,
) *Cliplist {
	return &Cliplist{
		ID:          GetIDGenerator().New(),
		UserID:      userID,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}
}
