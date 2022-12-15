package valueobject_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/valueobject"
)

func TestPassword_NewPassword(t *testing.T) {
	type testCase struct {
		test string
		plainText string
		expectedErr error
	}

	testCases := []testCase{
		{
			test: "Password is too short",
			plainText: "hoge",
			expectedErr: valueobject.ErrTooShortPassword,
		},
		{
			test: "Invalid password",
			plainText: "password",
			expectedErr: valueobject.ErrInvalidPassword,
		},
		{
			test: "Valid password",
			plainText: "password123",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := valueobject.NewPassword(tc.plainText)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
