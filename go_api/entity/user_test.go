package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

func TestUser_NewUser(t *testing.T) {
	type testCase struct {
		test string
		id string
		username string
		email string
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Empty Username validation",
			id: "",
			username: "",
			expectedErr: entity.ErrInvalidUser,
		},
		{
			test: "Valid User",
			username: "test",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := entity.NewUser(tc.username, valueobject.Password{})
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
