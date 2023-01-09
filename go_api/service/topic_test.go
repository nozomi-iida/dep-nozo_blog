package service_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/service"
	"github.com/nozomi-iida/nozo_blog/test"
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
