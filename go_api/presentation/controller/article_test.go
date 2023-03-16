package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestArticleController_PostRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	json := strings.NewReader(`{"title": "test", "content": "test", "isPublic": true, "tags": ["tag_1"]}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/articles", json)
	ac, err := controller.NewArticleController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	// refactorしたい
	token, err := us.UserId.Encode();
	r.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	middleware.AuthMiddleware(http.HandlerFunc(ac.PostRequest)).ServeHTTP(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("Response code is %v", w.Code)
	}
}



// 最初ほぼ同じだから共通化したい
func TestArticleController_DeleteRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	a := factories.CreateArticle(t, ts.Filename)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/articles/%v", a.ArticleID), nil)
	ac, err := controller.NewArticleController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	token, err := us.UserId.Encode();
	r.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	middleware.AuthMiddleware(http.HandlerFunc(ac.DeleteRequest)).ServeHTTP(w, r)
	if w.Code != http.StatusNoContent {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestArticleController_FindByIdRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	a := factories.CreateArticle(t, ts.Filename)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/articles/%v", a.ArticleID), nil)
	ac, err := controller.NewArticleController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	token, err := us.UserId.Encode();
	r.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	middleware.AuthMiddleware(http.HandlerFunc(ac.FindByIdRequest)).ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestArticleController_ListRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/articles", nil)
	ac, err := controller.NewArticleController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	token, err := us.UserId.Encode();
	r.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, token))
	middleware.AuthMiddleware(http.HandlerFunc(ac.ListRequest)).ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Response code is %v", w.Code)
	}
}
