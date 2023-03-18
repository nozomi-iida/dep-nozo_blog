package service_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/service"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestArticle_Post(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(ts.Filename),
	)

	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test        string
		title       string
		content     string
		isPublic    bool
		tags        []string
		authorId    uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Public article",
			title:       "Test",
			content:     "Test",
			isPublic:    true,
			tags:        []string{"tag"},
			authorId:    us.GetID(),
			expectedErr: nil,
		},
		{
			test:        "Private article",
			title:       "Test",
			content:     "Test",
			isPublic:    false,
			tags:        []string{"tag"},
			authorId:    us.GetID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ac, err := as.Post(tc.title, tc.content, tc.tags, tc.isPublic, tc.authorId, nil)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if tc.isPublic && ac.PublishedAt.IsZero() {
				t.Error("isPublic is true but publishedAt is null")
			}
		})
	}
}

func TestArticle_Delete(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ac := factories.CreateArticle(t, ts.Filename)
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(ts.Filename),
	)
	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test        string
		articleID   uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to delete article",
			articleID:   ac.ArticleID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := as.Delete(tc.articleID)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestArticle_FindById(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ac := factories.CreateArticle(t, ts.Filename)
	pa := factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(ts.Filename),
	)
	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test        string
		articleID   uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to find by id article",
			articleID:   ac.ArticleID,
			expectedErr: nil,
		},
		{
			test:        "Failed to find by id article",
			articleID:   pa.ArticleID,
			expectedErr: article.ErrArticleNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := as.FindById(tc.articleID)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestArticle_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	engineerAC := factories.CreateArticle(t, ts.Filename, factories.SetTitle("engineer"))
	factories.CreateArticle(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(ts.Filename),
	)
	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test          string
		query         article.ArticleQuery
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Success to list articles",
			query:         article.ArticleQuery{},
			expectedCount: 3,
			expectedErr:   nil,
		},
		{
			test:          "List from keyword",
			query:         article.ArticleQuery{Keyword: engineerAC.Title},
			expectedCount: 1,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			res, err := as.List(tc.query)

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedCount != len(res.Articles) {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(res.Articles))
			}
		})
	}
}
