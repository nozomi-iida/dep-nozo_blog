package adminservice

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type topicConfiguration func(as *TopicService) error

type TopicService struct {
	ap topic.TopicRepository
}

type TopicDto struct {
	TopicID     uuid.UUID `json:"topicId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type TopicListDto struct {
	Topics []TopicDto `json:"topics"`
}

func topicEntitiesToListDto(topics []entity.Topic) TopicListDto {
	t := TopicListDto{}
	for _, topic := range topics {
		topicDto := TopicDto{
			TopicID:     topic.TopicID,
			Name:        topic.Name,
			Description: topic.Description,
		}

		t.Topics = append(t.Topics, topicDto)
	}

	return t
}

func NewTopicService(cfgs ...topicConfiguration) (*TopicService, error) {
	aas := &TopicService{}

	for _, cfg := range cfgs {
		err := cfg(aas)
		if err != nil {
			return nil, err
		}
	}

	return aas, nil
}

func WithSqliteTopicRepository(fileString string) topicConfiguration {
	return func(aas *TopicService) error {
		s, err := sqlite.New(fileString)
		if err != nil {
			return err
		}

		aas.ap = s

		return nil
	}
}

func (as *TopicService) List() (TopicListDto, error) {
	t, err := as.ap.List()
	topicDtos := topicEntitiesToListDto(t)
	return topicDtos, err
}

func (as *TopicService) Create(name string, description string) error {
	tp, err := entity.NewTopic(entity.TopicArgument{Name: name, Description: description})
	err = as.ap.Create(tp)
	if err != nil {
		return err
	}

	return nil
}

func (as *TopicService) Update(topicID uuid.UUID, name string, description string) error {
	tp, err := entity.NewTopic(entity.TopicArgument{TopicID: topicID, Name: name, Description: description})
	err = as.ap.Update(tp)
	if err != nil {
		return err
	}

	return nil
}
