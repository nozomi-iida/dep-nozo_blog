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

type PublicFindByNameQuery struct {
	AssociatedWith AssociatedType
}

type TopicDto struct {
	TopicID uuid.UUID `json:"topicId"`
	Name string `json:"name"`
	Description string `json:"description"`
	Articles []entity.Article `json:"articles"`
}

type TopicListDto struct {
	Topics []TopicDto `json:"topics"`
}

type TopicQueryService interface {
	PublicList(query TopicQuery) (TopicListDto, error)
	PublicFindByName(name string, query PublicFindByNameQuery) (TopicDto, error)
}
