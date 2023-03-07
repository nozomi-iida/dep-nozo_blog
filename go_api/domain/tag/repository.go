package tag

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrFailedToListTags = errors.New("failed to get the tags to the repository")
)

type TagQuery struct {
	Keyword string
}

type TagRepository interface {
	Create(tag entity.Tag) (entity.Tag, error)
	List(TagQuery) ([]entity.Tag, error)
}
