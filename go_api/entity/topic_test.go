package entity_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/entity"
)

func TestTopic_NewTopic(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Name blank",
			name:        "",
			description: "test",
			expectedErr: entity.ErrInvalidTopic,
		},
		{
			test:        "Valid topic",
			name:        "test",
			description: "test",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := entity.NewTopic(entity.Topic{Name: tc.name, Description: tc.description})
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
