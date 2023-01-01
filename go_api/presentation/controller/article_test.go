package controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/middleware"
	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestArticleController_PostRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	json := strings.NewReader(`{"title": "test", "content": "test", "isPublic": true, "tags": ["tag_1"]}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/articles/post", json)
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
