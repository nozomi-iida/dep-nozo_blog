package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/domain/tag"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/tag/sqlite"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestTagSqlite_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	factories.CreateTag(t, ts.Filename)
	factories.CreateTag(t, ts.Filename, factories.SetTagName("hoge"))
	factories.CreateTag(t, ts.Filename, factories.SetTagName("tag 3"))
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test          string
		query         tag.TagQuery
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "list tags",
			query:         tag.TagQuery{},
			expectedCount: 3,
			expectedErr:   nil,
		},
		{
			test: "list tags with keyword",
			query: tag.TagQuery{
				Keyword: "test",
			},
			expectedCount: 1,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tags, err := sq.List(tc.query)
			if err != tc.expectedErr {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
			if len(tags) != tc.expectedCount {
				t.Errorf("expected count: %v, got: %v", tc.expectedCount, len(tags))
			}
		})
	}
}

func TestTagSqlite_Create(t *testing.T) {
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
			test:        "create tag",
			name:        "tag 1",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			tag, err := entity.NewTag(tc.name)
			_, err = sq.Create(tag)
			if err != tc.expectedErr {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
