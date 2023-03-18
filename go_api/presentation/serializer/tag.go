package serializer

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type TagJson struct {
	TagID uuid.UUID `json:"tagId"`
	Name  string    `json:"name"`
}

func Tag2Json(tag entity.Tag) TagJson {
	return TagJson{
		TagID: tag.TagID,
		Name:  tag.Name,
	}
}
