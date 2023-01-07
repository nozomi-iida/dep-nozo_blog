package controller_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestTopicController_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	json := strings.NewReader(`{"name": "test", "description": "test"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/topics", json)
	tc, err := controller.NewTopicController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	tc.CreteRequest(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("Response code is %v", w.Code)
	}
}
