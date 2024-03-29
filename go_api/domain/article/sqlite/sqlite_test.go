package sqlite_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/article"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/article/sqlite"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestArticleSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	u := test.CreateUser(t, ts.Filename)
	tp := factories.CreateTopic(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}
	type testCase struct {
		test        string
		article     entity.Article
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Success to create user",
			article: entity.Article{
				ArticleID: uuid.New(),
				Title:     "test",
				Content:   "test",
				AuthorID:  u.GetID(),
				Tags: []entity.Tag{
					{TagID: uuid.New(), Name: "tag_1"},
					{TagID: uuid.New(), Name: "tag_2"},
				},
				TopicID: &tp.TopicID,
			},
			expectedErr: nil,
		},
		{
			test: "Success to create null topic article",
			article: entity.Article{
				ArticleID: uuid.New(),
				Title:     "test",
				Content:   "test",
				AuthorID:  u.GetID(),
				Tags: []entity.Tag{
					{TagID: uuid.New(), Name: "tag_1"},
					{TagID: uuid.New(), Name: "tag_2"},
				},
				TopicID: nil,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err = sq.Create(tc.article)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			sq.FindById(tc.article.ArticleID)
		})
	}
}

func TestArticleSqlite_Update(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	a := factories.CreateArticle(t, ts.Filename)
	topic := factories.CreateTopic(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test           string
		updatedArticle entity.Article
		expectedErr    error
	}
	a.SetTitle("update title")
	a.SetTopicID(&topic.TopicID)
	a.SetTags([]string{"update"})

	testCases := []testCase{
		{
			test:           "Success to update article",
			updatedArticle: a,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := sq.Update(tc.updatedArticle)
			ac, err := sq.FindById(tc.updatedArticle.ArticleID)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.updatedArticle.Title != ac.Title {
				t.Errorf("Expected id %v, got %v", tc.updatedArticle.Title, ac.Title)
			}
			fmt.Printf("topic: %v\n", ac.Topic)
			if err == nil && tc.updatedArticle.Tags[0].Name != ac.Tags[0].Name {
				t.Errorf("Expected tag %v, got %v", tc.updatedArticle.Tags[0].Name, ac.Tags[0].Name)
			}
		})
	}
}

func TestArticleSqlite_Delete(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	a := factories.CreateArticle(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test        string
		articleId   uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Success to delete article",
			articleId:   a.ArticleID,
			expectedErr: nil,
		},
		{
			test:        "Failed to delete article",
			articleId:   uuid.New(),
			expectedErr: article.ErrArticleNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := sq.Delete(tc.articleId)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestArticleSqlite_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	engineerArticle := factories.CreateArticle(t, ts.Filename, factories.SetTitle("engineer"))
	time1900 := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	oldestArticle := factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(&time1900), factories.SetTitle("oldest article"))
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test          string
		query         article.ArticleQuery
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "get Article",
			query:         article.ArticleQuery{},
			expectedCount: 2,
			expectedErr:   nil,
		},
		{
			test:          "get keyword engineer",
			query:         article.ArticleQuery{Keyword: engineerArticle.Title},
			expectedCount: 1,
			expectedErr:   nil,
		},
		{
			test:          "get all Article",
			query:         article.ArticleQuery{WithDraft: true},
			expectedCount: 3,
			expectedErr:   nil,
		},
		{
			test:          "get oldest by published_at",
			query:         article.ArticleQuery{OrderBy: article.PublishedAtAsc},
			expectedCount: 2,
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			rs, err := sq.List(tc.query)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && len(rs.Articles) != tc.expectedCount {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(rs.Articles))
			}
			if err == nil && tc.query.OrderBy == article.PublishedAtAsc && rs.Articles[0].Title != oldestArticle.Title {
				t.Errorf("Expected title %v, got %v", oldestArticle.Title, rs.Articles[0].Title)
			}
		})
	}
}

func TestArticleSqlite_FindById(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	a := factories.CreateArticle(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Errorf("sqlite error: %v", err)
	}

	type testCase struct {
		test        string
		articleId   uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Get article",
			articleId:   a.ArticleID,
			expectedErr: nil,
		},
		{
			test:        "Not found article",
			articleId:   uuid.New(),
			expectedErr: article.ErrArticleNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ac, err := sq.FindById(tc.articleId)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.articleId != ac.ArticleID {
				t.Errorf("Expected id %v, got %v", tc.articleId, ac.ArticleID)
			}
		})
	}
}
