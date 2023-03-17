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
	Tags []Tag
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

	return Article{
		ArticleID: uuid.New(),
		Title: aa.Title,
		Content: aa.Content,
		PublishedAt: aa.PublishedAt,
		Tags: aa.Tags,
		AuthorID: aa.AuthorID,
		TopicID: aa.TopicID,
	}, nil
}

func (a *Article) Public() {
	now := time.Now()
	a.PublishedAt = &now
}

func (a *Article) SetID(id uuid.UUID)  {
	if id.ID() < 0 {
		return
	}
	a.ArticleID = id	
}

func (a *Article) SetTitle(title string)  {
	if title == "" {
		return
	}
	a.Title = title
}

func (a *Article) SetContent(content string)  {
	if content == "" {
		return
	}
	a.Content = content
}

// ちゃんとエラー返すようにするべきな気がする
func (a *Article) SetTags(tagNames []string) {
	var tags []Tag
	if len(tagNames) > 3 {
		return
	}
	for _, tag := range tagNames {
		nt, err := NewTag(tag)
		if err != nil {
			return
		}
		tags = append(tags, nt)
	}
	a.Tags = tags
}

func (a *Article) SetTopicID(id *uuid.UUID)  {
	if id == nil || id.ID() < 0 {
		return
	}
	a.TopicID = id	
}
