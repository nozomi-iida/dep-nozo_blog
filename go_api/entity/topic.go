package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidTopic = errors.New("A Topic has to have an a valid topic")
)

type Topic struct {
	TopicID     uuid.UUID
	Name        string
	Description string
}

func NewTopic(topic Topic) (Topic, error) {
	if topic.Name == "" {
		return Topic{}, ErrInvalidTopic
	}

	if topic.TopicID.ID() == 0 {
		topic.TopicID = uuid.New()
	}

	return topic, nil
}

func (t *Topic) SetTopicId(id uuid.UUID) {
	t.TopicID = id
}

func (t *Topic) SetName(name string) {
	if name == "" {
		return
	}
	t.Name = name
}

func (t *Topic) SetDescription(description string) {
	if description == "" {
		return
	}
	t.Description = description
}
