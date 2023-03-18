package service_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/service"
	"github.com/nozomi-iida/nozo_blog/test"
)

func TestAuth_SignUp(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	as, err := service.NewAuthService(
		service.WithSqliteUserRepository(ts.Filename),
	)
	if err != nil {
		t.Fatal(err)
	}

	username := "nozomi"
	password := "password123"
	_, err = as.SignUp(username, password)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuth_SignIn(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	as, err := service.NewAuthService(
		service.WithSqliteUserRepository(ts.Filename),
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = as.SignIn(us.GetUsername(), "password123")
	if err != nil {
		t.Fatal(err)
	}
}
