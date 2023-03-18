package entity_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/entity"
)

func TestTag_NewTag(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: entity.ErrInvalidTag,
		},
		{
			test:        "Valid tag",
			name:        "test",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := entity.NewTag(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
