package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type SqliteRepository struct {
	db *sql.DB
}

type sqliteUser struct {
	ID uuid.UUID `json:"id"`
	Username string `json:"username"`
}

func (sc sqliteUser) ToEntity() entity.User  {
	u := entity.User{}	

	u.SetID(sc.ID)
	u.SetUsername(sc.Username)

	return u
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

func (sr *SqliteRepository) Create(u entity.User) error {
		_, err := sr.db.Exec("INSERT INTO user(username) VALUES (?)", u.Username)
		if err != nil {
			return err
		}

		return nil
}
