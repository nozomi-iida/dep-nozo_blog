package sqlite

import (
	"database/sql"

	"github.com/nozomi-iida/nozo_blog_go_api/domain/tag"
	"github.com/nozomi-iida/nozo_blog_go_api/entity"
)

type sqliteRepository struct {
	db *sql.DB
}

func New(fileString string) (*sqliteRepository, error) {
	db, err := sql.Open("sqlite3", fileString)

	if err != nil {
		return nil, err
	}

	return &sqliteRepository{
		db,
	}, err
}

func (sr *sqliteRepository) Create(t entity.Tag) (entity.Tag, error) {
	_, err := sr.db.Exec("INSERT INTO tags(tag_id, name) VALUES (?, ?)", t.TagID, t.Name)
	if err != nil {
		return entity.Tag{}, err
	}

	return t, nil
}

func (sr *sqliteRepository) List(query tag.TagQuery) ([]entity.Tag, error) {
	var ts []entity.Tag

	rows, err := sr.db.Query(`
		SELECT
			tags.tag_id,
			tags.name
		FROM
			tags
		WHERE
			tags.name LIKE ?
	`, "%"+query.Keyword+"%")
	if err != nil {
		return []entity.Tag{}, tag.ErrFailedToListTags
	}
	defer rows.Close()

	for rows.Next() {
		var t entity.Tag
		err := rows.Scan(&t.TagID, &t.Name)
		if err != nil {
			return []entity.Tag{}, tag.ErrFailedToListTags
		}
		ts = append(ts, t)
	}

	return ts, nil
}
