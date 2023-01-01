package controller_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/presentation/controller"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestAuthController_SignUpRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	json := strings.NewReader(`{"username": "nozomi", "password": "password123"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sign_up", json)
	ac, err := controller.NewAuthController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	ac.SignUpRequest(w, r)
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestAuthController_SignInRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	test.CreateUser(t, ts.Filename)
	json := strings.NewReader(`{"username": "nozomi", "password": "password123"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sign_in", json)
	ac, err := controller.NewAuthController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	ac.SignInRequest(w, r)
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
