package service

import (
	"errors"

	"github.com/nozomi-iida/nozo_blog/domain/user"
	"github.com/nozomi-iida/nozo_blog/domain/user/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
	"github.com/nozomi-iida/nozo_blog/valueobject"
)

var (
	ErrDuplicateUsername = errors.New("Duplicate username")
)

type AuthConfiguration func(as *AuthService) error

type AuthService struct {
	users user.UserRepository
}

type AuthResponse struct {
	user entity.User
	token string
}

func NewAuthService(cfgs ...AuthConfiguration) (*AuthService, error)  {
	os := &AuthService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithSqliteUserRepository(fileString string) AuthConfiguration  {
	return func(as *AuthService) error {
		u, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		as.users = u

		return nil
	}
}

func (as *AuthService) SignUp(username string, password string) (AuthResponse, error)  {
	ps, err := valueobject.NewPassword(password)
	u, err := entity.NewUser(username, ps)
	if err != nil {
		return AuthResponse{}, err
	}

	user, err := as.users.Create(u)
	if err != nil {
		return AuthResponse{}, err
	}

	tokenString, _ := valueobject.NewJwtToken(user.ID)
	token, err := tokenString.Encode();
	if err != nil {
		return AuthResponse{}, err
	}
	
	return AuthResponse{user: user, token: token}, nil
}
