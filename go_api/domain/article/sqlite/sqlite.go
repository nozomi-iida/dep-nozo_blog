package sqlite

import (
	"database/sql"

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

func (sr *SqliteRepository) Create(a entity.Article) (entity.Article, error)  {
	tx, _ := sr.db.Begin()
	_, err := sr.db.Exec("INSERT INTO articles(article_id, title, content, published_at, author_id) VALUES (?, ?, ?, ?, ?)", a.ArticleID, a.Title, a.Content, a.PublishedAt, a.ArticleID)
	for _, tag := range a.Tags {
		_, err = sr.db.Exec("INSERT INTO article_tags(article_id, name) VALUES (?, ?)", a.ArticleID, tag)	
	}
	if err != nil {
		tx.Rollback()
		return entity.Article{}, article.ErrFailedToCreateArticle
	}

	tx.Commit()
	return a, nil
}
