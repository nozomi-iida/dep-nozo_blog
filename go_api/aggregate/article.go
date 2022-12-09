package aggregate

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrInvalidArticle = errors.New("A Article has to have an a valid article")
	ErrTooManyTags = errors.New("The maximum number of tags an article can have is 5")
	ErrNoAuthor = errors.New("")
)

type Article struct {
	ID uuid.UUID
	Title string
	Content string
	PublishedAt *time.Time
	Author *entity.User
	Topic *entity.Topic
	Tags *[]string
}

func NewArticle(title string, content string, tags []string, author entity.User, topic *entity.Topic) (Article, error)  {
	if title == "" || content == "" {
		return Article{}, ErrInvalidArticle
	}

	if len(tags) > 5 {
		return Article{}, ErrTooManyTags
	}	

	return Article{
		Title: title,
		Content: content,
		Topic: topic,
		Tags: &tags,
		Author: &author,
	}, nil
}
