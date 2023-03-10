package service_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/tag"
	"github.com/nozomi-iida/nozo_blog/service"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestTag_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	nts, err := service.NewTagService(
		service.WithSqliteTagRepository(ts.Filename),
	)
	factories.CreateTag(t, ts.Filename)
	factories.CreateTag(t, ts.Filename, factories.SetTagName("hoge"))
	factories.CreateTag(t, ts.Filename, factories.SetTagName("tag 3"))
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test string
		query tag.TagQuery
		expectedCount int
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "list tags",
			query: tag.TagQuery{},
			expectedCount: 3,
			expectedErr: nil,
		},
		{
			test: "list tags with keyword",
			query: tag.TagQuery{
				Keyword: "test",
			},
			expectedCount: 1,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			topics, err := nts.List(tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err != nil && len(topics) == tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(topics))
			}
		})
	}
}
