package article

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrArticleNotFound = errors.New("the article was not found in the repository")
	ErrFailedToCreateArticle = errors.New("failed to create the article to the repository")
)

type ArticleRepository interface {
	Create(entity.Article) (entity.Article, error) 
}
