package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type SqliteRepository struct {
	db *sql.DB
}

type sqliteUser struct {
	id uuid.UUID `db:"id"`
	username string `db:"username"`
}

func (sc sqliteUser) ToEntity() entity.User  {
	u := entity.User{}	

	u.SetID(sc.id)
	u.SetUsername(sc.username)

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

func (sr *SqliteRepository) FindById(id uuid.UUID) (entity.User, error)  {
	rows, err := sr.db.Query("SELECT id, username FROM users WHERE users.id == ?", id)
	var su sqliteUser
	for rows.Next() {
		err := rows.Scan(&su.id, &su.username)
		if err != nil {
			return entity.User{}, user.ErrUserNotFound
		}
	}
	if err != nil {
		return entity.User{}, user.ErrUserNotFound
	}
	defer rows.Close()
	rows.Scan(su)

	u := su.ToEntity()

	return u, nil
}

func (sr *SqliteRepository) Create(u entity.User) (entity.User, error) {
	if sr.exist(u) {
		return entity.User{}, user.ErrUserNotFound
	}

	_, err := sr.db.Exec("INSERT INTO users(id, username, password) VALUES (?, ?, ?)", u.GetID(), u.GetUsername(), u.GetPassword()); 
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (sr *SqliteRepository) findByUsername(username string) (entity.User, error)  {
	var su sqliteUser
	err := sr.db.QueryRow("SELECT id, username FROM users WHERE users.username == ?", username).Scan(&su.id, &su.username)
	if err != nil {
		return entity.User{}, user.ErrUserNotFound
	}
	u := su.ToEntity()

	return u, nil
}

func (sr *SqliteRepository)exist(user entity.User) bool  {
	us, err := sr.findByUsername(user.GetUsername())
	fmt.Printf("get user %v, error: %v \n", us, err)
	if err != nil {
		return false
	} else {
		return true
	}
}
