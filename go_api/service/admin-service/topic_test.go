package adminservice_test

import (
	"testing"

	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestTopic_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	factories.CreateTopic(t, ts.Filename)
	type testCase struct {
		test          string
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Success to list topic",
			expectedCount: 3,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
		})
	}
}

func TestTopicService_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	tps, err := adminservice.NewTopicService(
		adminservice.WithSqliteTopicRepository(ts.Filename),
	)

	if err != nil {
		t.Errorf("service error: %v", err)
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
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := tps.Create(tc.name, "test")
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
