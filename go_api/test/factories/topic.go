package factories

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type topicOptions func(*entity.Topic)

func SetTopicName(name string) topicOptions {
	return func(t *entity.Topic) {
		t.SetName(name)
	}
}

func CreateTopic(t *testing.T, fileName string, options ...topicOptions) entity.Topic {
	topic, err := entity.NewTopic(entity.Topic{Name: "test topic", Description: "description"})
	for _, op := range options {
		op(&topic)
	}

	sq, err := sqlite.New(fileName)
	_, err = sq.Create(topic)
	if err != nil {
		t.Error("create topic err:", err)
	}

	return topic
}
