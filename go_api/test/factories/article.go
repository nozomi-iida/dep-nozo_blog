package factories

import (
	"fmt"
	"testing"
	"time"

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

func SetPublishedAt(time *time.Time) articleOptions {
	return func(a *entity.Article) {
		a.PublishedAt = time
	}	
}

var called = 0

func CreateArticle(t *testing.T, fileName string, options ...articleOptions) entity.Article {
	user := test.CreateUser(t, fileName, test.SetUsername(fmt.Sprintf("user%v", called)))
	topic := CreateTopic(t, fileName, SetTopicName(fmt.Sprintf("user%v", called)))
	now := time.Now()
	a, err := entity.NewArticle(entity.ArticleArgument{
		Title: "test article", 
		Content: "content", 
		Tags: []string{
			"tag_1",
			"tag_2",
		}, 
		PublishedAt: &now,
		AuthorID: user.GetID(),
		TopicID: &topic.TopicID,
	})
	for _, op := range options {
		op(&a)
	}
	
	sq, err := sqlite.New(fileName)
	_, err = sq.Create(a)
	if err != nil {
		t.Error("create user err:", err)
	}

	called = called + 1
	return a
}

