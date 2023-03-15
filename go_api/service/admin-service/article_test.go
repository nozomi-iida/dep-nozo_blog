package adminservice_test

import (
	"testing"

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
		test string
		query article.ArticleQuery
		expectedCount int
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "Success to list articles",
			query: article.ArticleQuery{},
			expectedCount: 3,
			expectedErr: nil,
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

