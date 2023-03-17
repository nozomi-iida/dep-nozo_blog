package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type sqliteRepository struct {
	db *sql.DB
}

type sqliteTopic struct {
	topicId uuid.UUID
	name string
	description string
}

func (st sqliteTopic) toEntity() entity.Topic {
	t := entity.Topic{}	
	t.SetTopicId(st.topicId)
	t.SetName(st.name)
	t.SetDescription(st.description)

	return t
}

func New(fileString string) (*sqliteRepository, error)  {
	db, err := sql.Open("sqlite3", fileString)
	
	if err != nil {
		return nil, err
	}

	return &sqliteRepository{
		db,
	}, err
}

func (sr *sqliteRepository) Create(t entity.Topic) (entity.Topic, error) {
	if sr.exist(t.Name) {
		return entity.Topic{}, topic.ErrTopicAlreadyExist
	}

	_, err := sr.db.Exec("INSERT INTO topics(topic_id, name, description) VALUES (?, ?, ?)", t.TopicID, t.Name, t.Description)
	if err != nil {
		return entity.Topic{}, err
	}

	return t, nil
}

func (sr *sqliteRepository) PublicList(q topic.TopicQuery) (topic.TopicListDto, error)  {
	var tld topic.TopicListDto

	rows, err := sr.db.Query(`
		SELECT
			topics.topic_id,
			topics.name,
			topics.description
		FROM
			topics
		WHERE topics.name LIKE ?;
	`, "%" + q.Keyword + "%")

	if err != nil {
		return topic.TopicListDto{}, topic.ErrFailedToListTopics 
	}

	for rows.Next() {
		var td topic.TopicDto

		err = rows.Scan(
			&td.TopicID,
			&td.Name,
			&td.Description,
		)	
		if err != nil {
			return topic.TopicListDto{}, topic.ErrFailedToListTopics 
		}

		if q.AssociatedWith == topic.Article {
			articleRows, err := sr.db.Query(`
				SELECT
					articles.article_id,
					articles.title,
					articles.content,
					articles.published_at,
					articles.author_id,
					articles.topic_id
				FROM
					articles
				WHERE
					articles.topic_id == ?;
			`, td.TopicID)
			if err != nil {
				return topic.TopicListDto{}, topic.ErrFailedToListTopics 
			}
			for articleRows.Next() {
				var ac entity.Article
				err = articleRows.Scan(
					&ac.ArticleID,
					&ac.Title,
					&ac.Content,
					&ac.PublishedAt,
					&ac.AuthorID,
					&ac.TopicID,
				)
				if err != nil {
					return topic.TopicListDto{}, topic.ErrFailedToListTopics 
				}
				var tg entity.Tag
				tagRows, err := sr.db.Query(`
					SELECT
						tags.tag_id,
						tags.name
					FROM
						tags
					INNER JOIN
						article_tags
					ON
						tags.tag_id == article_tags.tag_id
					WHERE	
						article_tags.article_id == ?;
				`, ac.ArticleID)	
				if err != nil {	
					return topic.TopicListDto{}, topic.ErrFailedToListTopics
				}
				for tagRows.Next() {
					err = tagRows.Scan(
						&tg.TagID,
						&tg.Name,
					)
					if err != nil {
						return topic.TopicListDto{}, topic.ErrFailedToListTopics
					}
					ac.Tags = append(ac.Tags, tg)
				}
				if ac.PublishedAt != nil {
					td.Articles = append(td.Articles, ac)
				}
			}
		}

		tld.Topics = append(tld.Topics, td)	
	}
	return tld, nil
}

func (sr *sqliteRepository) findByName(name string) (entity.Topic, error)  {
	rows, err := sr.db.Query("SELECT name FROM topics WHERE topics.name == ?", name)	
	var st sqliteTopic
	for rows.Next() {
		err := rows.Scan(&st.name)
		if err != nil {
			return entity.Topic{}, topic.ErrTopicNotFound
		}
	}
	defer rows.Close()
	t := st.toEntity()
	if err != nil || t.Name == "" {
		return entity.Topic{}, topic.ErrTopicNotFound
	}
	return t, nil
}

func (sr *sqliteRepository) exist(name string) bool {
	t, _ := sr.findByName(name)
	return t.Name != ""
}
