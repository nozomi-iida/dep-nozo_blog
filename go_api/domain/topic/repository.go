package topic

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrTopicNotFound = errors.New("the topic was not found in the repository")
	ErrTopicAlreadyExist = errors.New("topic already exits")
	ErrFailedToListTopics = errors.New("failed to get the articles to the repository")
)

type TopicQuery struct {
	Keyword string
}

type TopicRepository interface {
	List() ([]entity.Topic, error)
	Create(topic entity.Topic) (entity.Topic, error)
}
