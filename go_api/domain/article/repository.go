package article

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrArticleNotFound = errors.New("the article was not found in the repository")
	ErrFailedToCreateArticle = errors.New("failed to create the article to the repository")
	ErrFailedToUpdateArticle = errors.New("failed to update the article to the repository")
	ErrFailedToGetArticle = errors.New("failed to get the article to the repository")
	ErrFailedToListArticle = errors.New("failed to get the articles to the repository")
	ErrFailedToDeleteArticle = errors.New("failed to delete the article to the repository")
)

type ArticleDto struct {
	ArticleID uuid.UUID `json:"articleId"`
	Title string `json:"title"`
	Content string `json:"content"`
	PublishedAt sql.NullString `json:"publishedAt,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Topic sql.NullString `json:"topic,omitempty"`
	AuthorName string `json:"author"`
}

type ListArticleDto struct {
	Articles []ArticleDto `json:"articles"`
}

type ArticleQuery struct {
	Keyword string
}

// repositoryからはentityは返さない方が良い気がするけど、良い詰替え方が分からないのでこのまま
// 詰め替えたい理由
// 1. jsonタグをentityにつけたくない
// 2. repositoryの後にentityを維持れるのってどうなの？
type ArticleRepository interface {
	List(query ArticleQuery) (ListArticleDto, error)
	Create(entity.Article) (entity.Article, error) 
	FindById(id uuid.UUID) (ArticleDto, error)
	Update(entity.Article) (entity.Article, error) 
	Delete(id uuid.UUID) error 
}
