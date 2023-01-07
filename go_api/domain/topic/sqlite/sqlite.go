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
