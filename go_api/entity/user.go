package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/valueobject"
	"golang.org/x/crypto/bcrypt"
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
	id uuid.UUID
	username string
	password string
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
		id: uuid.New(),
		username: username,
		password: encryptedPassword,
	}, nil
}

func (u *User) GetID() uuid.UUID {
	return u.id
}

func (u *User) SetID(id uuid.UUID) {
	u.id = id
}

func (u *User) GetUsername() string {
	return u.username
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) IsMatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))	
	return err != bcrypt.ErrMismatchedHashAndPassword
}
