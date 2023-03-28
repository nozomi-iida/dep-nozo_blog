package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog_go_api/valueobject"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUser      = errors.New("A User has to have an a valid user")
	ErrTooShortPassword = errors.New("Password is too short")
)

func isValidEmailFormat(email string) bool {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return regex.MatchString(email)
}

// もしかしたらjsonにするのはpresentation層かもな
type User struct {
	// encodeをuserに実装すればいいと思う
	UserId   valueobject.JwtToken `json:"userId"`
	Username string               `json:"username"`
	password string
}

func NewUser(username string, password valueobject.Password) (User, error) {
	if username == "" {
		return User{}, ErrInvalidUser
	}

	encryptedPassword, err := password.Encrypt()
	if err != nil {
		return User{}, err
	}
	jt, err := valueobject.NewJwtToken(uuid.New())
	if err != nil {
		return User{}, err
	}

	return User{
		UserId:   jt,
		Username: username,
		password: encryptedPassword,
	}, nil
}

func (u *User) GetID() uuid.UUID {
	return u.UserId.ID
}

func (u *User) SetID(id uuid.UUID) {
	jw, err := valueobject.NewJwtToken(id)
	if err != nil {
		return
	}
	u.UserId = jw
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) IsMatchPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
	return err != bcrypt.ErrMismatchedHashAndPassword
}
