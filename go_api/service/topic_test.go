package service_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/domain/topic"
	"github.com/nozomi-iida/nozo_blog_go_api/service"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestTopicService_PublicList(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	nts, err := service.NewTopicService(
		service.WithSqliteTopicRepository(ts.Filename),
	)
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 2"))
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("topic 3"))

	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test          string
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Get 3 topics",
			expectedCount: 3,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			topics, err := nts.PublicList(topic.TopicQuery{})
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err != nil && len(topics.Topics) == tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(topics.Topics))
			}
		})
	}
}

func TestTopicService_PublicFindByName(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	nts, err := service.NewTopicService(
		service.WithSqliteTopicRepository(ts.Filename),
	)
	targetedTopic := factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))

	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test        string
		name        string
		query       topic.PublicFindByNameQuery
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Get 3 topics",
			name:        targetedTopic.Name,
			query:       topic.PublicFindByNameQuery{},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			to, err := nts.PublicFindByName(tc.name, topic.PublicFindByNameQuery{})
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && to.Name != tc.name {
				t.Errorf("Expected name %v, got %v", tc.name, to.Name)
			}
		})
	}
}
