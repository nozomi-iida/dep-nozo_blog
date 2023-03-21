package service

import (
	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
)

type topicConfigurations func(tp *TopicService) error

type TopicService struct {
	tp topic.TopicRepository
	tc topic.TopicQueryService
}

func NewTopicService(cfgs ...topicConfigurations) (*TopicService, error) {
	ts := &TopicService{}

	for _, cfg := range cfgs {
		err := cfg(ts)
		if err != nil {
			return nil, err
		}
	}

	return ts, nil
}

func WithSqliteTopicRepository(fileString string) topicConfigurations {
	return func(ts *TopicService) error {
		s, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		ts.tp = s
		ts.tc = s

		return nil
	}
}

func (ts *TopicService) PublicList(query topic.TopicQuery) (topic.TopicListDto, error) {
	topics, err := ts.tc.PublicList(query)
	if err != nil {
		return topic.TopicListDto{}, err
	}

	return topics, nil
}

func (ts *TopicService) PublicFindByName(name string, query topic.PublicFindByNameQuery) (topic.TopicDto, error) {
	tp, err := ts.tc.PublicFindByName(name, query)
	if err != nil {
		return topic.TopicDto{}, err
	}

	return tp, nil
}
