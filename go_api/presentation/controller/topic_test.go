package controller_test

import (
	"fmt"
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
