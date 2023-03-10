package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/test/factories"
)

func TestTagController_ListRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	factories.CreateTag(t, ts.Filename)
	factories.CreateTag(t, ts.Filename, factories.SetTagName(("tag_2")))
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPatch, "/articles", nil)
	ac, err := controller.NewTagController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	ac.ListRequest(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Response code is %v", w.Code)
	}
}
