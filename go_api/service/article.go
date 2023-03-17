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

func (as *ArticleService) Post(title string, content string, tagNames []string, isPublic bool, authorId uuid.UUID, topicID *uuid.UUID) (entity.Article, error)  {
	var tags []entity.Tag
	for _, t := range tagNames {
		tag, err := entity.NewTag(t)
		if err != nil {
			return entity.Article{}, err
		}
		tags = append(tags, tag)
	}

	a, err := entity.NewArticle(entity.ArticleArgument{Title: title, Content: content, Tags: tags, PublishedAt: nil, AuthorID: authorId, TopicID: topicID})
	
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



func (as *ArticleService) Delete(articleID uuid.UUID) error {
	err := as.ap.Delete(articleID)
	if err != nil {
		return err
	}

	return nil
}

func (as *ArticleService) FindById(id uuid.UUID) (article.ArticleDto, error) {
	a, err := as.ap.FindById(id)
	if a.PublishedAt == nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}
	return a, err
}


func (as *ArticleService) List(query article.ArticleQuery) (article.ListArticleDto, error) {
	a, err := as.ap.List(query)
	return a, err
}
