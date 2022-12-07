package user

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrUserNotFound = errors.New("the user was not found in the repository")
	ErrFailedToCreateUser = errors.New("failed to create the user to the repository")
	ErrFailedToUpdateUser = errors.New("failed to update the user to the repository")
)

type UserRepository interface {
	FindById() (entity.User, error)
	Create(entity.User) error
	Update(entity.User) error
}

