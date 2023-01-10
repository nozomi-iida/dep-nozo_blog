package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidArticle = errors.New("A Article has to have an a valid article")
	ErrTooManyTags = errors.New("The maximum number of tags an article can have is 5")
)

type Article struct {
	ArticleID uuid.UUID
	Title string
	Content string
	Tags []string
	PublishedAt *time.Time
	AuthorID uuid.UUID	
	// 小文字にしたい？
	TopicID *uuid.UUID
}

func NewArticle(ac Article) (Article, error)  {
	if ac.Title == "" || ac.Content == "" || ac.AuthorID.ID() <= 0{
		return Article{}, ErrInvalidArticle
	}	

	if len(ac.Tags) > 3 {
		return Article{}, ErrTooManyTags
	}	
	ac.ArticleID = uuid.New()

	return ac, nil
}
