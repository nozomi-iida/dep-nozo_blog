package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
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
