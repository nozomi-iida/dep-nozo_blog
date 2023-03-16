package adminservice

import (
	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/domain/article/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
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

func (as *ArticleService) FindById(id uuid.UUID) (article.ArticleDto, error) {
	a, err := as.ap.FindById(id)
	return a, err
}

func (as *ArticleService) Update(articleID uuid.UUID, title string, content string, tags []string, topicID *uuid.UUID, isPublic bool) (entity.Article, error)  {
	// TODO: findByIdでentity.Articleを取得して、そのentity.ArticleをUpdateするようにする
	a := entity.Article{}	
	a.SetID(articleID)
	a.SetTitle(title)
	a.SetContent(content)
	a.SetTags(tags)
	a.SetTopicID(a.TopicID)
	if(isPublic) {
		a.Public()
	}
	_, err := as.ap.Update(a)
	if err != nil {
		return entity.Article{}, err
	}
	return a, nil
}
