package admincontroller_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/presentation"
	"github.com/nozomi-iida/nozo_blog/presentation/serializer"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestArticleController_ListRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, testServer.URL+"/api/v1/admin/articles", nil)
	token, err := us.UserId.Encode();
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test string
		expectedCount int
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "get all Article",
			expectedCount: 3,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			var got article.ListArticleDto
			err = json.NewDecoder(resp.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedCount != len(got.Articles) {
				t.Errorf("Expected count %v, got %v", tc.expectedCount, len(got.Articles))
			}
		})
	}
}

func TestArticleController_FindByIdRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	ca := factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, testServer.URL+fmt.Sprintf("/api/v1/admin/articles/%v", ca.ArticleID), nil)
	token, err := us.UserId.Encode();
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test string
		expectedArticleId uuid.UUID
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "get Article",
			expectedArticleId: ca.ArticleID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			var got serializer.ArticleJson
			err = json.NewDecoder(resp.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedArticleId != got.ArticleID {
				t.Errorf("Expected id %v, got %v", tc.expectedArticleId, got.ArticleID)
			}
		})
	}
}

func TestArticleController_PatchRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	var r, _ = presentation.NewRouter(ts.Filename)
	testServer := httptest.NewServer(r)
	t.Cleanup(func() {
		testServer.Close()
	})
	us := test.CreateUser(t, ts.Filename)
	ca := factories.CreateArticle(t, ts.Filename, factories.SetPublishedAt(nil))
	updatedTitle := "update test"
	body := strings.NewReader(fmt.Sprintf(`{"title": "%s", "content": "test", "isPublic": true, "tags": ["tag_1"]}`, updatedTitle))
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPatch, testServer.URL+fmt.Sprintf("/api/v1/admin/articles/%v", ca.ArticleID), body)
	token, err := us.UserId.Encode();
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test string
		expectedArticleTitle string
		expectedErr error
	}

	testCases := []testCase {
		{
			test: "update Article",
			expectedArticleTitle: updatedTitle,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			resp, err := cli.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			var got entity.Article
			err = json.NewDecoder(resp.Body).Decode(&got)
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil && tc.expectedArticleTitle != got.Title {
				t.Errorf("Expected title %v, got %v", tc.expectedArticleTitle, got.Title)
			}
		})
	}
}
