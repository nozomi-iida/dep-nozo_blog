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
	ErrFailedToUpdateArticle = errors.New("failed to update the article to the repository")
	ErrFailedToGetArticle = errors.New("failed to get the article to the repository")
	ErrFailedToListArticle = errors.New("failed to get the articles to the repository")
	ErrFailedToDeleteArticle = errors.New("failed to delete the article to the repository")
	ErrInvalidOrderType = errors.New("invalid order type")
)

type ArticleDto struct {
	ArticleID uuid.UUID
	Title string
	Content string
	PublishedAt *time.Time
	Tags []entity.Tag
	Topic *entity.Topic
	Author entity.User
}

type ListArticleDto struct {
	Articles []ArticleDto
}

type OrderType string

const (
	PublishedAtDesc OrderType = "published_at_desc"
	PublishedAtAsc  OrderType = "published_at_asc"
)

func NewOrderType(order string) (OrderType, error) {
	switch order {
	case "published_at_desc":
		return PublishedAtDesc, nil
	case "published_at_asc":
		return PublishedAtAsc, nil
	default:
		return "", ErrInvalidOrderType
	}
}

type ArticleQuery struct {
	Keyword string
	WithDraft bool
	OrderBy OrderType
}

// repositoryからはentityは返さない方が良い気がするけど、良い詰替え方が分からないのでこのまま
// 詰め替えたい理由
// 1. jsonタグをentityにつけたくない
// 2. repositoryの中にentityを維持れるのってどうなの？
type ArticleRepository interface {
	List(query ArticleQuery) (ListArticleDto, error)
	Create(entity.Article) (entity.Article, error) 
	FindById(id uuid.UUID) (ArticleDto, error)
	Update(entity.Article) error
	Delete(id uuid.UUID) error 
}
