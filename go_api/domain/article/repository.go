package article

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrArticleNotFound = errors.New("the article was not found in the repository")
	ErrFailedToCreateArticle = errors.New("failed to create the article to the repository")
	ErrFailedToGetArticle = errors.New("failed to get the article to the repository")
)

type ArticleDto struct {
	ArticleID uuid.UUID `json:"article_id" db:"article_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Topic string `json:"topic,omitempty"`
}

type ArticleRepository interface {
	Create(entity.Article) (entity.Article, error) 
	FindById(id uuid.UUID) (ArticleDto, error)
}
