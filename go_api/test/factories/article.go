package factories

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog/domain/article/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/test"
)

type articleOptions func(*entity.Article)

func SetTitle(title string) articleOptions {
	return func(a *entity.Article) {
		a.Title = title
	}	
}

func CreateArticle(t *testing.T, fileName string, options ...articleOptions) entity.Article {
	user := test.CreateUser(t, fileName)
	topic := CreateTopic(t, fileName)
	a, err := entity.NewArticle(entity.Article{
		Title: "test article", 
		Content: "content", 
		Tags: []string{"tag_1", "tag_2"}, 
		AuthorID: user.GetID(),
		TopicID: &topic.TopicID,
	})
	article, err := entity.NewArticle(a)	
	sq, err := sqlite.New(fileName)
	_, err = sq.Create(article)
	if err != nil {
		t.Error("create user err:", err)
	}

	return article
}

