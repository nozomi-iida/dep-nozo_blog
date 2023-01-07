package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidTopic = errors.New("A Topic has to have an a valid topic")
)

type Topic struct {
	TopicID uuid.UUID
	Name string
	Description string
}

type TopicOption func(t *Topic)

func SetDescription(description string) TopicOption {
	return func(t *Topic) {
		t.Description = description	
	}
}

// description to optional argument
func NewTopic(name string, opts ...TopicOption) (Topic, error)  {
	if name == "" {
		return Topic{}, ErrInvalidTopic
	}

	topic := Topic{
		TopicID: uuid.New(),
		Name: name,
	}

	for _, opt := range opts {
		opt(&topic)
	}

	return topic, nil
}

func (t *Topic) SetTopicId(id uuid.UUID)  {
	t.TopicID = id
}

func (t *Topic) SetName(name string)  {
	if name == "" {
		return
	}
	t.Name = name
}

func (t *Topic) SetDescription(description string)  {
	if description == "" {
		return
	}
	t.Description = description
}
