package admincontroller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	admincontroller "github.com/nozomi-iida/nozo_blog/presentation/controller/admin-controller"
	"github.com/nozomi-iida/nozo_blog/presentation/middleware"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestArticleController_ListRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	factories.CreateArticle(t, ts.Filename)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/articles", nil)
	ac, err := admincontroller.NewArticleController(ts.Filename)
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

