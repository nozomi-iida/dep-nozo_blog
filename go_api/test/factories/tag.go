package factories

import (
	"testing"

	"github.com/nozomi-iida/nozo_blog_go_api/domain/tag/sqlite"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
)

type tagOptions func(*entity.Tag)

func SetTagName(name string) tagOptions {
	return func(t *entity.Tag) {
		t.SetName(name)
	}
}

func CreateTag(t *testing.T, fileName string, options ...tagOptions) entity.Tag {
	tag, err := entity.NewTag("testTag")
	for _, op := range options {
		op(&tag)
	}

	sq, err := sqlite.New(fileName)
	_, err = sq.Create(tag)
	if err != nil {
		t.Error("create tag err:", err)
	}

	return tag
}
