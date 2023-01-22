package sqlite

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/libs"
	"github.com/simukti/sqldb-logger/logadapter/zapadapter"
	sqldblogger "github.com/simukti/sqldb-logger"
)

type SqliteRepository struct {
	db *sql.DB
}

func New(fileString string) (*SqliteRepository, error)  {
	db, err := sql.Open("sqlite3", fileString)
	
	db = sqldblogger.OpenDriver(fileString, db.Driver(), zapadapter.New(libs.ZipLogger()))
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

func (sr *SqliteRepository) List(q article.ArticleQuery) (article.ListArticleDto, error)  {
	var rs article.ListArticleDto
	var articleIDs []interface{}
	articleMap := make(map[uuid.UUID]article.ArticleDto) 

	rows, err := sr.db.Query(`
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
		WHERE articles.published_at is NOT NULL
		AND articles.title LIKE ?;
	`, "%" + q.Keyword + "%")

	if err != nil {
		return article.ListArticleDto{}, article.ErrFailedToListArticle
	}

	for rows.Next() {
		var ad article.ArticleDto
		err = rows.Scan(
			&ad.ArticleID, 
			&ad.Title, 
			&ad.Content, 
			&ad.PublishedAt, 
			&ad.Topic,
			&ad.AuthorName,
		)
		if err != nil {
			return article.ListArticleDto{}, article.ErrFailedToListArticle
		}

		articleIDs = append(articleIDs, ad.ArticleID.String())
		articleMap[ad.ArticleID] = ad
	}

	if len(articleMap) > 0 {
		repeat := strings.Repeat("?,", len(articleIDs)-1) +"?"
		rows, err = sr.db.Query(`
			SELECT
				tags.name,
				article_tags.article_id
			FROM
				article_tags
			INNER JOIN
				tags
			ON
				article_tags.tag_id = tags.tag_id
			WHERE article_tags.article_id IN ( `+ repeat +`);
		`, articleIDs... 
		)

		if err != nil {
			return article.ListArticleDto{}, article.ErrFailedToListArticle
		}
		defer rows.Close()
		for rows.Next() {
			var tagName string
			var articleID uuid.UUID 
			err = rows.Scan(&tagName, &articleID)
			if err != nil {
				return article.ListArticleDto{}, article.ErrFailedToListArticle
			}
			ar := articleMap[articleID]
			ar.Tags = append(ar.Tags, tagName)
			articleMap[articleID] = ar
		}
	}

	for _, ac := range articleMap {
		rs.Articles = append(rs.Articles, ac)
	}

	return rs, nil
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
