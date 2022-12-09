package aggregate_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/aggregate"
	"github.com/nozomi-iida/nozo_blog/entity"
)

func TestArticle_NewArticle(t *testing.T) {
	type testCase struct {
		test string
		title string
		content string
		author *entity.User
		topic *entity.Topic
		tags []string
		expectedErr error
	}

	user, _ := entity.NewUser("Nozomi")

	testCases := []testCase{
		{
			test: "Empty Title validation",
			title: "",
			content: "",	
			author: &user,
			tags: nil,
			expectedErr: aggregate.ErrInvalidArticle,
		},
		{
			test: "Too many tags",
			title: "test",
			content: "test",	
			author: &user,
			tags: []string{"a", "b", "c", "d", "e", "f"},
			expectedErr: aggregate.ErrTooManyTags,
		},
		{
			test: "valid article",
			title: "test",
			content: "test",	
			author: &user,
			tags: nil,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewArticle(tc.title, tc.content, tc.tags, *tc.author, tc.topic)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
