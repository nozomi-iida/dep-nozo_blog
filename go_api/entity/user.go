package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

var (
	ErrInvalidUser = errors.New("A User has to have an a valid user")
)

func isValidEmailFormat(email string) bool {
  regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  return regex.MatchString(email)
}

type User struct {
	ID uuid.UUID
	Username string
}

func NewUser(username string) (User, error)  {
	if username == "" {
		return User{}, ErrInvalidUser
	}

	return User{
		Username: username,
	}, nil
}

func (u *User) SetID(id uuid.UUID)  {
	u.ID = id
}

func (u *User) SetUsername(username string)  {
	u.Username = username
}
