package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type SqliteRepository struct {
	db *sql.DB
	users map[uuid.UUID]entity.User
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
		make(map[uuid.UUID]entity.User),
	}, err
}

func (sr *SqliteRepository) Create(u entity.User) (entity.User, error) {
	if sr.users == nil {
		sr.users = make(map[uuid.UUID]entity.User)
	}
	// このハンドリング方法あってるのかな？
	_, ok := sr.users[u.GetID()]
	if ok {
		return entity.User{},fmt.Errorf("user already exists: %w", user.ErrFailedToCreateUser)
	}

	encryptedPassword, err := u.Password.Encrypt()
	if err != nil {
		return entity.User{}, err 
	}

	if _, err := sr.db.Exec("INSERT INTO users(id, username, password) VALUES (?, ?, ?)", u.ID, u.Username, encryptedPassword.Value); err != nil {
		return entity.User{}, err
	}

	return u, nil
}
