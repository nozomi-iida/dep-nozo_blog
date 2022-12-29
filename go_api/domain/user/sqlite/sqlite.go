package sqlite

import (
	"database/sql"

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
	password string `db:"password"`
}

func (sc sqliteUser) ToEntity() entity.User  {
	u := entity.User{}	

	u.SetID(sc.id)
	u.SetUsername(sc.username)

	if sc.password != "" {
		u.SetPassword(sc.password)
	}

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

	u := su.ToEntity()

	return u, nil
}

func (sr *SqliteRepository) FindByUsername(username string) (entity.User, error)  {
	rows, err := sr.db.Query("SELECT id, username, password FROM users WHERE users.username == ?", username)
	var su sqliteUser
	for rows.Next() {
		err := rows.Scan(&su.id, &su.username, &su.password)
		if err != nil {
			return entity.User{}, user.ErrUserNotFound
		}
	}
	if err != nil {
		return entity.User{}, user.ErrUserNotFound
	}
	defer rows.Close()
	u := su.ToEntity()
	u.SetPassword(su.password)
	return u, nil	
}

func (sr *SqliteRepository) Create(u entity.User) (entity.User, error) {
	if sr.exist(u.GetUsername()) {
		return entity.User{}, user.ErrUserAlreadyExist
	}

	_, err := sr.db.Exec("INSERT INTO users(id, username, password) VALUES (?, ?, ?)", u.GetID(), u.GetUsername(), u.GetPassword()); 
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}

func (sr *SqliteRepository)exist(username string) bool  {
	us, _ := sr.FindByUsername(username)
	return us.GetUsername() != ""
}
