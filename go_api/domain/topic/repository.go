package topic

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrTopicNotFound = errors.New("the topic was not found in the repository")
	ErrTopicAlreadyExist = errors.New("topic already exits")
	ErrFailedToListTopics = errors.New("failed to get the topics to the repository")
	ErrInvalidAssociatedType = errors.New("invalid associated type")
)

type AssociatedType string

const (
	Article AssociatedType = "article"
)

func NewAssociatedType(associate string) (AssociatedType, error) {
	switch associate {
	case "article":
		return Article, nil
	default:
		return "", ErrInvalidAssociatedType
	}	
}

type TopicQuery struct {
	Keyword string
	AssociatedWith AssociatedType
}

type TopicRepository interface {
	List(query TopicQuery) ([]entity.Topic, error)
	Create(topic entity.Topic) (entity.Topic, error)
}
