package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/presentation/controller"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
	"github.com/nozomi-iida/nozo_blog_go_api/test/factories"
)

func TestTopicController_List(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("Get", "/topics", nil)
	tc, err := controller.NewTopicController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	tc.ListRequest(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestTopicController_FindByNameRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	factories.CreateTopic(t, ts.Filename, factories.SetTopicName("targeted"))
	defer ts.Remove()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("Get", "/topics/targeted", nil)
	tc, err := controller.NewTopicController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	tc.FindByNameRequest(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Response code is %v", w.Code)
	}
}
