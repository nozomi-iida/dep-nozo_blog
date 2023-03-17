package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestTopicSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}
	type testCase struct {
		test string
		name string
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "Success to create topic",
			name: "test",
			expectedErr: nil,
		},
		{
			test: "Failed to create topic, because topic already exist",
			name: "test",
			expectedErr: topic.ErrTopicAlreadyExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tp, err := entity.NewTopic(entity.Topic{Name: tc.name})
			_, err = sq.Create(tp)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestTopicSqlite_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 3"))
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test string
		query topic.TopicQuery
		expectedCount int 
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "Get 3 topics",
			query: topic.TopicQuery{},
			expectedCount: 3,
			expectedErr: nil,
		},
		{
			test: "Get targeted topic",
			query: topic.TopicQuery{Keyword: "targeted"},
			expectedCount: 1,
			expectedErr: nil,
		},
		{
			test: "Get topics with article association",
			query: topic.TopicQuery{AssociatedWith: topic.Article},
			expectedCount: 3,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			topics, err := sq.List(tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && len(topics) != tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(topics))
			}
			if err == nil && tc.query.AssociatedWith == topic.Article && len(topics[0]) == 0 {
				t.Error("Expected articles, got none")
			}
		})	
	}
}
