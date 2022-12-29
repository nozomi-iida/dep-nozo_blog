package auth_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nozomi-iida/nozo_blog/presentation/controller/auth"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestAuthController_SignUpRequest(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	json := strings.NewReader(`{"username": "nozomi", "password": "password123"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/sign_up", json)
	ac, err := auth.NewAuthController(ts.Filename)
	if err != nil {
		t.Errorf("Controller error %v", err)
	}
	ac.SignUpRequest(w, r)
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
