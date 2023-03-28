package topic

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrTopicNotFound         = errors.New("the topic was not found in the repository")
	ErrTopicAlreadyExist     = errors.New("topic already exits")
	ErrFailedToListTopics    = errors.New("failed to get the topics to the repository")
	ErrInvalidAssociatedType = errors.New("invalid associated type")
)

type TopicRepository interface {
	Create(topic entity.Topic) error
	Update(topic entity.Topic) error
	List() ([]entity.Topic, error)
}
