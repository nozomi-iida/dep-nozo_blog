package sqlite_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

// TODO: testCaseを増やす
func TestUserSqlite_FindById(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	us := test.CreateUser(t, ts.Filename)
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("create user err:", err)
	}	
	rs, err := sq.FindById(us.GetID())
	if rs.GetID() != us.GetID() {
		t.Errorf("expected id %v, got %v", us.GetID(), rs.GetID())
	}
}

func TestUserSqlite_FindByUsername(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	username := "hoge"
	us := test.CreateUser(t, ts.Filename, test.SetUsername(username))
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("sqlite new err:", err)
	}

	type testCase struct {
		test string
		name string
		expectedErr error
	}
	
	testCases := []testCase{
		{
			test: "Success to get user",	
			name: username,
			expectedErr: nil,
		},
		{
			test: "Failed to get user",	
			name: "fuga",
			expectedErr: user.ErrUserNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			rs, err := sq.FindByUsername(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			if err ==nil && rs.GetUsername() != us.GetUsername() {
				t.Error("find by username error", err)
			}
		})
	}
}

func TestUserSqlite_Create(t *testing.T) {
	ts := test.ConnectDB(t)
	defer ts.Remove()
	sq, err := sqlite.New(ts.Filename)
	if err != nil {
		t.Error("sqlite new err:", err)
	}	
	ps, err := valueobject.NewPassword("password123")
	us, err := entity.NewUser("nozomi", ps)
	if err != nil {
		t.Error("newUser err:", err)
	}
	type testCase struct {
		test string
		user entity.User
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Success to create user",	
			user: us,
			expectedErr: nil,
		},
		{
			test: "Failed to create user",
			user: us,
			expectedErr: user.ErrUserAlreadyExist,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err = sq.Create(tc.user)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
