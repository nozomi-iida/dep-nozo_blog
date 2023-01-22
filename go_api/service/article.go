package service

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
	as := &ArticleService{}

	for _, cfg := range cfgs {
		err := cfg(as)
		if err != nil {
			return nil, err
		}
	}

	return as, nil
}

func WithSqliteArticleRepository(fileString string) articleConfiguration {
	return func(as *ArticleService) error {
		u, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		as.ap = u

		return nil
	}
}

func (as *ArticleService) Post(title string, content string, tags []string, isPublic bool, authorId uuid.UUID) (entity.Article, error)  {
	a, err := entity.NewArticle(entity.ArticleArgument{Title: title, Content: content, Tags: tags, PublishedAt: nil, AuthorID: authorId})
	
	if err != nil {
		return entity.Article{}, err
	}
	if (isPublic) {
		a.Public()
	}
	a, err = as.ap.Create(a)
	if err != nil {
		return entity.Article{}, err
	}

	return a, nil
}

func (as *ArticleService) Update(articleId uuid.UUID, title string, content string, tags []string, isPublic bool) (entity.Article, error)  {
	_, err := as.ap.FindById(articleId)
	if err != nil {
		return entity.Article{}, err
	}
	return entity.Article{}, nil
}
