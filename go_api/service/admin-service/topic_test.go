package adminservice_test

import (
	"testing"

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
