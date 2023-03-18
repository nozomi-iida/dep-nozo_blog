package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidTag = errors.New("A Tag has to have an a valid tag")
)

type Tag struct {
	TagID uuid.UUID `json:"tag_id"`
	Name  string    `json:"name"`
}

func NewTag(name string) (Tag, error) {
	if name == "" {
		return Tag{}, ErrInvalidTag
	}

	return Tag{
		TagID: uuid.New(),
		Name:  name,
	}, nil
}

func (t *Tag) SetName(name string) {
	if name == "" {
		return
	}
	t.Name = name
}
