package factories

import (
	"fmt"
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

var topicCalled = 0

func CreateTopic(t *testing.T, fileName string, options ...topicOptions) entity.Topic {
	topic, err := entity.NewTopic(entity.Topic{Name: fmt.Sprintf("topic %v", topicCalled), Description: "description"})
	for _, op := range options {
		op(&topic)
	}

	sq, err := sqlite.New(fileName)
	_, err = sq.Create(topic)
	if err != nil {
		t.Error("create topic err:", err)
	}

	topicCalled++
	return topic
}
