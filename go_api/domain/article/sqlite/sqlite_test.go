package sqlite_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestArticleSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	u := test.CreateUser(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}
	fmt.Printf("user: %v\n", u.GetID())
	type testCase struct {
		test string
		article entity.Article
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Success to create user",
			article: entity.Article{ArticleID: uuid.New(), Title: "test", Content: "test", AuthorID: u.GetID(), Tags: []string{"tag_1", "tag_2"}},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err = sq.Create(tc.article)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
