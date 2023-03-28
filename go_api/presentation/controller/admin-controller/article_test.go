package admincontroller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/domain/article"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation"
	admincontroller "github.com/nozomi-iida/nozo_blog_go_api/presentation/controller/admin-controller"
	"github.com/nozomi-iida/nozo_blog_go_api/presentation/serializer"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
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
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test          string
		expectedCount int
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "get all Article",
			expectedCount: 3,
			expectedErr:   nil,
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
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test              string
		expectedArticleId uuid.UUID
		expectedErr       error
	}

	testCases := []testCase{
		{
			test:              "get Article",
			expectedArticleId: ca.ArticleID,
			expectedErr:       nil,
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
	areq := admincontroller.ArticleRequest{
		Title:    "update title",
		Content:  "update content",
		IsPublic: true,
		TagNames: []string{"tag_1"},
	}
	jsonBody, err := json.Marshal(areq)
	body := bytes.NewBuffer(jsonBody)
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPatch, testServer.URL+fmt.Sprintf("/api/v1/admin/articles/%v", ca.ArticleID), body)
	token, err := us.UserId.Encode()
	req.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	if err != nil {
		t.Fatal(err)
	}
	cli := &http.Client{}
	type testCase struct {
		test           string
		articleRequest admincontroller.ArticleRequest
		expectedErr    error
	}

	testCases := []testCase{
		{
			test:           "update Article",
			articleRequest: areq,
			expectedErr:    nil,
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
			if err == nil && tc.articleRequest.Title != got.Title {
				t.Errorf("Expected title %v, got %v", tc.articleRequest.Title, got.Title)
			}
			if err == nil && len(got.Tags) <= 0 {
				t.Errorf("Expected tags %v, got %v", tc.articleRequest.TagNames, got.Tags)
			}
		})
	}
}
