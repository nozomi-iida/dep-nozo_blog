package sqlite

import (
	"database/sql"
	"fmt"

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
	_, err = tx.Exec("INSERT INTO articles(article_id, title, content, published_at, author_id, topic_id) VALUES (?, ?, ?, ?, ?, ?)", a.ArticleID, a.Title, a.Content, a.PublishedAt, a.AuthorID, a.TopicID)
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

	err := sr.db.QueryRow(`
		SELECT 
			articles.article_id,
			articles.title,
			articles.content,
			articles.published_at,
			topics.name as topic,
			authors.username as authorName
		FROM 
			articles 
		INNER JOIN
			topics
		ON
			topics.topic_id = articles.topic_id
		INNER JOIN
			users as authors
		ON
			authors.user_id = articles.author_id
		WHERE articles.article_id = ?
	`, id).Scan(
		&ad.ArticleID, 
		&ad.Title, 
		&ad.Content, 
		&ad.PublishedAt, 
		&ad.Topic,
		&ad.AuthorName,
	)
	if err != nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}

	rows, err := sr.db.Query(`
		SELECT
			tags.name
		FROM
			article_tags as tags
		WHERE tags.article_id = ?
	`, id)

	if err != nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}
	defer rows.Close()
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return article.ArticleDto{}, article.ErrArticleNotFound
		}
		ad.Tags = append(ad.Tags, tag)
	}
	
	return ad, nil	
}
