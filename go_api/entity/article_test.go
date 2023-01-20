package entity_test

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

func TestArticle_NewArticle(t *testing.T) {
	type testCase struct {
		test string
		article entity.ArticleArgument
		expectedErr error
	}

	user, _ := entity.NewUser("Nozomi", valueobject.Password{})

	testCases := []testCase{
		{
			test: "Empty Title validation",
			article: entity.ArticleArgument{
				Title: "",
				Content: "test",
				AuthorID: user.GetID(),
			},
			expectedErr: entity.ErrInvalidArticle,
		},
		{
			test: "Empty AuthorId validation",
			article: entity.ArticleArgument{
				Title: "test",
				Content: "test",
			},
			expectedErr: entity.ErrInvalidArticle,
		},
		{
			test: "Too many tags",
			article: entity.ArticleArgument{
				Title: "test",
				Content: "test",
				AuthorID: user.GetID(),
				Tags: []string{
					"tag_1",
					"tag_2",
					"tag_3",
					"tag_4",
				},
			},
			expectedErr: entity.ErrTooManyTags,
		},
		{
			test: "valid article",
			article: entity.ArticleArgument{
				Title: "test",
				Content: "test",
				AuthorID: user.GetID(),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := entity.NewArticle(tc.article)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
