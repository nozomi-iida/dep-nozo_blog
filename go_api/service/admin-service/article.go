package adminservice

import (
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/domain/article/sqlite"
)

type articleConfiguration func(as *ArticleService) error

type ArticleService struct {
	ap article.ArticleRepository
}

func NewArticleService(cfgs ...articleConfiguration) (*ArticleService, error) {
	aas := &ArticleService{}
	
	for _, cfg := range cfgs {
		err := cfg(aas)
		if err != nil {
			return nil, err
		}
	}

	return aas, nil
}

func WithSqliteArticleRepository(fileString string) articleConfiguration {
	return func(aas *ArticleService) error {
		s, err := sqlite.New(fileString)
		if err != nil {
			return err
		}

		aas.ap = s

		return nil
	}
}

func (as *ArticleService) List() (article.ListArticleDto, error)  {
	aq := article.ArticleQuery{WithDraft: true}
	a, err := as.ap.List(aq)
	return a, err
}
