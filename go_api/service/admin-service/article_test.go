package adminservice_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	adminservice "github.com/nozomi-iida/nozo_blog/service/admin-service"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestArticle_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	factories.CreateArticle(t, ts.Filename, factories.SetTitle("engineer"))
	factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	factories.CreateArticle(t, ts.Filename)
	as, err := adminservice.NewArticleService(
		adminservice.WithSqliteArticleRepository(ts.Filename),
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
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			res, err := as.List()

			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedCount != len(res.Articles) {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(res.Articles))
			}
		})
	}
}

func TestArticle_FindById(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ac := factories.CreateArticle(t, ts.Filename)
	as, err := adminservice.NewArticleService(
		adminservice.WithSqliteArticleRepository(ts.Filename),
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

func TestArticle_Update(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ac := factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	var tags []string
	for _, tag := range ac.Tags {
		tags = append(tags, tag.Name)
	}
	as, err := adminservice.NewArticleService(
		adminservice.WithSqliteArticleRepository(ts.Filename),
	)
	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test        string
		articleID   uuid.UUID
		title       string
		content     string
		isPublic    bool
		tags        []string
		topicID     *uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Update article",
			articleID:   ac.ArticleID,
			title:       "Update Title",
			content:     ac.Content,
			isPublic:    true,
			tags:        tags,
			topicID:     nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err = as.Update(tc.articleID, tc.title, tc.content, tc.tags, tc.topicID, tc.isPublic)
			ac, err := as.FindById(tc.articleID)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && ac.Title != tc.title {
				t.Errorf("Expected title %v, got %v", tc.title, ac.Title)
			}
			if err == nil && tc.isPublic && ac.PublishedAt == nil {
				t.Errorf("Expected published")
			}
		})
	}
}
