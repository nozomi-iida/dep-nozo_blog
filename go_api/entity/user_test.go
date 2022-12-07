package entity_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/entity"
)

func TestUser_NewUser(t *testing.T) {
	type testCase struct {
		test string
		username string
		email string
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Empty Username validation",
			username: "",
			email: "test@test.com",
			expectedErr: entity.ErrInvalidUser,
		},
		{
			test: "Empty Email validation",
			username: "test",
			email: "",
			expectedErr: entity.ErrInvalidUser,
		},
		{
			test: "Invalid Email",
			username: "test",
			email: "hoge",
			expectedErr: entity.ErrInvalidEmail,
		},
		{
			test: "Valid User",
			username: "test",
			email: "test@test.com",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := entity.NewUser(tc.username, tc.email)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
