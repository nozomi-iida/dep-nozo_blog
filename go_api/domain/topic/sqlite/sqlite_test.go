package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/domain/topic"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestTopicSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to create topic",
			name:        "test",
			expectedErr: nil,
		},
		{
			test:        "Failed to create topic, because topic already exist",
			name:        "test",
			expectedErr: topic.ErrTopicAlreadyExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tp, err := entity.NewTopic(entity.TopicArgument{Name: tc.name})
			err = sq.Create(tp)

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
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 2"))
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 3"))
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test          string
		query         topic.TopicQuery
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Get 3 topics",
			query:         topic.TopicQuery{},
			expectedCount: 3,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tld, err := sq.PublicList(tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && len(tld.Topics) != tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(tld.Topics))
			}
		})
	}
}

func TestTopicSqlite_PublicList(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	tt := factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))
	factories.CreateArticle(t, ts.Filename, factories.SetTopic(tt.TopicID))
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 3"))
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test          string
		query         topic.TopicQuery
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Get 3 topics",
			query:         topic.TopicQuery{},
			expectedCount: 3,
			expectedErr:   nil,
		},
		{
			test:          "Get targeted topic",
			query:         topic.TopicQuery{Keyword: "targeted"},
			expectedCount: 1,
			expectedErr:   nil,
		},
		{
			test:          "Get topics with article association",
			query:         topic.TopicQuery{AssociatedWith: topic.Article},
			expectedCount: 3,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tld, err := sq.PublicList(tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && len(tld.Topics) != tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(tld.Topics))
			}
		})
	}
}

func TestTopicSqlite_PublicFindByName(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, _ := sqlite.New(ts.Filename)
	targetTopic := factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))
	factories.CreateArticle(t, ts.Filename, factories.SetTopic(targetTopic.TopicID))

	type testCase struct {
		test        string
		name        string
		query       topic.PublicFindByNameQuery
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Get targeted topic",
			name:        targetTopic.Name,
			query:       topic.PublicFindByNameQuery{},
			expectedErr: nil,
		},
		{
			test:        "Get targeted topic with article association",
			name:        "targeted",
			query:       topic.PublicFindByNameQuery{AssociatedWith: topic.Article},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tp, err := sq.PublicFindByName(tc.name, tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tp.Name != tc.name {
				t.Errorf("Expected name %v, got %v", tc.name, tp.Name)
			}
			if err == nil && tc.query.AssociatedWith == topic.Article && len(tp.Articles) == 0 {
				t.Errorf("Expected article association, but got none")
			}
		})
	}
}

func TestTopicSqlite_Update(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, _ := sqlite.New(ts.Filename)
	targetTopic := factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))
	updatedTopic, _ := entity.NewTopic(
		entity.TopicArgument{
			TopicID:     targetTopic.TopicID,
			Name:        "updated",
			Description: targetTopic.Description,
		},
	)

	type testCase struct {
		test        string
		topic       entity.Topic
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Update targeted topic",
			topic:       updatedTopic,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := sq.Update(tc.topic)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
