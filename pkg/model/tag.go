package model

import "time"

type Tag struct {
	ID        string
	Name      string
	Color     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewTag(name, color string) *Tag {
	return &Tag{
		ID:    NewID(),
		Name:  name,
		Color: color,
	}
}
