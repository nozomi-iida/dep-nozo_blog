package topic

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type AssociatedType string

const (
	Article AssociatedType = "article"
)

func NewAssociatedType(associate string) (AssociatedType, error) {
	switch associate {
	case "article":
		return Article, nil
	default:
		return "", ErrInvalidAssociatedType
	}	
}

type TopicQuery struct {
	Keyword string
	// 配列の方が良いかも
	AssociatedWith AssociatedType
}

type TopicDto struct {
	TopicID uuid.UUID
	Name string
	Description string
	Articles []entity.Article
}

type TopicListDto struct {
	Topics []TopicDto
}

type TopicQueryService interface {
	PublicList(query TopicQuery) (TopicListDto, error)
}
