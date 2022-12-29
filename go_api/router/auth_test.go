package router_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/router"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestAuthRouter_HandleSignUpRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	json := strings.NewReader(`{"username": "nozomi", "password": "password123"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sign_up", json)

	ar, err := router.NewRouter(ts.Filename)
	if err != nil {
		t.Errorf("Router error %v", err)
	}
	ar.HandleSignUpRequest(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestAuthRouter_HandleSignInRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	test.CreateUser(t, ts.Filename)
	json := strings.NewReader(`{"username": "nozomi", "password": "password123"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sign_in", json)

	ar, err := router.NewRouter(ts.Filename)
	if err != nil {
		t.Errorf("Router error %v", err)
	}
	ar.HandleSignInRequest(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
