package entity_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/entity"
	"github.com/nozomi-iida/nozo_blog_go_api/valueobject"
)

func TestUser_NewUser(t *testing.T) {
	type testCase struct {
		test        string
		username    string
		password    string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Username validation",
			username:    "",
			password:    "",
			expectedErr: entity.ErrInvalidUser,
		},
		{
			test:        "Valid User",
			username:    "test",
			password:    "password",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			ps, err := valueobject.NewPassword(tc.password)
			_, err = entity.NewUser(tc.username, ps)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
