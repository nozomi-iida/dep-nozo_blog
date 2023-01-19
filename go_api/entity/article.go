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
	ArticleID uuid.UUID `json:"article_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Tags []string `json:"tags,omitempty"`
	AuthorID uuid.UUID `json:"authorID"`
	TopicID *uuid.UUID `json:"topicID,omitempty"`
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
