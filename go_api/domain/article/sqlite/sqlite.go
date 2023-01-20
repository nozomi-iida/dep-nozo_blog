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
	_, err = tx.Exec(`
		INSERT INTO 
			articles(
				article_id, 
				title, 
				content, 
				published_at, 
				author_id, 
				topic_id
			) 
		VALUES (?, ?, ?, ?, ?, ?)`, 
		a.ArticleID, a.Title, a.Content, a.PublishedAt, a.AuthorID, a.TopicID,
	)
	if err != nil {
		tx.Rollback()
		return entity.Article{}, article.ErrFailedToCreateArticle
	}
	for _, tag := range a.Tags {
		row, err := tx.Query(`
			SELECT
				tag_id
			FROM 
				tags
			WHERE tags.name = ?`, 
			tag.Name,
		)

		if !row.Next() || err != nil {
			_, err = tx.Exec(`
				INSERT INTO 
					tags(tag_id, name)
				VALUES(?, ?)`, 
				tag.TagID, tag.Name,
			)
			if err != nil {
				tx.Rollback()
				return entity.Article{}, article.ErrFailedToCreateArticle
			}
		}
		_, err = tx.Exec(`
			INSERT INTO 
				article_tags(article_id, tag_id) 
			VALUES (?, ?)`, 
			a.ArticleID, tag.TagID,
		)	
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
			article_tags
		INNER JOIN
			tags
		ON
			article_tags.tag_id = tags.tag_id
		WHERE article_tags.article_id = ?
	`, id)

	if err != nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}
	defer rows.Close()
	for rows.Next() {
		var tagName string
		if err := rows.Scan(&tagName); err != nil {
			return article.ArticleDto{}, article.ErrArticleNotFound
		}
		ad.Tags = append(ad.Tags, tagName)
	}
	
	return ad, nil	
}

func (sr *SqliteRepository) Update(a entity.Article) (entity.Article, error) {
	tx, err := sr.db.Begin()
	_, err = sr.db.Exec(`
		UPDATE
			articles
		SET 
			title = ?,
			content = ?,
			topic_id = ?
		WHERE
			articles.article_id = ?
	`, a.Title, a.Content, a.TopicID, a.ArticleID)
	if err != nil {
		tx.Rollback()
		return entity.Article{}, article.ErrFailedToUpdateArticle
	}

	tx.Commit()

	return a, nil	
}

func (sr *SqliteRepository) Delete(id uuid.UUID) error {
	result, err := sr.db.Exec("DELETE FROM articles WHERE articles.article_id = ?", id)
	if err != nil {
		return article.ErrFailedToDeleteArticle
	}

	ra, err := result.RowsAffected()
	if err != nil || ra == 0 {
		return article.ErrArticleNotFound
	}

	return nil	
}
