package article

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/aggregate"
)

var (
	ErrArticleNotFound = errors.New("the article was not found in the repository")
	ErrFailedToCreateArticle = errors.New("failed to create the article to the repository")
	ErrFailedToUpdateArticle = errors.New("failed to update the article to the repository")
	ErrFailedToDeleteArticle = errors.New("failed to delete the article to the repository")
)

type ArticleRepository interface {
	FindAll() ([]aggregate.Article, error)
	FindById() (aggregate.Article, error)
	Create(aggregate.Article) error
	Update(aggregate.Article) error
	Delete(id int) error
}
