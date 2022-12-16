package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

var (
	ErrInvalidUser = errors.New("A User has to have an a valid user")
	ErrTooShortPassword = errors.New("Password is too short")
)

func isValidEmailFormat(email string) bool {
  regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  return regex.MatchString(email)
}

type User struct {
	ID uuid.UUID
	Username string
	Password string
}

func NewUser(username string, password valueobject.Password) (User, error)  {
	if username == "" {
		return User{}, ErrInvalidUser
	}
	
	encryptedPassword, err := password.Encrypt()
	if err != nil {
		return User{}, err
	}
	
	return User{
		ID: uuid.New(),
		Username: username,
		Password: encryptedPassword,
	}, nil
}

func (u *User) GetID() uuid.UUID  {
	return u.ID
}

func (u *User) SetID(id uuid.UUID)  {
	u.ID = id
}

func (u *User) SetUsername(username string)  {
	u.Username = username
}
