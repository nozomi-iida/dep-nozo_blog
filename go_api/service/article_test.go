package service_test

import (
	"testing"

	"github.com/google/uuid"
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
		test string
		title string
		content string
		isPublic bool
		tags []string
		authorId uuid.UUID
		expectedErr error
	}	

	testCases := []testCase {
		{
			test: "Public article",
			title: "Test",
			content: "Test",
			isPublic: true,
			tags: []string{"tag"},
			authorId: us.GetID(),
			expectedErr: nil,
		},
		{
			test: "Private article",
			title: "Test",
			content: "Test",
			isPublic: false,
			tags: []string{"tag"},
			authorId: us.GetID(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ac, err := as.Post(tc.title, tc.content, tc.tags, tc.isPublic, tc.authorId)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if tc.isPublic && ac.PublishedAt.IsZero() {
				t.Error("isPublic is true but publishedAt is null")
			}
		})
	}
}

func TestArticle_Update(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	ac := factories.CreateArticle(t, ts.Filename)
	var tags []string
	for _, tag := range ac.Tags {
		tags = append(tags, tag.Name)
	}
	as, err := service.NewArticleService(
		service.WithSqliteArticleRepository(ts.Filename),
	)
	if err != nil {
		t.Errorf("service error: %v", err)
	}

	type testCase struct {
		test string
		articleID uuid.UUID
		title string
		content string
		isPublic bool
		tags []string
		expectedErr error
	}	

	testCases := []testCase {
		{
			test: "Update article",
			articleID: ac.ArticleID,
			title: "Update Title",
			content: ac.Content,
			isPublic: ac.PublishedAt.IsZero(),
			tags: tags,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ac, err = as.Update(tc.articleID, tc.title, tc.content, tc.tags, tc.isPublic)
		})
	}
}
