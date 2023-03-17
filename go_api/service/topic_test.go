package service_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/service"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestTopicService_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	tps, err := service.NewTopicService(
		service.WithSqliteTopicRepository(ts.Filename),
	)

	if err != nil {
		t.Errorf("service error: %v", err)
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
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := tps.Create(tc.name, "test")
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

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
		test string
		expectedCount int 
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "Get 3 topics",
			expectedCount: 3,
			expectedErr: nil,
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
