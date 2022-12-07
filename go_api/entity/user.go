package entity

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidUser = errors.New("A User has to have an a valid user")
	ErrInvalidEmail = errors.New("Email is Invalid")
)

func isValidEmailFormat(email string) bool {
  regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
  return regex.MatchString(email)
}

type User struct {
	ID int
	Username string
	Email string
}

func NewUser(username string, email string) (User, error)  {
	if username == "" || email == "" {
		return User{}, ErrInvalidUser
	}

	if !isValidEmailFormat(email) {
		return User{}, ErrInvalidEmail 
	}

	return User{
		Username: username,
		Email: email,
	}, nil
}
