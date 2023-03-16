package sqlite

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/article"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/libs"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zapadapter"
)

type SqliteRepository struct {
	db *sql.DB
}

type QueryArticleSqlite struct {
	ArticleID uuid.UUID 
	Title string 
	Content string 
	PublishedAt *time.Time
	Tags []entity.Tag 
	TopicID uuid.NullUUID
	TopicName sql.NullString
	TopicDescription sql.NullString
	Author entity.User 
}

func (qa QueryArticleSqlite) ToDto() article.ArticleDto {
	ad := article.ArticleDto{
		ArticleID: qa.ArticleID,
		Title: qa.Title,
		Content: qa.Content,
		PublishedAt: qa.PublishedAt,
		Tags: qa.Tags,
		Author: qa.Author,
	}	

	if qa.TopicID.Valid {
		var topic = entity.Topic{}
		topic.TopicID = qa.TopicID.UUID
		topic.Name = qa.TopicName.String
		topic.Description = qa.TopicDescription.String
		ad.Topic = &topic
	}


	return ad
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
	defer tx.Commit()
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
	err = createArticleTags(tx, a.ArticleID, a.Tags)

	return a, nil
}

// TODO: tagの更新処理をやる
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

	// rows, err := tx.Query(`
	// 	SELECT
	// 		article_tags.article_id,
	// 		article_tags.tag_id
	// 	FROM
	// 		article_tags
	// 	WHERE
	// 		article_tags.article_id = ?`,
	// 	a.ArticleID,
	// )
	// if err != nil {
	// 	tx.Rollback()
	// 	return entity.Article{}, article.ErrFailedToUpdateArticle
	// }

	// for rows.Next() {
	// 	var articleID uuid.UUID
	// 	var tagID uuid.UUID
	// 	err := rows.Scan(&articleID, &tagID)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return entity.Article{}, article.ErrFailedToUpdateArticle
	// 	}
	// 	_, err = tx.Exec(`
	// 		DELETE FROM 
	// 			article_tags
	// 		WHERE
	// 			article_tags.article_id = ? AND article_tags.tag_id = ?`,
	// 		articleID, tagID,
	// 	)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return entity.Article{}, article.ErrFailedToUpdateArticle
	// 	}
	// }

	// err = createArticleTags(tx, a.ArticleID, a.Tags)
	// if err != nil {
	// 	tx.Rollback()
	// 	return entity.Article{}, article.ErrFailedToUpdateArticle
	// }

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
	withDraftQuery := ""
	if !q.WithDraft {
		withDraftQuery = "AND articles.published_at IS NOT NULL"
	} 

	rows, err := sr.db.Query(`
		SELECT 
			articles.article_id,
			articles.title,
			articles.content,
			articles.published_at,
			topics.topic_id,
			topics.name,
			topics.description,
			authors.user_id,
			authors.username
		FROM 
			articles 
		LEFT JOIN
			topics
		ON
			articles.topic_id = topics.topic_id
		LEFT JOIN
			users as authors
		ON
		articles.author_id = authors.user_id
		WHERE articles.title LIKE ? ` + withDraftQuery + `;
	`, "%" + q.Keyword + "%")

	if err != nil {
		return article.ListArticleDto{}, article.ErrFailedToListArticle
	}

	for rows.Next() {
		var qa QueryArticleSqlite
		err = rows.Scan(
			&qa.ArticleID, 
			&qa.Title, 
			&qa.Content, 
			&qa.PublishedAt, 
			&qa.TopicID,
			&qa.TopicName,
			&qa.TopicDescription,
			&qa.Author.UserId.ID,
			&qa.Author.Username,
		)
		if err != nil {
			return article.ListArticleDto{}, article.ErrFailedToListArticle
		}

		articleIDs = append(articleIDs, qa.ArticleID.String())
		articleMap[qa.ArticleID] = qa.ToDto()
	}

	if len(articleMap) > 0 {
		repeat := strings.Repeat("?,", len(articleIDs)-1) +"?"
		rows, err = sr.db.Query(`
			SELECT
				tags.tag_id,
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
			var tag entity.Tag
			var articleID uuid.UUID 
			err = rows.Scan(&tag.TagID, &tag.Name, &articleID)
			if err != nil {
				return article.ListArticleDto{}, article.ErrFailedToListArticle
			}
			ar := articleMap[articleID]
			ar.Tags = append(ar.Tags, tag)
			articleMap[articleID] = ar
		}
	}

	for _, ac := range articleMap {
		rs.Articles = append(rs.Articles, ac)
	}

	return rs, nil
}

func (sr *SqliteRepository) FindById(id uuid.UUID) (article.ArticleDto, error) {
	var qa QueryArticleSqlite 

	err := sr.db.QueryRow(`
		SELECT 
			articles.article_id,
			articles.title,
			articles.content,
			articles.published_at,
			topics.topic_id,
			topics.name,
			topics.description,
			authors.user_id,
			authors.username
		FROM 
			articles 
		LEFT JOIN
			topics
		ON
		articles.topic_id = topics.topic_id
		LEFT JOIN
			users as authors
		ON
			articles.author_id = authors.user_id
		WHERE articles.article_id = ?;
	`, id).Scan(
		&qa.ArticleID, 
		&qa.Title, 
		&qa.Content, 
		&qa.PublishedAt, 
		&qa.TopicID,
		&qa.TopicName,
		&qa.TopicDescription,
		&qa.Author.UserId.ID,
		&qa.Author.Username,
	)

	if err != nil {
		return article.ArticleDto{}, article.ErrArticleNotFound
	}

	rows, err := sr.db.Query(`
		SELECT
			tags.tag_id,
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
		var tag entity.Tag
		if err := rows.Scan(&tag.TagID, &tag.Name); err != nil {
			return article.ArticleDto{}, article.ErrArticleNotFound
		}
		qa.Tags = append(qa.Tags, tag)
	}

	ad := qa.ToDto()
	
	return ad, nil	
}

func createArticleTags(tx *sql.Tx, ai uuid.UUID, tags []entity.Tag) error  {
	for _, tag := range tags {
		rows, err := tx.Query(`
			SELECT
				tag_id
			FROM 
				tags
			WHERE tags.name = ?`, 
			tag.Name,
		)

		if !rows.Next() {
			_, err := tx.Exec(`
				INSERT INTO 
					tags(tag_id, name)
				VALUES(?, ?)`, 
				tag.TagID, tag.Name,
			)
			if err != nil {
				tx.Rollback()
				return article.ErrFailedToCreateArticle
			}
		}

		_, err = tx.Exec(`
			INSERT INTO 
				article_tags(article_id, tag_id) 
			VALUES (?, ?)`, 
			ai, tag.TagID,
		)	
		if err != nil {
			tx.Rollback()
			return article.ErrFailedToCreateArticle
		}
	}

	return nil
}
