package service

import (
	"github.com/nozomi-iida/nozo_blog/domain/tag"
	"github.com/nozomi-iida/nozo_blog/domain/tag/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type tagConfigurations func(*TagService) error

type TagService struct {
	tg tag.TagRepository
}

func NewTagService(cfgs ...tagConfigurations) (*TagService, error) {
	tg := &TagService{}

	for _, cfg := range cfgs {
		err := cfg(tg)
		if err != nil {
			return nil, err 
		}
	}

	return tg, nil
}

func WithSqliteTagRepository(fileString string) tagConfigurations {
	return func(tg *TagService) error {
		s, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		tg.tg = s

		return nil
	}	
}

func (tg *TagService) List(query tag.TagQuery) ([]entity.Tag, error)  {
	tags, err := tg.tg.List(query)
	if err != nil {
		return []entity.Tag{}, err
	}

	return tags, nil
}
