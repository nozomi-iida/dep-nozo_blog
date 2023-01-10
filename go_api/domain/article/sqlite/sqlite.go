package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type SqliteRepository struct {
	db *sql.DB
}

func New(fileString string) (*SqliteRepository, error)  {
	db, err := sql.Open("sqlite3", fileString)
	
	if err != nil {
		return nil, err
	}
	return &SqliteRepository{
		db,
	}, err
}

func (sr *SqliteRepository) Create(a entity.Article) (entity.Article, error) {
	tx, err := sr.db.Begin()
	_, err = tx.Exec("INSERT INTO articles(article_id, title, content, published_at, author_id, topic_id) VALUES (?, ?, ?, ?, ?, ?)", a.ArticleID, a.Title, a.Content, a.PublishedAt, a.ArticleID, a.TopicID)
	if err != nil {
		tx.Rollback()
		return entity.Article{}, article.ErrFailedToCreateArticle
	}
	for _, tag := range a.Tags {
		_, err = tx.Exec("INSERT INTO article_tags(article_id, name) VALUES (?, ?)", a.ArticleID, tag)	
		if err != nil {
			tx.Rollback()
			return entity.Article{}, article.ErrFailedToCreateArticle
		}
	}

	tx.Commit()
	return a, nil
}

func (sr *SqliteRepository) FindById(id uuid.UUID) (article.ArticleDto, error) {
	var ad article.ArticleDto 

	err := sr.db.QueryRow("SELECT article_id FROM articles;", id).Scan(
		&ad.ArticleID, 
	)
	if err != nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}
	
	return ad, nil	
}
