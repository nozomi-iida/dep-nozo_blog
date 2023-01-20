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
	Tags []Tag `json:"tags,omitempty"`
	AuthorID uuid.UUID `json:"authorID"`
	TopicID *uuid.UUID `json:"topicID,omitempty"`
}

type ArticleArgument struct {
	Title string
	Content string
	PublishedAt *time.Time
	Tags []string
	AuthorID uuid.UUID
	TopicID *uuid.UUID
}

func NewArticle(aa ArticleArgument) (Article, error)  {
	if aa.Title == "" || aa.Content == "" || aa.AuthorID.ID() <= 0{
		return Article{}, ErrInvalidArticle
	}	

	if len(aa.Tags) > 3 {
		return Article{}, ErrTooManyTags
	}	
	var tags []Tag

	for _, tag := range aa.Tags {
		nt, err := NewTag(tag)
		if err != nil {
			return Article{}, ErrInvalidArticle
		}
		tags = append(tags, nt)
	}

	return Article{
		ArticleID: uuid.New(),
		Title: aa.Title,
		Content: aa.Content,
		PublishedAt: aa.PublishedAt,
		Tags: tags,
		AuthorID: aa.AuthorID,
		TopicID: aa.TopicID,
	}, nil
}
