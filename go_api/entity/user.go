package entity

import "errors"

var (
	ErrInvalidUser = errors.New("A User has to have an a valid user")
)

type User struct {
	ID int
	Username string
	Email string
}

func NewUser(username string, email string) (User, error)  {
	if username == "" || email == "" {
		return User{}, ErrInvalidUser
	}

	return User{
		Username: username,
		Email: email,
	}, nil
}
