package middleware_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/presentation/middleware"
	"github.com/nozomi-iida/nozo_blog_go_api/test"
)

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func TestAuthMiddleware(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	u := test.CreateUser(t, ts.Filename)
	token, err := u.UserId.Encode()
	if err != nil {
		t.Errorf("encode error: %v\n", err.Error())
	}
	type testCase struct {
		test  string
		token string
		code  int
	}

	testCases := []testCase{
		{
			test:  "Unauthorized error",
			token: "",
			code:  http.StatusUnauthorized,
		},
		{
			test:  "Authorized",
			token: token,
			code:  http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/articles", nil)
			r.Header.Set("Authorization", fmt.Sprintf(`Bearer %s`, tc.token))
			middleware.AuthMiddleware(http.HandlerFunc(mockHandler)).ServeHTTP(w, r)
			if w.Code != tc.code {
				t.Errorf("Response code is %v", w.Code)
			}
		})
	}
}
