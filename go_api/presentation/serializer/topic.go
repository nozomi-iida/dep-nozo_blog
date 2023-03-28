package serializer

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
)

type TopicJson struct {
	TopicID     uuid.UUID `json:"topicId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func Topic2Json(topic entity.Topic) TopicJson {
	return TopicJson{
		TopicID:     topic.TopicID,
		Name:        topic.Name,
		Description: topic.Description,
	}
}
