package user

import (
	"errors"

	"github.com/google/uuid"
	"github.com/nozomi-iida/nozo_blog/entity"
)

var (
	ErrUserNotFound = errors.New("the user was not found in the repository")
	ErrFailedToFindAllUser = errors.New("failed to findAll the user to the repository")
	ErrFailedToCreateUser = errors.New("failed to create the user to the repository")
	ErrFailedToUpdateUser = errors.New("failed to update the user to the repository")
)

type UserRepository interface {
	FindById(id uuid.UUID) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
}
