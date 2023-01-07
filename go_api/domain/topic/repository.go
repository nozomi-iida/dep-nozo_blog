package topic

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrTopicNotFound = errors.New("the topic was not found in the repository")
	ErrTopicAlreadyExist = errors.New("topic already exits")
)

type TopicRepository interface {
	Create(topic entity.Topic) (entity.Topic, error)
}
